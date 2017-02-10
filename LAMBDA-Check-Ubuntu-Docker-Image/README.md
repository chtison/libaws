# LAMBDA-Check-Ubuntu-Docker-Image

### data.json
```
{
	"Topic": "SNS-TOPIC-Main",
	"LambdaRole": "SNS-TOPIC-Main-ROLE-LAMBDA",
	"Schedules": [
		"cron(0 18 * * ? *)"
	]
}
```
- **Topic**: (**required**)
	* [Exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) whose value must be an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html).
- **LambdaRole**: (**required**)
	* [Exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) whose value must be an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html).
- **Schedules**: (**optional**)
 	* List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
