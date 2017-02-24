import boto3, json, os

DYNAMODB_TABLE_NAME = os.environ['DYNAMODB_TABLE_NAME_Users']

dynamodb = boto3.client('dynamodb')

def handler(event, context):
	claims = event['requestContext']['authorizer']['claims']
	response = dynamodb.get_item(
		TableName=DYNAMODB_TABLE_NAME,
		Key={
			'Id': {
				'S': claims['sub'],
			},
			'Email': {
				'S': claims['email'],
			},
		},
	)
	item = json.dumps(response.get('Item', {}), default=str)
	print item
	return {
		'statusCode': 200,
		'body': item,
	}
