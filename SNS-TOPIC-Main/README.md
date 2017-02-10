# SNS-TOPIC-Main

### Exported stack's output values

| Name | Type |
| ---- | ---- |
| SNS-TOPIC-Main | ARN of [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) |
| SNS-TOPIC-Main-POLICY | ARN of [AWS::IAM::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html) |
| SNS-TOPIC-Main-ROLE-LAMBDA | ARN of [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html) |

### `data.json` example:
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

#### Properties

| Name | Type |
| ---- | ---- |
| EndpointsSMS | List of telephone numbers in [E.164 format ](https://en.wikipedia.org/wiki/List_of_country_calling_codes) |
| EndpointsEmail | List of emails |
