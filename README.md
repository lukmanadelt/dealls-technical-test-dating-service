# Dealls Technical Test - Dating Service - Sign Up & Login

## Introduction

This project was created to finish the Dealls Technical Test for the Software Engineer candidate at Dealls. The goal of this project is to provide a backend service for the user sign up and login whose functionality can be accessed via REST APIs.

This project is a Go-based application that leverages Clean Architecture principles, the Factory Method design pattern, and PostgreSQL as its database. Docker is used for containerization, and a Makefile is provided for automation. Code quality is ensured with `golangci-lint`.

## Features

### Sign Up

- **Endpoint**: POST http://localhost:8080/dating/v1/users/signup
- **Sample request body**:
  ```
  {
    "email": "example@email.com",
    "password": "example",
    "name": "Example",
    "birth_date": "2024-05-01",
    "gender": "MALE",
    "location": "Indonesia"
  }
  ```
- **Sample response**:
  ```
  "Sign up successful"
  ```

### Login

- **Endpoint**: POST http://localhost:8080/dating/v1/users/login
- **Sample request body**:
  ```
  {
    "email": "example@email.com",
    "password": "example"
  }
  ```
- **Sample response**:
  ```
  {
    "token": "xxx"
  }
  ```

## Architecture

The project follows the Clean Architecture approach, ensuring a clear separation between different layers:

```
/dealls-technical-test-dating-service
├── cmd (This directory contains the entry points of the application)
│   └── app (The main package is within these subdirectories. The main.go file initializes the application and invokes the necessary components to start it)
├── internal (The internal directory encapsulates the core business logic and restricts access to other projects)
│   ├── domain (Contains domain-specific logic and entities)
│   │   ├── entity (Holds the core entities (models) of the application)
│   │   ├── repository (Defines repository interfaces)
│   │   ├── service (Implements domain services containing business logic)
│   │   └── usecase (Defines use cases (application services) that interact with the repositories and entities)
│   ├── infrastructure (Contains implementation details and external dependencies)
│   │   ├── auth (Authentication and authorization-related code)
│   │   ├── config (Configuration-related code)
│   │   ├── database (Database connection and setup)
│   │   ├── log (Logging setup and utilities)
│   │   └── repository (Implementation of the repository interfaces defined in the domain)
│   │   └── server (Server connection and setup)
│   └── interface (Adapters and interfaces for interacting with the outside world)
│       ├── controller (Handles HTTP requests, maps them to use cases, and returns responses)
│       └── routes (Handles web service routes)
└── pkg (The pkg directory is for code that's designed to be ideal place for utility packages and shared code that can be reused)
│   └── constant (Contains constants that doesn't belong to any specific layer of the architecture but is used across the application)
│   └── util (Contains utilities that doesn't belong to any specific layer of the architecture but is used across the application)
└── tests (Contains tests other than unit tests)
    └── integration (Contains setup code for integration tests and the actual integration test cases)
    └── postman (Contains postman collection and environment for postman tests)
```

## Installation

1. **Clone the repository**

   ```sh
   git clone https://github.com/lukmanadelt/dealls-technical-test-dating-service.git
   cd dealls-technical-test-dating-service
   ```

2. **Build the Docker image and run the Docker containers**

   ```sh
   make build
   ```

## Configuration

Configuration is managed via environment variables. You can find the configuration file `.env` in the root directory.

## Usage

1. **Run the application**

   ```sh
   make run
   ```

## Testing

1. **Run unit and integration tests**

   ```sh
   make test
   ```

2. **Lint the code**

   ```sh
   make lint
   ```

## Author

Lukman Adel Taufiqurahman @ May, 2024
