# Stack Name: `TEST-CUSTOM-Cognito-User-Pool-Client`

## Configuration

### [data.yaml](data.yaml)
```yaml
ServiceToken_CognitoUserPool: "!ImportValue CUSTOM-Cognito-User-Pool"
ServiceToken_CognitoUserPoolClient: "!ImportValue CUSTOM-Cognito-User-Pool-Client"
```

| Name | Type |
| ---- | ---- |
| ServiceToken_CognitoUserPool       | **Required**. ARN of [AWS::Lambda::Function](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html) |
| ServiceToken_CognitoUserPoolClient | **Required**. ARN of [AWS::Lambda::Function](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html) |
