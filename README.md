
# LogSphere

LogSphere is a comprehensive work log management system designed to facilitate easy logging of work activities, management of user roles, and reviews of work logs with built-in authentication and admin privileges.

## Features

- User authentication (login and registration)
- Role-based access control
- Admin dashboard for user and role management
- Work log submission and review

## Prerequisites

Before you begin, ensure you have met the following requirements:
- Docker and Docker Compose installed on your machine
- Basic knowledge of Docker container management

## Getting Started

These instructions will get your copy of the project up and running on your local machine for development and testing purposes.

### Clone the Repository

Start by cloning the repository to your local machine:

\```bash
git clone https://github.com/yourusername/LogSphere.git
cd LogSphere
\```

### Environment Setup

(Optional) If your project requires environment variables (e.g., for database credentials), you need to create a .env file at the root of the project and define the necessary variables as key-value pairs:

\```env
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=logsphere
DB_HOST=db
\```

### Running the Application

To build and start the application, run the following command in the root directory of the project:

\```bash
docker-compose up --build
\```

This command will start all services defined in `docker-compose.yml`, including the Go application and the PostgreSQL database.

### Accessing the Application

Once the application is running, you can access it in your web browser:

- Go application: [http://localhost:8080](http://localhost:8080)

### Stopping the Application

To stop and remove the containers, use the following command:

\```bash
docker-compose down
\```
