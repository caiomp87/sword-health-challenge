# Sword Health Challenge
API to manage a company's maintenance tasks.

## Description
This service is an API that manages the tasks performed during a working day. In this system, we have two types of users called Managers, and Technicians, and each one has exclusive roles and permissions on the platform.

## Usage

### Required envs
- `DB_CONNECTION_STRING` (MySQL connection string)
- `ACCESS_SECRET` (access secret to compose and build the JWT token)
- `CACHE_URL` (redis url server)

### Commands
- `make run` (execute the application)
- `make build` (generate application binary)

## Used tools
- Golang
- MySQL
- Redis
- Docker
- Kubernetes
