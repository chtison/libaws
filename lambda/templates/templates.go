package templates

// Python36 ...
const Python36 = `import json

def handler(event, context):
	body = {
		"message": "Hello World !",
		"event": event,
		"context": context,
	}
	return {
		"statusCode": 200,
		"body": json.dumps(body),
	}`
