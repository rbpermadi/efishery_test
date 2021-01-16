# product-api

## Description

Product Api is a part of repository for Efishery Test Assignment. It's using **NodeJs** as its programming language.

## Onboarding and Development Guide

### Prequisites
* [**NodeJS**](https://nodejs.org/)
* [**npm**](https://www.npmjs.com/)

### Setup

- Please install/clone the [Prequisites](#prequisites) and make sure it works perfectly on your local machine.

- Install dependencies

    ```
    cd path/to/product-api
    npm install
    ```

- Copy and edit(optional) `default.json.sample`

    ```
    cd path/to/product-api/config
    cp default.json.sample default.json
    ```

### Running the app

Finally, run **Efishery Test User Api** in your local machines.

    ```
    npm start
    ```

To kill the server you just need to hold `Ctrl + C`


### Test

- Run all tests

  ```
  npm test
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

### Scaffolding

Feathers has a powerful command line interface. Here are a few things it can do:

```
$ npm install -g @feathersjs/cli          # Install Feathers CLI

$ feathers generate service               # Generate a new Service
$ feathers generate hook                  # Generate a new Hook
$ feathers help                           # Show all commands
```


## FAQ

> Not available yet
