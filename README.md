# Virtual Vending Machine
This application is to support one to checkout electronic items from VMware's vending machines

## Dependencies
Run postgres container in your local environment. More on this [here](https://hub.docker.com/_/postgres).
```
docker run --name <container-name> -e POSTGRES_USER=<default-user> -e POSTGRES_PASSWORD=<default-password> -d <default-database>
```

Run following commands to set up your postgres to support the execution

- Get into docker container. **container-name** is the value provided for --name field when running docker container 
    ```
    docker exec -it <container-name> bash
    ```
- Connect to your db. **default-user** and **default-database** should be same as those provided for running docker container
    ```
    psql <default-user> -d <default-database>
    ```
- Execute sql schema statements
    1. create enums for role column of users table
    ```
        create type virtual_vending.UserRolesEnum as enum ('manager', 'engineer', 'admin');
    ```
    2. create user table   
    ```
    create table if not exists virtual_vending.users (
        id serial primary key,
        username varchar(63) not null,
        email varchar(63) not null,
        password varchar not null,
        address varchar,
        role userrolesenum not null,
        created_at timestamp default current_timestamp,
        update_at timestamp default current_timestamp,
        token varchar,
        constraint unique_username unique (username)
    ); 
    ```
    3. create products table   
    ```
    create table if not exists virtual_vending.products (
        id serial primary  key,
        name varchar not null,
        description varchar not null,
        cost int
        created_at timestamp default current_timestamp,
        updated_at timestamp default current_timestamp,
        constraint unique_product_name unique (name)
    );
    ```
  
 
## Backend Installation
1. **Creating docker image**

    Execute docker file to create backend docker image in the root directory
    ```
    docker build -t <image-name>:<image-tag> .
    ```
    NOTE: If image-tag is not provided, the latest will be taken as image tag by default
2. **Running docker image**

    ```
   docker run --name <container-name> -p 8080:8080 -e POSTGRES_DB=<default-database> -e POSTGRES_USER=<default-user> -e POSTGRES_PASSWORD=<default-password> -e POSTGRES_HOST=<docker-host-ip> <image-name>:<image-tag>
   ```
   1. *default-database*, *default-user* , *default-password* should be the one you provided while running postgres container
   2. *docker-host-ip* should be the ip of the docker host where the postgres is running
   3. *image-name* and *image-tag* should be those values that you provided when running 1 step of Backend Installation
   
   
## Available APIs
1. *Creating user*: POST
    ```/api/v1/user```

    Body:
    ```{
    "name": "Navami",
     "username": "kinin",
     "password": "VMware123!",
     "role": "admin",
     "address": "1774 JP Nagar 4th phase",
     "email": "kinin@vmware.com"
    }```
2. *Get all users*: GET
    ```/api/v1/user```
3. *Get user by username*: GET
    ```/api/v1/user/<username>```
4. *Delete user*: DELETE
    ```/api/v1/user/<username>```
5.*Create product*: POST
    ```/api/v1/product```

    Body:
    ```
   {
        "name":"macbook pro",
        "cost":"200000",
        "description":macbook pro"
    }
   ```

6. *login*: POST
    ```/login```
    Body:
    ```
    {
        "username": "kinin",
        "password": "kinin"
    }
   ```
   Response:
   ```
    {
        "id": 0,
        "created_at": "0001-01-01T00:00:00Z",
        "update_at": "0001-01-01T00:00:00Z",
        "name": "",
        "username": "madhava",
        "password": "",
        "role": "",
        "address": "",
        "email": "",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1hZGhhdmEifQ.QprsO-FkFhOUjiOsMfY8Ik5NNfQOnX9DVf8gT2xikpw"
    }
    ```
   
   **NOTE**: Use the token generated in the above response, in header for all other calls to backend including GET calls. 
   ex:
   ```
   "Authorization":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1hZGhhdmEifQ.QprsO-FkFhOUjiOsMfY8Ik5NNfQOnX9DVf8gT2xikpw"
   ```
   

