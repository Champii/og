.PHONY: all clean

COM_COLOR   = \033[0;34m
OBJ_COLOR   = \033[0;36m
OK_COLOR    = \033[0;32m
ERROR_COLOR = \033[0;31m
WARN_COLOR  = \033[0;33m
NO_COLOR    = \033[m

OK_STRING    = [OK]
ERROR_STRING = [ERROR]
WARN_STRING  = [WARNING]
COM_STRING   = [Compiling]

define run_and_test
	printf "%b" "- $(OBJ_COLOR)$(2)$(NO_COLOR)\r"; \
	$(1) > $@.out 2> $@.log; \
	RESULT=$$?; \
	if [ $$RESULT -ne 0 ]; then \
		printf '%-60b%b' '- $(OBJ_COLOR)$(2)' '$(ERROR_COLOR)$(ERROR_STRING)$(NO_COLOR)\n'; \
	elif [ -s $@.log ]; then \
		printf '%-60b%b' '- $(OBJ_COLOR)$(2)' '$(WARNING_COLOR)$(WARNING_STRING)$(NO_COLOR)\n'; \
	else \
		printf '%-60b%b' '- $(OBJ_COLOR)$(2)' '$(OK_COLOR)$(OK_STRING)$(NO_COLOR)\n'; \
	fi; \
	cat $@.log; \
	cat $@.out; \
	rm -f $@.log; \
	rm -f $@.out; \
	exit $$RESULT
endef

SRC_PATH=src/
SRC=$(wildcard $(SRC_PATH)*.og $(SRC_PATH)translator/*.og)
RES=lib/
EXE=og

all: grammar bootstrap

grammar: parser/*.go
parser/*.go: parser/Og.g4
	@$(call run_and_test,go generate,Generating grammar)

bootstrap: $(RES)
$(RES): $(SRC)
	@$(call run_and_test,./og -o lib src,Transforming [og -> go])
	@make build -s
	@make test -s

re: clean grammar
	@$(call run_and_test,og -o lib src,Transforming [og -> go] from previous build)
	@make build -s
	@make test -s
	@make bootstrap -s

build:
	@$(call run_and_test,go build,Compiling go files)

test: $(EXE)
	@$(call run_and_test,go test og/tests,Testing)

clean:
	@go clean
