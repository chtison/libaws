STACK-NAME  :=
S3-BUCKET   :=
S3-PREFIX    = cloudformation/$(STACK-NAME)

MAIN_TMPL  := cloudformation.tmpl.yaml
DATA_TMPL  := data.json
LAMBDA_FN  := lambda.py

OBJ_DIR    := objs
OBJ_STEP1  := cloudformation.template.yaml
OBJ_ZIP    := lambda.zip
OBJ_STEP2  := cloudformation.package.yaml

ZIP_ARGS :=

DEPLOY_FILE := $(OBJ_DIR)/$(OBJ_STEP2)
DEPLOY_ARGS := --capabilities CAPABILITY_IAM


.PHONY: template validate zip package deploy clean delete

help: ; $(info $(USAGE)) @true
define USAGE
usage: make COMMAND
make template          # execute template on $(MAIN_TMPL) with $(DATA_TMPL)
make validate          # validate the template
make zip               # create the zip to be uploaded to S3 in a lambda stack
make package           # upload the zip to S3 in a lambda stack
make deploy            # deploy stack named $(STACK-NAME) to your aws account
make delete            # delete stack named $(STACK-NAME) from your aws account
make clean             # remove local temporary files
endef

$(OBJ_DIR):
	mkdir -p $@

template: $(OBJ_DIR)/$(OBJ_STEP1)
$(OBJ_DIR)/$(OBJ_STEP1): $(OBJ_DIR) $(MAIN_TMPL) $(DATA_TMPL)
	tmpl $(MAIN_TMPL) $(DATA_TMPL) $@

validate: $(OBJ_DIR)/$(OBJ_STEP1)
	aws cloudformation validate-template --template-body file://$<

zip: $(OBJ_DIR)/$(OBJ_ZIP)
$(OBJ_DIR)/$(OBJ_ZIP): $(OBJ_DIR) $(LAMBDA_FN)
	rm -f -- $@
	cp -fT -- $(LAMBDA_FN) $(OBJ_DIR)/$(LAMBDA_FN)
	chmod 644 $(OBJ_DIR)/$(LAMBDA_FN)
	find $(OBJ_DIR)/$(PY_MODULES) -type d -execdir chmod 755 {} +
	find $(OBJ_DIR)/$(PY_MODULES) -type f -execdir chmod 644 {} +
	cd $(OBJ_DIR)/ && zip -r $(OBJ_ZIP) $(LAMBDA_FN) $(PY_MODULES) $(ZIP_ARGS)

package: $(OBJ_DIR)/$(OBJ_STEP2)
$(OBJ_DIR)/$(OBJ_STEP2): $(OBJ_DIR)/$(OBJ_STEP1) $(OBJ_DIR)/$(OBJ_ZIP)
	@[ "$(S3-BUCKET)" ] || (echo ERROR: S3-BUCKET variable is required && false)
	@[ "$(S3-PREFIX)" ] || (echo ERROR: S3-PREFIX variable is required && false)
	aws cloudformation package   \
		--template-file $<       \
		--s3-bucket $(S3-BUCKET) \
		--s3-prefix $(S3-PREFIX) \
		--output-template-file $@

deploy:
	@[ "$(STACK-NAME)" ] || (echo ERROR: STACK-NAME variable is required && false)
	aws cloudformation deploy          \
		--template-file $(DEPLOY_FILE) \
		--stack-name $(STACK-NAME)     \
		$(DEPLOY_ARGS)

clean:
	rm -rf -- $(OBJ_DIR)

delete:
	@[ "$(STACK-NAME)" ] || (echo ERROR: STACK-NAME variable is required && false)
	aws cloudformation delete-stack --stack-name $(STACK-NAME)
