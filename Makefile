# Capture AWS account values.
AWS_PROFILE=stoic
AWS_STAGE=dev

# Collect all functions that should be built and deployed.
functions := $(shell find functions -name \*main.go | awk -F'/' '{print $$2}')

build: ## Build golang binaries
	dep ensure -v
	@for function in $(functions) ; do \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$function functions/$$function/main.go ; \
	done

.PHONY: clean
clean: ## Remove golang binaries
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: clean build  ## Remove/build golang binaries
	sls deploy --verbose --aws-profile ${AWS_PROFILE} --stage ${AWS_STAGE}

.PHONY: remove
remove: ## Remove golang binaries from AWS
	sls remove --verbose --aws-profile ${AWS_PROFILE} --stage ${AWS_STAGE}
