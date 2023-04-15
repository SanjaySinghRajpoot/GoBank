# GoBank 

### Gobank is a scalable backend banking service developed in GoLang. The service allows users to create accounts, transfer money, and disable bank accounts through RESTful HTTP APIs built using the Gin framework. The following technical details are included in Gobank:

### Architecture and Deployment
The application uses the Model-View-Controller (MVC) architecture to facilitate easy code maintenance. The application is configured for deployment with robust unit testing and application configurations. To deploy the latest Docker images, a CI/CD pipeline is used to push images from the Container Registry to the Kubernetes cluster on AWS.

### User Authentication and API Security
Gobank uses JWT and PASETO tokens to authenticate users and secure the APIs. User authentication is implemented using the JWT token while PASETO tokens are used for API security. The application features token authentication, refresh, and auth token features for improved security.

### Testing and Documentation
Gobank has robust test cases for unit testing and end-to-end (e2e) testing. The API documentation is available through Swagger docs, which makes it easy to explore and test the APIs. Gobank also uses gRPC APIs for low-latency and high-throughput communication.


# Tech Stack 
- Backend - Go, Gin, Golang, PSQL 
- Tools - Docker, Kubernetes, JWT, Paseto 
- Database - Postgres 
- Cloud - AWS EKS, AWS ECR, Github Action, 

Features:

- Account Creation: Users can create their bank accounts by providing personal information and login credentials.
- Money Transfer: Users can transfer money to other bank accounts by providing the recipient's bank details and the amount to be transferred.
- Disable Bank Account: Users can disable their bank account in case of any fraudulent activity.
- RESTful HTTP APIs: The service provides easy-to-use RESTful HTTP APIs for seamless integration with frontend applications.
- Robust Unit Testing: Gobank has a comprehensive unit testing suite to ensure error-free execution and performance.
- App Configs: The service provides various app configs for easy deployment in production environments.
- Authentication and Authorization: Gobank uses JWT and Paseto tokens to authenticate users and secure the APIs from unauthorized access.
- CI/CD: Gobank uses Github Actions for continuous integration and deployment. It builds and pushes the latest Docker images to the Kubernetes cluster on AWS EKS.
- MVC Architecture: Gobank follows the MVC (Model-View-Controller) architecture to maintain the code easily and improve its maintainability.
- Swagger Docs: The service provides Swagger documentation for APIs for easy reference and integration.
- Secure APIs: Gobank uses JWT and Paseto tokens for authentication and authorization, and it also provides token refresh and auth token features for enhanced security.

# Installation 
You can easily setup the application with the help of docker file. Just clone the repo and run `docker compose up` from the root folder.
