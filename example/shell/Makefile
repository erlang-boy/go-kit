GITCOMMIT:=$(shell git describe --dirty --always)

CMDDIR:=src/cmd
CLI:=$(notdir $(wildcard $(CMDDIR)/*))

all: $(CLI)
	@echo "shell ==> done"

$(CLI):%:$(CMDDIR)/%
	@echo "shell ==> compiling $^"
	@./gb build -ldflags "-w -s -X shell/hsmain.GitCommit=$(GITCOMMIT) -X 'shell/hsmain.BuildDate=`date -u`'" cmd/$@

clean:
	@rm -rf ./pkg/*
	@rm -rf ./bin/*
