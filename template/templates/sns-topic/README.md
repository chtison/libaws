# `SNS-TOPIC`

## Configuration

```yaml
Subscriptions:
    - - email-json
      - "example@example.org"
    - - email
      - "example@example.com"
```

| Name          | Type                                                                                        |
| ------------- | ------------------------------------------------------------------------------------------- |
| Subscriptions | List of [Endpoint, Protocol](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html) |
