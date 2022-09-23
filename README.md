# CPFD Data server
미세입자유동 연구실 데이터 서버입니다.

## Requirements
- app.env 파일을 프로젝트의 root 폴더에 생성합니다.
- app.env 파일에 다음 아래와 같이 입력합니다.
```text
MYSQL_PWD="your_mysql_pwd"
MYSQL_USR="your_mysql_usrname"
MYSQL_DB="your_mysql_database"
MYSQL_HOST="your_mysql_host"
MYSQL_PORT="your_mysql_port"
REDIS_HOST="your_redis_host"
REDIS_PWD="your_redis_pwd"
REDIS_DB="your_redis_database"
REDIS_PORT="your_redis_port"
```

## Run
Docker 컨테이너 혹은 Local에서 서버를 실행할 수 있는데 위의 app.env를 각 상황에 알맞게 설정해야합니다.

### Run in docker container
```bash
docker network create "your_network"
docker run --name "mysql_container_name" -e MYSQL_ROOT_PASSWORD="your_mysql_pwd" -d --network "your_network" -p "mysql_port":"mysql_port" mysql:5.7
docker run --name "redis_container_name" --network "your_network" -d -p "redis_port":"redis_port" redis --requirepass "your_redis_pwd"
docker build --tag "your_image_name" .
docker run --network "your_network" -p 8080:8080 "your_image_name"
```
### Run in local
```bash
go run cmd/cpfd/*
```

## Usage


|    Name                              | Method  |     Uri                      | Description                   |
|    :----:                            |  :----: |    :----:                    | :----:                        |
| Get machine                          | GET     | /machines                    | 현재 등록되어 있는 기기를 반환        |
| Create a machine                     | POST    | /machines                    | 기기 등록                       |
| Get activity logs                    | GET     | /logs/activity               | 활동 로그를 JSON으로 반환         |
| Get activity logs to csv file        | GET     | /logs/file/activity          | 활동 로그를 CSV 파일로 반환         |
| Create an activity log               | POST    | /logs/activity               | 활동 로그 생성                    |
| Get current activity                 | GET     | /logs/activity/state         | 현재 연구실 활동 상태를 반환          |
| Get particle logs                    | GET     | /logs/particle               | 관측 미세먼지 농도를 JSON으로 반환   |
| Get particle logs to csv file        | GET     | /logs/file/particle          | 관측 미세먼지 농도를 CSV 파일로 반환 |
| Creat a particle log                 | POST    | /logs/particle               | 관측 미세먼지 농도 로그 생성        |
| Creat an indoor-property log         | POST    | /logs/indoor-property        | 실내 관측 데이터 로그 생성         |
| Get indoor-property logs             | GET     | /logs/indoor-property        | 실내 관측 데이터를 JSON으로 반환    |
| Get indoor-property logs to csv file | GET     | /logs/file/indoor-property   | 실내 관측 데이터를 CSV 파일로 반환   |


## 기술 스택
- gRPC
- protobuf
- MySQL
- Redis
- gin
- gorm
- viper
- logrus