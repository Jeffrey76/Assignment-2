# Assignment-2 

## Edufi 21. Messaging
## 1. Design consideration of your microservices

Following the design of microservices as a collection of small loosely coupled services to work together as an application,
In order to fufil the service of providing Messaging services for Students and Tutors, 2 Microservices were created in the development phase to be able to interact with each other and allow the conversation api to call onto replies api in order to get all messages in a conversation and show it to the user in a user intuitive manner 

These 2 microservices are defined through its business capabilities and functional area 
Conversation microservice to maintain and provide services related to Conversations
Replies microservice to maintain and provide services related to Replies
The User Interface (UI) will use these 2 APIs to get data from the databases to provide a web application to interact with the data directly and easily.

The microservices will provide services through the implementation of REST and REST APIs to transfer data in json format
between the front-end UI and the microservices.


Conversation API will run on localhost:9155 while Replies API will run on localhost:9156 to provide services to the UI
The UI will run on localhost:9150

The UI also makes use of other APIs that are developed externally, which includes student and tutors to get the usernames in order for the user to create new conversation based on their names instead of IDs.
 

## 2. Architecture diagram




## 3. Instructions for setting up and running your microservices
### Importing of additional packages
```sh
go get -u github.com/mitchellh/mapstructure
```

### Create Database
SQL Database Server Connection access set to localhost:3306
In Docker, Database Connection access set to Database15 Container within the same Docker Network
Run "Messages Setup.sql" to create Database Edufi with tables Conversation and Replies

### Running of applications (Conversation API, Repliesapi API, UI)
```sh
cd ../conversationapi
go run main.go
cd repliesapi
go run main.go
cd ../ui
go run main.go
```
