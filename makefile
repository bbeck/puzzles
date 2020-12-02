current_dir := $(shell ls -d1 cmd/*/[0-9][0-9]-[12] | tail -n1)
YEAR := $(word 2, $(subst /, ,$(subst -, ,$(current_dir))))
DAY := $(word 3, $(subst /, ,$(subst -, ,$(current_dir))))
PART := $(word 4, $(subst /, ,$(subst -, ,$(current_dir))))

.PHONY: help
help:
	@echo 'Usage: make [TARGET] [EXTRA_ARGUMENTS]'
	@echo ''
	@echo 'Targets:'
	@echo '  run    run the solution for the specified YEAR= DAY= and PART= arguments'
	@echo '  watch  watch for changes and rerun the solution for the specified YEAR= DAY= and PART= arguments'
	@echo '  next   create the solution template for the next day'
	@echo ''

.PHONY: watch
watch:
	@fswatch -d -e'.*~' -o aoc/ cmd/$(YEAR)/$(DAY)-$(PART) | \
     xargs -n1 -I{} $(MAKE) clear run

.PHONY: clear
clear:
	@clear

.PHONY: run
run:
	@echo -n '$(shell date +'%I:%M:%S %p'):  '
	go run cmd/$(YEAR)/$(DAY)-$(PART)/main.go

.PHONY: next
next:
    ifeq ($(PART),1)
		$(eval year=$(YEAR))
		$(eval day=$(DAY))
		$(eval part=2)
		$(eval copy=1)
    else ifneq ($(DAY),25)
		$(eval year=$(YEAR))
		$(eval day=$(shell printf "%02d" $(shell expr $(DAY) + 1)))
		$(eval part=1)
    else
		$(eval year=$(shell printf "%02d" $(shell expr $(YEAR) + 1)))
		$(eval day=1)
		$(eval part=1)
    endif

	@mkdir cmd/$(year)/$(day)-$(part)
	@if [ $(copy) ]; then                                                                                       \
	   cp -P cmd/$(YEAR)/$(DAY)-$(PART)/*.go cmd/$(year)/$(day)-$(part)/;                                       \
	else                                                                                                        \
	   touch cmd/$(year)/$(day)-$(part)/input.txt;                                                              \
   	   echo 'package main'                                                > cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo                                                              >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo 'import ('                                                   >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo '  "fmt"'                                                    >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo                                                              >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo '  "github.com/bbeck/advent-of-code/aoc"'                    >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo ')'                                                          >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo                                                              >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo 'func main() {'                                              >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo '  for _, line := range aoc.InputToLines($(year), $(day)) {' >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo '    fmt.Println(line)'                                      >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo '  }'                                                        >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo '}'                                                          >> cmd/$(year)/$(day)-$(part)/main.go; \
   	   echo                                                              >> cmd/$(year)/$(day)-$(part)/main.go; \
	fi;
