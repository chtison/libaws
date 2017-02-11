define USAGE
usage: make COMMAND
make template          # execute template on $(MAIN_TMPL) with $(DATA_TMPL)
make validate          # validate the template
make pytest            # run python tests
make zip               # create the zip to be uploaded to S3 in a lambda stack
make package           # upload the zip to S3 in a lambda stack
make deploy            # deploy stack named $(STACK-NAME) to your aws account
make delete            # delete stack named $(STACK-NAME) from your aws account
make clean             # remove local temporary files
endef

.PHONY: pytest zip package

PWD := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

include $(PWD)/Base.mk

PY_MODULES ?=
ZIP_ARGS   ?=

OBJ_PY_MODULES := $(addprefix $(OBJ_DIR)/, $(PY_MODULES))
LAMBDA_FN      ?= lambda.py
OBJ_ZIP        ?= lambda.zip
OBJ_PACKAGE    ?= cloudformation.package.yaml

DEPLOY_FILE    := $(OBJ_DIR)/$(OBJ_PACKAGE)

pytest:
	-python -m unittest discover -v -p \*_test.py
	rm -f -- *.pyc

$(OBJ_DIR)/$(OBJ_ZIP): $(OBJ_PY_MODULES)

$(OBJ_PY_MODULES):
	pip install -t $(OBJ_DIR) -r requirements.txt

zip: $(OBJ_DIR)/$(OBJ_ZIP)
$(OBJ_DIR)/$(OBJ_ZIP): $(OBJ_DIR) $(LAMBDA_FN)
	rm -f -- $@
	cp -fT -- $(LAMBDA_FN) $(OBJ_DIR)/$(LAMBDA_FN)
	chmod 644 $(OBJ_DIR)/$(LAMBDA_FN)
	find $(OBJ_DIR)/$(PY_MODULES) -type d -execdir chmod 755 {} +
	find $(OBJ_DIR)/$(PY_MODULES) -type f -execdir chmod 644 {} +
	cd $(OBJ_DIR)/ && zip -r $(OBJ_ZIP) $(LAMBDA_FN) $(PY_MODULES) $(ZIP_ARGS)

package: $(OBJ_DIR)/$(OBJ_PACKAGE)
$(OBJ_DIR)/$(OBJ_PACKAGE): $(OBJ_DIR)/$(OBJ_TEMPLATE) $(OBJ_DIR)/$(OBJ_ZIP)
	@[ "$(S3-BUCKET)" ] || (echo ERROR: S3-BUCKET variable is required && false)
	@[ "$(S3-PREFIX)" ] || (echo ERROR: S3-PREFIX variable is required && false)
	aws cloudformation package   \
		--template-file $<       \
		--s3-bucket $(S3-BUCKET) \
		--s3-prefix $(S3-PREFIX) \
		--output-template-file $@

deploy: $(DEPLOY_FILE)
