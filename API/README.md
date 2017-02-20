# Default Stack Name: `API`

## Exported variables

|  Name                  | Type                     |
| ---------------------- | ------------------------ |
| *${AWS::StackName}*    | URL of the deployed API. |

## Configuration

### [data.yaml](data.yaml)
```yaml
StageName: api
Lambda:
    Permissions:
        - Principal: apigateway.amazonaws.com
```

| Name      | Type                                                 |
| --------- | ---------------------------------------------------- |
| StageName | **Required**. Name for the deployment stage.         |
| Lambda    | [TEMPLATE-LAMBDA](../TEMPLATE-LAMBDA) configuration. |
