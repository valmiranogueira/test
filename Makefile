IMAGE ?= test:latest

.PHONY: build scan run clean

build:
	docker build -t $(IMAGE) .

scan:
	trivy image --scanners vuln $(IMAGE)

test:
	echo 'Running tests...'

run:
	docker run --rm $(IMAGE)

clean:
	docker image rm -f $(IMAGE) 2>/dev/null || true
