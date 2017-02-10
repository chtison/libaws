# LAMBDA-Check-Golang-Version

### data.json
```
{
	"Topic": "SNS-TOPIC-Main",
	"LambdaRole": "SNS-TOPIC-Main-ROLE-LAMBDA",
	"GolangVersion": "go1.7.5",
	"Schedules": [
		"cron(0 10 * * ? *)",
		"cron(0 20 * * ? *)"
	]
}
```
- **Topic**: (**required**)
	* [Exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) whose value must be an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html).
- **LambdaRole**: (**required**)
	* [Exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) whose value must be an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html).
- **GolangVersion**: (**required**)
 	* Sets to the VERSION variable in the [function.py](function.py) environment.
- **Schedules**: (**optional**)
 	* List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
