# Stack Name: `LAMBDA-Check-Golang-Version`

## Configuration

### [data.json](data.json)
```
{
	"Topic": "!ImportValue SNS-TOPIC-Main",
	"LambdaRole": "!ImportValue SNS-TOPIC-Main-ROLE-LAMBDA",
	"GolangVersion": "go1.7.5",
	"Schedules": [
		"cron(0 18 * * ? *)"
	]
}
```

| Name | Type |
| ---- | ---- |
| Topic | **Required**. ARN of [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) |
| LambdaRole | **Required**. ARN of [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html) |
| GolangVersion | **Required**. Expected [latest Golang version](https://golang.org/dl/) |
| Schedules | List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) |
