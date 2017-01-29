# SNS-TOPIC-Main

### Usage
At least one endpoint for the SNS topic must be filled in **data.json**.
###### data.json:
```
{
	"EndpointsSMS": [
		"+33611223344"
	],
	"EndpointsEmail": [
		"email@example.org",
		"email@example.com"
	]
}
```
##### Deploy the stack
```sh
make deploy
```
