# Default Stack Name: `SNS-TOPIC`

## Exported variables

|  Name                  | Type |
| ---------------------- | ---- |
| *${AWS::StackName}*    | ARN of [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) |

## Configuration

### [data.yaml](data.yaml)

```yaml
Subscriptions:
    - - email-json
      - "example@example.org"
    - - email
      - "example@example.com"
```

| Name          | Type |
| ------------- | ---- |
| Subscriptions | List of [[Endpoint, Protocol]]}(https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html) |
