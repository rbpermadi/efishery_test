# user_api

## Description

User Api is a part of repository for Efishery Test Assignment. Its a web-service app used as user flatform.  It's using **Go** as its programming language.

## Onboarding and Development Guide

### Documentation

- Database diagram
Currently it's only have 1 table.

  ```
  ----------------------------
  |           Users          |
  ----------------------------
  | id int (pk)              |
  | name varchar(50)         |
  | phone varchar(20)        |
  | role varchar(255)        |
  | password varchar(255)    |
  | created_at datetime      |
  | updated_at datetime      |
  ----------------------------

  ```

### Prequisites

* [**Go (1.12 or later)**](https://golang.org/doc/install)
* [**MySQL**](https://www.mysql.com/downloads/)
* [**Docker**](https://docs.docker.com/get-docker/)

### Setup

- Please install/clone the [Prequisites](#prequisites) and make sure it works perfectly on your local machine.

- Install dependencies

    ```
    make mod
    ```

- Copy and edit(optional) `env.sample`

    ```
    cp env.sample .env
    ```

- Migrate Database
  Make sure you already build docker image for standalone-migration. You can follow [Standalone Migration Docker Build](standalone-migration/README.md)

  Use `make db-setup` if it's your first time run this app, it will create the database and run the migration process.
    ```
    make db-setup
    ```

  If you already create database and you need to migrate, you can use `make db-migrate`.
    ```
    make db-migrate
    ```


### Running the app

Finally, run **Efishery Test User Api** in your local machines.

```
> make run
```

To kill the server you just need to hold `Ctrl + C`


### Test

- Run all tests

  ```
  make test
  ```

- Test coverage

  ```sh
  make coverage
  ```

### Contributing

1. Make new branch with descriptive name about your change(s) and checkout to that branch

   ````
   git checkout -b branch_name
   ````


2. Commit and push your change to upstream

   ````
   git commit -m "message"
   git push [remote_name] [branch_name]
   ````

3. Open pull request in `Github`

4. Ask someone to review your code.

5. If your code is approved, the pull request can be merged.

## FAQ

> Not available yet
