.DEFAULT_GOAL := install

.PHONY: distclean generate test install wfe

WFE_OS_NAME := $(shell uname -o 2>/dev/null || uname -s)

ifeq "$(WFE_OS_NAME)" "Cygwin"
	WFEOS := windows
	CMD := cmd /C
else
	ifeq "$(WFE_OS_NAME)" "Msys"
		WFEOS := windows
		CMD := cmd //C
	else
		ifneq (,$(findstring MINGW, $(WFE_OS_NAME)))
			WFEOS := windows
			CMD := cmd //C
		endif
	endif
endif

ifndef GOBIN
	ifeq "$(WFEOS)" "windows"
		GOBIN := $(shell $(CMD) "echo %GOPATH%| cut -d';' -f1")
		GOBIN := $(subst \,/,$(GOBIN))/bin
	else
        	GOBIN := $(shell echo $$GOPATH | cut -d':' -f1 )/bin
	endif
endif

distclean:
	go clean ./...
	rm -rf ${GOBIN}/wfe

generate:
	go list ./... | xargs go generate

test: 
	go list ./... | xargs go test

install: wfe

wfe: ${GOBIN}/wfe

${GOBIN}/wfe: $(shell /usr/bin/find . -type f -and -name '*.go')
	go install ./cmd/wfe
