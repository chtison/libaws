import json

def handler(event, context):
	print json.dumps(event)
	return {
		'statusCode': 200,
		'headers': {},
		'body': json.dumps(event),
	}
