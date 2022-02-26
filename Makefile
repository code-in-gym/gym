check:
	go vet ./... &&  markdownlint '**/*.md' && golangci-lint run && go test ./... -count=1

test:
	go test ./... -count=1
