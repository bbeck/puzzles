.DEFAULT_GOAL := help

mage := go run mage.go

## run the solution for the specified SITE/YEAR/DAY/PART
.PHONY: run
run:
	@$(mage) run

## run the solution for the specified SITE/YEAR/DAY/PART whenever a file changes
.PHONY: watch
watch:
	@$(mage) watch

## verify the solution output for the specified SITE/YEAR/DAY/PART
.PHONY: verify
verify:
	@$(mage) verify

## run the solution for the specified SITE/YEAR/DAY/PART whenever a file changes
.PHONY: next
next:
	@$(mage) next

## run all solutions for the specified SITE/YEAR
.PHONY: run-year
run-year:
	@$(mage) ListYear                                              | \
	while read year day part; do                                     \
	  printf "YEAR=%d DAY=%02d PART=%d " $${year} $${day} $${part};  \
	  YEAR=$${year} DAY=$${day} PART=$${part} $(mage) run;           \
	done

## verify the solution output of the specified SITE/YEAR
.PHONY: verify-year
verify-year:
	@$(mage) ListYear                                              | \
	while read year day part; do                                     \
	  YEAR=$${year} DAY=$${day} PART=$${part} $(mage) verify;        \
	done

## display this help message
.PHONY: help
help:
	@awk '                                                           \
	  BEGIN {                                                        \
	    printf "Usage:\n"                                            \
	  }                                                              \
	                                                                 \
	  /^##@/ {                                                       \
	    printf "\n\033[1m%s:\033[0m\n", substr($$0, 5)               \
	  }                                                              \
	                                                                 \
	  /^##([^@]|$$)/ && $$2 != "" {                                  \
	    $$1 = "";                                                    \
	    if (message == null) {                                       \
	      message = $$0;                                             \
	    } else {                                                     \
	      message = message "\n           " $$0;                     \
	    }                                                            \
	  }                                                              \
	                                                                 \
	  /^[a-zA-Z_-]+:/ && message != null {                           \
	    target = substr($$1, 0, length($$1)-1);                      \
	    printf "  \033[36m%-11s\033[0m %s\n", target, message;        \
	    message = null;                                              \
	  }                                                              \
	' $(MAKEFILE_LIST)
