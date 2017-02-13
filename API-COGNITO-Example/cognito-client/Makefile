.PHONY: default install build run clean

# OPTIONS for `run` command
# OPTIONAL if sets in src/main.ts
USER_POOL_ID :=
CLIENT_ID    :=
# REQUIRED
EMAIL        :=
PASSWORD     :=

ENV := USER_POOL_ID=$(USER_POOL_ID) CLIENT_ID=$(CLIENT_ID)

default: build

node_modules/:
	npm install

typings/:
	typings install

install: node_modules/ typings/

dist/:
	mkdir $@

dist/main.js: node_modules/ typings/ dist/ tsconfig.json src/main.ts
	tsc

build: dist/main.js

run: dist/main.js
	@[ "$(EMAIL)" ]    || (echo ERROR: EMAIL    variable is required && false)
	@[ "$(PASSWORD)" ] || (echo ERROR: PASSWORD variable is required && false)
	$(ENV) node $< signUp $(EMAIL) $(PASSWORD)
	$(ENV) node $< confirmRegistration $(EMAIL) $$(echo -n 'Enter the confirmation code and CTRL-D: ' >&2 ; cat)
	$(ENV) node $< signIn $(EMAIL) $(PASSWORD)

clean:
	rm -rf -- dist/ typings/ node_modules/
