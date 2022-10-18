.DEFAULT_GOAL := help

# What editor to open a new file in.
EDITOR := goland

# Retrieve our session cookie so that we can automatically download inputs.
SESSION := $(shell cat .session)

# Infer which solution we're working with based on which files exist locally.
# This inference will only execute in the situation where a particular solution
# part isn't specified as part of the make command.
last_dir := $(shell ls -d1 cmd/*/[0-9][0-9]-[12] | tail -n1 | awk -F'[/-]' '{ print $$2, $$3, $$4 }')
YEAR := $(word 1, $(last_dir))
DAY  := $(word 2, $(last_dir))
PART := $(word 3, $(last_dir))

# Canonicalize the YEAR/DAY/PART parameters so they're suitable to use in 
# directory or file names.  This will ensure that each parameter is padded to 
# the correct length with leading zeroes.
override YEAR := $(shell printf '%04d' $$((10\#$(YEAR))))
override DAY  := $(shell printf '%02d' $$((10\#$(DAY))))
override PART := $(shell printf '%01d' $$((10\#$(PART))))

# When downloading we need to know the day number without any leading zeroes.
DAY_NO_ZERO := $(shell printf '%d' $$((10\#$(DAY))))

# Determine the parameters for the next solution.  This is needed by the next 
# target.  We do this here because it requires conditionals and they are tricky
# to do within a target.
ifeq ($(DAY)-$(PART),25-1)
	NEXT_YEAR := $(shell printf "%04d" $$(($(YEAR) + 1)))
	NEXT_DAY  := 01
	NEXT_PART := 1
else ifeq ($(PART),1)
	NEXT_YEAR := $(YEAR)
	NEXT_DAY  := $(DAY)
	NEXT_PART := 2
else
	NEXT_YEAR := $(YEAR)
	NEXT_DAY  := $(shell printf "%02d" $$(($(DAY_NO_ZERO) + 1)))
	NEXT_PART := 1
endif

# When generating the source for the next day we need to know the next day 
# number without any leading zeroes.  If we don't then the day number will
# be interpreted as octal which is a problem for days 8 and 9.
NEXT_DAY_NO_ZERO := $(shell printf '%d' $$((10\#$(NEXT_DAY))))

## run the solution for the specified YEAR/DAY/PART
.PHONY: run
run: cmd/$(YEAR)/$(DAY)-1/input.txt
	@go run cmd/$(YEAR)/$(DAY)-$(PART)/*.go

# Download the input file for a particular day
cmd/$(YEAR)/$(DAY)-1/input.txt:
	@curl                                                          \
	  --fail                                                       \
	  --silent                                                     \
	  --cookie "session=$(SESSION)"                                \
	  --output cmd/$(YEAR)/$(DAY)-1/input.txt                      \
	  https://adventofcode.com/$(YEAR)/day/$(DAY_NO_ZERO)/input || \
	(echo "input.txt file not available" >&2; false)

## watch for changes and rerun the solution for the specified YEAR/DAY/PART
.PHONY: watch
watch:
	@find                                                          \
	    aoc                                                        \
	    cmd/$(YEAR)/$(DAY)-$(PART)                                 \
	    cmd/$(YEAR)/$(DAY)-1/input.txt                             \
	  -type f                                                    | \
	 entr -c make -s run YEAR=$(YEAR) DAY=$(DAY) PART=$(PART)

## create the solution template for the next DAY/PART
.PHONY: next
next:
	@mkdir -p cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)
	@if [[ $(NEXT_PART) -eq 2 ]]; then                                                                                                          \
	  cp -P cmd/$(YEAR)/$(DAY)-$(PART)/*.go cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/;                                                         \
	else                                                                                                                                        \
	  echo 'package main'                                                                  > cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo                                                                                >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo 'import ('                                                                     >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo '  "fmt"'                                                                      >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo                                                                                >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo '  "github.com/bbeck/advent-of-code/aoc"'                                      >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo ')'                                                                            >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo                                                                                >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo 'func main() {'                                                                >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo '  for _, line := range aoc.InputToLines($(NEXT_YEAR), $(NEXT_DAY_NO_ZERO)) {' >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo '    fmt.Println(line)'                                                        >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo '  }'                                                                          >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo '}'                                                                            >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	  echo                                                                                >> cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go; \
	fi
	@$(EDITOR) cmd/$(NEXT_YEAR)/$(NEXT_DAY)-$(NEXT_PART)/main.go >/dev/null 2>&1

## display this help message
.PHONY: help
help:
	@awk '                                                      \
	  BEGIN {                                                   \
	    printf "Usage:\n"                                       \
	  }                                                         \
	                                                            \
	  /^##@/ {                                                  \
	    printf "\n\033[1m%s:\033[0m\n", substr($$0, 5)          \
	  }                                                         \
	                                                            \
	  /^##([^@]|$$)/ && $$2 != "" {                             \
	    $$1 = "";                                               \
	    if (message == null) {                                  \
	      message = $$0;                                        \
	    } else {                                                \
	      message = message "\n           " $$0;                \
	    }                                                       \
	  }                                                         \
	                                                            \
	  /^[a-zA-Z_-]+:/ && message != null {                      \
	    target = substr($$1, 0, length($$1)-1);                 \
	    printf "  \033[36m%-8s\033[0m %s\n", target, message;   \
	    message = null;                                         \
	  }                                                         \
	' $(MAKEFILE_LIST)
