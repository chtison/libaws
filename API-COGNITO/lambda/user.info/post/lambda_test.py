import os, sys, unittest, yaml

with open('../data.yaml') as f:
	data = yaml.load(f)
os.environ['DYNAMODB_TABLE_NAME_Users'] = data['DYNAMODB_TABLE_NAME_Users']
execfile('lambda.py')

class TestLambdaFunction(unittest.TestCase):
	def test_handler(self):
		event = {
			'requestContext': {
				'authorizer': {
					'claims': {
						'sub': data['User']['Id'],
						'email': data['User']['Email'],
					},
				},
			},
		}
		with open('/dev/null', 'w') as f:
			stdout = sys.stdout
			sys.stdout = f
			response = handler(event, None)
			sys.stdout = stdout
		print
		print response['statusCode']
		print json.dumps(json.loads(response['body']), indent=4)
