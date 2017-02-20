LAMBDA_FILE    ?= lambda.py
LAMBDA_MODULES ?=
ZIP_NAME       ?= lambda.zip
ZIP_ARGS       ?=
OBJ_DIR        ?= obj

OBJ_LAMBDA_FILE    := $(OBJ_DIR)/$(LAMBDA_FILE)
OBJ_LAMBDA_MODULES := $(addprefix $(OBJ_DIR)/, $(LAMBDA_MODULES))
OBJ_ZIP            := $(OBJ_DIR)/$(ZIP_NAME)

.PHONY: default clean test zip

default: zip

$(OBJ_DIR):
	mkdir -p $@

clean:
	rm -rf -- $(OBJ_DIR)

test:
	if [ -f requirements.txt ] ; then pip install -r requirements.txt ; fi
	if [ -f requirements_test.txt ] ; then pip install -r requirements_test.txt ; fi
	-python -m unittest discover -v -p \*_test.py
	rm -f -- *.pyc

ifdef LAMBDA_MODULES
$(OBJ_LAMBDA_MODULES): requirements.txt
	pip install -t $(OBJ_DIR) -r requirements.txt

.PHONY: pip
pip: $(OBJ_LAMBDA_MODULES)
endif

$(OBJ_LAMBDA_FILE): $(LAMBDA_FILE)
	cp -fT -- $< $@

$(OBJ_ZIP): $(OBJ_DIR) $(OBJ_LAMBDA_FILE) $(OBJ_LAMBDA_MODULES)
	rm -f -- $@
	chmod 644 $(OBJ_LAMBDA_FILE)
	find $(OBJ_LAMBDA_MODULES) -type d -execdir chmod 755 {} +
	find $(OBJ_LAMBDA_MODULES) -type f -execdir chmod 644 {} +
	cd $(OBJ_DIR) && zip -r $(ZIP_NAME) $(LAMBDA_FILE) $(LAMBDA_MODULES) $(ZIP_ARGS)

zip: $(OBJ_ZIP)
