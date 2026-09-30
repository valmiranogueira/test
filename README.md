# Go Trivy Vulnerability Test

A simple Go project used to test vulnerability detection and automated dependency updates with Trivy.

The project intentionally includes a vulnerable Go dependency so that Trivy can detect it after building the Docker image.

## Requirements

* Go
* Docker
* Trivy
* Make

## Build

Build the Docker image:

```bash
make build
```

By default this creates:

```text
valmiranogueira/test:1.2.0
```

Override `IMAGE` to build another pipeline tag, for example:

```bash
make build IMAGE=valmiranogueira/test:1.2.0-1
```

## Run

Run the application:

```bash
make run
```

## Scan for Vulnerabilities

Scan the Docker image with Trivy:

```bash
make scan
```

Or directly:

```bash
trivy image --scanners vuln test:latest
```

The project intentionally uses:

```text
golang.org/x/text v0.38.0
```

This version contains a known vulnerability and is included for testing purposes.

## Fix the Vulnerability

Update the vulnerable dependency:

```bash
go get golang.org/x/text@v0.39.0
go mod tidy
```

Rebuild the image:

```bash
make build
```

Then scan it again:

```bash
make scan
```

## Clean Up

Remove the Docker image:

```bash
make clean
```

## Purpose

This repository is intended for testing workflows such as:

* Trivy vulnerability scanning
* Automated dependency updates
* Security pipelines
* Pull Request automation
* Automated releases
* Vulnerability remediation

## Weekly security pipeline fixture

The repository includes the minimal paths consumed by the weekly security build:

* `pkg/version/version.txt` provides the base operator version.
* `config/manager/` and `deploy/` contain image references rewritten by the pipeline.
* `deploy/cr.yaml` provides the pillar image used to derive `PILLAR_VERSION`.
* `e2e-tests/release_versions` tracks the image used by release tests.

> **Warning:** This repository intentionally contains a vulnerable dependency. It should only be used for testing and development purposes.
