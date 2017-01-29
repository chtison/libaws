PWD := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

.PHONY: run
run:
		docker run -it --rm -h docker -v $(PWD):/lambda -w /lambda chtison/workspace
