# Bloock Identity Managed API

This is an API for those who want to create, emit and offer verifiable credentials or VC's (https://www.w3.org/TR/vc-data-model-2.0/), using bloock product's and following the principles of self-sovereign identity and privacy by default. 

---

## Table of Contents

- [Installation](#installation)
    - [Docker Setup Guide](#docker-setup-guide)
        - [Option 1: Pull and Run the Docker Image](#option-1-pull-and-run-the-docker-image)
        - [Option 2: Use Docker Compose with Database Containers](#option-2-use-docker-compose-with-database-containers)
- [Configuration](#configuration)
    - [Variables](#variables)
    - [Configuration File](#configuration-file)
- [Database Support](#database-support)
- [Documentation](#documentation)
- [License](#license)

---

## Installation

You have one primary method to set up and run the Identity Bloock Managed API:

1. [Docker Setup Guide](#docker-setup-guide)

Each method has its advantages and use cases.

### Docker Setup Guide

Docker offers a convenient way to package and distribute the API, along with its required dependencies, in a self-contained environment. It's an excellent choice if you want a quick and hassle-free setup, or if you prefer isolation between your application and the host system.

### Option 1: Pull and Run the Docker Image

This option is straightforward and ideal if you want to get started quickly. Follow these steps:

1. **Pull the Docker Image:**

    - Open your terminal or command prompt.

    - Run the following command to pull the Docker image from [DockerHub](https://hub.docker.com/r/bloock/identity-managed-api):

      ```bash
      docker pull bloock/managed-api
      ```

      This command fetches the latest version of the Identity Bloock Managed API image from [DockerHub](https://hub.docker.com/r/bloock/identity-managed-api). We maintain a Docker repository with the latest releases of this repository.


2. **Create a `.env` File:**

    - In your project directory, create a `.env` file. You can use a text editor of your choice to create this file.

    - This file will contain the configuration for the API, including environment variables. Refer to the [Variables](#variables) section for a list of environment variables and their descriptions.

    - In the `.env` file, define the environment variables you want to configure for the API. Each environment variable should be set in the following format:
      ```txt
      VARIABLE_NAME=VALUE
      ```

    - Here's an example of what your `.env` file might look like:

      ```txt
      BLOOCK_DB_CONNECTION_STRING=file:bloock?mode=memory&cache=shared&_fk=1
      BLOOCK_API_KEY=your_api_key
      BLOOCK_WEBHOOK_SECRET_KEY=your_webhook_secret_key
      BLOOCK_PUBLIC_HOST=https://bloock.com/
      BLOOCK_MANAGED_KEY_ID=your_managed_key_id
      ```

      > **NOTE:** For the **BLOOCK_DB_CONNECTION_STRING** environment variable, you have the flexibility to specify your own MySQL or PostgreSQL infrastructure. Clients can provide their connection string for their database infrastructure. See the [Database](#database-support) section for available connections.

3. **Run the Docker Image with Environment Variables:**

    - Run the following command to start the Identity Bloock Managed API container while passing the `.env` file as an environment variable source:

     ```bash
     docker run --env-file .env -p 8080:8080 bloock/managed-api
     ```

    - This command maps the `.env` file into the container, ensuring that the API reads its configuration from the file. Viper automatically read these environment variables and make them accessible to the code.

    - By default, the above command runs the Docker container in the foreground, displaying API logs and output in your terminal. You can interact with the API while it's running.

   3.1. **Running Docker in the Background (Daemon Mode)**

    - Append the `-d` flag to the docker run command as follows:

    ```bash
    docker run -d --env-file config.txt -p 8080:8080 bloock/managed-api
    ```

   The `-d` flag tells Docker to run the container as a background process. You can continue using your terminal for other tasks while the Identity Bloock Managed API container runs silently in the background.


4. **Access the API:**

    - After running the Docker image, the Identity Bloock Managed API will be accessible at `http://localhost:8080`.

    - You can now make API requests to interact with the service.

By following these steps, you can quickly deploy the Identity Bloock Managed API as a Docker container with your customized configuration.

### Option 2: Use Docker Compose with Database Containers

If you need a more complex setup, such as using a specific database like **MySQL**, **Postgres** or **MemDB**, Docker Compose is your choice. Follow these steps:

1. **Choose the Docker Compose File:**

    - In our [repository](https://github.com/bloock/bloock-identity-managed-api), you will find Docker Compose files for different database types:

        - `docker-compose-mysql.yaml` for MySQL
        - `docker-compose-postgres.yaml` for PostgreSQL
        - `docker-compose.yaml for MemDB` (SQLite)


2. **Copy the Chosen Docker Compose File:**

    - Choose the Docker Compose file that corresponds to the database type you want to use. For example, if you prefer MySQL, copy `docker-compose-mysql.yaml`.


3. **Configure Environment Variables:**

    - Open the Docker Compose file in a text editor. Inside the file, locate the environment section for the api service. Here, you can specify environment variables that configure the API.

    - Refer to the [Variables](#variables) section for a list of environment variables and their descriptions.


4. **Set Environment Variables:**

    - In the `environment` section, you can set environment variables in the following format:
      ```yaml
      VARIABLE_NAME: "VALUE"
      ```

    - Here's an example of what your `environment` section might look like:

      ```yaml
      BLOOCK_DB_CONNECTION_STRING: "file:bloock?mode=memory&cache=shared&_fk=1"
      BLOOCK_API_KEY: "your_api_key"
      BLOOCK_WEBHOOK_SECRET_KEY: "your_webhook_secret_key"
      BLOOCK_PUBLIC_HOST: "https://bloock.com/"
      BLOOCK_MANAGED_KEY_ID: "your_managed_key_id"
      ```

5. **Run Docker Compose:**

    - Open your terminal, navigate to the directory where you saved the Docker Compose file, and run the following command:

    ```bash
     docker-compose -f docker-compose-mysql.yaml up
     ```

   Replace `docker-compose-mysql.yaml` with the name of the Docker Compose file you selected.

   5.1. **Running Docker in the Background (Daemon Mode)**

    - Append the `-d` flag to the docker run command as follows:

    ```bash
    docker-compose -f docker-compose-mysql.yaml up -d
    ```

   The `-d` flag tells Docker to run the container as a background process. You can continue using your terminal for other tasks while the Identity Bloock Managed API container runs silently in the background.


6. **Access the API:**

    - After running the Docker Compose command, the Identity Bloock Managed API will be accessible at http://localhost:8080. You can make API requests to interact with the service.

By following these steps, you can quickly set up the Identity Bloock Managed API with your chosen database type using the provided Docker Compose files.

---

## Configuration

The Identity Bloock Managed API leverages Viper, a powerful configuration management library, currently supporting environment variables and a YAML configuration file.

### Variables

Here are the configuration variables used by the Identity Bloock Managed API:

- **BLOOCK_API_KEY** (**REQUIRED**)
    - **Description**: Your unique [BLOOCK API key](https://docs.bloock.com/libraries/authentication/create-an-api-key).
    - **Purpose**: This [API key](https://docs.bloock.com/libraries/authentication/create-an-api-key) is required for authentication and authorization when interacting with the Bloock Identity Managed API. It allows you to securely access and use the API's features.
    - **[Create API Key](https://docs.bloock.com/libraries/authentication/create-an-api-key)**
- **BLOOCK_DB_CONNECTION_STRING** (***OPTIONAL***)
    - **Description**: Your custom database connection URL.
    - **Default**: "file:bloock?mode=memory&cache=shared&_fk=1"
    - **Purpose**: This variable allows you to specify your own [database](#database-support) connection string. You can use it to connect the API to your existing database infrastructure. The format depends on the [database](#database-support) type you choose.
    - **Required**: When docker database container or your existing database infrastructure provided.
- **BLOOCK_WEBHOOK_SECRET_KEY** (***REQUIRED***)
    - **Description**: Your [BLOOCK webhook secret key](https://docs.bloock.com/webhooks/overview).
    - **Purpose**: The [webhook secret key](https://docs.bloock.com/webhooks/overview) is used to secure and verify incoming webhook requests. It ensures that webhook data is received from a trusted source and has not been tampered with during transmission.
    - **[Create webhook](https://docs.bloock.com/webhooks/overview)**
- **BLOOCK_PUBLIC_HOST** (***REQUIRED***)
    - **Description**: Should contain the complete URL, including the protocol (`https://`) and domain or host name. It is essential to ensure that the provided URL is accessible and correctly points to you API's public endpoint. 
    - **Purpose**: Is used to specify the public host or URL of this deployed API. Allows other software clients applications (ex: PolygonID wallet) to make HTTP requests and API calls to interact with this service.
- **BLOOCK_LOCAL_PRIVATE_KEY** (***OPTIONAL***)
    - **Description**: Private key associated to your identity. The key pair must be of type [BJJ](https://iden3-docs.readthedocs.io/en/latest/iden3_repos/research/publications/zkproof-standards-workshop-2/baby-jubjub/baby-jubjub.html).
    - **Purpose**: If you want to sign data using your own local private key, you can specify it here. This private key is used for the generation of the issuer's [`did`](https://www.w3.org/TR/did-core/) and cryptographic operations to ensure data integrity and authenticity.
    - **Conflicts**: Conflicts with `BLOOCK_MANAGED_KEY_ID`. You must specify either a local key or a managed key.
- **BLOOCK_LOCAL_PUBLIC_KEY** (***OPTIONAL***)
    - **Description**: Public key associated to your identity.
    - **Purpose**: If you're using your own local public key, you should provide the corresponding public key here.
    - **Conflicts**: Conflicts with `BLOOCK_MANAGED_KEY_ID`. You must specify either a local key or a managed key.
- **BLOOCK_MANAGED_KEY_ID** (***OPTIONAL***)
    - **Description**: Key ID (UUID format) of type [BJJ](https://iden3-docs.readthedocs.io/en/latest/iden3_repos/research/publications/zkproof-standards-workshop-2/baby-jubjub/baby-jubjub.html). 
    - **Purpose**: If you're using your own managed key, you should provide the corresponding key id here. This managed key is used for the generation of the issuer's [`did`](https://www.w3.org/TR/did-core/) and cryptographic operations to ensure data integrity and authenticity. If you want to create a managed [BJJ](https://iden3-docs.readthedocs.io/en/latest/iden3_repos/research/publications/zkproof-standards-workshop-2/baby-jubjub/baby-jubjub.html) key with BLOOCK check the documentation [here](https://docs.bloock.com/keys/features/managed-keys). You can create either using our [SDK's](https://docs.bloock.com/keys/features/managed-keys) or [Dashboard UI](https://dashboard.bloock.com/).
    - **Conflicts**: Conflicts with `BLOOCK_LOCAL_PRIVATE_KEY` and `BLOOCK_LOCAL_PUBLIC_KEY`. You must specify either a local key or a managed key.
- **BLOOCK_API_HOST** (***OPTIONAL***)
    - **Description**: The API host IP address.
    - **Default**: 0.0.0.0
    - **Purpose**: This variable allows you to specify the IP address on which the Identity Bloock Managed API should listen for incoming requests. You can customize it based on your network configuration.
- **BLOOCK_API_PORT** (***OPTIONAL***)
    - **Description**: The API port number.
    - **Default**: 8080
    - **Purpose**: The API listens on this port for incoming HTTP requests. You can adjust it to match your preferred port configuration.
- **BLOOCK_API_DEBUG_MODE** (***OPTIONAL***)
    - **Description**:  Enable or disable debug mode.
    - **Default**: false
    - **Purpose**: When set to true, debug mode provides more detailed log information, which can be useful for troubleshooting and debugging. Set it to false for normal operation.

These configuration variables provide fine-grained control over the behavior of the Identity Bloock Managed API. You can adjust them to match your specific requirements and deployment environment.

### Configuration file

The configuration file should be named `config.yaml`. The service will try to locate this file in the root directory unless the BLOOCK_CONFIG_PATH is defined (i.e. `BLOOCK_CONFIG_PATH="app/conf/"`).

Sample content of `config.yaml`:

```yaml
BLOOCK_API_HOST: "0.0.0.0"
BLOOCK_API_PORT: "8080"
BLOOCK_API_DEBUG_MODE: "false"

BLOOCK_DB_CONNECTION_STRING: "file:bloock?mode=memory&cache=shared&_fk=1"

BLOOCK_API_KEY: ""
BLOOCK_WEBHOOK_SECRET_KEY: ""
BLOOCK_PUBLIC_HOST: ""

BLOOCK_LOCAL_PRIVATE_KEY: ""
BLOOCK_LOCAL_PUBLIC_KEY: ""
BLOOCK_MANAGED_KEY_ID: ""
```

### Database Support

The Identity Bloock Managed API is designed to be flexible when it comes to database integration. It supports three types of relational databases: **MemDB (SQLite)**, **MySQL**, and **Postgres**. The choice of database type depends on your specific requirements and infrastructure.

Here are the supported database types and how to configure them:

- **MySQL**: To connect to a MySQL database, you can use the following connection string format
   ````
   mysql://user:password@tcp(host:port)/database
   ````

Replace `user`, `password`, `host`, `port`, and `database` with your MySQL database credentials and configuration. This format allows you to specify the MySQL database you want to connect to.

- **Postgres**: For PostgreSQL database integration, use the following connection string format:

   ````
   postgresql://user:password@host/database?sslmode=disable
   ````

Similar to MySQL, replace `user`, `password`, `host`, and `database` with your PostgreSQL database details. Additionally, you can set the `sslmode` as needed. The `sslmode=disable` option is used in the example, but you can adjust it according to your PostgreSQL server's SSL requirements.

- **MemDB (SQLite)**: The API also supports in-memory SQLite databases. To use SQLite, you can specify the connection string as follows:

   ````
   file:dbname?mode=memory&cache=shared&_fk=1
   ````

In this format, `dbname` represents the name of your SQLite database. The API will create an in-memory SQLite database with this name.

If you already have an existing database infrastructure and want to use it with the Identity Bloock Managed API, you have the flexibility to provide your custom database connection string.

`Variable: BLOOCK_DB_CONNECTION_STRING`

The API provides a configuration variable called `BLOOCK_DB_CONNECTION_STRING` that allows you to specify your own database connection string independently of the way you run the API. Whether you run the API as a Docker container or as a standalone application, you can always set this variable to point to your existing database server.

---

## Documentation

You can access the following Postman collection where is the specification for the public endpoint of this API.

[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://www.postman.com/bloock/workspace/bloock-api/collection/16945237-1727027f-e3e7-4fe0-969d-afa295eaf2ca)

---

## License

See [LICENSE](LICENSE).

