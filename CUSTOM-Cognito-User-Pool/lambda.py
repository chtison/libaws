import boto3, httplib, json, requests

cognitoIDP = boto3.client('cognito-idp')

def handler(event, context):
	print 'REQUEST:'
	print json.dumps(event)
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
		if 'PhysicalResourceId' in event:
			data['PhysicalResourceId'] = event['PhysicalResourceId']
		data['Status'] = 'SUCCESS'
		f(event, data)
		data = json.dumps(data, separators=(',', ':'))
		print 'RESPONSE:'
		print data
		requests.put(event['ResponseURL'], data=data)
	return g

@decorator
def create(event, data):
	if 'ResourceProperties' not in event:
		event['ResourceProperties'] = {}
	if 'ServiceToken' in event['ResourceProperties']:
		del event['ResourceProperties']['ServiceToken']
	if 'PoolName' not in event['ResourceProperties']:
		event['ResourceProperties']['PoolName'] = data['StackId'].split('/')[1]+'-'+data['LogicalResourceId']
	try:
		response = cognitoIDP.create_user_pool(**event['ResourceProperties'])
		data['PhysicalResourceId'] = response['UserPool']['Id']
	except Exception as e:
		data['Status'] = 'FAILED'
		data['Reason'] = str(e)
		return
	return

@decorator
def update(event, data):
	return

@decorator
def delete(event, data):
	try:
		cognitoIDP.delete_user_pool(
			UserPoolId=data['PhysicalResourceId'],
		)
	except Exception as e:
		print str(e)
	return
