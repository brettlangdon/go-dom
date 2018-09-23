all:
	go run generate/*.go
.PHONY: all

clean:
	rm *.go
	rm -rf console/ css/ document/ window/
.PHONY: clean
