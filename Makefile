.PHONY: all build clean

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

define section_title
	printf "%b" "\n$(WARN_COLOR)-> $(OK_COLOR)$(1)$(NO_COLOR)\n";
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

SRC_PATH=src/
SRC=$(wildcard $(SRC_PATH)*.og $(SRC_PATH)translator/*.og)
RES=$(subst src/, lib/, $(SRC:.og=.go))
EXE=og

all: grammar build

grammar: parser/*.go
parser/*.go: parser/Og.g4
	@$(call section_title,Grammar)
	@$(call run_and_test,go generate,Generating parser from $<)
	@make clean build --no-print-directory

build:
	@$(call section_title,Oglang to Golang Compilation)
	@make $(EXE) --no-print-directory

$(EXE): $(RES)
	@$(call section_title,Building Binary)
	@$(call run_and_test,go build,Building go source)
	@make test --no-print-directory

re: clean
	@$(call section_title,Full Bbootstrap)
	@make re_ --no-print-directory

re_: grammar
	@$(call run_and_test,og -o lib src,Recompiling all sources from previous og version)
	@make build --no-print-directory
	@$(call section_title,Rebuilding with new og version)
	@make clean all --no-print-directory

test: all
	@$(call run_and_test,go test og/tests,Testing)

clean:
	@$(call section_title,Cleaning)
	@$(call run_and_test,rm -f $(RES),Delete src folder)

lib/%.go: src/%.og
	@$(call run_and_test,./og -o lib $?,Compiling $<)
