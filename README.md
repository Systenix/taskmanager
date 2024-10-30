# TaskManager

The taskmanager project is a task queue system implemented in Go, featuring an API and a worker component. It demonstrates a custom Dependency Injection (DI) system inspired by the Clean Architecture pattern, emphasizing separation of concerns, testability, and scalability.

## Table of Contents

- [Overview](#overview)
- [Disclaimer](#disclaimer)
- [Architecture](#architecture)
- [Custom DI Model](#custom-di-model)
- [Clean Architecture Application](#clean-architecture-application)
- [Deployment](#deployment)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
  - [Running the Services](#running-the-services)
- [Explanation of the Custom DI Model and Clean Architecture Application](#explanation-of-the-custom-di-model-and-clean-architecture-application)
  - [Custom Dependency Injection (DI) Model](#custom-dependency-injection-di-model)
  - [Application of Clean Architecture](#application-of-clean-architecture)

## Overview

The TaskManager system consists of:

- API Service (task_queue_system): Exposes RESTful endpoints for submitting tasks, checking status, and canceling tasks.
- Worker Service (task_worker): Processes tasks from a Redis-backed queue.
- Redis: Serves as the message broker and task storage.

## Disclaimer

Note: This project was developed in 2 days and serves as a playground to test new ideas in a concrete implementation. It is intended for educational purposes and may not be production-ready.

## Architecture

### Custom DI Model

The project implements a custom Dependency Injection system to manage component dependencies without relying on external DI frameworks. The DI system is designed with the following principles:

- Service Container: Acts as a registry for services and their dependencies.
- Factory Pattern: Uses factory interfaces and implementations to construct services and data stores.
- Configurability: Reads configuration from YAML files and validates them using CUE lang, allowing for environment-specific settings.
- Type Safety: Utilizes Go’s type assertions and interfaces to ensure components are correctly wired.

#### Key Components:

- container.Container: Interface defining methods to build use cases and manage services.
- servicecontainer.ServiceContainer: Concrete implementation of the container, holding the service mappings and configurations.
- Factories: Include data service factories and data store factories to create instances of services and data stores based on configuration.
- Use Cases: Business logic implementations that are constructed using the DI system.

### Clean Architecture Application

The project follows the Clean Architecture principles as proposed by Robert C. Martin (Uncle Bob):

- Domain Layer (domain/model): Contains enterprise-wide business rules and entities (e.g., Task).
- Use Cases (usecase): Implements application-specific business rules orchestrating the flow between entities and interfaces.
- Interface Adapters (interfaces): Converts data from the use case layer to formats suitable for frameworks (e.g., HTTP handlers in Gin).
- Frameworks and Drivers (infrastructure): Includes external agencies like databases, web frameworks, and external services (e.g., Redis data services).

## Deployment

### Prerequisites

- Docker and Docker Compose installed on your machine.
- Go (1.17 or later) if you plan to run the services without Docker.

### Setup

1. Clone the Repository:

```bash
git clone https://github.com/Systenix/taskmanager.git
cd taskmanager
```

2. Build and Start Services:

```bash
docker-compose up --build
```

This command builds the Docker images and starts the containers for the API, worker, and Redis.

### Running the Services

- API Service: The API is accessible at http://localhost:8080. The available endpoints are:
    - POST /api/v1/tasks: Submit a new task.
    - GET /api/v1/tasks/:task_id: Get the status of a task.
    - DELETE /api/v1/tasks/:task_id: Cancel a task.
- Worker Service: The worker listens for tasks on the Redis queue and processes them automatically.

### Test the API

#### Submit a New Task

```bash
curl -X POST -H "Content-Type: application/json" -d '{"type":"example","payload":"data"}' http://localhost:8080/api/v1/tasks
```

#### Expected Response:

```json
{"task_id":"<generated-task-id>"}
```

#### Check Task Status

Replace <generated-task-id> with the actual task ID received.

```bash
curl http://localhost:8080/api/v1/tasks/<generated-task-id>
```

#### Expected Response:

```json
{
  "task": {
    "id": "<generated-task-id>",
    "type": "example",
    "payload": "data",
    "status": "completed",
    "created_at": "2024-10-25T23:25:05Z",
    "updated_at": "2024-10-25T23:25:10Z",
    "attempts": 1,
    "max_attempts": 5
  }
}
```


## Explanation of the Custom DI Model and Clean Architecture Application

### Custom Dependency Injection (DI) Model

The custom DI system is designed to manage dependencies explicitly, without relying on external libraries. Here’s how it works:

- Service Registration and Resolution:
  - Container Interface: Defines methods for building use cases and managing services.
  - Service Container Implementation: Holds mappings of service keys to their instances and configurations.
  - Factories: Use factory interfaces to build data services (DataServiceFactory) and data stores (DataStoreFactory) based on configuration.
- Configuration Driven:
  - Configuration Files: Uses YAML files for configuration, validated by CUE lang for correctness.
  - Dynamic Building: Services and data stores are built dynamically at runtime based on the configuration, allowing for flexible and environment-specific setups.
- Type Safety and Abstraction:
  - Interfaces and Type Assertions: Go interfaces and type assertions ensure that components are correctly typed and reduce coupling.
  - Layered Factories: Factories are layered to handle the creation of different types of services and data stores, promoting separation of concerns.

### Application of Clean Architecture

The project applies Clean Architecture principles in the following ways:

- Layered Architecture:
  - Entities (domain/model): Define the core business objects (e.g., Task).
  - Use Cases (usecase): Implement application-specific business rules and orchestrate interactions between entities and interfaces.
  - Interface Adapters (interfaces): Convert use case outputs to HTTP responses, acting as controllers for the web framework.
  - Frameworks and Drivers (infrastructure): Handle external systems like databases and message queues (e.g., Redis).
- Dependency Rule:
  - Inner layers do not depend on outer layers. For example, the use case layer knows nothing about the web framework or the database implementation.
  - Dependencies are inverted using interfaces, allowing for implementations to be injected (e.g., data services).
- Testability and Maintainability:
  - Business logic can be tested without the need for external dependencies.
  - The system is easier to maintain and extend due to the clear separation of concerns.
