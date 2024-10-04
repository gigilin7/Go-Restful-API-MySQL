# Go-Restful-API-MySQL
Build a message board using Golang, Gin and MySQL. It can do CRUD for management users.

## Database
+ Download MySQL
  + [MySQL server](https://dev.mysql.com/downloads/mysql/) 9.0.1 Innovation, Windows (x86, 64-bit), MSI Installer
  + [MySQL workbench](https://dev.mysql.com/downloads/workbench/) 8.0.38, Windows (x86, 64-bit), MSI Installer
+ Open MySQL workbench and build MySQL connections
+ Create schema ( In MySQL, schema is equivalent to database )
  + In this project, schema name is  `go_database` ( You can see in `sql/connect.yaml` )

## Structure
```
.
├── controller
│   └── controller.go
├── model
│   └── model.go
├── repository
│   └── repository.go
├── router
│   └── router.go
└── sql
│   ├── connect.yaml
│   └── sql.go
├── go.mod
├── go.sum
└── main.go
```
Explain the individual functions and functions of the above folders:
+ `controller`：Check logic of CRUD operation.
+ `model`：The data object.
+ `repository`：Functions of CRUD operation with database. 
+ `router`：Set up website URL routing.
+ `sql`：The setting of database.

## Notice ⚠
+ Following changes in Go 1.16, functionality in `ioutil` has been moved to the `os` package. Therefore, `ioutil.ReadFile` should be changed to `os.ReadFile`
+ If you encounter problems like: `{"message":"Error 1146 (42S02): Table 'go_database.message' doesn't exist"}` when you want to query message.
Use the `AutoMigrate` method to automatically create data tables. ( You can see in `sql/sql.go` )

  ```go
  err = Connect.AutoMigrate(&model.Message{})
  if err != nil {
    return fmt.Errorf("error running AutoMigrate: %w", err)
  }
  ```

## Implement
```
 go run main.go
```
CRUD operation instructions:

```
curl http://localhost:8081/api/v1/message
```

```
curl http://localhost:8081/api/v1/message/:id
```

```
curl -X POST http://localhost:8081/api/v1/message \
-H "Content-Type: application/x-www-form-urlencoded" \
-d 'User_Id=1&Content=Hi there'
```

```
curl -X PATCH http://localhost:8081/api/v1/message/3 \
-H "Content-Type: application/x-www-form-urlencoded" \
-d 'Content=Test three!!!'
```

```
curl -X DELETE http://localhost:8081/api/v1/message/:id
```

## Result
### Query all messages - first time
<img src="https://github.com/gigilin7/Go-Restful-API-MySQL/blob/main/picture/mysql-getAll-origin.jpg" height=150>

### Create one message
<img src="https://github.com/gigilin7/Go-Restful-API-MySQL/blob/main/picture/mysql-create.jpg" height=300>

### Query all messages
<img src="https://github.com/gigilin7/Go-Restful-API-MySQL/blob/main/picture/mysql-getall-new.jpg" height=500>

### Update one message
<img src="https://github.com/gigilin7/Go-Restful-API-MySQL/blob/main/picture/mysql-update.jpg" height=300>

### Database in MySQL
<img src="https://github.com/gigilin7/Go-Restful-API-MySQL/blob/main/picture/mysql.jpg" height=500>


[Reference for learning](https://github.com/880831ian/go-restful-api-repository-messageboard)
