# Gator
This project was made as part of this course on boot.dev: https://www.boot.dev/courses/build-blog-aggregator-golang

## Introduction
Gator is a RSS Feed Aggregator. You can register multiple users and each user has their list of followed feeds, which can be fetched periodically.

## Setup (**Ubuntu/Debian**)
This project needs go, a running PostgresSQL database and a config file.

### 1. Install Go
Because this project is written in go, you'll need the go toolchain to build the executable. For a detailed explanation visit the official docs: https://go.dev/doc/install
1. install go
```sh
curl -sS https://webi.sh/golang | sh
```

2. run `go version` to check the installation.

### 2. Install PostgreSQL
1. install the postgres package
```sh
sudo apt update
sudo apt install postgresql postgresql-contrib
```

2. run `psql --version` to check the installation.

3. set a password
```sh
sudo passwd postgres
```
Once prompted, enter a password for the posgres user. `postgres` is a common choice.

4. Start the Postgres server in the background
```sh
sudo service postgresql start
```

### 3. (Recommended) Install Goose
This project is set up to use goose to run its database migrations, but you can also run the migrations without using Goose.
1. install goose
```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

2. run `goose --version` to check the installation.

### 4. (Optional) Run `setup.sh`
The setup script will take care of setting up the database, running the migrations and creating the config file, containing the database url. 
You can either run the script or do the setup manually as described in the following points. If you run the script you can skip steps 5-8.
Note, the setup script needs Goose to be installed.

By the default the script will create a datbase named `gator` with the password `postgres`, but these values can be changed by editing the variables in the script.
```bash
# user configurable values.
PASSWORD=postgres
DATABASE_NAME=gator
```

### 5. Create The Database
(This can be skipped, if you ran the setup script from step 4.)

1. Connect to the postgre server
```sh
sudo -u postgres psql
```

NOTE: To exit the psql client use `\q`.

2. Create a database, e.g. `gator`
```psql
CREATE DATABASE gator;
```

3. Connect to the new database
```psql
\c gator
```

4. Set the user password
```psql
ALTER USER postgres PASSWORD 'postgres';
```
This example sets the password to `postgres` (this is independent of the password for the postgres user, set before).

5. exit psql by entering `exit` or `\q`

### 6. Run The Migrations
(This can be skipped, if you ran the setup script from step 4.)

If you installed goose, move to `sql/schema` and run
```sh
goose postgres <connection-string> up
```
`<connection-string>` has the followin structure
```
protocol://username:password@host:port/database
```
Assuming you used the password `postgres` and named the databse `gator` the connection string would be
```
postgres://postgres:postgres@locathost:5432/gator
```
If you plan on connecting to the databse or running the migrations more often i would recommend putting the connection string into an .env file and sourcing it
```sh
cat > .env <<EOF
DB_URL=postgres://postgres:postgres@locathost:5432/gator
EOF
source .env
```

### 7. Add Configuration File
(This can be skipped, if you ran the setup script from step 4.)

The program expects a configuration file at `$HOME/.gatorconfig.json`, that contains the database url as a json item. 
This file will also be used by the program to store information, that outlives the program.

The connection string in this file needs to be expanded by appending `?sslmode=false` and placed in a JSON object under the key `db_url`.
Assuming the same connection string as step 6, this would result in:
```json
{
	"db_url": "postgres://postgres:postgres@locathost:5432/gator?sslmode=false",
}
```

### 8. Build The Project
(This can be skipped, if you ran the setup script from step 4.)

To install the project run
```sh
go install .
```

## Commands
- `gator login <username>`<br>
  Set the currently active user.<br>
  <username> Name of the user to login.

- `gator register <username>`<br>
  Add new user, and login.<br>
  <username> Name of the user to register.

- `gator reset`<br>
  Reset the database.

- `gator users`<br>
  List all registered users. Also indicates the active user.

- `gator agg <time between reqs>`<br>
  Start aggregation loop, fetching the followed feeds at the specified interval.<br>
  <time between reqs> String representation of an interval: e.g. `1h`.

- `gator addfeed <name> <url>`<br>
  Add a RSS feed to gator. Note, this does not automatically follow a feed.<br>
  <name> Alias for the feed.<br>
  <url> URL of the RSS feed.

- `gator feeds`<br>
  List all added feeds. These are the available feeds, not the followed feeds.

- `gator follow <url>`<br>
  Follow a feed. You can only follow feeds, added to gator. So you may need to use the `add` command first.
  Followed feeds are fetched by the `agg` command.<br>
  <url> URL of a RSS feed.

- `gator following`<br>
  List all followed feeds of the active user.

- `gator unfollow <feed_url>`<br>
  Unfollow a feed. This feed will no longer be fetched by the `agg` command.<br>
  <feed_url> URL of a RSS feed.

- `gator browse`<br>
  List posts of all followed feeds for the active user.
