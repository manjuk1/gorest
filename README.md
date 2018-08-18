# JWT based REST API's using Go

Following are the libraries used in this project

1. **Gin-Gonic** : An HTTP Web framework written in go.  
2. **Gorm** : ORM layer. 
3. **Postgres DB**: Database to host business data.
4. **Dep**: Dependecy management. (dep init; dep ensure) 
5. **Viper**: Configuration management(Support for both config file as well as env vars)
6. **Delve**: Runtime debugging using delve (dlv debug).
7. **jwt-go**: Library to support JWT based authentication.

# Directory Structure:

> ### Domain based structuring
```

└── app
│    └── users             // package user
│    	├── api.go         // request handlers for all User related http requests
│    	├── model.go       // data model definition
│    	├── serializer.go  // response formatting
│    	└── validator.go   // Request body validations
├── common                 // package common
│   ├── utils.go           // utility functions
├── config                 // package config
│   ├── config.go          // viper configuration 
│   └── config.yaml        // application configuration
├── db                     // package db
│   └── db.go              // DB connect manager using GORM
├── main.go                // package main
├── middlewares.go         // middleware for JWT auth
├── routes.go              // router binding

```
### Points to remember

> ### Domain based structuring

1. Go with domain based structuring only if we are sure that there cannot be interdependency between domains. It could lead to cirular imports if any dependecy arises.
2. Real world problems has depenedency between the domains. Ex: Get the number of orders placed by the users. In this case, User and Order domains are interdepndent. This leads to bottleneck in domain based structuring.
3. Code readabilty with domain based structuring. Ex: users.User to access User model, users.Create for creat user end point.

>  ### Rails style structuring
1. Rails style structuring solves problem mentioned with domain based structuring.
2. Models, views and controllers packages contains the code for all the model, view and controller respectively.
3. Code deadability with Rails style structuring. Ex: controllers.CreateUser for create user, models.User to access User model 

# API's

### Create User (Public API)
```
POST http://localhost:8080/api/v1/users

Request Body:

{
	"user": {
		"name": "NewName456",
		"email": "1234q56weqwe@32.com",
		"password": "1235456789" 
	}
}

Response: 201 Created

{
    "user": {
        "id": 17,
        "name": "NewName456",
        "email": "1234q56weqwe@32.com"
	"Token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzQ2ODIwNDcsImlkIjoxOX0.3BNwRO4aCFmrmBKjNcYgJ-JU0Y76ZpPK0tsNyLmWPTc
    }
}
```
### Get User Details (Protected API)
```
GET http://localhost:8080/api/v1/users/17

Header:
Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzQ2ODIwNDcsImlkIjoxOX0.3BNwRO4aCFmrmBKjNcYgJ-JU0Y76ZpPK0tsNyLmWPTc

Response: 200 OK

{
    "user": {
        "id": 17,
        "name": "NewName456",
        "email": "1234q56weqwe@32.com"
	"Token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzQ2ODIwNDcsImlkIjoxOX0.3BNwRO4aCFmrmBKjNcYgJ-JU0Y76ZpPK0tsNyLmWPTc
    }
}
````

### User Login (Protected API)

```
POST http://localhost:8080/api/v1/users/login

Request Body: 

{
	"user": {
		"email": "1234q56weqwe@32.com",
		"password": "123545678910" 
	}
}

Response:

200 OK
------

{
    "user": {
        "id": 17,
        "name": "NewName456",
        "email": "1234q56weqwe@32.com"
	"Token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzQ2ODIwNDcsImlkIjoxOX0.3BNwRO4aCFmrmBKjNcYgJ-JU0Y76ZpPK0tsNyLmWPTc
    }
}

401 Unauthorized 
----------------
{
    "errors": {
        "Authentication": "Unauthorized Access"
    }
}

422 Unprocessable entity
------------------------

{
    "errors": {
        "Email": "{key: required}"
    }
}

```
