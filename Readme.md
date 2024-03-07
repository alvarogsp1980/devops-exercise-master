# Store - Production Ready

# Introduction
This repository showcases the evolution of a simple order creation and retrieval application, making it production-ready from a Site Reliability Engineering (SRE) perspective. Starting from a basic setup, the project has been enhanced to include key aspects necessary for a robust, scalable, and maintainable production service.

# Features Implemented
- **tests**
- **logging**
- **monitoring/metrics with prometheus**
- **CI/CD**
- **Specific improvements** (more details here)

# Getting Started

## Prerequisites
- Docker and Docker Compose installed
- Access to Azure DevOps for CI/CD pipeline functionality
- Azure Subscription

## Running the Application (**from your laptop**)
To start the application and its dependencies, run:
```bash
docker-compose up
```

## Interacting with the Application (**from your laptop**)
Once the application is running, you can:
- Create a new order: `curl http://localhost:8080/create`
- Fetch an order: `curl http://localhost:8080/order/YOUR_ORDER_ID`

# CI/CD Pipeline
The CI/CD pipeline in Azure DevOps automates the testing, building, and deployment of the application. The pipeline consists of three main stages:
- **Build and Test**: Compiles the application and runs unit tests.
- **Docker Build and Push**: Builds the Docker image and pushes it to a Docker registry.
- **Deployment**: Deploys the Docker image to Azure App Service for Containers, making the application accessible via a public URL.


# Specific Improvements

To provide a clear and structured overview, I have organized these improvements in a table format, offering a concise summary of each key advancement made in the project:

| File                                            | Improvement Highlights                                                                                                                                                                                                                                                                                                                                                                              |
| ----------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **.env**                                        | Introduction of environment variables file for configuration management.                                                                                                                                                                                                                                                                                                                            |
| **.gitignore**                                  | Introduction of a `.gitignore` template tailored for Go projects to exclude unnecessary files from Git tracking.                                                                                                                                                                                                                                                                                    |
| **ARMTemplate-WebAppContainer/parameters.json** | New file introducing parameters for the ARM template, providing customization options for subscription ID, resource group, app name, and other Azure resource configurations.                                                                                                                                                                                                                       |
| **ARMTemplate-WebAppContainer/template.json**   | Introduction of an ARM template for deploying a Web App Container on Azure, including resource configurations like hosting plan, app service settings, and Docker registry information.                                                                                                                                                                                                             |
| **azure-pipelines.yml**                         | New CI/CD pipeline configuration for Azure DevOps focusing on automation of the Continuous Integration, Docker image creation and pushing, and deployment to Azure App Service.                                                                                                                                                                                                                     |
| **Dockerfile improvements**                     | - Simplified build stages and improved work directory structure<br>- Optimized `COPY` instructions and static binary compilation<br>- Updated base image for runtime and addition of `EXPOSE` instruction<br>- Enhanced comments for documentation and added time zone support                                                                                                                      |
| **docker-compose.yml improvements**             | - Explicit version specification and service-specific port syntax<br>- Addition of `restart` policy and volumes for Redis data persistence<br>- Explicit dependency declaration and optimization of Redis image<br>- Environment variables syntax improvement and removal of unnecessary ports exposure<br>- Explicit environment variables for Redis and utilization of Docker networking features |
| **main.go**                                     | - Introduction of structured logging with `logrus`<br>- Enhanced environment variable handling<br>- Metrics exposure with Prometheus<br>- Structured initialization process<br>- Consistent logging practices<br>- Enhanced code comments<br>- Dependency injection for Redis and HTTP server<br>- Prometheus integration for improved observability                                                |
| **pkg/redis/orders_repo.go**                    | - Documentation and comments for better readability<br>- Logging integration with `logrus`<br>- Improved variable naming and error handling with logging<br>- Structured logging with context<br>- Enhanced function signatures for dependency injection<br>- Consistent error handling                                                                                                             |
| **pkg/server/server.go**                        | - Logging integration and error handling enhancements<br>- Improved documentation and route setup refinement<br>- Dependency injection of logger<br>- Enhanced response writing and structured error logging<br>- Refined order ID extraction                                                                                                                                                       |
| **pkb/tests/integration/integration_test.go**   | - Test server setup and fixed responses for predictability<br>- Execution and recording of HTTP requests<br>- Assertion-based validation and JSON handling<br>- Dynamic route handling                                                                                                                                                                                                              |

This structured summary provides an at-a-glance overview of the advancements made in the project, highlighting the introduction of new features, enhancements in code quality and readability, and the integration of testing and deployment processes.