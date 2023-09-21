# Bloock Identity Managed API

This is an API for those who want to create, emit and offer verifiable credentials or VC's (https://www.w3.org/TR/vc-data-model-2.0/), using bloock product's and following the principles of self-sovereign identity and privacy by default. 

---

## Table of Contents

- [Installation](#installation)
    - [Using Docker](#using-docker)
        - [Option 1: Pull and Run the Docker Image](#option-1-pull-and-run-the-docker-image)
        - [Option 2: Use Docker Compose with Database Containers](#option-2-use-docker-compose-with-database-containers)
    - [Standalone Setup](#standalone-setup)
        - [Option 3: Clone the GitHub Repository](#option-3-clone-the-github-repository)
- [Configuration](#configuration)
    - [Variables](#variables)
    - [Configuration File](#configuration-file)
- [Database Support](#database-support)
- [Documentation](#documentation)
- [License](#license)

---

## Installation

You have two primary methods to set up and run the Bloock Identity Managed API: using Docker or deploying it as a standalone application. Each method has its advantages and use cases.

### Using Docker

Docker offers a convenient way to package and distribute the API, along with its required dependencies, in a self-contained environment. It's an excellent choice if you want a quick and hassle-free setup, or if you prefer isolation between your application and the host system.

### Option 1: Pull and Run the Docker Image

This option is straightforward and ideal if you want to get started quickly. Follow these steps:

1. **Pull the Docker Image:**

    - Open your terminal or command prompt.

    - Run the following command to pull the Docker image from [DockerHub](https://hub.docker.com/repository/docker/bloock/managed-api/general):

      ```bash
      docker pull bloock/managed-api
      ```
      
      This command fetches the latest version of the Bloock Identity Managed API image from [DockerHub](https://hub.docker.com/repository/docker/bloock/managed-api/general). We maintain a Docker repository with the latest releases of this repository.

2. **Create a `config.txt` File:**

    - In your project directory, create a `config.txt` file. You can use a text editor of your choice to create this file.

    - This file will contain the configuration for the API, including environment variables. Refer to the [Configuration](#variables) section for a list of environment variables and their descriptions.

    - In the `config.txt` file, define the environment variables you want to configure for the API. Each environment variable should be set in the following format:
      ```toml
      VARIABLE_NAME=VALUE
      ```

    - Here's an example of what your `config.txt` file might look like:

      ```toml
      BLOOCK_DB_CONNECTION_STRING=file:bloock?mode=memory&cache=shared&_fk=1
      BLOOCK_API_KEY=your_api_key
      BLOOCK_WEBHOOK_SECRET_KEY=your_webhook_secret_key
      BLOOCK_CLIENT_ENDPOINT_URL=https://bloock.com/endpoint/to/send/file
      ```

    Note: For the BLOOCK_DB_CONNECTION_STRING environment variable, you have the flexibility to specify your own MySQL or PostgreSQL infrastructure. Clients can provide their connection string for their database infrastructure. See the [Database](#database) section for available connections. 
    
    - Now, here's how Docker and Viper work together to use these environment variables
    
3. **Run the Docker Image with Environment Variables:**

    - Run the following command to start the Bloock Identity Managed API container while passing the `config.txt` file as an environment variable source:

     ```bash
     docker run --env-file config.txt -p 8080:8080 bloock/managed-api
     ```

    - This command maps the `config.txt` file into the container, ensuring that the API reads its configuration from the file. Viper automatically read these environment variables and make them accessible to the code.

4. **Access the API:**

    - After running the Docker image, the Bloock Identity Managed API will be accessible at `http://localhost:8080`.

    - You can now make API requests to interact with the service.

By following these steps, you can quickly deploy the Bloock Identity Managed API as a Docker container with your customized configuration.

### Option 2: Use Docker Compose with Database Containers

If you need a more complex setup, such as using a specific database like MySQL, Postgres or MemDB, Docker Compose is your choice. Follow these steps:

1. **Choose the Docker Compose File:**

    - In our [repository](https://github.com/bloock/bloock-identity-managed-api), you will find Docker Compose files for different database types:
   ````
   docker-compose-mysql.yaml for MySQL
   docker-compose-postgres.yaml for PostgreSQL
   docker-compose.yaml for MemDB (SQLite)
   ````

2. **Copy the Chosen Docker Compose File:**
  
    - Choose the Docker Compose file that corresponds to the database type you want to use. For example, if you prefer MySQL, copy docker-compose-mysql.yaml.

3. **Configure Environment Variables:**
    
    - Open the Docker Compose file in a text editor. Inside the file, locate the environment section for the api service. Here, you can specify environment variables that configure the API.

    - Refer to the [Configuration](#variables) section for a list of environment variables and their descriptions.

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
      BLOOCK_CLIENT_ENDPOINT_URL= "https://bloock.com/endpoint/to/send/file"
      ```

5. **Run Docker Compose:**

    - Open your terminal, navigate to the directory where you saved the Docker Compose file, and run the following command:

    ```bash
     docker-compose -f docker-compose-mysql.yaml up
     ```

   Replace docker-compose-mysql.yaml with the name of the Docker Compose file you selected.

6. **Access the API:**

    - After running the Docker Compose command, the Bloock Identity Managed API will be accessible at http://localhost:8080. You can make API requests to interact with the service.

By following these steps, you can quickly set up the Bloock Identity Managed API with your chosen database type using the provided Docker Compose files.

### Standalone Setup

Running the API as a standalone application provides more control and flexibility, allowing you to customize and integrate it into your specific environment. Choose this option if you have specific requirements or if you want to modify the API's source code.

### Option 3: Clone the GitHub Repository

You can also run this service as a common Golang binary if you need it.

#### Standalone Requirements

    - Makefile toolchain
    - Unix-based operating system (e.g. Debian, Arch, Mac OS X)
    - [Go](https://go.dev/) 1.20


To deploy the API as a standalone application, follow these steps:

1. **Clone the Repository:**

    - Open your terminal and navigate to the directory where you want to clone the repository.

    - Run the following command to clone the repository:

    ```bash
     git clone https://github.com/bloock/managed-api.git
     ```

2. **Navigate to the Repository:**

    - Change your current directory to the cloned repository:

    ```bash
     cd managed-api
     ```

3. **Set Up Configuration:**

    - Inside the repository, you'll find a config.yaml file.

    - Open config.yaml in a text editor and configure the environment variables as needed, following the format described in the [Configuration](#variables) section. For example:

    ```yaml
      BLOOCK_DB_CONNECTION_STRING: "file:bloock?mode=memory&cache=shared&_fk=1"
      BLOOCK_API_KEY: "your_api_key"
      BLOOCK_WEBHOOK_SECRET_KEY: "your_webhook_secret_key"
      BLOOCK_CLIENT_ENDPOINT_URL= "https://bloock.com/endpoint/to/send/file"
      ```

4. **Run the Application:** 

    - To run the application, execute the following command:

    ```bash
     go run cmd/main.go
     ```

   This command will start the Bloock Identity Managed API as a standalone application, and it will use the configuration provided in the config.yaml file.


5. **Access the API:**

    - After running the application, the Bloock Identity Managed API will be accessible at http://localhost:8080. You can make API requests to interact with the service.

---

### Configuration

The Bloock Identity Managed API leverages Viper, a powerful configuration management library, currently supporting environment variables and a YAML configuration file.

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

### Database Support

The Bloock Identity Managed API is designed to be flexible when it comes to database integration. It supports three types of relational databases: MemDB (SQLite), MySQL, and Postgres. The choice of database type depends on your specific requirements and infrastructure.

Here are the supported database types and how to configure them:

   - **MySQL**: To connect to a MySQL database, you can use the following connection string format
   ````
   mysql://user:password@tcp(host:port)/database
   ````

   Replace 'user', 'password', 'host', 'port', and 'database' with your MySQL database credentials and configuration. This format allows you to specify the MySQL database you want to connect to.

   - **Postgres**: For PostgreSQL database integration, use the following connection string format:

   ````
   postgres://user:password@tcp(host:port)/database?sslmode=disable
   ````

   Similar to MySQL, replace 'user', 'password', 'host', 'port', and 'database' with your PostgreSQL database details. Additionally, you can set the 'sslmode' as needed. The 'sslmode=disable' option is used in the example, but you can adjust it according to your PostgreSQL server's SSL requirements.

   - **MemDB (SQLite)**: The API also supports in-memory SQLite databases. To use SQLite, you can specify the connection string as follows:

   ````
   file:dbname?mode=memory&cache=shared&_fk=1
   ````

   In this format, 'dbname' represents the name of your SQLite database. The API will create an in-memory SQLite database with this name.

If you already have an existing database infrastructure and want to use it with the Bloock Identity Managed API, you have the flexibility to provide your custom database connection string.

**Variable: BLOOCK_DB_CONNECTION_STRING**

The API provides a configuration variable called **'BLOOCK_DB_CONNECTION_STRING'** that allows you to specify your own database connection string independently of the way you run the API. Whether you run the API as a Docker container or as a standalone application, you can always set this variable to point to your existing database server.

---

## Documentation

You can access the following Postman collection where is the specification for the public endpoint of this API.

[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://www.postman.com/bloock/workspace/bloock-api/documentation/16945237-1727027f-e3e7-4fe0-969d-afa295eaf2ca)

---

## License

See [LICENSE](LICENSE.md).

