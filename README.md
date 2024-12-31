# Multi-Language-Database

A comprehensive project demonstrating the integration of Go, C++, and Python for creating and managing a database application. Each language is utilized for its strengths: Go for backend server logic, C++ for high-performance components, and Python for automation, data processing, and analytics.

## Table of Contents

- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [Setup](#setup)
  - [Go Server Setup](#go-server-setup)
  - [C++ Components Setup](#c++-components-setup)
  - [Python Scripts Setup](#python-scripts-setup)
- [Usage](#usage)
  - [Running the Go Server](#running-the-go-server)
  - [Executing C++ Components](#executing-c++-components)
  - [Using Python Scripts](#using-python-scripts)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

- **Go**: Handles backend server logic, API endpoints, and database interactions.
- **C++**: Manages high-performance components.
- **Python**: Provides scripts for automation, data processing, and analytics.

### Directories

- `go-server/`: Contains Go server source code.
- `cpp-components/`: Contains C++ components.
- `python-scripts/`: Contains Python scripts.
- `schema.sql`: SQL file with the database schema.

## Database Schema

The project uses PostgreSQL as the database management system. The following schema is used:

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_posts_user_id ON posts(user_id);
