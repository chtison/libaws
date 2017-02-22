# `TEMPLATE-LAMBDA`

This is not a stack.

## Configuration

### [data.lambda.example.yaml](data.lambda.example.yaml)

```yaml
DefaultName: Default

Runtime: python2.7
Handler: handler.handler
Zip: lambda.zip

LogGroupRetentionInDays: 7
Timeout: 10

Environment:
    KEY: VALUE
    HELLO: WORLD

Policies:
    - Effect: Allow
      Action:
          - sns:Publish
      Resource: "!ImportValue SNS-TOPIC"

Permissions:
    - Action: lambda:Invoke*
      Principal: s3.amazonaws.com
      SourceAccount: "!Sub '${AWS::AccountId}'"
      SourceArn: "!Ref Bucket"
    - Principal: apigateway.amazonaws.com

Schedules:
    - 'cron(0 18 * * ? *)'
    - 'rate(2 minutes)'
```

| Name                    | Type                                  |
| ----------------------- | ------------------------------------- |
| DefaultName             | Default name used for resources.      |
| LogGroupRetentionInDays | Number of days for log retention.     |
| Timeout                 | Number of minutes for lambda timeout. |
| Environment             | Lambda environment dictionary.        |
| Policies                | List of [policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html). |
| Schedules               | List of [schedule expressions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html). |
