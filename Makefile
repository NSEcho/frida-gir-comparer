all:
	@go build

run:
	@go run main.go

clean:
	@rm -rf frida-core-*