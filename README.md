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

### Running of applications in local machine (Conversation API, Repliesapi API, UI)
```sh
cd ../conversationapi
go run main.go
cd repliesapi
go run main.go
cd ../ui
go run main.go
```


### Running of applications in Docker (Conversation API, Repliesapi API, UI)
```sh
docker pull jeffrey76/userinterface:latest
docker pull jeffrey76/conversationapi:latest
docker pull jeffrey76/repliesapi:latest
docker pull jeffrey76/database:latest
docker container run --detach --name database15 -e MYSQL_ROOT_PASSWORD=password -p 3306 --volumes-from databasedata15 jeffrey76/database
docker container run --name repliesapi15 -it -p 9156:9156 jeffrey76/repliesapi
docker container run --name conversationapi15 -it -p 9155:9155 jeffrey76/conversationapi
docker container run --name ui15 -it -p 9150:8000 jeffrey76/userinterface
```

