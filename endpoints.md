# ENDPOINTS

* GET("/books")
* GET("/books/:id")
* GET("/books/:id/content/:page")
* GET("/seed")

## Steps to Run project

Please download all dependencies:

```
    go mod download
```

Run the initial sql to create the tables found in folder migrations inside the internal folder. The sql file is named: book_setup.up.sql

Modify the local.env file, please fill all variables with the required data: db user name and password, db name along with output folder

Run the project and make a get request to /seed so that the tables are populated with data.

Retrieve data at will.
