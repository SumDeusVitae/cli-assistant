<h1 align="center">Server - CLI Assistant</h1>

<br />

## Introduction

The CLI Assistant server is the backend component that provides the core functionality for the CLI Assistant tool. It exposes a **RESTful API** that handles user registration, login, querying, and other key services that the client application interacts with. The server is containerized using Docker for easy deployment and runs on **Google Cloud Platform (GCP)**, making it scalable and efficient.

The server manages user data, handles authentication, and facilitates interactions between the client and a **GPT-based AI** model for answering user questions. The server is designed to handle **CRUD (Create, Read, Update, Delete)** operations securely and efficiently.

## Features

- **RESTful API** architecture for handling HTTP requests
- **User authentication and management** (register, login, whoami, reset)
- **Query processing**: Handling user questions using AI-powered responses
- **Health checks** to monitor the server's status
- **Environment variable configuration** to manage sensitive information like API keys and database URLs

## Architecture

This server is built using **Go** and leverages the **Chi** router to define the RESTful API endpoints. The API is designed around HTTP methods (GET, POST, etc.) to perform CRUD operations and interact with a database, all while providing the necessary authentication and authorization mechanisms.

The server interacts with a **database** for storing user information and provides API endpoints for user registration, login, and querying the AI system.

## API Endpoints

The server provides several key API endpoints that are consumed by the client:

### **User Authentication**
- **POST /v1/register**: Registers a new user.
- **POST /v1/login**: Logs in an existing user.
- **POST /v1/reset**: Resets user data.
- **GET /v1/whoami**: Checks the logged-in user.

### **Querying AI**
- **POST /v1/ask**: Accepts a question from the user and returns a response powered by the AI model (e.g., GPT).

### **Health Check**
- **GET /v1/healthz**: Returns the health status of the server.

## Server Architecture

### RESTful API Design

The server follows the principles of **REST (Representational State Transfer)**, using standard HTTP methods and status codes for communication between the client and the server. Here are the key points about how the server implements RESTful principles:

- **Resources**: Each endpoint (e.g., `/register`, `/login`, `/ask`) represents a specific resource that can be manipulated via HTTP methods.
- **HTTP Methods**: 
  - `POST` is used for creating new resources (e.g., `/register` to create a new user, `/ask` to submit a question).
  - `GET` is used for retrieving data (e.g., `/healthz` to check server health, `/whoami` to check the current user).
  - `PUT` or `PATCH` could be used for updating resources, but in this API, the `POST` and `GET` methods are most frequently used.
- **Stateless**: Each request is independent, meaning the server does not store information about the client's state. Authentication is managed through tokens or other forms of authorization.
- **JSON-based Communication**: All requests and responses are sent in JSON format for easy integration with modern web and mobile applications.

### Key Components in the Code

The code leverages several Go libraries to implement the RESTful API:
- **Chi Router**: Used to define the API routes and handle HTTP requests. It simplifies the routing logic, especially for RESTful services.
- **Go-SQL Database**: The server interacts with an SQL database for storing and retrieving user data. It uses **libsql** for database communication.
- **CORS Middleware**: The server uses the **CORS (Cross-Origin Resource Sharing)** middleware to allow cross-origin requests, enabling the client to interact with the server from different domains.
- **Environment Variables**: The server reads configuration values from environment variables (e.g., database URLs, API keys, etc.) using the **Godotenv** library.

