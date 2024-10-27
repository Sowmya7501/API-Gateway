# API-Gateway

API Gateway in Go

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Docker](#docker)
- [Kubernetes](#kubernetes)
- [Monitoring](#monitoring)

## Description

The API Gateway is a Go-based application that acts as a reverse proxy to route requests to different backend services. It provides a single entry point for clients and handles request routing.

## Features

- Reverse proxy for routing requests to backend services
- Load balancing
- Monitoring with Prometheus
- Easy deployment with Docker and Kubernetes

## Prerequisites

- Go 1.17 or later
- Docker
- Kubernetes (Minikube or any other Kubernetes cluster)
- kubectl
- Prometheus

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/API-Gateway.git
    ```

2. **Deploy all the yaml files in the folders inside K8s**

    ```sh
    kubectl apply -f xyz.yaml
    ```

## Usage

1. **Access the API Gateway:**

    The API Gateway will be running on `http://api-gateway.example.com`.

2. **Access Prometheus and Grafana application:**

    - Grafana dashboard will be accessible from `http://grafana.example.com`

    - Prometheus
    Use `kubectl port-forward` to access prometheus (if needed to construct promQL query):

    ```sh
    kubectl port-forward svc/prometheus 9090:9090
    ```

    Prometheus Dashboard will be accessible at `http://localhost:9090`.

## API Endpoints

- **Task Manager API:**
  - `POST /task`: Create a new task
    ```json
    # Example json
    {"assignee": "Jane", "assignor": "Dave", "name": "Task1"}
    ```
  - `GET /task`: Retrieve all tasks
  - `GET /task/{name}`: Retrieve a task by name
  - `PUT /task/{name}`: Update a task by name
  - `DELETE /task/{name}`: Delete a task by name

- **User Manager API:**
  - `POST /user`: Create a new user
    ```json
    # Example json
    {"userid": 1, "username": "Jane"}
    ```
  - `GET /user`: Retrieve all users
  - `GET /user/{name}`: Retrieve a user by name
  - `PUT /user/{name}`: Update a user by name
  - `DELETE /user/{name}`: Delete a user by name


1. **Deploy Prometheus:**

    ```sh
    kubectl apply -f monitoring/prometheus-configmap.yaml
    kubectl apply -f monitoring/prometheus-deployment.yaml
    ```

2. **Access Prometheus:**

    Use `kubectl port-forward` to access Prometheus:

    ```sh
    kubectl port-forward svc/prometheus 9090:9090 -n monitoring
    ```

    Prometheus will be accessible at `http://localhost:9090`.