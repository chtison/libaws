STACK_NAME       := API-COGNITO
TEMPLATE         := cloudformation.yaml ../TEMPLATE-LAMBDA/lambda.yaml ../TEMPLATE-API/api.yaml
DATA_YAML        := data.yaml
LAMBDA_DIR       := lambda
LAMBDA_FUNCTIONS := echo user.info/get user.info/post user.info/put
include ../misc/Cloudformation.mk
