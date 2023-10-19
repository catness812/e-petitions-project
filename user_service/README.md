# Description
User is a microservice built for the e-petitions project in order to provide a user-related functionality.

## Contents
- [Conventions](#conventions)
- [RPC Methods](#rpc-methods)
- [CreateUser](#createuser)
- [CreateUserOTP](#createuserotp)
- [UpdateUser](#updateuser)
- [DeleteUser](#deleteuser)
- [GetUserByEmail](#getuserbyemail)
- [GetUserEmailById](#getuseremailbyid)
- [Docker Image Creation](#docker-image-creation)

## Conventions
User is developed to use the gRPC protocol and will use messages from predefined proto files (internal/proto folder).

## RPC Methods
### CreateUser
Adds a new user entry to the database.
#### Parameters
- `UserRequest` is a message that contains the information needed to create the user. It should contain a valid email and a password.
```grpc
message UserRequest{
    string email = 1;
    string password = 2;
}
```
- `ResponseMessage` is a message that is returned after a successful creations of the user. It contains an informative message.
```grpc
message ResponseMessage {
  string message = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email": "example@isa.utm.md", "password" : "examplepass"}'\
 -import-path internal/proto   -proto internal/proto/user_svc.proto \
 localhost:50052 proto.UserService.CreateUser
```
#### Response
```json
{
  "message": "User register successfully"
}
```

### CreateUserOTP
Creates a user entry in the database exclusively for voting without an account.
#### Parameters
- `UserRequest` is a message that contains the information needed to create the user. It should contain a valid email and a password.
```grpc
message UserRequest{
    string email = 1;
    string password = 2;
}
```
- `ResponseMessage` is a message that is returned after a successful creations of the user. It contains an informative message.
```grpc
message ResponseMessage {
  string message = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email": "example2@isa.utm.md", "password" : "examplepass"}'\
 -import-path internal/proto   -proto internal/proto/user_svc.proto \
 localhost:50052 proto.UserService.CreateUserOTP
```
#### Response
```json
{
  "message": "User added successfully"
}
```

### UpdateUser
Updates the password of a user.
#### Parameters
- `UserRequest` is a message that contains the information needed to update the user's password. The 'password' field should contain the new password.
```grpc
message UserRequest{
    string email = 1;
    string password = 2;
}
```
- `ResponseMessage` is a message that is returned after the user's password was successfully updated. It contains an informative message.
```grpc
message ResponseMessage {
  string message = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email": "example2@isa.utm.md", "password" : "newexamplepass2"}'\
 -import-path internal/proto   -proto internal/proto/user_svc.proto \
 localhost:50052 proto.UserService.UpdateUser
```
#### Response
```json
{
  "message": "User updated successfully"
}
```

### DeleteUser
Deletes a user from the Database.
#### Parameters
- `DeleteUserRequest` is a message that contains the information needed to delete the user from the Database. The 'email' field should contain an existing email.
```grpc
message DeleteUserRequest {
    string email = 1;
}
```
- `ResponseMessage` is a message that is returned after the successful deletion of the user. It contains an informative message.
```grpc
message ResponseMessage {
  string message = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email": "example2@isa.utm.md"}'\
 -import-path internal/proto   -proto internal/proto/user_svc.proto \
 localhost:50052 proto.UserService.DeleteUser
```
#### Response
```json
{
  "message": "User deleted successfully"
}
```

### GetUserByEmail
Retrieves information about a user from the Database.
#### Parameters
- `GetUserByEmailRequest` is a message that contains the information needed to retrieve the user data from the Database. The 'email' field should contain an existing email.
```grpc
message GetUserByEmailRequest {
    string email = 1;
}
```
- `GetUserByEmailResponse` is a message that is returned after successfully invoking the method. It contains the data about the user from the Database.
```grpc
message GetUserByEmailResponse {
    uint32 id = 1;
    string email = 2;
    string password = 3;
    string role = 4;
    bool hasAccount = 5;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email": "example@isa.utm.md"}'\
 -import-path internal/proto   -proto internal/proto/user_svc.proto \
 localhost:50052 proto.UserService.GetUserByEmail
```
#### Response
```json
{
  "id": 6,
  "email": "example@isa.utm.md",
  "password": "encryptedpassword",
  "role": "user",
  "hasAccount": true
}
```

### GetUserEmailById
Retrieves the user email from the Database.
#### Parameters
- `GetUserEmailByIdRequest` is a message that contains the information needed to retrieve the user email from the Database. The 'id' field should contain a valid identifier from the Database that points to this user.
```grpc
message GetUserEmailByIdRequest{
    uint32 id = 1;
}
```
- `ResponseMessage` is a message that is returned after successfully invoking the method. The 'message' field contains the user email.
```grpc
message ResponseMessage{
    string message = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"id": 1}'\
 -import-path internal/proto   -proto internal/proto/user_svc.proto \
 localhost:50052 proto.UserService.GetUserEmailById
```
#### Response
```json
{
  "message": "example@isa.utm.md"
}
```

### AddAdmin
Updates a user with his role set to 'admin'.
#### Parameters
- `AddAdminRequest` is a message that contains the information needed to create a new Admin. The 'email' field should contain a valid email of an existing user.
```grpc
message AddAdminRequest{
    string email = 1;
}
```
- `ResponseMessage` is a message that is returned after successfully invoking the method. It serves as an informative message.
```grpc
message ResponseMessage{
    string message = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email": "example@isa.utm.md"}'\
 -import-path internal/proto   -proto internal/proto/user_svc.proto \
 localhost:50052 proto.UserService.AddAdmin
```
#### Response
```json
{
  "message": "User role updated successfully"
}
```

## Docker Image Creation
### Command
In order to build the image necessary for the Docker compose file, run this command:
```shell
docker build -t e-petitions-user:1.0 .
```