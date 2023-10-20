# Description
Petition is a service created for the e-petitions project to implement their data processing.

## Contents
- [Coneventions](#coneventions)
- [Rpc Methods](rpc-methods)
- [CreatePetition](#createpetition)
- [GetPetitionById](#getpetitionbyid)
- [GetPetitions](#getpetitions)
- [UpdatePetitionStatus](#updatepetitionstatus)
- [DeletePetition](#deletepetition)
- [ValidatePetitionId](#validatepetitionid)
- [CreateVote](#createvote)
- [GetUserPetitions](#getuserpetitions)
- [GetUserVotedPetitions](#getuservotedpetitions)
- [CheckIfPetitionsExpired](#checkifpetitionsexpired)
- [GetAllSimilarPetitions](#getallsimilarpetitions)
- [SearchByPetitionsTitle](#searchbypetitionstitle)

## Coneventions
Petition is developed to use the gRPC protocol and will use messages from predefined proto files (internal/proto folder).

## Rpc Methods
### CreatePetition
Check if the user exists in the database. If it exists, it validates the data and if everything is OK, the petition is created in the database.
#### Parameters
- `CreatePetitionRequest` is a message that contains the information needed to create a petition. It should contain an email stored in database.
```grpc
message CreatePetitionRequest {
  string title = 1;
  string description = 2;
  string image = 3;
  uint32 user_id = 4;
  string category = 5;
  uint32 vote_goal = 8;
  google.protobuf.Timestamp exp_date = 9;
}
```
- `PetitionId` is a message that is returned after a successful creation of the petition. It contain the petition id.
```grpc
message PetitionId {
  uint32 id = 1;
}
```
#### Rquest
```shell
grpcurl -plaintext -d '{"title":"title","description":"desc","image":"image","user_id":user_id,"category":"cat","vote_goal":vote_goal,"exp_date":"exp_date"}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.CreatePetition
```
#### Response
```json
{
    "id": 1
}
```

### GetPetitionById
Check if petition id exist in database. If it exists, it return thei data.
#### Parameters
- `PetitionId` is a message that contains the petition id.
```grpc
message PetitionId {
  uint32 id = 1;
}
```
- `Petition` is a message that is returned after a successful find in database of id. It contains the petition data.
```grpc
message Petition {
  uint32 id = 1;
  string title = 2;
  string category = 3;
  string description = 4;
  string image = 5;
  Status status = 6;
  uint32 user_id = 7;
  uint32 vote_goal = 8;
  uint32 current_votes = 9;
  google.protobuf.Timestamp exp_date = 10;
  google.protobuf.Timestamp updated_at = 11;
  google.protobuf.Timestamp created_at = 12;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"id": id}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.GetPetitionById
```
#### Response
```json
{
  "id": 1,
  "title": "title",
  "category": "category",
  "description": "description",
  "image": "iamage",
  "status": {
    "id": 1,
    "title": "status_title"
  },
  "userId": 1,
  "voteGoal": 1000,
  "expDate": "exp_date",
  "updatedAt": "updated_date",
  "createdAt": "created_date"
}
```

### GetPetitions
Check if exist any petitions and return them.
#### Parameters
- `GetPetitionsRequest` is a message that contains pages number and limit of petition on page.
```grpc
message CreateVoteRequest{
  uint32 user_id = 1;
  uint32 petition_id = 2;
}
```
- `GetPetitionsResponse` is a message that is returned after find of petitions in database. It contains array of petitions.
```grpc
message GetPetitionsResponse {
  repeated Petition petitions = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"limit":limit, "page":page}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.GetPetitions
```
#### Response
```json
{
  "petitions": [
    {
        "id": 1,
        "title": "title",
        "category": "category",
        "description": "description",
        "image": "iamage",
        "status": {
            "id": 1,
            "title": "status_title"
        },
        "userId": 1,
        "voteGoal": 1000,
        "expDate": "exp_date",
        "updatedAt": "updated_date",
        "createdAt": "created_date"
    },
    "..."
  ]
}
```

### UpdatePetitionStatus
Check if the petition id exist in database. If it exist, update their status.
#### Parameters
- `UpdatePetitionStatusRequest` is a message that contains the petition id and status.
```grpc
message UpdatePetitionStatusRequest{
  uint32 id = 1;
  string status = 2;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"id": id, "status": "petition_status"}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.UpdatePetitionStatus
```
#### Response
Response is empty.

### DeletePetition
Check if the petition id exist in database. If it exist, delete them.
#### Parameters
- `PetitionId` is a message that contains the petition id.
```grpc
message PetitionId{
  uint32 id = 1;
}
```
#### Rquest
```shell
grpcurl -plaintext -d '{"id": id}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.DeletePetition
```
#### Response
Response is empty

### ValidatePetitionId
Check if petition exist in database. If it do not exist, return error.
#### Parameters
- `PetitionId` is a message that contains the petition id.
```grpc
message PetitionId{
  uint32 id = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"id": id}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.ValidatePetitionId
```
#### Response
Response is empty


### CreateVote
Check if petition exist in database. If it exist, check if user exist and has not voted. If it exist and has not voted, add a vote to petition to the current votes.
#### Parameters
- `CreateVoteRequest` is a message that contains the petition id.
```grpc
message CreateVoteRequest{
    uint32 user_id = 1;
    uint32 petition_id = 2;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"petition_id": petition_id, "user_id": user_id}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.CreateVote
```
#### Response
Response is empty.

### GetUserPetitions
Return all petitions created by user.
#### Parameters
- `GetUserPetitionsRequest` is a message that contains user id, pages number and limit of petition on page.
```grpc
message GetUserPetitionsRequest {
    uint32 user_id = 1;
    uint32 page = 2;
    uint32 limit = 3;
}
```
- `GetUserPetitionsResponse` is a message that is returned after find petitions created by user in database. It contains array of petitions.
```grpc
message GetUserPetitionsResponse {
    repeated Petition petitions = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"limit":limit, "page":page, "user_id":user_id}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.GetUserPetitions
```
#### Response
```json
{
    "petitions": [
        {
            "id": 1,
            "title": "tite",
            "category": "category",
            "description": "description",
            "image": "image",
            "status": "status",
            "user_id": 1,
            "vote_goal": 1000,
            "current_votes": 0,
            "exp_date": "exp_date",
            "updated_at": "update_date",
            "created_at": "created_date"
        }
    ]
}
```


### GetUserVotedPetitions
Return all petitions voted by this user.
#### Parameters
- `GetUserVotedPetitionsRequest` is a message that contains user id, pages number and limit of petition on page.
```grpc
message GetUserVotedPetitionsRequest {
  uint32 user_id = 1;
  uint32 page = 2;
  uint32 limit = 3;
}
```
- `GetUserVotedPetitionsResponse` is a message that is returned after find petitions voted by user in database. It contains array of petitions.
```grpc
message GetUserPetitionsResponse {
  repeated Petition petitions = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"limit":limit, "page":page, "user_id":user_id}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.GetUserVotedPetitions
```
#### Response
```json
{
    "petitions": [
        {
            "id": 1,
            "title": "tite",
            "category": "category",
            "description": "description",
            "image": "image",
            "status": "status",
            "user_id": 1,
            "vote_goal": 1000,
            "current_votes": 0,
            "exp_date": "exp_date",
            "updated_at": "update_date",
            "created_at": "created_date"
        }
    ]
}
```

### CheckIfPetitionsExpired
Check that the expiration date has not passed. If it passed, delete petition.
#### Parameters
- `Petition` is a message that contains the information of petition.
```grpc
message Petition {
  uint32 id = 1;
  string title = 2;
  string category = 3;
  string description = 4;
  string image = 5;
  Status status = 6;
  uint32 user_id = 7;
  uint32 vote_goal = 8;
  uint32 current_votes = 9;
  google.protobuf.Timestamp exp_date = 10;
  google.protobuf.Timestamp updated_at = 11;
  google.protobuf.Timestamp created_at = 12;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"id":id,"title":"title","category":"cat","description":"desc","image":"image","status":{"id":status_id,"title":"status_title"},"user_id":user_id,"vote_goal":vote_goal,"current_votes":current_votes,"exp_date":"exp_date","updated_at":"updated_date","created_at":"created_date"}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.CheckIfPetitionsExpired
```
#### Response
Response is empty.

### GetAllSimilarPetitions
Check if in database exist similar petitions and return them.
#### Parameters
- `PetitionSuggestionRequest` is a message that contains title of petition.
```grpc
message PetitionSuggestionRequest {
  string title = 1;
}
```
- `PetitionSuggestionResponse` is a message that is returned after find of similar petitions. It contain array of PetitionInfo
```grpc
message PetitionInfo {
    uint32 id = 1;
    string title = 2;
    uint32 user_id = 3;
}

message PetitionSuggestionResponse {
    repeated PetitionInfo SuggestedPetitions = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"title": "title"}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.GetAllSimilarPetitions
```
#### Response
```json
{
    "SuggestedPetitions": [
        {
            "id": 1,
            "title": "title",
            "user_id": 1
        },
        "..."
    ]
}
```

### SearchByPetitionsTitle
Check if petition title exist in database. If it exist, return them.
#### Parameters
- `SearchPetitionsByTitRequest` is a message that contains title of petition, page number and limit of petition on page.
```grpc
message SearchPetitionsByTitRequest {
  string title = 1;
  uint32 page = 2;
  uint32 limit = 3;
}
```
- `PetitionSuggestionResponse` is a message that is returned after find of similar petitions. It contain array of PetitionInfo
```grpc
message PetitionInfo {
    uint32 id = 1;
    string title = 2;
    uint32 user_id = 3;
}

message PetitionSuggestionResponse {
    repeated PetitionInfo SuggestedPetitions = 1;
}
```
#### Request
```shell
grpcurl -plaintext -d '{"limit":1, "page": 1, "title": "title"}'\
 -import-path internal/proto   -proto internal/proto/petition_svc.proto \
 localhost:50050 proto.PetitionService.GetAllSimilarPetitions
```
#### Response
```json
{
    "SuggestedPetitions": [
        {
            "id": 1,
            "title": "title",
            "user_id": 1
        },
        "..."
    ]
}
```