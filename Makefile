.PHONY: run

PWD := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

run:
	docker run -it --rm -h docker -v $(PWD):/libaws -w /libaws chtison/workspace
