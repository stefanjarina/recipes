# Recipes SvelteKit SPA (Javascript)

- SvelteKit
- TailwindCSS
- Flowbite

## Docker

### Prepare database

It is adviced to use just scripts to create and setup database, but you can do it manually as well.

#### With script

```bash
# from root of this repository run
just db-up
just db-create
```

#### Manually

- Create network for containers

```bash
docker network create recipes
```

- Create Database container for PostgreSQL

```bash
docker run -d --net recipes -e "POSTGRES_PASSWORD=postgres" -v recipes-db-data:/var/lib/postgresql/data --name recipes-db -p 5432:5432 postgres:latest
```

- Copy SQL script to container

```bash
docker cp ./_db/* recipes-db:/
```

- Execute SQL script to create user and database

```bash
docker exec recipes-db /bin/sh -c 'psql -h localhost -U postgres -a -f /create_db.sql > /tmp/user_create.log 2>&1'
```

### Build image

```bash
# from root of this repository run
cd backendss/aspnet
docker-compose up --build
```

### Build/Run APPLICATION with Docker-Compose

```bash
docker-compose up
docker-compose up --build  # to rebuild
```
