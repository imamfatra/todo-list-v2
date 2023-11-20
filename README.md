# To-Do List API

This API service serves as the front-end for a to-do list application, providing the following functionalities:

## Features

1. **User Authentication:**
   - System to create and manage user accounts.
   
2. **To-Do List Management:**
   - Create new to-do lists.
   - Modify existing to-do lists.
   - Delete unwanted to-do lists.

## Setup infrastructure
- Create the todo-network:
    ```bash
    make network
- Start postgres container:
    ```bash
    make postgres
- Create todo_list database:
    ```bash
    make createdb
- Run db migration up 1 version:
    ```bash
    make migrateup
- Run db migration down 1 version:
    ```bash
    make migratedown
## Generate code
- Generate SQL CRUD with sqlc:
    ```bash
    make sqlc
- Run Test
    ```bash
    make test
- Run Server
    ```bash
    make server
## Deploy to Docker-Compose
- Build Docker-Compose
    ```bash
    docker compose up