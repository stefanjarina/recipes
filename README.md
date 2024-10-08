# The Recipes Project

Showcase of simple Recipes application.

Written in several back-end frameworks and languages as well as in several modern front-end frameworks.

## Types of projects

- **frontend**: frameworks such as svelte, vue, react, blazor etc...
- **backend**: frameworks such as aspnet, gin, rails, phoenix, etc...
- **fullstack**: frameworks such as sveltekit, next, nuxt, astro, solidstart, etc...
- **baas**: config files for BaaS (e.g. supabase, appwrite)

## Required software for tooling (not frameworks themselves)

- [just](https://just.systems/) (task runner, crossplatform make alternative)
- [python 3.10+](https://www.python.org/)
- [PowerShell 7+](https://github.com/PowerShell/PowerShell)
- [Docker](https://www.docker.com/products/docker-desktop/)

## JUST tasks

- create database

```bash
just db-up
```

- scaffold new project

```bash
# Create new project
just new frontend svelte
# Create new project with different port for docker compose
just new frontend svelte -p 3456
just new backend go
```

- remove project

```bash
just remove backend go
# Above is destructive action, confirmation will be needed, you can also pass --yes after just command
```

- prepare database and user

```bash
# below is automatically called by 'just new' task
# this only works for backends
just db-create go
just db-prepare go
just db-seed go
```

- drop database

```bash
# below is automatically called by 'just remove' task
# this only works for backends
just db-drop go
# Above is destructive action, confirmation will be needed, you can also pass --yes after just command
```

- remove docker stuff

```bash
just db-down # removes both container and volume
just db-rm-c # removes only container
just db-rm-v # removes only volume
# Above are all destructive actions, confirmation will be needed, you can also pass --yes after just command
```

- list of available tasks

```bash
$ just
Available recipes:
    db-create NAME            # Create database and user
    db-down                   # Remove postgres docker container and volume
    db-drop NAME              # Drop database
    db-prepare NAME           # Prepare database schema
    db-rm-c                   # Remove postgres docker container
    db-rm-v                   # Remove docker volume
    db-seed NAME              # Seed database
    db-start                  # Start postgres docker container
    db-stop                   # Stop postgres docker container
    db-up                     # Create postgres docker container
    new TYPE NAME PORT='3000' # Create new project
    remove TYPE NAME          # Remove project
```

## Database

### Docker

If not using BaaS platform, then latest PostgreSQL is created and started in docker with default user and password

### Databases

For backend projects database with 'bm_%NAME%' is created with database, user names and password being the same
