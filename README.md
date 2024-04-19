# Aries - Push Notification System

Aries is a simple and extendible push notification system. It is designed with  simplicity, scalability and efficiency in the mind. You can modify individual components based on your needs.

## Project Structure

The project is organized into several directories, each with a specific purpose:

- **`/cmd`**: Contains the entry points for the services. Each service has its own directory with a `main.go` file.
  - `/service1`
    - `main.go`  # Entry point for service 1
  - `/service2`
    - `main.go`  # Entry point for service 2
- **`/internal`**: Core application logic, organized by service.
  - `/service1`
    - `/app`        # Application business rules (core logic)
    - `/infra`      # Infrastructure (database, network code)
    - `/interfaces` # Interfaces (HTTP handlers, gRPC APIs)
  - `/service2`
    - `/app`        # Core application logic for service 2
    - `/infra`      # Infrastructure for service 2
    - `/interfaces` # Interfaces for service 2
- **`/pkg`**: Shared libraries and utilities used across services.
  - `/commons`     # Shared libraries and utilities
- **`/configs`**: Configuration files and environment specifics.
- **`/scripts`**: Scripts for building, running, and maintenance.
- **`/deployments`**: Dockerfiles and Kubernetes manifests for deployment.
- **`/docs`**: Documentation for the project.
- **`/tests`**: Tests, including integration and e2e tests.

## Getting Started

### Prerequisites

- Go 1.15 or higher
- Docker (for containerization)
- Kubernetes (for orchestration, optional)

### Setting Up Your Local Development Environment

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/my-go-project.git
    cd my-go-project
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up environment variables:
    ```bash
    cp configs/example.env configs/.env
    edit configs/.env  # Modify the environment variables if necessary
    ```

### Building the Application

To build all services, run the following command from the project root:

```bash
./scripts/build.sh

