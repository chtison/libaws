import boto3, json

dynamodb = boto3.client('dynamodb')

def handler(event, context):
	print json.dumps(event)
	return
