# Valhalla project API 📚

This valhalla API is responsible for performing operations related to project. It is a standalone service so that it can be scalable and maintainable without causing problems for maintainers.

# Start database

To start this API, you will need to have the database up and running. To do so, go to the database directory in this proyect and create a .env file with the following parameters.

```bash
IP_MONGODB=mongodb-ip           # your mongodb docker ip
MASK_MONGODB=ip-mask            # your mongodb ip mask

MONGO_ADMIN_USERNAME=username   # your mongodb username
MONGO_ADMIN_PASSWORD=password   # your mongodb password
MONGO_ADMIN_SERVER=serverhost   # your mongodb server host
MONGO_DATA_PATH=data-path       # data path to persist mongodb

MONGO_WEB_USERNAME=username     # username for mongo-express web access
MONGO_WEB_PASSWORD=password     # password for mongo-express web access
```

then start the database and express instance with the following command:

```bash
docker compose up 
```

# Prepare API

Just create a env file with the following params:

```shell
# API common data
IP=127.0.0.1        # 0.0.0.0 if it is in a docker container
PORT=4444           # the port you want the api to start in
VERSION=v1          # the API version
API_NAME=project    # the API name
ENV=release         # debug mode
SECRET=!Sup3rs3cr3t # secret for cryptography

# Mongo database configuration
MONGO_USER=user             # your mongodb username
MONGO_PASSWORD=password     # your mongodb password
MONGO_SERVER=serverhost     # your mongodb server host
MONGO_PORT=mongo-port       # your mongodb port
```

# Start API with docker

To start the API from the docker container just run

```shell
docker compose up
```

# Start API with local go

To start the API from go just run

```shell
sh start.sh
```


