import sys
from utilities import call_command, call_command_and_capture

BASE_NAME="recipes"
CONTAINER_NAME=f"{BASE_NAME}-db"
VOLUME_NAME=f"{BASE_NAME}-db-data"
NETWORK_NAME=BASE_NAME
POSTGRES_PASSWORD="postgres"
POSTGRES_VERSION="latest"
POSTGRES_PORT="5432"
POSTGRES_PASSWORD="postgres"

def get_postgres():
    """Function to get the status of the PostgreSQL container"""
    return call_command_and_capture("docker container ls -a -f 'name=" + CONTAINER_NAME + "' --format '{{.Status}}'")

_, network_status, _ = call_command_and_capture("docker network ls --filter 'name=^" + BASE_NAME + "$' --format '{{.Name}}'")

if not network_status:
    print(f"Network '{NETWORK_NAME}' does not exist yet. Creating...")
    call_command(f"docker network create {NETWORK_NAME}")

_, status, _ = get_postgres()

if status:
    print("Container is already running. Skipping container creation...")
else:
    command = (f"docker run -d --net {NETWORK_NAME} -h db -e POSTGRES_PASSWORD={POSTGRES_PASSWORD} "
               f"-v {VOLUME_NAME}:/var/lib/postgresql/data --name {CONTAINER_NAME} "
               f"-p {POSTGRES_PORT}:5432 postgres:{POSTGRES_VERSION}")
    rc, _, _ = call_command_and_capture(command)
    if rc == 0:
        print("Container created")
    else:
        print("Error creating container. Exiting...")
        sys.exit(1)
