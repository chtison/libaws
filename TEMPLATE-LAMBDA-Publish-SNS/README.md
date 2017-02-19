# `TEMPLATE-LAMBDA-Publish-SNS`

### Configuration

```json
{
	"Topic": "!ImportValue SNS-TOPIC",
	"LogGroupRetentionInDays": 7,
	"Timeout": 10,
	"Environment": [
		["KEY1", "VALUE1"],
		["KEY2", "VALUE2"]
	],
	"Schedules": [
		"cron(0 18 * * ? *)",
		"rate(2 minutes)"
	]
}
```

| Name | Type |
| ---- | ---- |
| Topic | **Required**. ARN of [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html). |
| LogGroupRetentionInDays | Number in days of log retention. |
| Timeout | Number in minutes of lambda's timeout. |
| Environment | List of ["KEY", "VALUE"] for lambda environment. |
| Schedules | List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) |
