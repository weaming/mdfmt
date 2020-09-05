GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o mdfmt-darwin-amd64 -a
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o mdfmt-linux-amd64 -a
