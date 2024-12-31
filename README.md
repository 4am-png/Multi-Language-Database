# Multi-Language-Database

A comprehensive project demonstrating the integration of Go, C++, and Python for creating and managing a database application. Each language is utilized for its strengths: Go for backend server logic, C++ for high-performance components, and Python for automation, data processing, and analytics.

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [Setup](#setup)
  - [General Setup](#general-setup)
  - [Go Server Setup](#go-server-setup)
  - [C++ Components Setup](#c-components-setup)
  - [Python Scripts Setup](#python-scripts-setup)
- [Usage](#usage)
  - [Running the Go Server](#running-the-go-server)
  - [Executing C++ Components](#executing-c-components)
  - [Using Python Scripts](#using-python-scripts)
- [Contributing](#contributing)
- [License](#license)
- [Contact Information](#contact-information)

## Introduction

The Multi-Language-Database project showcases how to leverage the unique strengths of different programming languages to build a robust database application. This project serves as an educational tool for developers looking to understand the interoperability of Go, C++, and Python in a single project.

## Project Structure

The project is divided into three main components, each implemented in a different programming language:

- **Go**: This component handles the backend server logic, including API endpoints and database interactions. Go is chosen for its concurrency capabilities and efficient performance in handling web servers.
- **C++**: This component manages high-performance tasks that require intensive computations. C++ is selected for its execution speed and fine-grained control over system resources.
- **Python**: This component is used for automation, data processing, and analytics. Python is ideal for these tasks due to its extensive libraries and ease of writing scripts.

### Directory Structure

The project is divided into three main components, each implemented in a different programming language:
```
Multi-Language-Database/
├── go-server/
│   ├── main.go
│   ├── go.mod
│   └── go.sum
├── cpp-components/
│   ├── fetch_users.cpp
│   ├── Makefile
│   └── README.md
├── python-scripts/
│   ├── fetch_users.py
│   ├── requirements.txt
│   └── README.md
├── schema.sql
├── LICENSE
├── README.md
└── .gitignore
```


### Detailed Directory Breakdown

- **`go-server/`**: Contains the Go source code for the backend server.
  - `main.go`: The main entry point for the Go server.
  - `go.mod` and `go.sum`: Dependency management files for Go modules.
- **`cpp-components/`**: Contains the C++ source code for high-performance components.
  - `fetch_users.cpp`: A sample C++ program to fetch user data from the database.
  - `Makefile`: A file to automate the build process for C++ components.
  - `README.md`: Instructions for setting up and running the C++ components.
- **`python-scripts/`**: Contains the Python scripts for automation and data processing.
  - `fetch_users.py`: A sample Python script to fetch user data from the database.
  - `requirements.txt`: A list of dependencies required for the Python scripts.
  - `README.md`: Instructions for setting up and running the Python scripts.
- **`schema.sql`**: SQL script to set up the PostgreSQL database schema.
- **`LICENSE`**: License information for the project.
- **`README.md`**: Main README file for the project.
- **`.gitignore`**: Git ignore file specifying which files and directories to ignore in the repository.

## Database Schema

The project uses PostgreSQL as the database management system. The following schema is used:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL, -- Додано для зберігання хешу пароля
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_posts_user_id ON posts(user_id);
CREATE INDEX idx_posts_title ON posts(title);

```

