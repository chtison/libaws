import sys, unittest
execfile('lambda.py')

class TestLambdaFunction(unittest.TestCase):

	def log(self, data):
		print data, '',
		sys.stdout.flush()

	def test_checkType_1(self):
		data = {}
		self.assertTrue(checkType(42, int, data, '.AnInteger'), data)
		self.assertTrue(len(data) == 0)
		self.log(data)

	def test_checkType_2(self):
		data = {}
		self.assertFalse(checkType('42', int, data, '.AnInteger'), data)
		self.assertFalse(len(data) == 0)
		self.log(data)

	def test_normalizeParameter_1(self):
		data = {}
		properties = { 'Policies': None }
		self.assertFalse(normalizeParameter(['Policies', 'PasswordPolicy', 'MinimumLength'], int, properties, data))
		self.assertFalse(len(data) == 0)
		self.log(data)

	def test_normalizeParameter_2(self):
		data = {}
		properties = { 'Policies': { 'PasswordPolicy': { 'MinimumLength': None }}}
		self.assertFalse(normalizeParameter(['Policies', 'PasswordPolicy', 'MinimumLength'], int, properties, data))
		self.assertFalse(len(data) == 0)
		self.log(data)

	def test_normalizeParameter_3(self):
		data = {}
		properties = { 'Policies': { 'PasswordPolicy': { 'MinimumLength': 42 }}}
		self.assertTrue(normalizeParameter(['Policies', 'PasswordPolicy', 'MinimumLength'], int, properties, data))
		self.assertTrue(len(data) == 0)
		self.log(data)
