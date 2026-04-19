# 99-technical-assignment

## Requirement
- make
- Go Version 1.24+

## How To Run
```shel
make run
```

## Create User
```curl
curl --location 'localhost:8080/users' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=Willy'
```

## Get All User
```curl
curl --location 'localhost:8080/users?page_num=1&page_size=10'
```

## Get Specific User
```curl
curl --location 'localhost:8080/users/1'
```