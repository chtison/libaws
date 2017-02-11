import json, os, unittest

with open('data.json') as f:
	data = json.load(f)
os.environ['TOPIC'] = 'X'
os.environ['VERSION'] = data['TorVersion']
execfile('lambda.py')

class TestLambdaFunction(unittest.TestCase):
	def test_handler(self):
		handler(None, None)
