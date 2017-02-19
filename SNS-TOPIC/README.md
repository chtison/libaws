# Default Stack Name: `SNS-TOPIC`

## Exported variables

|  Name | Type |
| ---------------------- | ---- |
| *${AWS::StackName}* | ARN of [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) |

## Configuration

### [data.json](data.json)

```json
{
	"Subscriptions": [
		["email", "email@example.org"],
		["email-json", "email@example.com"]
	]
}
```

| Name | Type |
| ---- | ---- |
| Subscriptions | List of [[Endpoint, Protocol]](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html) |
