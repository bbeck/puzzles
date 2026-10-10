.DEFAULT_GOAL := help
mage := go tool mage

## Solving

.PHONY: run
run:  ## run SITE/YEAR/DAY/PART
	@$(mage) run

.PHONY: verify
verify:  ## run SITE/YEAR/DAY/PART and verify the output
	@$(mage) verify

.PHONY: watch
watch:  ## run SITE/YEAR/DAY/PART whenever files change
	@$(mage) watch

.PHONY: next
next:  ## setup the source for the next puzzle
	@$(mage) next

.PHONY: wait-until-start-time
wait-until-start-time:  ## wait until the start time for SITE
	@$(mage) WaitUntilStartTime

## Verifying

.PHONY: run-day
run-day:  ## run the entire SITE/YEAR/DAY
	@$(mage) ListDay                                               | \
	while read year day part; do                                     \
	  printf "YEAR=%d DAY=%02d PART=%d " $${year} $${day} $${part};  \
	  YEAR=$${year} DAY=$${day} PART=$${part} $(mage) run;           \
	done

.PHONY: run-year
run-year:  ## run the entire SITE/YEAR
	@$(mage) ListYear                                              | \
	while read year day part; do                                     \
	  printf "YEAR=%d DAY=%02d PART=%d " $${year} $${day} $${part};  \
	  YEAR=$${year} DAY=$${day} PART=$${part} $(mage) run;           \
	done

.PHONY: verify-day
verify-day:  ## run the entire SITE/YEAR/DAY and verify the output
	@$(mage) ListDay                                               | \
	while read year day part; do                                     \
	  YEAR=$${year} DAY=$${day} PART=$${part} $(mage) verify;        \
	done

.PHONY: verify-year
verify-year:  ## run the entire SITE/YEAR and verify the output
	@$(mage) ListYear                                              | \
	while read year day part; do                                     \
	  YEAR=$${year} DAY=$${day} PART=$${part} $(mage) verify;        \
	done

## Library

.PHONY: test
test:  ## run the unit tests
	@go test "./lib/..."

.PHONY: help
help:
	@awk '                                                           \
	  BEGIN {                                                        \
	    reset = "\033[0m";                                           \
	    header_color = "\033[1;35m";                                 \
	    target_color = "\033[1;1m";                                  \
	    comment_color = "\033[90m";                                  \
	    indent = "  ";                                               \
	  }                                                              \
	                                                                 \
	  NR == FNR {                                                    \
	    if (/^[a-zA-Z_-]+:.*## /) {                                  \
	      t = $$0; sub(/:.*/, "", t);                                \
	      if (length(t) > width) width = length(t);                  \
	    }                                                            \
	    next;                                                        \
	  }                                                              \
	                                                                 \
	  /^## / {                                                       \
	    sub(/^## /, "");                                             \
	    printf "\n%s%s%s\n", header_color, $$0, reset;               \
	    next;                                                        \
	  }                                                              \
	                                                                 \
	  /^[a-zA-Z_-]+:.*## / {                                         \
	    t = $$0; sub(/:.*/, "", t);                                  \
	    c = $$0; sub(/^[^#]*## */, "", c);                           \
	    printf "%s%s%-*s%s %s- %s%s\n",                              \
	      indent, target_color, width, t, reset,                     \
	      comment_color, c, reset;                                   \
	  }                                                              \
	' $(MAKEFILE_LIST) $(MAKEFILE_LIST)
