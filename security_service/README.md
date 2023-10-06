# Description
Security is a microservice built for the e-petitions project in order to provide a layer of safety for our users.

## Conventions
Security is developed to use the gRPC protocol and will use messages from predefined proto files (internal/proto folder).

## RPC Methods
### Login
```shell
grpcurl -plaintext -d '{"email": "example@email.com", "password" : "examplepass"}'\
 -import-path internal/proto   -proto internal/proto/security_svc.proto   \
 localhost:9002 proto.SecurityService.Login
```
#### Request
#### UserCredentials:
- `UserCredentials` is a message that contains the information needed to log in the user. It should contain an existent email that is stored in the database and a valid password.
```grpc
message UserCredentials {
  string email = 1;
  string password = 2;
}
```
#### Response
#### Tokens:
- `Tokens` is a message that is returned after a successful identification of the user. It contains a set of tokens: an access token and a refresh token.
```grpc
message Tokens {
  string access_token = 1;
  string refresh_token = 2;
}
```
