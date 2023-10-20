# Description
The gRPC Gateway for Microservices in Golang is a critical component of our system that serves as a unified entry point for accessing various microservices through a single HTTP API interface. This gateway acts as a translation layer, enabling clients to communicate with gRPC-based microservices using RESTful HTTP requests. It abstracts the complexities of gRPC while providing a consistent and user-friendly interface to our distributed services.


## Contents
- [User Service](#user-service-endpoints)
    - [Create User](#create-user)
    - [Create User OTP](#create-user-otp)
    - [Get User By Email](#get-user-by-email)
    - [Get User By ID](#get-user-by-id)
    - [Update User](#update-user)
    - [Update User](#update-user)
    - [Delete User](#delete-user)
    - [Add Admin](#add-admin)
- [Petition Service](#petition-service-endpoints)
    - [CreatePetition](#create-petition)
    - [Get Petition By ID](#get-petition-by-id)
    - [Get All Petition](#get-all-petition-)
    - [Update Petition Status](#update-petition-status)
    - [Delete Petition](#delete-petition)
    - [Create Vote](#create-vote)
    - [Get Petition By Title](#get-petitions-by-title)
    - [Get All Similar Petitions](#get-all-similar-petitions)
    - [Get User Petitions](#get-user-petitions)
    - [Get User Voted Petitions](#get-user-voted-petitions)
- [Docker Image](#docker-image-creation)
    - [Build](#command)

## Conventions
Gateway is developed to redirect http request to grpc request and send responses from grpc microservices to http. 

## User Service Endpoints

### Create User

#### Permissions

- User
- Admin

#### Description

This endpoint allows users with USER or ADMIN permissions to create a new user.

#### Request Body
- `email` is the email needed for user registration
- `password` the password needed for user registration , password requires to be minim 8 character long , have minimum 1 symbol character, 1 upper and lower letter character and minimum 1 number 
```json
{
  "email": "example@isa.utm.md",
  "password": "examplepass"
}
```
#### Request
```curl
curl -L 'http://localhost:1337/user/' \
-H 'Content-Type: application/json' \
--data-raw '{
    "email": <your_email>,
    "password": <your_password>
}'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `message` is the specification of operation success or fail
- `error` is the state of process
```json
{
  "error": true,
  "message": "User already Exists"
}
```
### Create User OTP

#### Permissions

- User
- Admin

#### Description

This endpoint allows users with USER or ADMIN permissions to create a new user without password using otp.

#### Request Body
- `email` is the email needed for user registration
- `password` the password needed for user registration , password requires to be minim 8 character long , have minimum 1 symbol character, 1 upper and lower letter character and minimum 1 number
```json
{
  "email": "example@isa.utm.md",
  "password": "examplepass"
}
```
#### Request
```curl
curl -L 'http://localhost:1337/user/otp' \
-H 'Content-Type: application/json' \
--data-raw '{
    "email": <your_email>,
    "password": <your_password>
}'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `message` is the specification of operation success or fail
- `error` is the state of process
```json
{
  "error": true,
  "message": "User already Exists"
}
```


### Get User By Email

#### Permissions

- Admin

#### Description

The purpose of this endpoint is to retrieve user information. It allows users with "ADMIN" permissions to access and view user details. This could include information about a specific user account, such as their username, email, role, or any other relevant user-related information.
#### Request Headers
- `Authorization`: acces_token
#### Request Body
- `email` is the email needed for getting user information
```json
{
  "email": "example@isa.utm.md"
}
```
#### Request
```curl
curl --L --request GET 'http://localhost:1337/user' \
-H 'Content-Type: application/json' \
-H 'Authorization: <your_token>\
--data-raw '{
    "email": <your_email>
    }'
}'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `message` is the specification of operation success or fail
- `error` is the state of process
```json
{
  "error": true,
  "message": "failed to validate token"
}
```
### Get User by ID

#### Permissions

- User

#### Description

This endpoint allows users with USER and ADMIN permissions to get user information.

#### Request Headers
- `Authorization`: acces_token
#### Endpoint URL Parameters

`uid`: This is a URL parameter representing the unique identifier of the user to be shown the information. Administrators need to specify the ID of the user they want to get information.

#### Request
```curl
curl -L -X GET 'http://localhost:1337/user/uid' \
-H 'Content-Type: application/json' \
-H 'Authorization: <your_token> \

```
#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `message` is the specification of operation success or fail
- `error` is the state of process
```json
{
  "error": true,
  "message": "failed to validate token"
}
```

### Update User

#### Permissions

- User

#### Description

This endpoint allows users with USER permissions to update his user information.

#### Request Headers
- `Authorization`: acces_token
#### Request Body
- `email` is the email needed for user to update user information
```json
{
  "email": "example@isa.utm.md"
}
```

#### Request
```curl
curl -L 'http://localhost:1337/user/update' \
-H 'Authorization: <your_token>' \
-H 'Content-Type: application/json' \
--data-raw '{
    "email" : "<your_email>",
    "password": "<your_password>"
}'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `message` is the specification of operation success or fail
- `error` is the state of process
```json
{
  "error": true,
  "message": "failed to validate token"
}
```

### Delete User

#### Permissions

- Admin
- User

#### Description

This endpoint allows users with USER or ADMIN permissions to delete a user.

#### Request Headers
- `Authorization`: acces_token
#### Request Body
- `email` is the email needed for user to delete a user
```json
{
  "email": "example@isa.utm.md"
}
```
#### Request
```curl
curl -L -X DELETE 'http://localhost:1337/user' \
-H 'Authorization: <your_token>' \
-H 'Content-Type: application/json' \
--data-raw '{
    "email" : "<your_email>"
}'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `message` is the specification of operation success or fail
- `error` is the state of process
```json
{
  "error": true,
  "message": "failed to validate token"
}
```

### Add Admin

#### Permissions

- Admin

#### Description

This endpoint allows users with  ADMIN permissions to assign a new admin.

#### Request Headers
- `Authorization`: acces_token
#### Request Body
- `email` is the email needed for admin to assign a new user
```json
{
  "email": "example@isa.utm.md"
}
```

#### Request
```curl
curl -L 'http://localhost:1337/user/admin' \
-H 'Authorization: <your_token>' \
-H 'Content-Type: application/json' \
--data-raw '{
    "email":"<your_email>"
}'
```


#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `message` is the specification of operation success or fail
- `error` is the state of process
```json
{
  "error": true,
  "message": "failed to validate token"
}
```

### Petition Service Endpoints

### Create Petition

#### Permissions

- Admin
- User

#### Description

The purpose of this endpoint is to enable users to create a new petition. It's designed for users who want to initiate a new petition request.
#### Request Headers
- `Authorization`: acces_token
#### Request Body
- `title`(string) is the title of the petition
- `category`(string) is the category of the petition
- `description`(string) is the body of the petition
- `image`(string) is the path to the petition image
- `user_id`(uint) is the id of creator off the petition
- `vote_goal`(uint) is the number of votes to be reached for the petition to address to the consignee
- `exp_date`(string) is the date when petition expires , date needs to be formatted in RFC3339 format
```json
{
  "title": "examplee",
  "category": "Student Dorms",
  "description": "example",
  "image": "path/to/the/image",
  "user_id": 1,
  "vote_goal": 10,
  "exp_date": "2023-12-31T00:00:00Z"
}
```

#### Request
```curl
curl -L 'http://localhost:1337/petition' \
-H 'Content-Type: application/json' \
-d '{
    "title": "<your_title>",
    "category": "<your_category>",
    "description": "your_description",
    "image": "path/to/the/image",
    "user_id": <your_id>,
    "vote_goal": <your_vote_goal>,
    "exp_date": "<your_time>"
}'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
- `petition_id` is the id of the new petition
```json
{
  "error": false,
  "message": "Petition retrived successfully",
  "petition_id": 1 
}
```

### Get Petition by ID

#### Permissions

- User
- Admin

#### Description

This endpoint allows users with USER and ADMIN permissions to get petition information.


#### Endpoint URL Parameters

-`pid`: This is a URL parameter representing the unique identifier of the petition to be shown the information. User need to specify the ID of the petition they want to get information.

#### Request
```curl
curl -L 'http://localhost:1337/petition/:pid'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
- `petition` is the whole petition information
```json
{
  "error": true,
  "message": "failed to validate token",
  "petition": {"petition_id": 1,
               "title": "exampleee",
               "category": "null",
               "description": "example",
               "image": "",
               "status": {
                 "id": 4,
                 "status": "DRAFT"
               },
               "user_id": 4,
               "vote_goal": 10,
               "current_votes": 0,
               "exp_date": "2023-12-31 00:00:00",
               "updated_at": "2023-10-19 12:43:24",
               "created_at": "2023-10-19 12:43:24"
              }
}
```

### Get All Petition 

#### Permissions

- User
- Admin

#### Description

This endpoint allows users with USER and ADMIN permissions to get all petitions information.

#### Endpoint URL Parameters

- `page`: This is a URL parameter representing the page number
- `limit`: This is a URL parameter representing the number of petitions displaied on the page

#### Request
```curl
curl -L 'http://localhost:1337/petition/all/:page/:limit'
```


#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
- `petitions` is the array of petitions information
```json
{
  "error": true,
  "message": "failed to validate token",
  "petitions": [{"petition_id": 1,
               "title": "exampleee",
               "category": "null",
               "description": "example",
               "image": "",
               "status": {
                 "id": 4,
                 "status": "DRAFT"
               },
               "user_id": 4,
               "vote_goal": 10,
               "current_votes": 0,
               "exp_date": "2023-12-31 00:00:00",
               "updated_at": "2023-10-19 12:43:24",
               "created_at": "2023-10-19 12:43:24"
              },
              {
              "petition_id": 5,
              "title": "exampleee",
              "category": "null",
              "description": "example",
              "image": "",
              "status": {
                "id": 4,
                "status": "DRAFT"
              },
              "user_id": 1,
              "vote_goal": 10,
              "current_votes": 0,
              "exp_date": "2023-12-31 00:00:00",
              "updated_at": "2023-10-19 19:55:17",
              "created_at": "2023-10-19 19:55:17"
              }
  ]
}
```

### Update Petition Status

#### Permissions

- Admin

#### Description

The purpose of this endpoint is to allow users with ADMIN permissions to update the status of a petition. It provides administrative control over the status of petitions within the system.

#### Request Body
- `id`(uint) is the id of the petition
- `status`(string) is the updated status of the petition
```json
{
  "id": 4,
  "status": "DRAFT"
}
```
#### Request
```curl
curl -L 'http://localhost:1337/petition/status/' \
-H 'Content-Type: application/json' \
-d '{
    "id": 4,
    "status": "DRAFT"
}'
```
#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
```json
{
  "error": false,
  "message": "Petition status updated successfully"
  
}
```

### Delete Petition

#### Permissions

- Admin

#### Description

The purpose of this endpoint is to allow users to delete their own petitions or, for administrators, to delete unposted petitions. Users can use this endpoint to remove a petition they created or to remove unposted petitions from the system.

#### Endpoint URL Parameters

`pid`: This is a URL parameter representing the unique identifier of the petition to be deleted. Users or administrators need to specify the ID of the petition they want to delete.


#### Request
```curl
curl -L -X DELETE 'http://localhost:1337/petition/:pid'
```
#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
```json
{
  "error": true,
  "message": "failed to validate token"
}
```


### Create Vote

#### Permissions

- User

#### Description

The purpose of this endpoint is to allow users to sign a petition. Users can express their support for a specific petition by "signing" it, which is a common feature in many petition systems.
#### Endpoint URL Parameters
- `uid`: This is a URL parameter representing the unique identifier of the user what is signing the petition. Users  need to specify own ID if they want to sign.
- `pid`: This is a URL parameter representing the unique identifier of the petition to be signed. Users  need to specify the ID of the petition they want to sign.

#### Request
```curl
curl -L 'http://localhost:1337/petition/sign/:uid/:pid'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
```json
{
  "error": true,
  "message": "User email not found"
}
```

### Get Petitions By Title

#### Permissions

- User
- Admin

#### Description

This endpoint allows users with USER and ADMIN permissions to get all similar title petitions.

#### Endpoint URL Parameters

- `page`: This is a URL parameter representing the page number
- `limit`: This is a URL parameter representing the number of petitions displaied on the page

#### Request Body
- `title`(string) is the title of which we want to find similar petitions
```json
{
  "title": "Bad Service"
}
```
#### Request
```curl
curl -L -X GET 'http://localhost:1337/petition/search/:page/:limit' \
-H 'Content-Type: application/json' \
-d '{
    "title": "<your_title>"
}'
```


#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
- `petitions` is the array of petitions information
```json
{
  "error": false,
  "message": "Petitions retrived successfully",
  "petitions": [{"petition_id": 1,
               "title": "example",
               "category": "null",
               "description": "example",
               "image": "",
               "status": {
                 "id": 4,
                 "status": "DRAFT"
               },
               "user_id": 4,
               "vote_goal": 10,
               "current_votes": 0,
               "exp_date": "2023-12-31 00:00:00",
               "updated_at": "2023-10-19 12:43:24",
               "created_at": "2023-10-19 12:43:24"
              },
              {
              "petition_id": 5,
              "title": "exampleee",
              "category": "null",
              "description": "example",
              "image": "",
              "status": {
                "id": 4,
                "status": "DRAFT"
              },
              "user_id": 1,
              "vote_goal": 10,
              "current_votes": 0,
              "exp_date": "2023-12-31 00:00:00",
              "updated_at": "2023-10-19 19:55:17",
              "created_at": "2023-10-19 19:55:17"
              }
  ]
}
```

### Get All Similar Petitions

#### Permissions

- User
- Admin

#### Description

This endpoint allows users with USER and ADMIN permissions to get all similar petitions.

#### Request Body
- `title`(string) is the title of which we want to find similar petitions
```json
{
  "title": "Bad Service"
}
```
#### Request
```curl
curl -L -X GET 'http://localhost:1337/petition/similar' \
-H 'Content-Type: application/json' \
-d '{
    "title": "<your_title>"
}'
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
- `petitions` is the array of petitions information
```json
{
  "error": false,
  "message": "Petitions retrived successfully",
  "petitions": [{"petition_id": 1,
               "title": "example",
               "category": "null",
               "description": "example",
               "image": "",
               "status": {
                 "id": 4,
                 "status": "DRAFT"
               },
               "user_id": 4,
               "vote_goal": 10,
               "current_votes": 0,
               "exp_date": "2023-12-31 00:00:00",
               "updated_at": "2023-10-19 12:43:24",
               "created_at": "2023-10-19 12:43:24"
              },
              {
              "petition_id": 5,
              "title": "exampleee",
              "category": "null",
              "description": "example",
              "image": "",
              "status": {
                "id": 4,
                "status": "DRAFT"
              },
              "user_id": 1,
              "vote_goal": 10,
              "current_votes": 0,
              "exp_date": "2023-12-31 00:00:00",
              "updated_at": "2023-10-19 19:55:17",
              "created_at": "2023-10-19 19:55:17"
              }
  ]
}
```
### Get User Petitions

#### Permissions

- User

#### Description

This endpoint allows users with USER permissions to get all created by himself petitions.

#### Endpoint URL Parameters

- `uid`: This is a URL parameter representing the unique identifier of the user what created petitions. Users  need to specify own ID if they want to get all created petitions.
- `page`: This is a URL parameter representing the page number
- `limit`: This is a URL parameter representing the number of petitions displaied on the page

#### Request
```curl
curl -L 'http://localhost:1337/user/petitions/:uid/:page/:limit' 
```

#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
- `petitions` is the array of petitions information
```json
{
  "error": false,
  "message": "Petitions retrived successfully",
  "petitions": [{"petition_id": 1,
               "title": "example",
               "category": "null",
               "description": "example",
               "image": "",
               "status": {
                 "id": 4,
                 "status": "DRAFT"
               },
               "user_id": 4,
               "vote_goal": 10,
               "current_votes": 0,
               "exp_date": "2023-12-31 00:00:00",
               "updated_at": "2023-10-19 12:43:24",
               "created_at": "2023-10-19 12:43:24"
              },
              {
              "petition_id": 5,
              "title": "exampleee",
              "category": "null",
              "description": "example",
              "image": "",
              "status": {
                "id": 4,
                "status": "DRAFT"
              },
              "user_id": 1,
              "vote_goal": 10,
              "current_votes": 0,
              "exp_date": "2023-12-31 00:00:00",
              "updated_at": "2023-10-19 19:55:17",
              "created_at": "2023-10-19 19:55:17"
              }
  ]
}
```

### Get User Voted Petitions

#### Permissions

- User

#### Description

This endpoint allows users with USER permissions to get all signed by himself petitions.

#### Endpoint URL Parameters

- `uid`: This is a URL parameter representing the unique identifier of the user what signed petitions. Users need to specify own ID if they want to get all signed petitions.
- `page`: This is a URL parameter representing the page number
- `limit`: This is a URL parameter representing the number of petitions displaied on the page

#### Request
```curl
curl -L 'http://localhost:1337/user/voted/:uid/:page/:limit' 
```
#### Response

Upon execution, this endpoint returns a JSON response with the following structure:
- `error` is the state of process
- `message` is the specification of operation success or fail
- `petitions` is the array of petitions information
```json
{
  "error": false,
  "message": "Petitions retrived successfully",
  "petitions": [{"petition_id": 1,
               "title": "example",
               "category": "null",
               "description": "example",
               "image": "",
               "status": {
                 "id": 4,
                 "status": "DRAFT"
               },
               "user_id": 4,
               "vote_goal": 10,
               "current_votes": 0,
               "exp_date": "2023-12-31 00:00:00",
               "updated_at": "2023-10-19 12:43:24",
               "created_at": "2023-10-19 12:43:24"
              },
              {
              "petition_id": 5,
              "title": "exampleee",
              "category": "null",
              "description": "example",
              "image": "",
              "status": {
                "id": 4,
                "status": "DRAFT"
              },
              "user_id": 1,
              "vote_goal": 10,
              "current_votes": 0,
              "exp_date": "2023-12-31 00:00:00",
              "updated_at": "2023-10-19 19:55:17",
              "created_at": "2023-10-19 19:55:17"
              }
  ]
}
```

## Docker Image Creation
### Command
In order to build the image necessary for the Docker compose file, run this command:
```shell
docker build -t e-petitions-gateway:1.0 .
```


