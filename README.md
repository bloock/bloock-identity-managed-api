# Bloock Identity Managed API

This is an API for those who want to create, emit and offer verifiable credentials or VC's (https://www.w3.org/TR/vc-data-model-2.0/), using bloock product's and following the principles of self-sovereign identity and privacy by default. 

---

## Installation

There are two options for running this service:

1. [Docker](#docker-guide)
2. [Standalone](#standalone-guide)

### Docker Guide

Running this API using Docker allows for a quick setup and/or deployment.

#### Locally

To start the service using Docker Compose, you can follow the following steps:

To start the service with MySQL:

```
docker compose -f docker-compose-mysql.yaml up
```

To start the service with Postgres:

```
docker compose -f docker-compose-postgres.yaml up
```

To start the service with MemDB:

```
docker compose -f docker-compose.yaml up
```

> **NOTE:** Remember to update the [configuration](#configuration) variables as required.

### Docker image

We mantain a Docker repository with the latest releases of this repository in [DockerHub](https://hub.docker.com/repository/docker/bloock/managed-api/general).

---

### Standalone Guide

You can also run this service as a common Golang binary if you need it.

#### Standalone Requirements

- Makefile toolchain
- Unix-based operating system (e.g. Debian, Arch, Mac OS X)
- [Go](https://go.dev/) 1.20

#### Standalone Setup

1. Add the required [configuration](#configuration) variables.
2. Run `go run cmd/main.go`

---

### Configuration

The service uses viper as a configuration manager, currently supporting environment variables and a YAML configuration file.

##### Variables

- **BLOOCK_API_PORT**: The main API port; default is 8080.
- **BLOOCK_API_HOST**: The main API host IP; default is 10.0.5.23.
- **BLOOCK_API_KEY**: Your Bloock API key.
- **BLOOCK_WEBHOOK_SECRET_KEY**: Your webhook secret key.
- **BLOOCK_LOCAL_PRIVATE_KEY**: If you want to set a local key, you should provide your private key.
- **BLOOCK_LOCAL_PUBLIC_KEY**: If you want to set a local key, you should provide your public key.
- **BLOOCK_MANAGED_KEY_ID**: If you want to set a managed key, ypu should provide you api key id (UUID format).
- **BLOOCK_PUBLIC_HOST:** You API public host.
- **BLOOCK_ISSUER_DID_METHOD**: Advanced. If you want a different DID method type allowed by BLOOCK. By default, 'polygonid' it's used.
- **BLOOCK_ISSUER_DID_BLOCKCHAIN**: Advanced. If you want a different DID blockchain type allowed by BLOOCK. By default, 'polygon' blockchain it's used.
- **BLOOCK_ISSUER_DID_NETWORK**: Advanced. If you want a different DID blockchain type allowed by BLOOCK. By default, 'mumbai' network it's used.
- **BLOOCK_DB_CONNECTION_STRING**: Your database URL; e.g., mysql://username:password@localhost:3306/mydatabase.
- **BLOOCK_API_DEBUG_MODE**: debug mode prints more log information; true or false.

##### Configuration file

The configuration file should be named `config.yaml`. The service will try to locate this file in the root directory unless the BLOOCK_CONFIG_PATH is defined (i.e. `BLOOCK_CONFIG_PATH="app/conf/"`).

Sample content of `config.yaml`:

```yaml
BLOOCK_API_HOST: "0.0.0.0"
BLOOCK_API_PORT: "8080"
BLOOCK_API_DEBUG_MODE: "false"

BLOOCK_DB_CONNECTION_STRING: "file:bloock?mode=memory&cache=shared&_fk=1"

BLOOCK_API_KEY: ""
BLOOCK_WEBHOOK_SECRET_KEY:  ""

BLOOCK_PUBLIC_HOST: ""

BLOOCK_LOCAL_PRIVATE_KEY: ""
BLOOCK_LOCAL_PUBLIC_KEY: ""
BLOOCK_MANAGED_KEY_ID: ""

BLOOCK_ISSUER_DID_METHOD: ""
BLOOCK_ISSUER_DID_BLOCKCHAIN: ""
BLOOCK_ISSUER_DID_NETWORK: ""
```

### Database
The service supports three types of relational databases: MemDB (SQLite), MySQL, and Postgres. You only need to provide the database URL in the following format:

````
MySQL: <user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True
Postgres: postgresql://username:password@localhost:5432/mydatabase
MemDB: file:dbname?mode=memory&cache=shared&_fk=1
````

---

## Documentation

You can access the following Postman collection where is the specification for the public endpoint of this API.

[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://www.postman.com/bloock/workspace/bloock-api/documentation/16945237-1727027f-e3e7-4fe0-969d-afa295eaf2ca)

---

## License

See [LICENSE](LICENSE.md).

