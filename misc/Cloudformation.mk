STACK_NAME            ?=
S3_BUCKET             ?=
S3_PREFIX             ?= cloudformation/$(STACK_NAME)
LAMBDA_DIR            ?= .
LAMBDA_FUNCTIONS      ?=
INCLUDED_TEMPLATES    ?=
TEMPLATE              ?= cloudformation.yaml
DATA_JSON             ?= # data.json # empty to remove tmpl preprocessing
OBJ_DIR               ?= obj
DEPLOY_ARGS           ?= --capabilities CAPABILITY_IAM
OBJ_TEMPLATE          ?= $(OBJ_DIR)/cloudformation.yaml
OBJ_LAMBDA_ZIP        := $(addprefix $(OBJ_DIR)/, $(addsuffix .zip, $(LAMBDA_FUNCTIONS)))
ifneq "$(LAMBDA_FUNCTIONS)$(INCLUDED_TEMPLATES)" ""
OBJ_PACKAGED_TEMPLATE ?= $(OBJ_DIR)/cloudformation.packaged.yaml
endif

.PHONY: default clean template validate deploy delete

default: deploy

$(OBJ_DIR):
	mkdir -p $@

clean:
	rm -rf -- $(OBJ_DIR)
ifdef LAMBDA_FUNCTIONS
	WD=$$(pwd) ; for FN in $(LAMBDA_FUNCTIONS) ; do \
		cd $$WD/$(LAMBDA_DIR)/$$FN && $(MAKE) --no-print-directory clean ; \
	done
endif

$(OBJ_TEMPLATE): $(OBJ_DIR) $(TEMPLATE) $(DATA_JSON)
ifdef DATA_JSON
	tmpl $(TEMPLATE) $(DATA_JSON) $@
else
	cp -fT -- $(TEMPLATE) $@
endif

template: $(OBJ_TEMPLATE)

validate: $(OBJ_TEMPLATE)
	aws cloudformation validate-template --template-body file://$<

ifdef OBJ_PACKAGED_TEMPLATE
$(OBJ_PACKAGED_TEMPLATE): $(OBJ_TEMPLATE) $(INCLUDED_TEMPLATES)
	@[ "$(S3_BUCKET)" ] || (echo ERROR: S3_BUCKET variable is required && false)
	@[ "$(S3_PREFIX)" ] || (echo ERROR: S3_PREFIX variable is required && false)
ifdef LAMBDA_FUNCTIONS
	WD=$$(pwd) ; for FN in $(LAMBDA_FUNCTIONS) ; do \
		cd $$WD/$(LAMBDA_DIR)/$$FN && \
		$(MAKE) --no-print-directory zip && \
		cp -fT -- obj/lambda.zip $$WD/$(OBJ_DIR)/$$FN.zip  ; \
	done
endif
	aws cloudformation package \
		--template-file $< \
		--s3-bucket $(S3_BUCKET) \
		--s3-prefix $(S3_PREFIX) \
		--output-template-file $@

.PHONY: package
package: $(OBJ_PACKAGED_TEMPLATE)

deploy: $(OBJ_PACKAGED_TEMPLATE)
else
deploy: $(OBJ_TEMPLATE)
endif
	@[ "$(STACK_NAME)" ] || (echo ERROR: STACK_NAME variable is required && false)
	aws cloudformation deploy \
		--template-file $< \
		--stack-name $(STACK_NAME) \
		$(DEPLOY_ARGS)

delete:
	@[ "$(STACK_NAME)" ] || (echo ERROR: STACK_NAME variable is required && false)
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
