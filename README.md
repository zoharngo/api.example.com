# api.example.com

## Project Setup

### Init

- $ sudo apt install golang-go go-dep
- $ dep init

- $ dep ensure --add github.com/joho/godotenv
- $ dep ensure --add github.com/go-xorm/xorm
- $ dep ensure --add _ "github.com/go-sql-driver/mysql"
- $ dep ensure --add github.com/gorilla/mux

### DB Setup

- $ sudo apt update && sudo apt install mysql-server -y
- $ sudo mysql_secure_installation
- $ sudo mysql

- mysql> SELECT user,authentication_string,plugin,host FROM mysql.user;
- mysql> ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
- mysql> FLUSH PRIVILEGES;

- mysql -u root -p

- mysql> CREATE USER 'sammy'@'localhost' IDENTIFIED BY 'password';
- mysql> GRANT ALL PRIVILEGES ON *.* TO 'sammy'@'localhost' WITH GRANT OPTION;
- mysql> exit

- $ systemctl status mysql.service

### DB migration

- $ sudo npm install -g db-migrate db-migrate-mysql

- $ db-migrate create add-users
- $ db-migrate up
- $ db-migrate down

#### [test] scope
- $ db-migrate up:test 
- $ db-migrate down:test


- mysql> show tables;
- mysql> explain users;




