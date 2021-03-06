# Quote-service

Quote service is a part of payment system which can be used to create quotes by users for money transactions

Altough this is my first Golang project after quite long time I have tried to structure the project as 'clean' as I can.

Reference: <https://nurcahyaari.medium.com/how-i-implement-clean-code-architecture-on-golang-projects-68be58830621>

If you want to get the project up and running:

## Minimal requirements

1) Golang installed and configured properly

2) PostgreSQL database server installed

## Steps to follow

1) Download project

2) Create Database with name of your choice

3) Create `.env` file in the project's root directory and add your database connection params into. Please refer to `.env.example` file for required params

4) Project entry point is main.go file. So, use `go run main.go` command

## Todo

- [ ] Add [swagger](https://github.com/swaggo/gin-swagger) API documentation
- [ ] Unit tests to prevent any bugs

## Steps to manually testing application

Firstly, create a user in the system via `POST` request to `http://localhost:5000/api/v1/users`. The endpoint expects 2 fields shown below as request body to create a new user profile:

```json
{
    "name":"Bakhtiyar Garashov",
    "email":"bakhtiyar.garashov@ut.ee"
}
```

Firing `GET` request to the same endpoint returns list of all users and their quote objects in nested structure. For example

```json
{
    "data": [
        {
            "id": 1,
            "name": "Bakhtiyar Garashov",
            "email": "bakhtiyar.garashov@gut.ee",
            "quotes": [
                {
                    "id": 1,
                    "currency_source": "EUR",
                    "currency_target": "AZN",
                    "amount": 1000,
                    "fee": 60.47,
                    "estimated_delivery_time": "2021-12-20T07:19:18+02:00",
                    "user_id": 1
                }
            ]
        }
    ],
    "message": "All users",
    "success": "true"
}
```

Next phase is creating a quote for already existing user profile. Use `POST` request to `http://localhost:5000/api/v1/quotes` for that. This part reflects 1st functionality of assignment. Request body should be in the structure below:

```json
{
    "source_currency":"EUR",
    "target_currency":"AZN",
    "amount":1000,
    "user_id":1
}
```

Note that if you use an user id which is not exist it will handle this situation gracefully.

According to the 2nd functionality it requires limiting the user creating no more than 10 quotes per minute. This part is a little bit challenging and can be solved in multiple ways. I believe [**API rate limiting**](https://nordicapis.com/everything-you-need-to-know-about-api-rate-limiting/) concept is one of the best solutions for a feature like that. So, I have used a rate limiting go [package](https://github.com/julianshen/gin-limiter). This is very small package and provides a comfortable way of defining new limiter. I have added it as a middleware and only being called by the controller which is responsible for creating a new quote in the system. Package provides a way to easily establish limiter based on any parameter you want. For example, user's IP, api_key or any other custom parameter. It is a perfect way to prevent abusing any endpoint.

For testing 2nd functionality, please hit quote creating endpoint for the same user (by using the same user_id in your request body) 10 times. When you hit request for 11th times it will return `429 Too many requests` error.

Furthermore, as a response for newly created quote the json schema as below will be returned:

```json
{
    "data": {
        "id": 1,
        "fee": 60.47,
        "estimated_delivery_time": "2021-12-20T07:19:18+02:00"
    },
    "message": "Quote created",
    "success": "true"
}
```

Side note: The properties `fee` and `estimated_delivery_time` were generated randomly by util functions as the application is not giving instruction about how to calculate those parameters.
