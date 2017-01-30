STACK-NAME := LAMBDA-Check-Golang-Version
S3-BUCKET  :=
S3-PREFIX  := lambda/$(STACK-NAME)

.PHONY: step1 step2 pip zip deploy clean

all: deploy

objs/:
	mkdir -p objs/

step1: objs/aws-sam.step1.yaml
objs/aws-sam.step1.yaml: objs/ aws-sam.tmpl.yaml data.json
	tmpl aws-sam.tmpl.yaml data.json $@

validate: objs/aws-sam.step1.yaml
	aws cloudformation validate-template --template-body file://$<

pip: objs/bs4/
objs/bs4/:
	pip install -t objs/ beautifulsoup4
	rm -rf -- objs/beautifulsoup4-*.dist-info

zip: objs/lambda.zip
objs/lambda.zip: function.py objs/bs4/
	rm -f -- $@
	cp -fT -- function.py objs/function.py
	chmod 644 objs/function.py
	find objs/bs4 -type d -execdir chmod 755 {} +
	find objs/bs4 -type f -execdir chmod 644 {} +
	cd objs/ && zip -r lambda.zip function.py bs4/ -x bs4/tests/\*

step2: objs/aws-sam.step2.yaml
objs/aws-sam.step2.yaml: objs/aws-sam.step1.yaml objs/lambda.zip
	@[ "$(S3-BUCKET)" ] || (echo ERROR: S3-BUCKET variable is required && false)
	@[ "$(S3-PREFIX)" ] || (echo ERROR: S3-PREFIX variable is required && false)
	aws cloudformation package   \
		--template-file $<       \
		--s3-bucket $(S3-BUCKET) \
		--s3-prefix $(S3-PREFIX) \
		--output-template-file $@

deploy: objs/aws-sam.step2.yaml
	@[ "$(STACK-NAME)" ] || (echo ERROR: STACK-NAME variable is required && false)
	aws cloudformation deploy      \
		--template-file $<         \
		--stack-name $(STACK-NAME) \
		--capabilities CAPABILITY_IAM

clean:
	rm -rf -- objs/
