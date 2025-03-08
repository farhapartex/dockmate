## DOCKMATE

### Project code, Dockerfile & Docker Compose Generator

### Overview

This Golang script generates Dockerfiles and Docker Compose configurations for Python-based frameworks (Django+DRF, FastAPI, Flask). Users can specify their framework, port, database (if needed), Redis (if needed), and Celery (if needed). The script automates project creation, Dockerfile setup, and optional docker-compose configuration.

### Usage

```bash
go run main.go --framework=<django|fastapi|flask> [--db=<postgres|mysql>] [--redis] [--celery]
```

### Features

* Generates a project structure for Django, FastAPI, or Flask.

* Creates a Dockerfile for containerization.

* Optionally generates docker-compose.yaml for databases, Redis, and Celery.

* Sets up virtual environments for Django projects and installs dependencies automatically.

### Example
Run the following command to generate a Django project with PostgreSQL and Redis:

```bash
go run main.go --framework=django
```

### Future Enhancements
* Support for more frameworks (Node.js, Gin, etc.).
* Customizable Dockerfile and Compose templates.
* Advanced service configurations.