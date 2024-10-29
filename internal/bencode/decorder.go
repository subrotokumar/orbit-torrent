package bencode

import (
	"fmt"
	"io"
	"unicode/utf8"
)

func Decode(bencodedString string, start int) (interface{}, int, error) {
	if start == len(bencodedString) {
		return nil, start, io.ErrUnexpectedEOF
	}
	first := bencodedString[start]

	switch {
	case first == 'i':
		return DecodeInteger(bencodedString, start)
	case first == 'l':
		return DecodeList(bencodedString, start)
	case first == 'd':
		return DecodeDict(bencodedString, start)
	case first >= '0' && first <= '9':
		return DecodeString(bencodedString, start)
	default:
		return nil, start, fmt.Errorf("bad bencode format")
	}
}

// Lists are encoded as l<bencoded_elements>e.
//
// For example, ["hello", 52] would be encoded as l5:helloi52ee. Note that there are no separators between the elements.
func DecodeList(bencodedString string, start int) (result []interface{}, index int, err error) {
	index = start
	index++

	result = make([]interface{}, 0)

	for {
		if index >= len(bencodedString) {
			return nil, start, fmt.Errorf("bad list")
		}
		if bencodedString[index] == 'e' {
			index++
			break
		}
		var x interface{}
		x, index, err = Decode(bencodedString, index)
		if err != nil {
			return nil, index, err
		}

		result = append(result, x)
	}
	return result, index, nil
}

// Integers are encoded as i<number>e.
//
// For example, 42 is encoded as i42e and -42 is encoded as i-42e.
func DecodeInteger(bencodedString string, start int) (int, int, error) {
	index := start + 1

	isNegetive := false
	if bencodedString[index] == '-' {
		isNegetive = true
		index++
	}

	result := 0
	for index < len(bencodedString) && bencodedString[index] >= '0' && bencodedString[index] <= '9' {
		result = result*10 + int(bencodedString[index]-'0')
		index++
	}

	if index == len(bencodedString) || bencodedString[index] != 'e' {
		return 0, start, fmt.Errorf("bad int")
	}

	index++

	if isNegetive {
		result = -result
	}

	return result, index, nil
}

// Strings are encoded as <length>:<contents>.
//
// For example, the string "hello" is encoded as "5:hello".
func DecodeString(bencodedString string, start int) (string, int, error) {
	runes := []rune(bencodedString)
	index := start
	length := 0
	for index < len(runes) && runes[index] >= '0' && runes[index] <= '9' {
		length = length*10 + (int(runes[index]) - '0')
		index++
	}

	if index == len(runes) || runes[index] != ':' {
		return "", start, fmt.Errorf("bad string")
	}
	index++

	if index+length > len(runes) {
		return "", start, fmt.Errorf("bad string: out of bounds")
	}

	result := string(runes[index : index+length])
	normalizeString := func(input string) string {
		var result []rune

		for _, r := range input {
			if r == '\ufffd' || !utf8.ValidRune(r) {
				result = append(result, '\ufffd') // Replace with Unicode replacement character
			} else {
				result = append(result, r)
			}
		}

		return string(result)
	}
	result = normalizeString(result)
	index += length
	return string(result), index, nil
}

// A dictionary is encoded as d<key1><value1>...<keyN><valueN>e. <key1>, <value1> etc. correspond to the bencoded keys & values. The keys are sorted in lexicographical order and must be strings.
//
// For example, {"hello": 52, "foo":"bar"} would be encoded as: d3:foo3:bar5:helloi52ee (note that the keys were reordered).
func DecodeDict(bencodedString string, start int) (result map[string]interface{}, index int, err error) {
	index = start + 1
	result = make(map[string]interface{})
	for {
		if index >= len(bencodedString) {
			return nil, start, fmt.Errorf("bad list")
		}
		if bencodedString[index] == 'e' {
			index++
			break
		}

		pairst := index
		var key, val interface{}
		key, index, err = Decode(bencodedString, index)
		if err != nil {
			return nil, index, err
		}
		keys, ok := key.(string)
		if !ok {
			return nil, pairst, fmt.Errorf("dict key is not a string: %q", key)
		}
		val, index, err = Decode(bencodedString, index)
		if err != nil {
			return nil, index, err
		}
		fmt.Printf("%s = %v", key, val)
		result[keys] = val
	}
	return result, index, nil
}
