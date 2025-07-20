#!/bin/bash

# user configurable values.
PASSWORD=postgres
DATABASE_NAME=test

# do not modify below this line.
# ------------------------------
DB_URL=postgres://postgres:$PASSWORD@localhost:5432/$DATABASE_NAME
CONFIG_DB_URL=$DB_URL?sslmode=disable

# create tabel and setup password.
sudo -u postgres psql <<EOF
CREATE DATABASE $DATABASE_NAME;
ALTER USER postgres PASSWORD '$PASSWORD';
EOF

# this will not override an existing file, but if a file already existed, the setup program will probably not work correctly.
cat >> "$HOME/.test_gatorconfig.json" <<EOF
{
	"db_url": "$CONFIG_DB_URL",
}
EOF

SETUP_DIR=$(dirname $(realpath "$0"))
cd "$SETUP_DIR/sql/schema"
goose postgres "$DB_URL" up
