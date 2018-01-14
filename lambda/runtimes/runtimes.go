package runtimes

// Runtime ...
type Runtime struct {
	Extension string
	Template  string
}

// Runtimes ...
var Runtimes = map[string]*Runtime{
	"python2.7": &Runtime{
		Extension: "py",
		Template: `import json

def handler(event, context):
	body = {
		"message": "Hello World !",
		"event": event,
		"context": context,
	}
	return {
		"statusCode": 200,
		"body": json.dumps(body),
	}
`,
	},
	"python3.6": &Runtime{
		Extension: "py",
		Template: `import json

def handler(event, context):
	body = {
		"message": "Hello World !",
		"event": event,
		"context": context,
	}
	return {
		"statusCode": 200,
		"body": json.dumps(body),
	}
`,
	},
}
