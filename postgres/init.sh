#!/bin/bash

# Immediately exits if any error occurs during the script
# execution. 
set -o errexit


# Creating an array that defines the environment variables
# that must be set. This can be consumed later via arrray
# variable expansion ${REQUIRED_ENV_VARS[@]}.
readonly REQUIRED_ENV_VARS=(
  "DB_USER"
  "DB_PASSWORD"
  "DB_DATABASE"
  "POSTGRES_USER")


# Main execution:
# - verifies if all environment variables are set
# - runs the SQL code to create user and database
# - creates initial table
main() {
  check_env_vars_set
  init_user_and_db
  init_table
}


# Checks if all of the required environment
# variables are set. If one of them isn't,
# echoes a text explaining which one isn't
# and the name of the ones that need to be
check_env_vars_set() {
  for required_env_var in ${REQUIRED_ENV_VARS[@]}; do
    if [[ -z "${!required_env_var}" ]]; then
      echo "Error:
    Environment variable '$required_env_var' not set.
    Make sure you have the following environment variables set:
      ${REQUIRED_ENV_VARS[@]}
Aborting."
      exit 1
    fi
  done
}


# Performs the initialization in the already-started PostgreSQL
# using the preconfigured POSTGRES_USER user.
init_user_and_db() {
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
     CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';
     CREATE DATABASE $DB_DATABASE;
     GRANT ALL PRIVILEGES ON DATABASE $DB_DATABASE TO $DB_USER;
EOSQL
}

# Create the iniital table for storing of posts
init_table() {
  psql -v ON_ERROR_STOP=1 --username "$DB_USER" --dbname "$DB_DATABASE" <<-EOSQL
      CREATE TABLE posts (
        id serial PRIMARY KEY,
        title VARCHAR NOT NULL,
        content VARCHAR NOT NULL,
        posted TIMESTAMP WITH TIME ZONE NOT NULL
      );
EOSQL
}

# Executes the main routine with environment variables
# passed through the command line. We don't use them in
# this script but now you know 🤓
main "$@"