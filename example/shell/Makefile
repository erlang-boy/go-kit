GITCOMMIT:=$(shell git describe --dirty --always)

CMDDIR:=src/cmd
CLI:=$(notdir $(wildcard $(CMDDIR)/*))

CHECKS:=check godeps
VERBOSE:=-v

all: $(CLI)
	@echo "hive_shell ==> done"

$(CLI):%:$(CMDDIR)/%
	@echo "hive_shell ==> compiling $^"
	@./gb build -ldflags "-w -s -X hive_shell/hsmain.GitCommit=$(GITCOMMIT) -X 'hive_shell/hsmain.BuildDate=`date -u`'" cmd/$@

clean:
	@rm -rf ./pkg/*
	@rm -rf ./bin/*
