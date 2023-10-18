# e-petitions-project
# Description
A modern solution to make users' voices matter!

## Contents
- [Conventions](#conventions)
- [RPC Methods](#rpc-methods)
- [Login](#login)
- [RefreshSession](#refreshsession)
- [ValidateToken](#validatetoken)
- [SendOTP](#sendotp)
- [ValidateOTP](#validateotp)
- [Docker Image Creation](#docker-image-creation)

## How to use
In order to use the app, run the "docker compose up" in the root of your project. This will run the services inside a containerized environment. Requests should be sent to the Gateway service port. See Gateway Readme file for instructions.
```shell
docker compose up
```
