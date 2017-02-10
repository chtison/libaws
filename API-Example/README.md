# Stack Name: `API-Example`

| Exported Variable Name | Type |
| ---------------------- | ---- |
| API-Example | URL of the deployed API |

## Configuration

### [data.json](data.json)
```
{
	"LambdaRole": "!ImportValue ROLE-LAMBDA-Default",
	"ApiTitle": "API-Example",
	"StageName": "api"
}
```

| Name | Type |
| ---- | ---- |
| LambdaRole | **Required**. ARN of [AWS::IAM::Role](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html) |
| ApiTitle | **Required**. Title for the API |
| StageName | **Required**. Name for the deployment stage |
