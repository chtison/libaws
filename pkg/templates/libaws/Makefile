.PHONY: default template clean

STACK_NAME := StackName

TMPL_DEPS := cloudformation.tmpl.yaml cloudformation.data.yaml

OUT_DIR  := out
OUT_FILE := cloudformation.yaml
OUT_PATH := $(OUT_DIR)/$(OUT_FILE)

default: template

$(OUT_DIR):
	@mkdir -p $@

template: $(OUT_PATH)
$(OUT_PATH): $(OUT_DIR) $(TMPL_DEPS)
	libaws template run > $@

clean:
	rm -rf -- $(OUT_DIR)

deploy: $(OUT_PATH)
	aws cloudformation deploy \
		--template-file $(OUT_PATH) \
		--stack-name $(STACK_NAME)
