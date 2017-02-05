# LAMBDA-Check-Ubuntu-Docker-Image

### data.json
```
{
	"Topic": "SNS-TOPIC-Main",
	"PolicyPublishToTopic": "SNS-TOPIC-Main-POLICY",
	"Schedules": [
		"cron(0 18 * * ? *)"
	]
}

```
- **Topic**: (**required**)
	* [Exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) whose value must be an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html).
- **PolicyPublishToTopic**: (**required**)
	* [Exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) whose value must be an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::SNS::ManagedPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html).<br>
	**PolicyPublishToTopic** must grant permission to publish to **Topic**.
- **Schedules**: (**optional**)
 	* List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).
