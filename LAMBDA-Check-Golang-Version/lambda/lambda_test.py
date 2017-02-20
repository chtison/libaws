import yaml, os, unittest

with open('../data.yaml') as f:
	data = yaml.load(f)
os.environ['TOPIC'] = 'X'
os.environ['GOVERSION'] = data['Environment']['GOVERSION']
execfile('lambda.py')

class TestLambdaFunction(unittest.TestCase):
	def test_handler(self):
		handler(None, None)
