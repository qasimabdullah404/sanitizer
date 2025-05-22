# Log Sanitizer API

[![Go Report Card](https://goreportcard.com/badge/github.com/qasimabdullah404/sanitizer)](https://goreportcard.com/report/github.com/qasimabdullah404/sanitizer)  

---

## Overview

**Log Sanitizer** is a lightweight, high-performance Go web service that ingests log files or strings and automatically obfuscates sensitive information like passwords, API keys, tokens, and secrets. It’s designed to help you securely process logs before sharing, storing, or analyzing them.

Built with clean modular architecture using:

- **Go** — concurrency and performance
- **Zap** — structured logging
- **Mux** — HTTP router
- **Middleware** — request logging & metrics

---

## Features

- **Upload logs via HTTP POST** — simple multipart file upload interface
- **Obfuscate secrets automatically** — pattern-based redaction of credentials
- **Read logs from PostgreSQL database** (optional)
- **Configurable via environment variables** — port, log level, etc.
- **Dockerized for easy deployment** with HTTPS support (self-signed or Let’s Encrypt)
- **Extensible & testable** — modular code with comprehensive unit tests

---

## Getting Started

### Prerequisites

- [Go 1.20+](https://golang.org/dl/)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)

### RUN

Clone the repo:

```bash
git clone https://github.com/qasimabdullah404/sanitizer.git
cd sanitizer
docker compose up --build -d
