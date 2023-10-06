# Description
Security is a microservice built for the e-petitions project in order to provide a layer of safety for our users.

## Contents
- [Conventions](#conventions)
- [RPC Methods](#rpc-methods)
- [Login](#login)

## Conventions
Security is developed to use the gRPC protocol and will use messages from predefined proto files (internal/proto folder).

## RPC Methods
### Login
Checks if a user exists in the database. In case of existence, it returns a set of tokens.
#### Parameters
- `UserCredentials` is a message that contains the information needed to log in the user. It should contain an existent email that is stored in the database and a valid password.
```grpc
message UserCredentials {
  string email = 1;
  string password = 2;
}
```
- `Tokens` is a message that is returned after a successful identification of the user. It contains a set of tokens: an access token and a refresh token.
```grpc
message Tokens {
  string access_token = 1;
  string refresh_token = 2;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email": "example@email.com", "password" : "examplepass"}'\
 -import-path internal/proto   -proto internal/proto/security_svc.proto   \
 localhost:9002 proto.SecurityService.Login
```
#### Response
```json
{
"accessToken": "access token here",
"refreshToken": "refresh token here"
}
```
### Refresh Session
Regenerates a set of new tokens for the user session. Checks if the refresh token sent is contained in the Redis repository.
#### Parameters
- `RefreshRequest` is a message that contains the information needed to return a new set of tokens. It should contain a valid refresh token that is stored in the Redis Database.
```grpc
message RefreshRequest {
  string token = 1;
}
```
- `Tokens` is a message that is returned after a successful identification of the user. It contains a set of tokens: an access token and a refresh token.
```grpc
message RefreshResponse{
  map<string, string> tokens = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"token":"put your token here"}'\
 -import-path internal/proto   -proto internal/proto/security_svc.proto    \
 localhost:9002 proto.SecurityService.RefreshSession
```
#### Response
```json
{
  "tokens": {
    "refresh_token": "refresh token here",
    "access_token": "access token here"
  }
}
```
