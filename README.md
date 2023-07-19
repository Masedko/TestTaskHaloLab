# Go RESTful API Test Task for Halo Lab 

## Run using docker-compose

```shell
git clone https://github.com/Masedko/TestTaskHaloLab.git

cd TestTaskHaloLab

mv .env.example .env  # This will work on Linux and MacOS, for Windows simply use ren .env.example .env

docker build -t gotesttaskhalolab:1.0 -f Dockerfile .

docker-compose up -d
```

## Access Swagger docs

```shell
http://localhost:8080/
```


## Data generating configuration

You can also edit data generating configurations via `internal/data/static` files.
