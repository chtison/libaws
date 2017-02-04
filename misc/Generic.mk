define USAGE
usage: make COMMAND
make template          # execute template on $(MAIN_TMPL) with $(DATA_TMPL)
make validate          # validate the template
make deploy            # deploy stack named $(STACK-NAME) to your aws account
make delete            # delete stack named $(STACK-NAME) from your aws account
make clean             # remove local temporary files
endef

include ../misc/Base.mk

DEPLOY_FILE := $(OBJ_DIR)/$(OBJ_STEP1)

deploy: $(DEPLOY_FILE)