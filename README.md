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

---

## Introduction

The **Multi-Language-Database** project demonstrates how to use the unique strengths of Go, C++, and Python to build a robust and efficient database application. It serves as both an educational resource and a practical example of multi-language interoperability in a single project.

---

## Project Structure

The project is divided into three main components:

- **Go**: Backend server for API endpoints and database interaction.
- **C++**: High-performance modules for computationally intensive tasks.
- **Python**: Scripts for automation, data analysis, and additional utilities.

### Directory Structure

```plaintext
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

Detailed Directory Breakdown
	•	go-server/: Contains Go source files for the backend server.
	•	main.go: Entry point for the server.
	•	go.mod/go.sum: Go module files for dependency management.
	•	cpp-components/: C++ source code for high-performance operations.
	•	fetch_users.cpp: Fetches user data from the database.
	•	Makefile: Automates the build process for C++ components.
	•	python-scripts/: Python scripts for data fetching and analytics.
	•	fetch_users.py: Fetches user data in a structured format.
	•	requirements.txt: Dependencies for Python scripts.
	•	schema.sql: SQL file for initializing the PostgreSQL database.
	•	.gitignore: Specifies files/directories to be ignored by Git.
	•	LICENSE: License information for the project.

Database Schema

The project uses PostgreSQL with the following schema:

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL, -- Storing hashed passwords
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

Setup

General Setup
	1.	Clone the repository:

git clone https://github.com/your-username/Multi-Language-Database.git
cd Multi-Language-Database


	2.	Set up the PostgreSQL database using schema.sql:

psql -U your-username -d your-database-name -f schema.sql



Go Server Setup
	1.	Navigate to the go-server directory and install dependencies:

cd go-server
go mod tidy


	2.	Run the server:

go run main.go



C++ Components Setup
	1.	Navigate to the cpp-components directory and build the components:

cd cpp-components
make


	2.	Execute the compiled binary:

./fetch_users.out



Python Scripts Setup
	1.	Navigate to the python-scripts directory and install dependencies:

cd python-scripts
pip install -r requirements.txt


	2.	Run the scripts:

python fetch_users.py

Usage

Running the Go Server
	•	Start the server by running main.go in the go-server directory.
	•	API endpoints will be accessible at http://localhost:8080.

Executing C++ Components
	•	After building, execute the C++ binary files for specific tasks.

Using Python Scripts
	•	Use Python scripts for data analytics or fetching user information.

Contributing

Contributions are welcome! Please follow the standard fork-and-pull request model.
	1.	Fork the repository.
	2.	Create a new branch for your feature or fix.
	3.	Commit your changes and open a pull request.

License

This project is licensed under the MIT License. See the LICENSE file for details.

Contact Information

For any inquiries or support, please contact:
	•	Name: Your Name
	•	Email: your.email@example.com
	•	GitHub: your-username