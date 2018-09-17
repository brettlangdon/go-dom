GO_FILES = $(shell find generate/*.json | ack '/(?P<fname>.*?)\.' --output '$$+{fname}.go')

all:
	go generate
.PHONY: all

clean:
	rm $(GO_FILES)
.PHONY: clean
