## gin-rest-api-sample
Golang REST API sample with MariaDB integration using Gin and GORM. (This project IS NOT a starter kit, it is just an example project.)

This project is a sample project that contains following features:

- REST API server with [Gin Framework](https://github.com/gin-gonic/gin)
- Modular Routes
- Database integration using [GORM](http://gorm.io/)
- Live Reload using [codegangsta/gin](https://github.com/codegangsta/gin)
- JWT Token based Authentication
- [Supported REST API Documentation](https://documenter.getpostman.com/view/723994/RWTeVNA4) (Postman)


## Project Setup

```
$ dep ensure
$ go get github.com/jinzhu/gorm
$ go get github.com/codegangsta/gin
```

GORM should be installed via `go get` since installation via `dep` is imperfect (it does not download dialects directory).

[codegangsta/gin](https://github.com/codegangsta/gin) is an optional package to install if you want to make usage of live reloading feature of server (just like nodemon in Node.js environment). 

### MariaDB Configuration

This project uses MariaDB to store data. Install MariaDB and create a sample database and a user account.

#### Install MariaDB

- [macOS](https://mariadb.com/kb/en/library/installing-mariadb-on-macos-using-homebrew/)
- [Windows](https://mariadb.com/kb/en/library/installing-mariadb-msi-packages-on-windows/)
- [Ubuntu](https://www.itzgeek.com/how-tos/linux/ubuntu-how-tos/install-mariadb-on-ubuntu-16-04.html)


#### Create Database / Account
```sql
CREATE DATABASE sample;
GRANT ALL PRIVILEGES ON sample.* to sample@'%' IDENTIFIED BY 'samplepass';
GRANT ALL PRIVILEGES ON sample.* to sample@'localhost' IDENTIFIED BY 'samplepass';
```

### Configure Environment Variables

Open .env file and edit the values if you need to. This project uses [godotenv](https://github.com/joho/godotenv) to read and use .env file. 

Database config string is formatted in [go-sql-driver format](https://github.com/go-sql-driver/mysql#parameters).

## Start Project

```
$ go run main.go
```

To explicitly compile the code before you run the server:

```
$ go build main.go
$ ./main
```

To use live-reloading in development environment, 

```
$ ./scripts/start-dev
```