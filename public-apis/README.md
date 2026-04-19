# 99-technical-assignment

## Requirement
- make
- Go Version 1.24+

## How To Run
```shel
make run
```

## Get All Listing
```curl
curl --location 'localhost:8088/public-api/listings?page_num=1&page_size=10'
```

## Create User
```curl
curl --location 'localhost:8088/public-api/users' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Messi"
}'
```

## Create Listing
```curl
curl --location 'localhost:8088/public-api/listings' \
--header 'Content-Type: application/json' \
--data '{
    "user_id": 2,
    "listing_type": "rent",
    "price": 4500
}'
```