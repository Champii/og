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
		printf '%-60b%b' '$(ERROR_COLOR)$(ERROR_STRING)$(NO_COLOR)\n'; \
	elif [ -s $@.log ]; then \
		printf '%-60b%b' '$(WARNING_COLOR)$(WARNING_STRING)$(NO_COLOR)\n'; \
	else \
		printf '%-60b%b' '- $(OBJ_COLOR)$(2)' '$(OK_COLOR)$(OK_STRING)$(NO_COLOR)\n'; \
	fi; \
	cat $@.log; \
	cat $@.out; \
	rm -f $@.log; \
	rm -f $@.out; \
	exit $$RESULT
endef

all: generate build bootstrap test

generate:
	@$(call run_and_test,go generate,Generating grammar)

bootstrap:
	@$(call run_and_test,./scripts/bootstrap.sh,Transforming og -> go)
	@$(call run_and_test,go build,Recompiling og -> go)

build:
	@$(call run_and_test,go build,Compiling go src)

test:
	@$(call run_and_test,go test og/tests,Testing)

clean:
	@go clean
