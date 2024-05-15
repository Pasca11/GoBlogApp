.SILENT:

export GOPROXY=https://proxy.golang.org

export GOSUMDB=on

.PHONY:

run:
	go run cmd/blog/main.go
