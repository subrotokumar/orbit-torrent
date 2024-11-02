version:
	@go run cmd/main.go version

decode:
	@go run cmd/main.go decode --input i45e

info:
	@go run cmd/main.go info --file sample.torrent

peers:
	@go run cmd/main.go peers --file sample.torrent

test:
	@go test ./... -cover