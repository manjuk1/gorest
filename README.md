# REST API's with Go

Libraries and its purpose:
Gin-Gonic
Gorm
Postgres DB
Dep for dependecy management
Viper for configuration management - envprefix - PREFIX_NAME in env!!!
dlv debug
jwt-go 

Directory Structure:
Lessons learnt!!!


Dep init
dep ensure

API's built:

Public API
POST http://localhost:8080/api/v1/users

Request Body:

{
"user": {
	"name": "NewName456",
	"email": "1234q56weqwe@32.com",
	"password": "1235456789" 
}
}

Response Body: 201 Created

{
    "user": {
        "id": 17,
        "name": "NewName456",
        "email": "1234q56weqwe@32.com"
    }
}

GET http://localhost:8080/api/v1/users/17

Response: 200 OK

{
    "user": {
        "id": 17,
        "name": "NewName456",
        "email": "1234q56weqwe@32.com"
    }
}


POST http://localhost:8080/api/v1/users/login

Body: 

{
"user": {
	"email": "1234q56weqwe@32.com",
	"password": "123545678910" 
}
}

Response:

401 Unauthorized 
{
    "errors": {
        "Authentication": "Unauthorized Access"
    }
}

200 OK

{
    "message": "User Authenticated Successfully"
}

422 Unprocessable entity

{
    "errors": {
        "Email": "{key: required}"
    }
}