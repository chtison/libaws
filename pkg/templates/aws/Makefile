.PHONY: run clean fclean

run: generated.go

generated.go: generate.go _tmp/CloudFormationResourceSpecification.json
	go run generate.go
	@echo "package aws" > generated.go
	@echo "// Templates ..." >> generated.go
	@echo "const Templates = \``cat _tmp/generated/*/*.tmpl.yaml`\`" >> generated.go
	@gofmt -w generated.go
	@echo "👍  generated.go successfully created"

_tmp/:
	@mkdir -p $@

_tmp/CloudFormationResourceSpecification.json: _tmp/CloudFormationResourceSpecification.gz
	gzip -dc $< > $@

_tmp/CloudFormationResourceSpecification.gz: _tmp/
	curl 'https://d1uauaxba7bl26.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json' -o $@

clean:
	rm -rf -- _tmp

fclean: clean
	rm -f -- generated.go
