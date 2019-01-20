![](https://img.shields.io/badge/lang-golang-blue.svg) ![](https://img.shields.io/badge/framework-serverless-blue.svg)

Who hasn't grown up wanting to terraform a planet in a remote universe? Well now you can with the `Stoic Terraform Tool` using **AWS Lambda**, **Go**, and **Serverless** framework.

To quickly start terraforming your very own planet, follow the basic steps below and post questions/issues here on the repo. We're constantly looking for better ways to terraform planets, so you'll see near constant feature updates in the future.

# Some Considerations for the API

## Future Feature Work
This is an ongoing project and will receive continued feature work. In it's current state, you can only **Create** and **Read** basic planetary colonization actions. In addition, there is no backing data layer so your hard work won't persist yet. **DynamoDB** is coming, along with **Update** and **Delete** planetary colonization actions. Then, **Tests** and **Custom API Authorizers** will help to ensure the quality and security of our Colony service. Lastly, we'll wrap our Colony service efforts up with **CodePipeline** templates that you can use to ensure a _hands off_ approach to executing more planetary colonization efforts.

## Structure of the Service & API

For the structure of our project, we're going to be following a **Microservice style** approach. Each AWS Lambda function is responsible for one action only. For example, there are 4 separate Lambda functions for reading, creating, updating and deleting a colony within our Colony service.

## Working with API Gateway Events in Lambda

The easiest way in **Go** to provide the responses that the AWS API Gateway needs is to install the `github.com/aws/aws-lambda-go/events` package:

```bash
go get github.com/aws/aws-lambda-go/events
```

This provides a couple of useful types (`APIGatewayProxyRequest` and `APIGatewayProxyResponse`) which contain information about incoming HTTP requests and allow us to construct responses that the API Gateway understands.

Explore this package to better understand the types and what they offer if you plan to do any customization to the terraforming service.

# Dependencies for the Solution

1. Install the [Go Programming Languge](https://golang.org/doc/install) for your OS.

2. Install [dep](https://golang.github.io/dep/docs/installation.html) to support Dependency management for Go.

3. Install the [Serverless Framework](https://serverless.com/framework/docs/getting-started/) and dependencies required by framework.

4. Edit the variables in the `Makefile` file at project root for your own use.

5. Make sure you have an AWS account and a proper configuration in `.aws/credentials` for your OS.

6. Edit the configuration options in the `serverless.yml` file at project root for your own use.

Once you've gone through the above steps, you're ready to begin your career as a _Jr. Engineer_ in a _Single-Planet Terraform Corporation_ by continuing on with the commands in the sections following.

# Commands for the Solution

Build the project artifacts.

```bash
make build
```

Clean the project artifacts.

```bash
make clean
```

Deploy the project artifacts to AWS.

```bash
make deploy

rm -rf ./bin ./vendor Gopkg.lock
dep ensure -v
Root project is "aws-golang-terraform-colonies"
 3 transitively valid internal packages
 2 external packages imported from 1 projects
(0)   ✓ select (root)
(1)     ? attempt github.com/aws/aws-lambda-go with 2 pkgs; 14 versions to try
(1)         try github.com/aws/aws-lambda-go@v1.8.1
(1)     ✓ select github.com/aws/aws-lambda-go@v1.8.1 w/4 pkgs
  ✓ found solution with 4 packages from 1 projects

Solver wall times by segment:
  b-list-versions: 1.867151425s
      b-list-pkgs: 370.903803ms
           b-gmal: 238.742657ms
  b-source-exists:  95.338517ms
         new-atom:     615.52µs
      select-root:    238.967µs
          satisfy:    121.959µs
      select-atom:     90.621µs
            other:     15.632µs

  TOTAL: 2.573219101s

(1/1) Wrote github.com/aws/aws-lambda-go@v1.8.1
sls deploy --verbose --aws-profile stoic --stage dev
Serverless: Packaging service...
Serverless: Excluding development dependencies...
Serverless: Creating Stack...
Serverless: Checking Stack create progress...
CloudFormation - CREATE_IN_PROGRESS - AWS::CloudFormation::Stack - Stoic-Athena-Colonies-Stack
...
Service Information
service: Colonies
stage: dev
region: us-west-2
stack: Stoic-Athena-Colonies-Stack
api keys:
  None
endpoints:
  POST - https://6s6prv1gfe.execute-api.us-west-2.amazonaws.com/dev/colonies
  GET - https://6s6prv1gfe.execute-api.us-west-2.amazonaws.com/dev/colonies/{id}
functions:
  create: Colonies-dev-create
  read: Colonies-dev-read
layers:
  None

Stack Outputs
CreateLambdaFunctionQualifiedArn: arn:aws:lambda:us-west-2:750444023825:function:Stoic-Athena-Colonies-Dev-Create-Func:2
ServiceEndpoint: https://6s6prv1gfe.execute-api.us-west-2.amazonaws.com/dev
ServerlessDeploymentBucketName: stoic-athena-colonies-st-serverlessdeploymentbuck-enuzc6d7yab4
ReadLambdaFunctionQualifiedArn: arn:aws:lambda:us-west-2:750444023825:function:Stoic-Athena-Colonies-Dev-Read-Func:2
```

Once deployed to AWS, now test that the `create` function is working. Make sure you noted the API endpoints generated as outputs in the `make deploy` command.

```bash
curl -X POST \
  https://6s6prv1gfe.execute-api.us-west-2.amazonaws.com/dev/colonies \
  -H 'Content-Type: application/json' \
  -d '{"name":"ABCDEF-0123456789XYZ","planet":"SIRIUS","corporation":"117-TETRISCORP","coords":"100.2342356234, -345.7823425345"}'

{
    "message": "Colony created successfully",
    "transmission": {
        "name": "ABCDEF-0123456789XYZ",
        "planet": "SIRIUS",
        "corporation": "117-TETRISCORP",
        "coords": "100.2342356234, -345.7823425345"
    }
}
```

Test that the `read` function is working.

```bash
curl -X GET \
  https://6s6prv1gfe.execute-api.us-west-2.amazonaws.com/dev/colonies/ABCDEF-0123456789XYZ \
  -H 'Content-Type: application/json'

{
    "message": "Colony read successfully",
    "transmission": {
        "name": "ABCDEF-0123456789XYZ",
        "planet": "SIRIUS",
        "corporation": "117-TETRISCORP",
        "coords": "1.02345, -2.23456"
    }
}
```

Remove the project artifacts from AWS.

```bash
make remove
```

# In Parting

It's just me here. Holler if you want to join the movement. I hope you had some fun in learning the technology stack used in this project. This solution was meant to provide a learning platform for myself and others that are interested in keeping the fun in development and also embracing continuous learning and self-improvement.
