# API Gateway

Simple api-gateway in go

## Table of contents
- [Description](#description)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Grafana](#grafana)

## Description
The API Gateway is a Go-based application that acts as a reverse proxy to route requests to different backend services. It provides a single entry point for clients and handles request routing.

## Features
- Reverse proxy for routing requests to backend services
- Load balancing
- Monitoring with Prometheus
- Easy deployment with Docker and Kubernetes

## Prerequisites
- Go 1.23.2 or later
- Docker
- Kubernetes (Minikube or any other Kubernetes cluster)
- kubectl
- Prometheus

## Installation

- **Clone the repository**
```bash
    git clone https://github.com/yourusername/API-Gateway.git
```
- **Deploy all the yaml files in the folders inside K8s**
```bash
    kubectl apply -f xyz.yaml
```

## Usage

### Access API Gateway
- The API Gateway will be running on `http://api-gateway.example.com`.
### API Endpoints
- **Task Manager API:**
```
  - `POST /task`: Create a new task
    
    # Example json
    {"assignee": "Jane", "assignor": "Dave", "name": "Task1"}
   
  - `GET /task`: Retrieve all tasks
  - `GET /task/{name}`: Retrieve a task by name
  - `PUT /task/{name}`: Update a task by name
  - `DELETE /task/{name}`: Delete a task by name
```
- **User Manager API:**
```
  - `POST /user`: Create a new user
 
    # Example json
    {"userid": 1, "username": "Jane"}
   
  - `GET /user`: Retrieve all users
  - `GET /user/{name}`: Retrieve a user by name
  - `PUT /user/{name}`: Update a user by name
  - `DELETE /user/{name}`: Delete a user by name
```

### Monitoring

#### Access Prometheus (optional):
- `kubectl port-forward` to access prometheus (if needed to construct promQL query)
- Prometheus Dashboard will be accessible at `http://localhost:9090`.
```bash
  kubectl port-forward svc/prometheus 9090:9090
```
#### Access Grafana
  The Dashboard will be accessible from `http://grafana.example.com`

## Grafana
- CPU and Memory Usage by pods
<img width="604" alt="grafana_pr_cpu_mem" src="https://github.com/user-attachments/assets/0b0e5dc8-bc6a-47ca-b110-1d717ee12c6f">

- Metrics on APIs
<img width="599" alt="grafana_pr_api_reqres" src="https://github.com/user-attachments/assets/b1636bdb-8632-43bb-a90e-b638426ef0e3">



