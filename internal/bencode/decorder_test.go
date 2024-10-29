package bencode

import (
	"testing"
)

func TestInvalidBencode(t *testing.T) {
	sampleData := []string{"Idsahdgse", "sadasd", "", "i34de"}
	for _, key := range sampleData {
		t.Run("InvalidBencode:"+key, func(t *testing.T) {
			_, _, err := Decode(key, 0)
			if err == nil {
				t.Errorf("expected error for invalid bencode input: %v", key)
			}
		})
	}
}

func TestBencodeInteger(t *testing.T) {
	validSamples := map[string]int{
		"i42e":    42,
		"i75653e": 75653,
		"i-4343e": -4343,
		"i0e":     0,
	}
	for input, expected := range validSamples {
		t.Run("BencodeIntegerValid:"+input, func(t *testing.T) {
			res, _, err := Decode(input, 0)
			if err != nil {
				t.Fatalf("unexpected error for input %v: %v", input, err)
			}
			if result, ok := res.(int); !ok || result != expected {
				t.Errorf("got %v, expected %v", result, expected)
			}
		})
	}

	invalidSamples := []string{
		"i4d2e",
		"i756sdsd53e",
		"i-43sdsa43e",
	}
	for _, input := range invalidSamples {
		t.Run("BencodeIntegerInvalid:"+input, func(t *testing.T) {
			_, _, err := Decode(input, 0)
			if err == nil {
				t.Errorf("expected error for invalid integer input: %v", input)
			}
		})
	}
}

func TestBencodeString(t *testing.T) {
	validSamples := map[string]string{
		"5:hello":        "hello",
		"0:":             "",
		"11:hello world": "hello world",
	}
	for input, expected := range validSamples {
		t.Run("BencodeStringValid:"+input, func(t *testing.T) {
			res, _, err := Decode(input, 0)
			if err != nil {
				t.Fatalf("unexpected error for input %v: %v", input, err)
			}
			if result, ok := res.(string); !ok || result != expected {
				t.Errorf("got %v, expected %v", result, expected)
			}
		})
	}

	invalidSamples := []string{
		"5:hell",
		"6:32322",
		"11:world",
		":5757",
		":sahdh",
		"5orbit",
	}
	for _, input := range invalidSamples {
		t.Run("BencodeStringInvalid:"+input, func(t *testing.T) {
			_, _, err := Decode(input, 0)
			if err == nil {
				t.Errorf("expected error for invalid string input: %v", input)
			}
		})
	}
}

func TestBencodeList(t *testing.T) {
	validSamples := map[string][]interface{}{
		"l5:helloi52ee":   {"hello", 52},
		"le":              {},
		"l4:novai-14eiee": {"nova", -14, 0},
	}
	for input, expected := range validSamples {
		t.Run("BencodeListValid:"+input, func(t *testing.T) {
			res, _, err := Decode(input, 0)
			if err != nil {
				t.Fatalf("unexpected error for input %v: %v", input, err)
			}
			if result, ok := res.([]interface{}); !ok || len(result) != len(expected) {
				t.Fatalf("got %v, expected %v", result, expected)
			} else {
				for i := range result {
					if result[i] != expected[i] {
						t.Errorf("at index %d, got %v, expected %v", i, result[i], expected[i])
					}
				}
			}
		})
	}

	invalidSamples := []string{
		"l5:hellodi52ee",
		"li67e",
		"l4:novai-14eie",
	}
	for _, input := range invalidSamples {
		t.Run("BencodeListInvalid:"+input, func(t *testing.T) {
			_, _, err := Decode(input, 0)
			if err == nil {
				t.Errorf("expected error for invalid list input: %v", input)
			}
		})
	}
}

func TestBencodeDict(t *testing.T) {
	validSamples := map[string]map[string]interface{}{
		"d3:foo3:bare":             {"foo": "bar"},
		"d3:foo3:bar3:bazi42ee":    {"foo": "bar", "baz": 42},
		"d4:name5:apple3:agei25ee": {"name": "apple", "age": 25},
		"de":                       {},
	}

	for input, expected := range validSamples {
		t.Run("BencodeDictValid:"+input, func(t *testing.T) {
			res, end, err := Decode(input, 0)
			t.Logf("end \t %v \t %v", end, len(input))
			if err != nil {
				t.Fatalf("unexpected error for input %v: %v", input, err)
			}
			result, ok := res.(map[string]interface{})
			if !ok {
				t.Fatalf("expected map[string]interface{}, got %T", res)
			}

			if len(result) != len(expected) {
				t.Fatalf("expected %d keys, got %d", len(expected), len(result))
			}

			for key, expectedValue := range expected {
				value, exists := result[key]
				if !exists {
					t.Errorf("key %v missing in result", key)
				}
				if value != expectedValue {
					t.Errorf("for key %v, expected %v, got %v", key, expectedValue, value)
				}
			}
		})
	}

	invalidSamples := []string{
		// "d3:foo3:barextra",
		"d5orbite",
		"d3:foo3:bar3:bazi42e",
		"d3:foo3:bar3:bazi42",
		"d4:name5:apple3:age",
		"d3:foo",
	}

	for _, input := range invalidSamples {
		t.Run("BencodeDictInvalid:"+input, func(t *testing.T) {
			_, _, err := Decode(input, 0)
			if err == nil {
				t.Errorf("expected error for invalid dictionary input: %v", input)
			}
		})
	}
}
