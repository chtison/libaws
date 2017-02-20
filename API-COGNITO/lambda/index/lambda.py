import json

def handler(event, context):
	body = json.dumps(event)
	print body
	return {
		'statusCode': 200,
		'headers': {},
		'body': body,
	}
