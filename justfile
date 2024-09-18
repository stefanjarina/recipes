set windows-shell := ["pwsh.exe", "-NoLogo", "-Command"]

container_name := "recipes-db"
volume_name := "recipes-db-data"

[private]
default:
  @just --list --justfile {{justfile()}}

# Create new project
@new TYPE NAME PORT='3000': && (db-create NAME) (db-prepare NAME) (db-seed NAME)
  python3 ./_tools/scripts/init_new_project.py {{TYPE}} {{NAME}} -p {{PORT}}

# Remove project
[unix, confirm('This will remove the project. Are you sure?')]
@remove TYPE NAME: (db-drop NAME)
  if [ -d "./{{ TYPE }}s/{{NAME}}" ]; then rm -rf "./{{ TYPE }}s/{{NAME}}"; fi

# Remove project
[windows, confirm('This will remove the project. Are you sure?')]
@remove TYPE NAME: (db-drop NAME)
  if (Test-Path "./{{ TYPE }}s/{{NAME}}") { Remove-Item -Recurse -Force -Path "./{{ TYPE }}s/{{NAME}}" }

# Create postgres docker container
@db-up:
  python3 ./_tools/scripts/create_postgres.py

# Stop postgres docker container
@db-stop:
  docker container stop {{ container_name }}

# Start postgres docker container
@db-start:
  docker container start {{ container_name }}

# Remove postgres docker container and volume
[confirm('This will remove both docker container and volume. Are you sure?')]
@db-down:
  docker container rm {{ container_name }} --force
  docker volume rm {{ volume_name }}

# Remove postgres docker container
[confirm('This will remove docker container. Are you sure?')]
@db-rm-c:
  docker container rm {{ container_name }} --force

# Remove docker volume
[confirm('This will remove docker volume. Are you sure?')]
@db-rm-v:
  docker volume rm {{ volume_name }}

# Create database and user
@db-create NAME:
  python3 ./_tools/scripts/run_sql_script.py -n {{NAME}} -s "create_db.sql"

# Prepare database schema
@db-prepare NAME:
  python3 ./_tools/scripts/run_sql_script.py -n {{NAME}} -s "prepare_db.sql" -d

# Seed database
@db-seed NAME:
  python3 ./_tools/scripts/run_sql_script.py -n {{NAME}} -s "seed_db_base.sql" -d
  python3 ./_tools/scripts/run_sql_script.py -n {{NAME}} -s "seed_db_recipe.sql" -d

# Drop database
[confirm('This will remove database and user inside postgres. Are you sure?')]
@db-drop NAME:
  python3 ./_tools/scripts/run_sql_script.py -n {{ NAME }} -s "drop_db.sql"
