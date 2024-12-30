# Multi-Language Database Project

This project is designed to demonstrate the integration of three programming languages (Go, C++, and Python) for creating and managing a database application. Each language is used for different aspects of the project to leverage their strengths.

## Table of Contents

- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [Setup](#setup)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

- **Go**: Handles backend server logic, API endpoints, and database interactions.
- **C++**: Manages high-performance components.
- **Python**: Provides scripts for automation, data processing, and analytics.

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
