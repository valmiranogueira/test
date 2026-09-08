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

This creates:

```text
test:latest
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

> **Warning:** This repository intentionally contains a vulnerable dependency. It should only be used for testing and development purposes.
