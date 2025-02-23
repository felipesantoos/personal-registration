
# Personal Registration - Multicontainer Application

This project is a multicontainer Go-based application for personal registration, using Docker and PostgreSQL.

## Prerequisites

Before you begin, ensure you have the following installed on your machine:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Project Structure

- `.docker/`: Contains the Dockerfile and the docker-compose.
- `.env`: contains the environment variables.
- `main.go`: Entry point for the Go application.
- `src/`: Source code for the application, including database connection and logic.

## Setup and Run

1. **Clone the Repository**:

   ```bash
   git clone <https://github.com/Lutero95/personal-registration.git>

   cd personal-registration
   ```

2. **Set Up Environment Variables**:

   Update the `.env` file with the following values:

   ```env
   DB_HOST=database
   DB_PORT=5432(or another port)
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=your_database_name
   ```

3. **Build and Run Containers**:

   Use Docker Compose to build and start the containers:

   ```bash
   docker-compose up --build
   ```

   This command will:

    - Build the API container from the Dockerfile located in `.docker/Dockerfile.api`.
    - Start a PostgreSQL container with the database initialized from `src/connection/init.sql`.

4. **Access the API**:

   The API will be accessible at `http://localhost:8000`.

## Troubleshooting

- **Port Already in Use**: If you encounter a port allocation error, ensure that ports `5432` (PostgreSQL) and `8000` (API) are not being used by other applications.
- **Database Connection Issues**: Verify that the environment variables in `.env` match the database configuration.

## Development

To run the Go application locally without Docker:

1. Set up a local PostgreSQL instance and update the `.env` file accordingly.
2. Run the application using:

   ```bash
   go run main.go
   ```
