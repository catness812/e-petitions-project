# E-Petitions Microservice Architecture

[![GoLang](https://img.shields.io/badge/Programming%20Language-GoLang-blue)](https://golang.org/)
[![Gin-Gonic](https://img.shields.io/badge/Framework-Gin--Gonic-green)](https://github.com/gin-gonic/gin)
[![PostgreSQL](https://img.shields.io/badge/Data%20Storage-PostgreSQL-blue)](https://www.postgresql.org/)
[![GORM](https://img.shields.io/badge/ORM-GORM-red)](https://gorm.io/)
[![Redis](https://img.shields.io/badge/Cache-Redis-red)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Deployment-Docker-blue)](https://www.docker.com/)
[![Cloud Deployment](https://img.shields.io/badge/Deployment-Cloud%20(in%20progress)-blue)](https://cloud.google.com/)
[![Domain](https://img.shields.io/badge/Domain-epetitii.co-lightgrey)](https://epetitii.co/)
[![REST API](https://img.shields.io/badge/Communication-REST%20API-orange)](https://en.wikipedia.org/wiki/Representational_state_transfer)

# Description
A modern solution to make users' voices matter!

# Problem Statement
In the Republic of Moldova, there is a scarcity of user-friendly platforms for individuals to express their opinions effectively. The existing websites are often poorly designed, inadequately developed, and do not sufficiently prioritize user preferences.

# Solution
To address this issue, we propose an enhanced platform for petitioning in Moldova. This solution incorporates robust security features, streamlined login processes, intuitive drafting tools, and efficient archiving mechanisms. Our platform is dedicated to facilitating every aspect associated with petitions and providing a seamless experience for expressing your opinion.

NOTE: As a MVP, the solution is aimed for the students of UTM as a priority.

## How to use
1. Clone the repository:
```shell
git clone https://github.com/catness812/e-petitions-project.git
```
2. In order to use the app, run the "docker compose up" in the root of your project. This will run the services inside a containerized environment. Requests should be sent to the Gateway service port (``http_port: :1337``). See Gateway Readme file for instructions.
```shell
docker compose up
```

## How to use (for regular users):
Still in development (currently finishing the front)

## Tech Stack

This project leverages the following technologies and frameworks:

- **Programming Language**: [GoLang](https://golang.org/) with [Gin-Gonic framework](https://github.com/gin-gonic/gin) for fast, feature-rich, and clean code.

- **Data Storage and Querying**: [PostgreSQL](https://www.postgresql.org/) with [GORM](https://gorm.io/) for Go (ORM) for efficient data management. [Redis](https://redis.io/) is used for OTP & Security Services, providing faster response times and time-dependent variables.

- **Deployment**: Utilizes [Docker](https://www.docker.com/) for seamless deployment. Ongoing work on cloud deployment with a domain registered at [epetitii.co](https://epetitii.co/) via [porkbun.com](https://porkbun.com/).

- **Protocols and Communication**: The gateway communicates with the web client (frontend) through a classic [REST API](https://en.wikipedia.org/wiki/Representational_state_transfer). Other services, including internal communication with the gateway, employ [gRPC](https://grpc.io/) for superior performance.

## Microservices Involved
As seen in the repo, there are 6 main services:
1. [Gateway](gateway) - for external communication
2. [Mail Service](mail_service) - for notifying users and admin about changes involving petitions and personal information.
3. [Petition Service](petition_service) - a service for managing all information referring to petitions.
4. [User Service](user_service) - same as the last one, but about users.
5. [Security Service](security_service) - this service is a middleware, responsible for secure registering and logging the user and other aspects.
6. [OTP Service](security_service) - OTP service is a part of security service, but it is responsible for signing a petition.
7. [Spam Filter](https://github.com/grumpycatyo-collab/spam_filter_epetitions.git) - yeah, there are 7 of them, not 6, because spam filer service is a middleware, from an external repo, that manages the censoring and grammar correction of information inputted by the user (it works in Python and communicating with the front via web-sockets).
   
NOTE: every microservice has its documentation so don't hesitate to click on it and see how it's working.

### Front Development
Frontend has a different repository developed in React TS, with Chakra UI. [Here](https://github.com/valeis/e_petitions_project_front.git) you can see more about it, and even launch a test to see how it's working.


### Contributors
![Contributors](https://img.shields.io/github/contributors/catness812/e-petitions-project)

Please note that not all contributors are active, some of them have left the team.
All contributions, bug reports, bug fixes, documentation improvements, enhancements, and ideas are welcome.

We don't have a contributing guide or a issues Git page, but we are on the way to make them.


