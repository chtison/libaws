# Stack Name: `TEST-CUSTOM-Cognito-User-Pool`

## Configuration

### [data.json](data.json)
```
{
	"ServiceToken": "!ImportValue CUSTOM-Cognito-User-Pool"
}
```

| Name | Type |
| ---- | ---- |
| ServiceToken | **Required**. ARN of [AWS::Lambda::Function](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html) |
