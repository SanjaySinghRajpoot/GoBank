# GoBank 

A backend banking service in GoLang. 
- A scalable banking service in GoLang. Users can Create Accounts, transfer money and disable bank account. 
- RESTful HTTP APIs using Gin framework. 
- Robust Unit testing and app configs for deployment. Authenticating users, and securing the APIs with JWT and PASETO 
- CI/CD to build and push latest docker images from Container Registry to kubernetes Cluster on AWS. 
- Run entire application using a single docker command 
- Applied MVC architecture and followed best practices to easily maintain the code 
- Test cases for Unit testing and e2e testing
- Swagger docs added for APIs 
- gRPC APIs used for low latency and high throughput communication. 
- Secured APIs using JWT and Paseto token Authentication, Refresh and auth token feature added 


# Tech Stack 
- Backend - Go, Gin, Golang, PSQL 
- Tools - Docker, Kubernetes, JWT, Paseto 
- Database - Postgres 
- Cloud - AWS EKS, AWS ECR, Github Action, 

# Installation 
You can easily setup the application with the help of docker file. Just clone the repo and run `docker compose up` from the root folder.