# Stack Name: `TEST-CUSTOM-Cognito-User-Pool`

## Configuration

### [data.yaml](data.yaml)
```yaml
ServiceToken_CognitoUserPool: "!ImportValue CUSTOM-Cognito-User-Pool"
```

| Name | Type |
| ---- | ---- |
| ServiceToken_CognitoUserPool | **Required**. ARN of [AWS::Lambda::Function](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html) |
