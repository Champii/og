.PHONY: all clean

all:
	@go generate
	@go build

run: all
	@./Og exemples/test.og

clean:
	@go clean
