## EPay Database Migrator

This is a simple migrator for EPay. It is based on the [gorm](https://gorm.io) library.

## Usage

First, you need to build the binary:

```bash
go build -o migrator cmd/migrator/main.go
```

### Migrate

You need to set the following environment variables:

```bash
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=epay
export DB_PASSWORD=[password]
export DB_NAME=epay
```

you can also set by using the `.env` file

Then you can run the migrator:

```bash
./migrator --action migrate
```

If you want to debug the migrator, you can set the `--debug` flag:

```bash
./migrator --action migrate --debug
```
