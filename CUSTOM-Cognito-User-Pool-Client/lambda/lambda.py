import boto3, httplib, json, requests

cognitoIDP = boto3.client('cognito-idp')

def handler(event, context):
	def log(s):
		print s
		print json.dumps(event)
	if event['RequestType'] == 'Create':
		log('CREATE')
		create(event)
	elif event['RequestType'] == 'Update':
		log('UPDATE')
		update(event)
	elif event['RequestType'] == 'Delete':
		log('DELETE')
		delete(event)

def decorator(f):
	def g(event):
		data = {}
		data['StackId'] = event['StackId']
		data['RequestId'] = event['RequestId']
		data['LogicalResourceId'] = event['LogicalResourceId']
		if 'PhysicalResourceId' in event:
			data['PhysicalResourceId'] = event['PhysicalResourceId']
		else:
			data['PhysicalResourceId'] = 'X'
		data['Status'] = 'SUCCESS'
		f(event, data)
		data = json.dumps(data, separators=(',', ':'), default=str)
		print 'RESPONSE:'
		print data
		requests.put(event['ResponseURL'], data=data)
	return g

@decorator
def create(event, data):
	if 'ResourceProperties' not in event:
		event['ResourceProperties'] = {}
	if 'Properties' not in event['ResourceProperties']:
		event['ResourceProperties']['Properties'] = {}
	properties = event['ResourceProperties']['Properties']
	if not checkType(properties, dict, data, ''): return
	if not normalizeParameter(['GenerateSecret'], Bool(), properties, data): return
	if not normalizeParameter(['RefreshTokenValidity'], int, properties, data): return
	try:
		response = cognitoIDP.create_user_pool_client(**properties)
		data['PhysicalResourceId'] = response['UserPoolClient']['ClientId']
		if 'ClientSecret' in response['UserPoolClient']:
			data['Data'] = {
				'ClientSecret': response['UserPoolClient']['ClientSecret'],
			}
		print 'CREATE_USER_POOL_CLIENT:'
		print json.dumps(response['UserPoolClient'], default=str)
	except Exception as e:
		data['Status'] = 'FAILED'
		print str(e)
		data['Reason'] = str(e)
		return
	return

class Bool:
	__name__ = 'bool'
	def __call__(self, s):
		return s.lower() == 'true'

def checkType(var, t, data, varPath):
	if type(var) is not t:
		data['Status'] = 'FAILED'
		data['Reason'] = 'Properties{} is not of type {}'.format(varPath, t.__name__)
		return False
	return True

def normalizeParameter(keys, fn, properties, data, keyPath=None):
	key = keys.pop(0)
	if key not in properties:
		return True
	keyPath = keyPath + '.' + str(key) if keyPath else '.' + str(key)
	if len(keys) == 0:
		try:
			properties[key] = fn(properties[key])
		except Exception as e:
			data['Status'] = 'FAILED'
			data['Reason'] = 'Properties{} = "{}" and cannot be converted to type {}'.format(keyPath, properties[key], fn.__name__)
			return False
		return True
	if not checkType(properties[key], dict, data, keyPath): return False
	return normalizeParameter(keys, fn, properties[key], data, keyPath)

@decorator
def update(event, data):
	data['Status'] = 'FAILED'
	data['Reason'] = 'Update not implemented.'
	return

@decorator
def delete(event, data):
	if event['PhysicalResourceId'] == 'X':
		return
	try:
		cognitoIDP.delete_user_pool_client(
			UserPoolId=event['ResourceProperties']['Properties']['UserPoolId'],
			ClientId=event['PhysicalResourceId'],
		)
	except Exception as e:
		print str(e)
	return
