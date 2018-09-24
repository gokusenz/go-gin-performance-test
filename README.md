# Service | Performance Criteria Test

## Introduction

We are always welcome new languages and new frameworks to be used in our microservice universe. Anyway, we need to define some criteria as our performance standard. So, If you are willing to use any of your interested frameworks just make a demo that can pass our criteria then you are most welcome :)


## Target & Objectives
- We are going to build a new service API using specific framework and language
- Our service has to connect with Fastwork staging database
- Our service has to provide an API route to get product by it’s ID
- Our service must response complete data regarding expected response below
- We will use `ab` (Apache benchmark) tool  for do a performance test


## Protocol
- HTTP-2 - REST API
  - We can test with `HTTP-1` instead for the easier setup process


## Request body
- Endpoint
  - domain/products/:id

- HTTP method
  - GET


## Response body
- no need to be exactly matched with example response below
```json
{
  "id": "uuid",
  "title": "title",
  "slug": "slug",
  "description": "description",
  "about_seller": "about_seller",
  "base_price": 0,
  "extras": "json_string",
  "extra_description": "extra_description",
  "images": [
    {
      "id": "uuid",
      "image_medium": "image_medium",
      "is_cover_photo": true
    }
  ]
}
```

## Test command

- Product ID: f967f5f1-1c68-459f-9674-cc3cb926ca4b
```bash
ab -c 100 -n 1000 http://localhost:4000/api/products/f967f5f1-1c68-459f-9674-cc3cb926ca4b
```
