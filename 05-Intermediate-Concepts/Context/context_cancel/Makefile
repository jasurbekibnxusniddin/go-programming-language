.DEFAULT_GOAL := build

.PHONY: build run-ok run-cancel

build:
	go build

run-ok: build
	./context_cancel false
	rm context_cancel

run-cancel: build
	./context_cancel true
	rm context_cancel

