# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests

## Build and deploy

1.  cd lambda
2.  GOOS=linux GOARCH=amd64 go build -o bootstrap
3.  Zip bootstrap into function.zip
4.  cd ..
5.  cdk diff
6.  cdk deploy
