# VOAUTH
The Voauth application is a web-based tool designed to show developers how to create their own OAuth provider platforms, similar to those of Google, Twitter, Giithub, and more.



- [VOAUTH](#voauth)
- [SETUP](#setup)
  - [Backend](#backend)
  - [Frontend](#frontend)
   

# SETUP

## Backend
### Prerequisites

- Docker 
- Docker-compose

### Steps

```sh
cd backend

cp .env.example .env

make api
```


## Frontend
### Prerequisites

- Node
- Go
  
### Steps

```sh
cd frontend

npm i

cp .env.sample .env

npm run dev
```