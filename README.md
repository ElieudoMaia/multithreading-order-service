Here's the updated README with MySQL included:

---

# Multithreading Order Service

This project is a **Golang-based application** designed to process orders using multithreading with RabbitMQ. It includes a RabbitMQ consumer, an HTTP server, and integrates MySQL for data storage. The project is containerized using Docker and includes monitoring with Prometheus and Grafana.

## Features

- **Multithreaded RabbitMQ Consumer:** Efficiently processes messages from RabbitMQ using multiple threads.
- **HTTP Server:** Provides an API for managing and tracking orders.
- **MySQL Database:** Stores order data and other relevant information.
- **Dockerized Environment:** Uses Docker to set up and run the application in a consistent environment.
- **Monitoring:** Integrates Prometheus and Grafana for real-time monitoring and metrics visualization.

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/elieudomaia/multithreading-order-service.git
cd multithreading-order-service
```

### Build and Run with Docker

To build and run the application using Docker:

```bash
docker-compose up -d
```

This command will:

- Build the Docker images for the application.
- Start the RabbitMQ consumer, HTTP server, and MySQL database.
- Set up Prometheus and Grafana for monitoring.

### Accessing the Application

- **HTTP Server:** The server will be accessible at `http://localhost:8181`.
- **MySQL Database:** Accessible within the Docker network at `mysql:3306`. You can connect using a MySQL client.
- **Prometheus:** Metrics will be available at `http://localhost:9090`.
- **Grafana:** Monitoring dashboards can be accessed at `http://localhost:3000` (default login is `admin/admin`).

### Configuration

You can modify the environment variables in the `docker-compose.yml` file to adjust the configuration of RabbitMQ, MySQL, Prometheus, and other services.
