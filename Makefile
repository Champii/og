.PHONY: all clean

all:
	@go generate
	@go build

run: all
	@./og_new tests/toto.og

clean:
	@go clean
