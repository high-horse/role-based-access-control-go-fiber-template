# Role-Permission-Based Authentication API

This is a Go-based backend API template for role and permission-based access control with authentication. It provides a robust structure for managing users, roles, permissions, and posts, ensuring secure access to resources based on user roles.

---

## Table of Contents
1. [Features](#features)
2. [Project Structure](#project-structure)
3. [Prerequisites](#prerequisites)
4. [Setup](#setup)
5. [Database Migrations](#database-migrations)
6. [Running the Application](#running-the-application)
7. [API Endpoints](#api-endpoints)
8. [Makefile Commands](#makefile-commands)
9. [Contributing](#contributing)
10. [License](#license)

---

## Features
- **Role-Based Access Control (RBAC)**: Users are assigned roles, and roles have specific permissions.
- **Authentication**: JWT-based authentication for secure API access.
- **Database Migrations**: Uses `goose` for managing database schema migrations.
- **Structured Codebase**: Clean and modular project structure for scalability.
- **Logging**: Logs are stored in the `storage/logs` directory for debugging and monitoring.

---

## Project Structure

```bash
.
├── cmd
│ └── main.go
├── db
│ ├── config
│ │ └── config.go
│ ├── migrations
│ │ ├── 20250306161717_create_users_table.sql
│ │ ├── 20250306162104_create_posts_table.sql
│ │ ├── 20250306164140_create_roles_table.sql
│ │ ├── 20250306164152_create_permissions_table.sql
│ │ └── 20250306164202_create_role_has_permissions_table.sql
│ └── pool
│ └──── pool.go  // for database connecction pool
├── go.mod
├── go.sum
├── helpers
│ └── httpResponse.go
├── internal
│ ├── auth
│ │ ├── authHandler.go
│ │ ├── authService.go
│ │ └── authUserService.go
│ ├── post
│ │ ├── postHandler.go
│ │ └── postService.go
│ └── router
│ └── authRouter.go
├── Makefile
├── middleware
│ └── auth.go
├── pkg
│ ├── models
│ │ ├── permissionModel.go
│ │ ├── postModel.go
│ │ ├── register.go
│ │ ├── roleHasPermissionModel.go
│ │ ├── roleModel.go
│ │ └── userModel.go
│ └── utils
│ ├── jwt.go
│ ├── log.go
│ ├── password.go
│ └── validation.go
├── README.md
├── storage
│ ├── app
│ └── logs
│ ├── 2025-03-08.log
│ └── 2025-03-09.log
└── tmp
├── build-errors.log
└── main
```

---

## Prerequisites
- Go (version 1.20 or higher)
- PostgreSQL (or any compatible database)
- Goose (for database migrations)
- Make (optional, but recommended for running commands)

---

## Setup
1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/role-permission-auth.git
   cd role-permission-auth
   ```
   
2. **Install dependencies**:
   ```bash
   go mod download
   ```
   
3. **Set up the database**:
- Create a PostgreSQL database (e.g., attempt2).
- Update the database connection string in the Makefile or wherever it's used:
   ```bash
   postgres://postgres:root@localhost:5432/attempt2?sslmode=disable
   ```

4. **Run migrations:**:
   ```bash
    make goose-up
   ```
   
---

## Database Migrations
This project uses goose for database migrations. Here are the available commands:
- Create a new migration:
   ```bash
   make create-migration TABLE_NAME=your_table_name
   ```
- Apply migrations:
   ```bash
   make goose-up
   ```
   
- Rollback migrations:
   ```bash
   make goose-down
   ```
- Check migration status:
   ```bash
   make goose-status
   ```
---

## Running the Application
- Start the server:
   ```bash
   air
   ```

---

## Makefile Commands
- Create a new migration:
```bash
make create-migration TABLE_NAME=your_table_name
```

- Apply migrations:
```bash
make goose-up
```

- Rollback migrations:
```bash
make goose-down
```

- Check migration status:
```bash
make goose-status
```