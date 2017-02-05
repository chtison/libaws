# LAMBDA-Check-Golang-Version

### data.json
```
{
	"Topic": "SNS-TOPIC-Main",
	"PolicyPublishToTopic": "SNS-TOPIC-Main-POLICY",
	"GolangVersion": "go1.7.5",
	"Schedules": [
		"cron(0 10 * * ? *)",
		"cron(0 20 * * ? *)"
	]
}
```
- **Topic**: (**required**)
	* [Exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) whose value must be an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html).
- **PolicyPublishToTopic**: (**required**)
	* [Exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) whose value must be an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::SNS::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html).<br>
	**PolicyPublishToTopic** must grant permission to publish to **Topic**.
- **GolangVersion**: (**required**)
 	* Sets to the VERSION variable in [function.py](function.py) environment.
- **Schedules**: (**optional**)
 	* List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
