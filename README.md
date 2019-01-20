![](https://img.shields.io/badge/lang-golang-blue.svg) ![](https://img.shields.io/badge/framework-serverless-blue.svg)

Who hasn't grown up wanting to terraform a planet in a remote universe? Well now you can with the `Stoic Terraform Tool` using **AWS Lambda**, **Go**, and **Serverless** framework.

To quickly start terraforming your very own planet, follow the basic steps below and post questions/issues here on the repo. We're constantly looking for better ways to terraform planets, so you'll see near constant feature updates in the future.

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
```

Once deployed to AWS, now test that the `create` function is working.

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
