import os, unittest

os.environ['TOPIC'] = 'X'
execfile('lambda.py')

class TestLambdaFunction(unittest.TestCase):
	def test_handler(self):
		handler(None, None)
