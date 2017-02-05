import boto3, httplib, json, requests

cognitoIDP = boto3.client('cognito-idp')

def handler(event, context):
	if event['RequestType'] == 'Create':
		create(event)
	elif event['RequestType'] == 'Update':
		update(event)
	elif event['RequestType'] == 'Delete':
		delete(event)

def decorator(f):
	def g(event):
		data = {}
		data['StackId'] = event['StackId']
		data['RequestId'] = event['RequestId']
		data['LogicalResourceId'] = event['LogicalResourceId']
		data['PhysicalResourceId'] = event.get('PhysicalResourceId', data['StackId'].split('/')[1]+'-'+data['LogicalResourceId']+'-'+data['RequestId'])
		f(event, data)
		data = json.dumps(data)
		requests.put(event['ResponseURL'], data=data)
	return g

@decorator
def create(event, data):
	data['Status'] = 'SUCCESS'

@decorator
def update(event, data):
	data['Status'] = 'SUCCESS'

@decorator
def delete(event, data):
	data['Status'] = 'SUCCESS'
