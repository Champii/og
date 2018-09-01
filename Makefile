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

all: grammar build bootstrap

grammar:
	@$(call run_and_test,go generate,Generating grammar)

bootstrap:
	@$(call run_and_test,./og -o lib src,Transforming [og -> go])
	@$(call run_and_test,go build,Recompiling new go files)
	@$(call run_and_test,go test og/tests,Testing)

rebootstrap:
	@$(call run_and_test,og -o lib src,Transforming [og -> go] from previous build)
	@$(call run_and_test,go build,Compiling new go files)
	@$(call run_and_test,go test og/tests,Testing)
	@$(call run_and_test,./og -o lib src,Transforming [og -> go])
	@$(call run_and_test,go build,Recompiling new go files)
	@$(call run_and_test,go test og/tests,Testing)

build:
	@$(call run_and_test,go build,Building go src)

test:
	@$(call run_and_test,go test og/tests,Testing)

install:
	@$(call run_and_test,cp ./og /usr/bin,Installing into /usr/bin/og)

clean:
	@go clean
