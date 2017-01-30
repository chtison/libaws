# LAMBDA-Check-Golang-Version

### data.json
```
{
	"Topic": "SNS-TOPIC-Main",
	"GolangVersion": "go1.7.5",
	"Schedules": [
		"cron(0 10 * * ? *)",
		"cron(0 20 * * ? *)"
	]
}
```
- **Topic**: (**required**)
  - Will be [imported](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html) as an [exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) pointing to an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) object.
- **GolangVersion**: (**required**)
  - Sets to the VERSION variable in [function.py](function.py) environment.
- **Schedules**: (**optional**)
  - List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html).

### Deploy the stack
```sh
make deploy S3-BUCKET=bucket-name
```
