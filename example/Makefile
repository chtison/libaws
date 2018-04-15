.PHONY: template lambdas test package deploy clean cf-validate cf-resources cf-events cf-describe cf-delete

STACK_NAME := LibawsExample
S3_BUCKET  :=
S3_PREFIX  := lambdas

-include secrets.mk

TEMPLATES := cloudformation.tmpl.yaml
DATAS     := cloudformation.data.yaml

LAMBDAS      := $(shell find lambdas -type d)
LAMBDAS_OUT  := $(addprefix _out/, $(LAMBDAS))

help:
	$(info $(USAGE)) @true
define USAGE
usage: make COMMAND
make template          # run libaws template engine with [$(TEMPLATES)] and [$(DATAS)]
make lambdas           # build lambdas functions: [$(LAMBDAS)]
make test              # run tests for lambdas functions
make package           # run aws cloudformation package on cloudformation template
make deploy            # deploy the packaged cloudformation template to a stack named "$(STACK_NAME)"
make clean             # remove the _out directory
make cf-validate       # Validate cloudformation template
make cf-resources      # List aws stack named "$(STACK_NAME)" resources
make cf-describe       # Describe aws stack named "$(STACK_NAME)"
make cf-delete         # Delete aws stack named "$(STACK_NAME)"
endef

_out/cloudformation.yaml: _out $(DATAS) $(TEMPLATES)
	libaws run $(foreach DATA, $(DATAS),-d $(DATA)) $(TEMPLATES) > $@

template: _out/cloudformation.yaml

_out/lambdas/%: lambdas/%/*
	GOOS=linux GOARCH=amd64 go build -ldflags='-s -w ' -o $@ $^ && chmod 511 $@
	rm -f $@.zip
	cd _out/lambdas && zip $(notdir $@).zip $(notdir $@)
	chmod 400 $@.zip

lambdas: $(LAMBDAS_OUT)

test:
	for l in $(LAMBDAS); do go test ./$$l; done

_out/cloudformation.pack.yaml: _out/cloudformation.yaml $(LAMBDAS_OUT)
	aws cloudformation package \
		--template-file  $< \
		--s3-bucket $(S3_BUCKET) \
		--s3-prefix "$(S3_PREFIX)" \
		--output-template-file $@

package: _out/cloudformation.pack.yaml

deploy: _out/cloudformation.pack.yaml
	aws cloudformation deploy \
		--template-file $< \
		--capabilities CAPABILITY_IAM \
		--stack-name $(STACK_NAME)

_out _out/lambdas:
	@mkdir -p $@

clean:
	rm -rf -- _out

cf-validate: _out/cloudformation.yaml
	aws cloudformation validate-template --template-body file://$<

cf-resources:
	aws cloudformation list-stack-resources --stack-name $(STACK_NAME)

cf-events:
	aws cloudformation describe-stack-events --stack-name $(STACK_NAME) | grep ResourceStatusReason

cf-delete:
	aws cloudformation delete-stack --stack-name $(STACK_NAME)
	aws cloudformation wait stack-delete-complete --stack-name $(STACK_NAME)

cf-describe:
	aws cloudformation describe-stacks --stack-name $(STACK_NAME)
