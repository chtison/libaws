# Stack Name: `LAMBDA-Check-Tor-Version`

## Configuration

### [data.json](data.json)
```
{
	"Topic": "!ImportValue SNS-TOPIC-Main",
	"LambdaRole": "!ImportValue SNS-TOPIC-Main-ROLE-LAMBDA",
	"TorVersion": "0.2.9.9",
	"Schedules": [
		"cron(0 18 * * ? *)"
	]
}
```

| Name | Type |
| ---- | ---- |
| Topic | **Required**. ARN of [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) |
| LambdaRole | **Required**. ARN of [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html) |
| TorVersion | **Required**. Expected [latest Tor version](https://gitweb.torproject.org/tor.git/plain/ReleaseNotes) |
| Schedules | List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) |
