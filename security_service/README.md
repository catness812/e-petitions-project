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
## Test Cases
### Login: Successful
In order for a login request to return the expected data, the request should include an **_email_** and a **_password_** of a currently registered user. A registered user is considered the one present in the database.
Suppose we have a user registered with the **_email_** set to _**example@email.com**_ and the **_password_** set to **_examplepass_**. If sent, this request would return a set of tokens: access and refresh:

#### Request Message
```json
{
  "email": "example@email.com",
  "password" : "examplepass"
}
```
#### Expected Response Message
```json
{
  "accessToken": "access token here",
  "refreshToken": "refresh token here"
}
```

### Login: Failed
A login request will return an error from the gRPC method with a status _**"NotFound"**_ error and a _**"failed to login user"**_ error.

#### Request Message
```json
{
  "email": "failexample@email.com",
  "password" : "failexamplepass"
}
```
#### Expected Response Message
```grpc
ERROR:
Code: NotFound
Message: failed to login user
```

### Refresh: Successful
In order for a refresh request to return the expected data, the request should include a correctly generated **_refresh token_** that was returned after a **_successful login_**.
Suppose the request contains a correctly generated refresh token with a valid signature. If sent, this request would return a new set of tokens: access and refresh:

#### Request Message
```json
{
  "token": "correctly generated with a valid signature refresh token here"
}
```
#### Expected Response Message
```json
{
  "tokens": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTcxODMwMDksInVzZXJFbWFpbCI6ImV4YW1wbGVAZW1haWwuY29tIn0.5aQNUv8_mEXMt3eYoAa_ymUWfTcbiYuVS2wSSFMXJ94",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTcyMDM3MDksInVzZXJFbWFpbCI6ImV4YW1wbGVAZW1haWwuY29tIn0.rQuTpVE5n2Ti81-YMSocZSQZOBU5Zqj6LR5xw8al1x8"
  }
}
```

### Refresh: Failed
A refresh request will return an error from the gRPC method with a status _**"Unauthenticated"**_ error and a _**"failed to refresh user session"**_ error.

#### Request Message
```json
{
  "token": "incorrectly generated with a invalid signature refresh token here"
}
```
#### Expected Response Message
```grpc
ERROR:
Code: Unauthenticated
Message: failed to refresh user session
```

### ValidateToken: Successful
In order for a validate request to return the expected data, the request should include a correctly generated **_token_** that was returned after a **_successful login_**.
Suppose the request contains a correctly generated token with a valid signature. If sent, this request would return the token and the subject included in the claims of this token:

#### Request Message
```json
{
  "token": "correctly generated with a valid signature token here"
}
```
#### Expected Response Message
```json
{
  "token": "correctly generated with a valid signature token here",
  "email": "example@email.com"
}
```

### ValidateToken: Failed
A validate request will return an error from the gRPC method with a status _**"Unauthenticated"**_ error and a _**"failed to validate token"**_ error.

#### Request Message
```json
{
  "token": "incorrectly generated with a invalid signature token here"
}
```
#### Expected Response Message
```grpc
ERROR:
Code: Unauthenticated
Message: failed to validate token
```

### SendOTP: Successful
In order for a send OTP request to return the expected data, the request should include an **_email_**.
Suppose the request contains an email. If sent, this request would return the sent otp and the recipient.

#### Request Message
```json
{
  "email": "example@email.com"
}
```
#### Expected Response Message
```json
{
  "OTP": "12345",
  "email": "example@email.com"
}
```

### ValidateOTP: Successful
In order for a send OTP request to return the expected data, the request should include an **_email_**.
Suppose the request contains an email. If sent, this request would return the otp sent and the recipient.

#### Request Message
```json
{
  "email": "example@email.com",
  "OTP": "12345"
}
```
#### Expected Response Message
```json
{
  "validated": true
}
```
### ValidateOTP: Failed
A validate OTP request will return an error from the gRPC method with a status _**"InvalidArgument"**_ error and a _**"failed to validate otp"**_ error, if the otp sent in the request is not present in the Redis Database.


#### Request Message
```json
{
  "email": "example@email.com",
  "OTP": "12345"
}
```
#### Expected Response Message
```grpc
ERROR:
Code: InvalidArgument
Message: failed to validate otp
```