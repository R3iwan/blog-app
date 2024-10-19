# Blog Application

This is a simple blog application that uses a PostgreSQL database and a backend service. The services are managed using Docker Compose. A RESTful blog API that allows users to create, read, update, and delete blog posts. Features JWT authentication to restrict access, and PostgreSQL is used for storing post data. Designed with a clear separation between handlers, services, and data models.

## Prerequisites

- Docker
- Docker Compose

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/blog-app.git
    cd blog-app
    ```

2. Create a `.env` file in the root directory and add the following environment variables:
    ```env
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=yourpassword
    POSTGRES_DB=blogdb
    DB_HOST=db
    DB_PORT=5432
    DB_USER=blog_app_user
    DB_PASSWORD=yourpassword
    DB_NAME=blog_app
    JWT_SECRET=yourjwtsecret
    ```

## Usage

1. Start the services using Docker Compose:
    ```sh
    docker-compose up
    ```

2. The backend service will be available at `http://localhost:8080`.

## Services

- **db**: PostgreSQL database service.
- **backend**: Backend service for the blog application.

## Ports

- PostgreSQL: `5432:5432`
- Backend: `8080:8080`

## Dependencies

- The backend service depends on the db service.

## License

This project is licensed under the MIT License.
