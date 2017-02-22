import json

def handler(event, context):
	print json.dumps(event, default=str)
	return {
		"message": "Bye World !"
	}
