STACK-NAME := SNS-TOPIC-Main

.PHONY: all step1 deploy clean

all: deploy

step1: aws-cf.step1.yaml
aws-cf.step1.yaml: aws-cf.tmpl.yaml data.json
	tmpl $< $(word 2, $^) $@

deploy: aws-cf.step1.yaml
	@[ "$(STACK-NAME)" ] || (echo ERROR: STACK-NAME variable is required && false)
	aws cloudformation deploy      \
		--template-file $<         \
		--stack-name $(STACK-NAME) \
		--capabilities CAPABILITY_IAM

clean:
	rm -f -- aws-cf.step1.yaml
