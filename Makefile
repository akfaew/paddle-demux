PROJECT = your_project_name
TEST_ARGS = -failfast -tags=$(ALL_TAGS)

## development
run:
	dev_appserver.py prod.yaml --enable_host_checking=false

update:
	go get -u
	go mod tidy
	go mod verify

fmt:
	go fmt ./...

lint:
	golangci-lint run

test: fmt lint
	go test $(TEST_ARGS) ./...

## release
check-workspace: fmt
	# check if workspace is clean
	@if ! git diff-index --quiet HEAD --; then \
		echo "You have unstaged changes"; \
		git status; \
		exit 1; \
	fi

deploy: check-workspace test
	echo y | gcloud app deploy prod.yaml --project=$(PROJECT)

deploy-safe: check-workspace test
	gcloud app deploy prod.yaml --project=$(PROJECT) --no-promote

release: deploy push

push: test
	git push origin

## misc
clean:
	rm -f coverage.out coverage.html callvis.dot callvis.png
