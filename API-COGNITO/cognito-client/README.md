# COGNITO-CLIENT

Here is a client for retrieving an ID token from COGNITO service and test it on the secured API.

### Usage
##### Build
```
make build
```
##### Run
```
make run USER_POOL_ID=<> CLIENT_ID=<> EMAIL=<> PASSWORD=<>
```
##### Run step by step
Fill [userPoolId](src/main.ts#L7) and [clientId](src/main.ts#L8), then:
```
node dist/main.js usage
node dist/main.js signUp EMAIL PASSWORD
node dist/main.js confirmRegistration EMAIL CONFIRMATION-CODE
node dist/main.js signIn EMAIL PASSWORD
```
