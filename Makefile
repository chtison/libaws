.PHONY: install dev

install:
	go install github.com/chtison/libaws

dev:
	$(MAKE) -C pkg/templates/ run
	go install -v github.com/chtison/libaws
