# SNS-TOPIC-Main

### data.json:
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
At least one endpoint for the SNS topic must be filled in **data.json**.
- **EndpointsSMS**
	- list of telephone numbers in [E.164](https://en.wikipedia.org/wiki/List_of_country_calling_codes) format
- **EndpointsEmail**
	- list of emails

### Deploy the stack
```sh
make deploy
```
