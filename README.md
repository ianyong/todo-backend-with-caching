# To-do Backend

## Getting Started

1. Install Go `>= 1.17` by following the instructions [here](https://golang.org/doc/install).
1. Install PostgreSQL `>= 12` by following the instructions [here](https://www.postgresql.org/download/).
1. Make a copy of `.env.development` as `.env.development.local`.
   For development on your local system, it is recommended that you connect to PostgreSQL via Unix-domain sockets so that there is no need for database server credentials, allowing you to leave `DB_USER` and `DB_PASSWORD` empty.
   This can be done by setting the `DB_HOST` to the following locations:
   * Linux: `/var/run/postgresql`
   * macOS: `/tmp`

   Note that if you are connecting via `localhost`, you might need to set `DB_SSLMODE=disable`.
1. Create the database.
   ```sh
   $ make createdb
   ```
1. Migrate the database.
   ```sh
   $ make migratedb
   ```
1. Start the server.
   ```sh
   $ make run
   ```

## Linting

1. Install `golangci-lint` by following the instructions [here](https://golangci-lint.run/usage/install/#local-installation).
1. Optionally, you can integrate the linter with your IDE if it is supported by following the instructions [here](https://golangci-lint.run/usage/integrations/).
   Otherwise, you will need to run `make fmt` and `make lintfix` to automatically format and fix any lint violations before you commit any changes, or add a pre-commit Git hook that does it for you automatically.

Note that `gosec` is a supported linter in `golangci-lint`.
As such, there is no need to separately install `gosec` for local development.
