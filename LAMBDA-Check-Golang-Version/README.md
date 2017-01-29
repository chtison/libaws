# LAMBDA-Check-Golang-Version

### data.json
```
{
	"Topic": "SNS-TOPIC-Main",
	"GolangVersion": "go1.7.5",
	"CronHours": [
      10,
      20
	]
}

```
- **Topic**: **string** (**required**)
  - Will be [imported](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html) as an [exported variable](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-exports.html) pointing to an [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) object.
- **GolangVersion**: **string** (**required**)
  - Sets to the VERSION variable in [function.py](function.py) environment.
- **CronHours**: **[0 <= x <= 23]** (**required**)
  - Sets the hours (UTC) at which lambda function will be called every day.

### Deploy the stack
```sh
make deploy S3-BUCKET=bucket-name
```
