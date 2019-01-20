# Dependencies for the Solution

1. Install the [Go Programming Languge](https://golang.org/doc/install) for your OS.

2. Install [dep](https://golang.github.io/dep/docs/installation.html) to support Dependency management for Go.

3. Install the [Serverless Framework](https://serverless.com/framework/docs/getting-started/) and dependencies required by framework.

4. Edit the variables in the `Makefile` at project root for your own use.

5. Make sure you have an AWS account and a proper configuration in `.aws/credentials` for your OS.

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
