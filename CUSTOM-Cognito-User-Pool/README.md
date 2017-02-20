# Stack Name: `CUSTOM-Cognito-User-Pool`

## Exported variables

|  Name                  | Type |
| ---------------------- | ---- |
| *${AWS::StackName}*    | ARN of [AWS::Lambda::Function](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html) |

## Usage

This custom resource is a binding to [create_user_pool](https://boto3.readthedocs.io/en/latest/reference/services/cognito-idp.html#CognitoIdentityProvider.Client.create_user_pool).</br>
See [test](test/) for an example of a stack using this custom resource.
