# BLOOCK Identity Managed API

This is an API for those who want to create, emit and offer verifiable credentials or VC's (https://www.w3.org/TR/vc-data-model-2.0/), using bloock product's and following the principles of self-sovereign identity and privacy by default. 

---

## Table of Contents

- [Installation](#installation)
    - [Docker Setup Guide](#docker-setup-guide)
        - [Option 1: Pull and Run the Docker Image](#option-1-pull-and-run-the-docker-image)
        - [Option 2: Use Docker Compose with Database Containers](#option-2-use-docker-compose-with-database-containers)
- [Configuration](#configuration)
    - [Variables](#variables)
- [Database Support](#database-support)
- [Documentation](#documentation)
- [License](#license)

---

## Installation

You have one primary method to set up and run the Identity BLOOCK Managed API:

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
      docker pull bloock/identity-managed-api
      ```

      This command fetches the latest version of the Identity BLOOCK Managed API image from [DockerHub](https://hub.docker.com/r/bloock/identity-managed-api). We maintain a Docker repository with the latest releases of this repository.


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
      BLOOCK_BLOOCK_API_KEY=your_api_key
      BLOOCK_BLOOCK_WEBHOOK_SECRET_KEY=your_webhook_secret_key
      ```

      > **NOTE:** For the **BLOOCK_DB_CONNECTION_STRING** environment variable, you have the flexibility to specify your own MySQL or PostgreSQL infrastructure. Clients can provide their connection string for their database infrastructure. See the [Database](#database-support) section for available connections.

3. **Run the Docker Image with Environment Variables:**

    - Run the following command to start the Identity BLOOCK Managed API container while passing the `.env` file as an environment variable source:

     ```bash
     docker run --env-file .env -p 8080:8080 bloock/identity-managed-api
     ```

    - This command maps the `.env` file into the container, ensuring that the API reads its configuration from the file. Viper automatically read these environment variables and make them accessible to the code.

    - By default, the above command runs the Docker container in the foreground, displaying API logs and output in your terminal. You can interact with the API while it's running.

   3.1. **Running Docker in the Background (Daemon Mode)**

    - Append the `-d` flag to the docker run command as follows:

    ```bash
    docker run -d --env-file config.txt -p 8080:8080 bloock/identity-managed-api
    ```

   The `-d` flag tells Docker to run the container as a background process. You can continue using your terminal for other tasks while the Identity BLOOCK Managed API container runs silently in the background.


4. **Access the API:**

    - After running the Docker image, the Identity BLOOCK Managed API will be accessible at `http://localhost:8080`.

    - You can now make API requests to interact with the service.

By following these steps, you can quickly deploy the Identity BLOOCK Managed API as a Docker container with your customized configuration.

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
      environment:
        BLOOCK_DB_CONNECTION_STRING: "file:bloock?mode=memory&cache=shared&_fk=1"
        BLOOCK_BLOOCK_API_KEY: "your_api_key"
        BLOOCK_BLOOCK_WEBHOOK_SECRET_KEY: "your_webhook_secret_key"
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

   The `-d` flag tells Docker to run the container as a background process. You can continue using your terminal for other tasks while the Identity BLOOCK Managed API container runs silently in the background.


6. **Access the API:**

    - After running the Docker Compose command, the Identity BLOOCK Managed API will be accessible at http://localhost:8080. You can make API requests to interact with the service.

By following these steps, you can quickly set up the Identity BLOOCK Managed API with your chosen database type using the provided Docker Compose files.

---

## Configuration

The Identity BLOOCK Managed API leverages Viper, a powerful configuration management library, currently supporting environment variables and a YAML configuration file.

### Variables

Here are the configuration variables used by the Identity BLOOCK Managed API.

Basic configuration:

- **BLOOCK_BLOOCK_API_KEY** (**REQUIRED**)
    - **Description**: Your unique [BLOOCK API key](https://docs.bloock.com/libraries/authentication/create-an-api-key).
    - **Purpose**: This [API key](https://docs.bloock.com/libraries/authentication/create-an-api-key) is required for authentication and authorization when interacting with the BLOOCK Identity Managed API. It allows you to securely access and use the API's features.
    - **[Create API Key](https://docs.bloock.com/libraries/authentication/create-an-api-key)**
    - **Example**: no9rLf9dOMjXGvXQX3I96a39qYFoZknGd6YHtY3x1VPelr6M-TmTLpAF-fm1k9Zi
- **BLOOCK_BLOOCK_WEBHOOK_SECRET_KEY** (***REQUIRED***)
    - **Description**: Your [BLOOCK webhook secret key](https://docs.bloock.com/webhooks/overview).
    - **Purpose**: The [webhook secret key](https://docs.bloock.com/webhooks/overview) is used to secure and verify incoming webhook requests. It ensures that webhook data is received from a trusted source and has not been tampered with during transmission.
    - **[Create webhook](https://docs.bloock.com/webhooks/overview)**
    - **Example**: ew1b2d5qf7WeUOPy1u1CW6FXro6j5plS
- **BLOOCK_API_PUBLIC_HOST** (***REQUIRED***)
    - **Description**: Should contain the complete URL, including the protocol (`https://`) and domain or host name. It is essential to ensure that the provided URL is accessible and correctly points to you API's public endpoint.
    - **Purpose**: Is used to specify the public host or URL of this deployed API. Allows other software clients applications (ex: PolygonID wallet) to make HTTP requests and API calls to interact with this service.
    - **Example**: https://1039-94-132-61-84.ngrok-free.app
- **BLOOCK_DB_CONNECTION_STRING** (***OPTIONAL***)
    - **Description**: Your custom database connection URL.
    - **Default**: "file:bloock?mode=memory&cache=shared&_fk=1"
    - **Purpose**: This variable allows you to specify your own [database](#database-support) connection string. You can use it to connect the API to your existing database infrastructure. The format depends on the [database](#database-support) type you choose.
    - **Required**: When docker database container or your existing database infrastructure provided.
- **BLOOCK_AUTH_SECRET** (***OPTIONAL***)
    - **Description**: If you want to add control in your API calls with Bearer Token you can add a secret here. A Bearer token is a type of token used for authentication and authorization and is used in web applications and APIs to hold user credentials and indicate authorization for requests and access.
    - **Purpose**: The idea is that you can set a secret, which will be the same that you will have to pass in the headers of your requests in order to validate yourself.
    - **Example**: 0tEStdP(dg5=VU4iX4+7}e((HVd^ShVm
- **BLOOCK_API_HOST** (***OPTIONAL***)
    - **Description**: The API host IP address.
    - **Default**: 0.0.0.0
    - **Purpose**: This variable allows you to specify the IP address on which the Identity BLOOCK Managed API should listen for incoming requests. You can customize it based on your network configuration.
- **BLOOCK_API_PORT** (***OPTIONAL***)
    - **Description**: The API port number.
    - **Default**: 8080
    - **Purpose**: The API listens on this port for incoming HTTP requests. You can adjust it to match your preferred port configuration.
- **BLOOCK_API_DEBUG_MODE** (***OPTIONAL***)
    - **Description**:  Enable or disable debug mode.
    - **Default**: false
    - **Purpose**: When set to true, debug mode provides more detailed log information, which can be useful for troubleshooting and debugging. Set it to false for normal operation.
    
Advanced configuration. Please only edit these variables if you are familiar with the BLOOCK digital identity protocol.

In case you want to deploy an issuer with local keys (i.e. not managed by BLOOCK services) you must set the following variables in order to create your issuer together with the API deployment:

- **BLOOCK_ISSUER_KEY_KEY** (***REQUIRED***)
    - **Description**: Represents a private key of type [Baby JubJub](https://docs.iden3.io/getting-started/babyjubjub/).
    - **Purpose**: This private key will be used to create your issuer. In addition, for all operations where the issuer's signature is required, the same will be used to perform such operations.
    - **Required**: If you want to use your issuer locally, you only need to omit the `issuer_key` query when executing your requests.
    - **Example**: bf5e13dd8d9f784aee781b4de7836caa3499168514553eaa3d892911ad3c115t
- **BLOOCK_ISSUER_PUBLISH_INTERVAL** (***REQUIRED***)
    - **Description**: This is the frequency at which the state of your local issuer will be transacted to blockchain.
    - **Purpose**: This variable will allow you to choose the time interval you want to spend to execute the transaction and the economic cost you want to assume.
    - **Options**: You must pass one of the following integers: 1, 5, 15 or 60. Representing every 1 minute, 5 minutes, 15 minutes or 60 minutes.
    - **Example**: 1
- **BLOOCK_ISSUER_NAME** (***OPTIONAL***)
    - **Description**: The issuer name.
    - **Purpose**: Simply to identify your issuer by name.
    - **Example**: Test Issuer Name
- **BLOOCK_ISSUER_DESCRIPTION** (***OPTIONAL***)
    - **Description**: The issuer description.
    - **Purpose**: Simply to add a description of you issuer.
    - **Example**: this is my first issuer creation
- **BLOOCK_ISSUER_IMAGE** (***OPTIONAL***)
    - **Description**: You can set up an image for you issuer. You will see that image issuer on your [BLOOCK management dashboard](https://dashboard.bloock.com/login).
    - **Purpose**: You will have to pass an image in base64url to be able to decode it later.
    - **Example**: iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAEAAAAAApiSv5AAAHQklEQVR4nOydQW7luA4Afz5y_yv3rCZ-C4IiTdrpQVWtGm1bUpKCQEgi9f3nz_8E...

If you want to perform verification, you can edit parameters of your verification process:

- **BLOOCK_VERIFICATION_EXPIRATION** (***OPTIONAL***)
    - **Description**: You can edit how long a verification is valid for (by specifying the number in minutes), i.e. when you start a verification process by default that session id that is created has an expiration after 60 minutes. You can add more time or reduce it. More features will be added soon.
    - **Default**: 60 minutes.
    - **Example**: 120

**Only for BLOOCK development environments:**

- **BLOOCK_API_POLYGON_TEST_ENABLED** (***OPTIONAL***)
    - **Description**: Boolean to activates the development environment.
    - **Purpose**: Basically it changes all the variables concerning the Polygon ID identity protocol by referencing the Mumbai network.
    - **Example**: false

These configuration variables provide fine-grained control over the behavior of the Identity BLOOCK Managed API. You can adjust them to match your specific requirements and deployment environment.

### Database Support

The Identity BLOOCK Managed API is designed to be flexible when it comes to database integration. It supports three types of relational databases: **MemDB (SQLite)**, **MySQL**, and **Postgres**. The choice of database type depends on your specific requirements and infrastructure.

Here are the supported database types and how to configure them:

- **MySQL**: To connect to a MySQL database, you can use the following connection string format
   ````
   mysql://user:password@tcp(host:port)/database
   ````

Replace `user`, `password`, `host`, `port`, and `database` with your MySQL database credentials and configuration. This format allows you to specify the MySQL database you want to connect to.

- **Postgres**: For PostgreSQL database integration, use the following connection string format:

   ````
   postgres://user:password@host/database?sslmode=disable
   ````

Similar to MySQL, replace `user`, `password`, `host`, and `database` with your PostgreSQL database details. Additionally, you can set the `sslmode` as needed. The `sslmode=disable` option is used in the example, but you can adjust it according to your PostgreSQL server's SSL requirements.

- **MemDB (SQLite)**: The API also supports in-memory SQLite databases. To use SQLite, you can specify the connection string as follows:

   ````
   file:dbname?mode=memory&cache=shared&_fk=1
   ````

In this format, `dbname` represents the name of your SQLite database. The API will create an in-memory SQLite database with this name.

If you already have an existing database infrastructure and want to use it with the Identity BLOOCK Managed API, you have the flexibility to provide your custom database connection string.

`Variable: BLOOCK_DB_CONNECTION_STRING`

The API provides a configuration variable called `BLOOCK_DB_CONNECTION_STRING` that allows you to specify your own database connection string independently of the way you run the API. Whether you run the API as a Docker container or as a standalone application, you can always set this variable to point to your existing database server.

---

## Documentation

You can access the following Postman collection where is the specification for the public endpoint of this API.

[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://www.postman.com/bloock/workspace/bloock-api/collection/16945237-1727027f-e3e7-4fe0-969d-afa295eaf2ca)

---

## License

See [LICENSE](LICENSE).

