.PHONY: all build clean

COM_COLOR   = \033[0;34m
OBJ_COLOR   = \033[0;36m
MAGENTA_COLOR  = \033[0;35m
CYAN_COLOR  = \033[0;36m
OK_COLOR    = \033[0;32m
ERROR_COLOR = \033[0;31m
WARN_COLOR  = \033[0;33m
NO_COLOR    = \033[m

OK_STRING    = [OK]
ERROR_STRING = [ERROR]
WARN_STRING  = [WARNING]
COM_STRING   = [Compiling]

define title
	printf "%b" "\n$(CYAN_COLOR)=> $(WARN_COLOR)$(1)$(NO_COLOR)\n";
endef

define section_title
	printf "%b" "$(MAGENTA_COLOR)-> $(OK_COLOR)$(1)$(NO_COLOR)\n";
endef

define run_and_test
	printf "%b" "- $(OBJ_COLOR)$(2)$(NO_COLOR)\r"; \
	$(1) 2> $@.log; \
	RESULT=$$?; \
	if [ $$RESULT -ne 0 ]; then \
		printf '%-60b%b' '- $(OBJ_COLOR)$(2)' '$(ERROR_COLOR)$(ERROR_STRING)$(NO_COLOR)\n'; \
	elif [ -s $@.log ]; then \
		printf '%-60b%b' '- $(OBJ_COLOR)$(2)' '$(WARNING_COLOR)$(WARNING_STRING)$(NO_COLOR)\n'; \
	else \
		printf '%-60b%b' '- $(OBJ_COLOR)$(2)' '$(OK_COLOR)$(OK_STRING)$(NO_COLOR)\n'; \
	fi; \
	cat $@.log; \
	rm -f $@.log; \
	exit $$RESULT
endef

SRC_PATH=lib/
SRC=$(wildcard $(SRC_PATH)og/*.og $(SRC_PATH)translator/*.og $(SRC_PATH)ast/*.og $(SRC_PATH)ast/walker/*.og $(SRC_PATH)ast/common/*.og)
RES=$(SRC:.og=.go)
EXE=og
CC=og

all: grammar build

grammar: parser/*.go
parser/*.go: parser/Og.g4
	@$(call section_title,Grammar)
	@$(call run_and_test,go generate,Generating parser from $<)

build:
	@$(call title,Building from `$(CC) -v`)
	@$(CC) lib
	@make test -s --no-print-directory

$(EXE): $(SRC) $(RES)
	@make test --no-print-directory
# 	@$(call section_title,Building Binary)
# 	@go build

re:
	@$(call title,Full Re-Bootstrap)
	@make re_ --no-print-directory

re_: grammar
	@make clean build --no-print-directory
	@make CC='./og' clean new --no-print-directory

new:
	@make CC='./og' all --no-print-directory

test:
	@$(call section_title,Testing)
	@go test ./tests

doc:
	@docsify serve ./docs

clean:
	@$(call section_title,Cleaning src folder)
	@rm -f $(RES)
	@rm -f tests/exemples/*.go

# lib/%.go: src/%.og
# 	@$(call run_and_test,$(CC) -o lib $?,Compiling $<)
