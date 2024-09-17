import argparse
import os
import sys
from time import sleep
from utilities import call_command, call_command_and_capture, get_strings

CONTAINER_NAME="recipes-db"

parser = argparse.ArgumentParser(description='Script to create a database for a specified backend.')
parser.add_argument("-n", "--name", dest="name", type=str, help="Name of the backend")
parser.add_argument("-s", "--script", dest="script", type=str, help="Name of the SQL script")
parser.add_argument("-d", "--database", dest="database", action='store_true', help="Run against database")

if len(sys.argv) == 1:
    parser.print_help()
    sys.exit(1)

args = parser.parse_args()

root_path = os.path.join(os.path.dirname(__file__), '../../')
backend_path = os.path.join(root_path, 'backends', args.name)

text_start, text_end = get_strings(args.script)

if not os.path.exists(backend_path):
    print(f"'{args.name}' backend does not exist. Skipping...")
    sys.exit(0)

sql_file = os.path.abspath(os.path.join(backend_path, '_db', args.script))

print(f"{text_start.capitalize()} database using '{sql_file}'")

call_command(f"docker cp {sql_file} {CONTAINER_NAME}:/{args.script}")

if args.script == "create_db.sql":
    sleep(5)

command = f"docker exec {CONTAINER_NAME} /bin/sh -c 'psql -h localhost" 
if args.database:
    command = command + f" -U recipes_{args.name}"
    command = command + f" -d recipes_{args.name}"
else:
    command = command + " -U postgres"

command = command + f" -a -f /{args.script}'"

print(command)

rc, _, err = call_command_and_capture(command)

if rc == 0:
    print(f"Database {text_end} successfully")
else:
    print(f"Error {text_start} database. Exiting...")
    print(err)
    sys.exit(1)
