# Stack Name: `LAMBDA-Check-Ubuntu-Docker-Image`

## Configuration

### [data.json](data.json)
```
{
	"Topic": "SNS-TOPIC-Main",
	"LambdaRole": "SNS-TOPIC-Main-ROLE-LAMBDA",
	"Schedules": [
		"cron(0 18 * * ? *)"
	]
}
```

| Name | Type |
| ---- | ---- |
| Topic | **Required**. Exported variable name -> ARN of [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) |
| LambdaRole | **Required**. Exported variable name -> ARN of [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html) |
| Schedules | List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) |
