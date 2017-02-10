# LibAWS
![libaws](misc/libaws.png)

LibAWS is a collection of AWS cloudformation stacks.

### Dependencies
- [docker](https://www.docker.com/)
	* [chtison/workspace](https://github.com/chtison/workspace)

### Usage
Launch the docker environment dedicated to this project:
```
make
```
Authenticate to AWS:
```
aws configure
```
Deploy a stack:
```sh
cd <STACK-FOLDER>
# Makefile usage
make
# It is sometimes required to specify an S3 bucket (e.g., lambda functions)
make deploy [S3-BUCKET=<BUCKET-NAME>]
```
