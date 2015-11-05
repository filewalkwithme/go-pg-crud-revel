# go-pg-crud-revel

This is an example of CRUD web application writen using the web mvc framework Revel. It makes use of Postgres as database.

![01.png](https://github.com/maiconio/go-pg-crud/blob/master/screenshots/01.png)

![02.png](https://github.com/maiconio/go-pg-crud/blob/master/screenshots/02.png)

# setup

You will need to create a database as describe in the file `info.sql`.
After that you will need to setup Revel:

```
go get github.com/revel/revel
go get github.com/revel/cmd/revel
```

# running

```
revel run github.com/maiconio/go-pg-crud-revel
```

The application will be running at http://localhost:9000
