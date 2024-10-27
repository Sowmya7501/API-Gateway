import requests
import json
import logging
import argparse
import time

# Configure logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# Base URLs for the APIs
BASE_URL = "http://api-gateway.example.com"

# Function to send a POST request to create a task
def create_task(task):
    url = f"{BASE_URL}/task"
    logging.info(f"Sending POST request to {url} with data: {task}")
    response = requests.post(url, json=task)
    logging.info(f"Create Task Response: {response.status_code}, {response.json()}")

# Function to send a GET request to retrieve all tasks
def get_all_tasks():
    url = f"{BASE_URL}/task"
    logging.info(f"Sending GET request to {url}")
    response = requests.get(url)
    logging.info(f"Get All Tasks Response: {response.status_code}, {response.json()}")

# Function to send a GET request to retrieve a task by name
def get_task_by_name(name):
    url = f"{BASE_URL}/task/{name}"
    logging.info(f"Sending GET request to {url}")
    response = requests.get(url)
    logging.info(f"Get Task by Name Response: {response.status_code}, {response.json()}")

# Function to send a DELETE request to delete a task by name
def delete_task_by_name(name):
    url = f"{BASE_URL}/task/{name}"
    logging.info(f"Sending DELETE request to {url}")
    response = requests.delete(url)
    logging.info(f"Delete Task by Name Response: {response.status_code}, {response.json()}")

# Function to send a PUT request to update a task by name
def update_task_by_name(name, updated_task):
    url = f"{BASE_URL}/task/{name}"
    logging.info(f"Sending PUT request to {url} with data: {updated_task}")
    response = requests.put(url, json=updated_task)
    logging.info(f"Update Task by Name Response: {response.status_code}, {response.json()}")

# Function to send a POST request to create a user
def create_user(user):
    url = f"{BASE_URL}/user"
    logging.info(f"Sending POST request to {url} with data: {user}")
    response = requests.post(url, json=user)
    logging.info(f"Create User Response: {response.status_code}, {response.json()}")

# Function to send a GET request to retrieve all users
def get_all_users():
    url = f"{BASE_URL}/user"
    logging.info(f"Sending GET request to {url}")
    response = requests.get(url)
    logging.info(f"Get All Users Response: {response.status_code}, {response.json()}")

# Function to send a GET request to retrieve a user by name
def get_user_by_name(name):
    url = f"{BASE_URL}/user/{name}"
    logging.info(f"Sending GET request to {url}")
    response = requests.get(url)
    logging.info(f"Get User by Name Response: {response.status_code}, {response.json()}")

# Function to send a DELETE request to delete a user by name
def delete_user_by_name(name):
    url = f"{BASE_URL}/user/{name}"
    logging.info(f"Sending DELETE request to {url}")
    response = requests.delete(url)
    logging.info(f"Delete User by Name Response: {response.status_code}, {response.json()}")

# Function to send a PUT request to update a user by name
def update_user_by_name(name, updated_user):
    url = f"{BASE_URL}/user/{name}"
    logging.info(f"Sending PUT request to {url} with data: {updated_user}")
    response = requests.put(url, json=updated_user)
    logging.info(f"Update User by Name Response: {response.status_code}, {response.json()}")

# Function to perform the API operations
def perform_operations():
    # Task operations
    new_task = {"assignee": "John", "assignor": "Doe", "name": "Task1"}
    create_task(new_task)
    get_all_tasks()
    get_task_by_name("Task1")
    update_task_by_name("Task1", {"assignee": "Jane", "assignor": "Doe", "name": "Task1"})
    delete_task_by_name("Task1")

    # User operations
    new_user = {"username": "john_doe"}
    create_user(new_user)
    get_all_users()
    get_user_by_name("john_doe")
    update_user_by_name("john_doe", {"userid": 1, "username": "john_doe_updated"})
    delete_user_by_name("john_doe")

# Main function to handle command-line arguments and run the operations in a loop
def main():
    parser = argparse.ArgumentParser(description="Run API operations in a loop.")
    parser.add_argument("--num", type=int, default=1, help="Number of times to run the operations")
    args = parser.parse_args()

    for i in range(args.num):
        logging.info(f"Running iteration {i + 1}/{args.num}")
        perform_operations()
        time.sleep(1)  # Optional: Add a delay between iterations

if __name__ == "__main__":
    main()