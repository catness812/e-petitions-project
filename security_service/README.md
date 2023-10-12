# Description
Security is a microservice built for the e-petitions project in order to provide a layer of safety for our users.

## Contents
- [Conventions](#conventions)
- [RPC Methods](#rpc-methods)
- [Login](#login)
- [RefreshSession](#refreshsession)
- [ValidateToken](#validatetoken)
- [SendOTP](#sendotp)
- [ValidateOTP](#validateotp)
- [Docker Image Creation](#docker-image-creation)

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
 -import-path internal/proto   -proto internal/proto/security_svc.proto \
 localhost:9002 proto.SecurityService.Login
```
#### Response
```json
{
"accessToken": "access token here",
"refreshToken": "refresh token here"
}
```

### RefreshSession
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
grpcurl -plaintext -d '{"token":"put your refresh token here"}'\
 -import-path internal/proto   -proto internal/proto/security_svc.proto \
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

### ValidateToken
Validates the token and returns the subject contained in the token if the token is valid.
#### Parameters
- `Token` is a message that contains the token sent to be validated. It should contain a valid token according to our jwt specification.
```grpc
message Token {
  string token = 1;
}
```
- `ValidateTokenResponse` is a message that is returned after a successful validation of the token. It contains the token that was initially sent in the request and the subject contained in the claims of the token.
```grpc
message ValidateTokenResponse {
  string token = 1;
  string email = 2;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"token":"put your access token here"}'\
 -import-path internal/proto   -proto internal/proto/security_svc.proto \
 localhost:9002 proto.SecurityService.ValidateToken
```
#### Response
```json
{
  "token" : "the token that was initially sent",
  "email" : "example@email.com"
}
```

### SendOTP
Generates and sends an OTP by publishing a message in RabbitMQ. The message published contains a link to verify the OTP sent. The message is consumed by the Mail Service.
#### Parameters
- `OTPInfo` is a message that contains the both the OTP and user email. The "SendOTP" method will only use the "email" field of this message to send the OTP to the specified email. After successfully sending the OTP, the method will return a "OTPInfo" message that will contain the generated OTP and the recipient's email.
```grpc
message OTPInfo{
  string OTP = 1;
  string email = 2;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email":"example@email.com"}'\
 -import-path internal/proto   -proto internal/proto/security_svc.proto \
 localhost:9002 proto.SecurityService.SendOTP
```
#### Response
```json
{
  "OTP": "27062",
  "email": "example@email.com"
}
```

### ValidateOTP
Verifies the OTP by checking if the Redis Database contains an entry where key is equal to the email sent in the request and its value is equal to the OTP contained in the request.
#### Parameters
- `OTPInfo` is a message that contains the both the OTP and user email.
```grpc
message OTPInfo{
  string OTP = 1;
  string email = 2;
}
```
- `IsOTPValidated` is a message that contains the confirmation whether the OTP was validated or not.
```grpc
message IsOTPValidated{
  bool validated = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"email":"example@email.com", "otp" : "12345"}'\
 -import-path internal/proto   -proto internal/proto/security_svc.proto \
 localhost:9002 proto.SecurityService.ValidateOTP
```
#### Response
```json
{
  "validated": true
}
```

## Docker Image Creation
### Command
In order to build the image necessary for the Docker compose file, run this command:
```shell
docker build -t e-petitions-security:1.0 .
```
