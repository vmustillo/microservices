# Credit Card Sample Microservices

- This project contains an Account Service, Card Service, Auth Service, Payment Service, and Gateway Service
- All services are talking to postgres containers to store their data, and the services all consume GRPC.
- The Gateway service is used as a rest Endpoint, which uses GRPC-Gateway to translate REST to GRPC and routes the request to the correct microservice
- All services contain Makefiles to generate the needed protobuffer code as well as swagger docs

### Account Service

- Holds information regarding an account, such as ID, Credit Limit, Balance, and Owner ID
- The Account Service can be queried to create an account or retrieve an account based on ID

### Auth Service

- Holds information regarding users such as username, password, Name, and account ID
- This service can be used to create a user or to login
- Login returns a signed JWT
- When creating a user, the password is hashed with a random 32 byte string using the Argon 2 Algorithm. The hashed password and salt are stored in the database

### Card Service

- Holds the user, card number, and account number that is associated to the card
- Service can be used to create a new card or get a card
- Card numbers all contain the same first four numbers and the last 12 are random. All card numbers are checked in the DB to make sure they are unique before they are assigned

### Payment Service

- Can be used to make a payment
- The payment service uses goroutines to make requests to the card service and account service before approving. The requests make sure the card exists and the amount of the purchase will not push the user over their credit limit
