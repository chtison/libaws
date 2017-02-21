import json

def handler(event, context):
	print json.dumps(event, default=str)
	if 'Error' in event:
		raise Exception(event['Error'])
	return "Bye World !"
