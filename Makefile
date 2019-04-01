COMMIT_REVISION := $(shell git log --pretty=%h -1)
REVISION_FLAG := "-X main.revision=${COMMIT_REVISION}"
TARGET := validate-commit-msg

ifeq ($(OS),Windows_NT)
	GOOS := windows
	COPY := copy
else
	COPY := cp
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		GOOS := linux
	else ifeq ($(UNAME_S),Darwin)
		GOOS := darwin
	endif
endif

all: ${TARGET}
clean:
	rm -rf ${TARGET}

validate-commit-msg: main.go arguments.go git.go
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=amd64 go build -o $@ -ldflags ${REVISION_FLAG}

.PHONY: all clean
