include ../misc/Makefile

STACK-NAME := LAMBDA-Check-Golang-Version

deploy: $(DEPLOY_FILE)

PY_MODULES := bs4
ZIP_ARGS   := -x bs4/tests/\*

$(OBJ_DIR)/bs4/:
	pip install -t $(OBJ_DIR) beautifulsoup4

$(OBJ_DIR)/$(OBJ_ZIP): $(OBJ_DIR)/bs4/
