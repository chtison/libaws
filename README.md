# LibAWS
![libaws](misc/libaws.png)

LibAWS is a collection of AWS cloudformation stacks.

### Dependencies
- [docker](https://www.docker.com/)
	* [chtison/workspace](https://github.com/chtison/workspace)

### Usage
Launch the docker environment dedicated to this project:
```
make run
```
Authenticate to AWS:
```
aws configure
```
Go to a stack folder:
```sh
cd <STACK-FOLDER>
```
Print available makefile commands:
```
make help
```
Deploy a stack:
```sh
# It is sometimes required to specify an S3 bucket (e.g., lambda functions)
make deploy [S3_BUCKET=<BUCKET-NAME>]
```
Delete a stack:
```
make delete
```
