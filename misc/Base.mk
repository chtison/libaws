STACK-NAME ?=
S3-BUCKET  ?=
S3-PREFIX  ?= cloudformation/$(STACK-NAME)

MAIN_TMPL  ?= cloudformation.tmpl.yaml
DATA_TMPL  ?= data.json

OBJ_DIR      ?= objs
OBJ_TEMPLATE ?= cloudformation.template.yaml

DEPLOY_FILE ?=
DEPLOY_ARGS ?= --capabilities CAPABILITY_IAM

.PHONY: help template validate deploy clean delete

help: ; $(info $(USAGE)) @true

$(OBJ_DIR):
	mkdir -p $@

template: $(OBJ_DIR)/$(OBJ_TEMPLATE)
$(OBJ_DIR)/$(OBJ_TEMPLATE): $(OBJ_DIR) $(MAIN_TMPL) $(DATA_TMPL)
	tmpl $(MAIN_TMPL) $(DATA_TMPL) $@

validate: $(OBJ_DIR)/$(OBJ_TEMPLATE)
	aws cloudformation validate-template --template-body file://$<

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
