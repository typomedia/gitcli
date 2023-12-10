build: tidy
	go build -ldflags "-s -w" -o dist/ .

run: tidy
	go run main.go

tidy:
	go mod tidy

check:
	go install github.com/client9/misspell/cmd/misspell@latest
	misspell -error .
	go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	gocyclo -over 14 .
	go install honnef.co/go/tools/cmd/staticcheck@latest
	staticcheck ./...
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -quiet --severity high ./...
	go install github.com/sonatype-nexus-community/nancy@latest
	go list -json -deps ./... | nancy sleuth

loc:
	go install github.com/boyter/scc/v3@latest
	scc --exclude-dir vendor --exclude-dir bin .

compile: tidy
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o dist/gitcli-linux-arm64 .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/gitcli-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/gitcli-macos-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/gitcli-win-amd64.exe .
