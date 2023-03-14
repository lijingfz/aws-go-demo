# aws-go-demo
DO test for Suirui Lambda use Go to add eventbridge rule

[ec2-user@ip-172-31-6-0 test0314]$ go mod init test0314
go: creating new go.mod: module test0314
go: to add module requirements and sums:
        go mod tidy
[ec2-user@ip-172-31-6-0 test0314]$ go mod tidy
go: finding module for package github.com/aws/aws-lambda-go/lambda
go: finding module for package github.com/aws/aws-sdk-go-v2/config
go: finding module for package github.com/aws/aws-sdk-go-v2/service/eventbridge
go: finding module for package github.com/aws/aws-sdk-go-v2/service/eventbridge/types
go: found github.com/aws/aws-lambda-go/lambda in github.com/aws/aws-lambda-go v1.38.0
go: found github.com/aws/aws-sdk-go-v2/config in github.com/aws/aws-sdk-go-v2/config v1.18.16
go: found github.com/aws/aws-sdk-go-v2/service/eventbridge in github.com/aws/aws-sdk-go-v2/service/eventbridge v1.18.6
go: found github.com/aws/aws-sdk-go-v2/service/eventbridge/types in github.com/aws/aws-sdk-go-v2/service/eventbridge v1.18.6
[ec2-user@ip-172-31-6-0 test0314]$ GOOS=linux GOARCH=amd64 go build -o main main.go
[ec2-user@ip-172-31-6-0 test0314]$ ls -al
total 12012
drwxrwxr-x  2 ec2-user ec2-user       61 Mar 14 01:07 .
drwx------ 16 ec2-user ec2-user     4096 Mar 14 01:05 ..
-rw-rw-r--  1 ec2-user ec2-user      962 Mar 14 01:06 go.mod
-rw-rw-r--  1 ec2-user ec2-user     4198 Mar 14 01:06 go.sum
-rwxrwxr-x  1 ec2-user ec2-user 12276423 Mar 14 01:07 main
-rw-rw-r--  1 ec2-user ec2-user     1572 Mar 14 01:05 main.go
[ec2-user@ip-172-31-6-0 test0314]$ zip main.zip main
  adding: main (deflated 48%)
[ec2-user@ip-172-31-6-0 test0314]$ aws lambda create-function --function-name my-function314 --runtime go1.x --role arn:aws:iam::890717383483:role/lambda-ex --handler main --zip-file fileb://main.zip
{
    "FunctionName": "my-function314",
    "FunctionArn": "arn:aws:lambda:us-west-2:890717383483:function:my-function314",
    "Runtime": "go1.x",
    "Role": "arn:aws:iam::890717383483:role/lambda-ex",
    "Handler": "main",
    "CodeSize": 6424483,
    "Description": "",
    "Timeout": 3,
    "MemorySize": 128,
    "LastModified": "2023-03-14T01:07:59.358+0000",
    "CodeSha256": "mZSpL8TZl9mUyC3B3njmEZzqWZPbrxIMGcpydnhX6K0=",
    "Version": "$LATEST",
    "TracingConfig": {
        "Mode": "PassThrough"
    },
    "RevisionId": "3d73020c-f761-4759-ab11-490c10ca406a",
    "State": "Pending",
    "StateReason": "The function is being created.",
    "StateReasonCode": "Creating",
    "PackageType": "Zip",
    "Architectures": [
        "x86_64"
    ],
    "EphemeralStorage": {
        "Size": 512
    },
    "SnapStart": {
        "ApplyOn": "None",
        "OptimizationStatus": "Off"
    }
}
[ec2-user@ip-172-31-6-0 test0314]$
