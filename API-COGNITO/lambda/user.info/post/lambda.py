import boto3, json, os

DYNAMODB_TABLE_NAME = os.environ['DYNAMODB_TABLE_NAME_Users']

dynamodb = boto3.client('dynamodb')

def handler(event, context):
	claims = event['requestContext']['authorizer']['claims']
	response = dynamodb.update_item(
		TableName=DYNAMODB_TABLE_NAME,
		Key={
			'Id': {
				'S': claims['sub'],
			},
			'Email': {
				'S': claims['email'],
			},
		},
		ReturnValues='UPDATED_NEW',
		ReturnConsumedCapacity='INDEXES',
		ReturnItemCollectionMetrics='SIZE',
		UpdateExpression='ADD i :one',
		ExpressionAttributeValues={
			':one': {
				'N': '1',
			}
		},

	)
	del response['ResponseMetadata']
	body = json.dumps(response, default=str)
	print body
	return {
		'statusCode': 200,
		'body': body,
	}
