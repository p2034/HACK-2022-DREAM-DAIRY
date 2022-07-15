**/register** (POST)

Request:
```json
{
  "username": "Steve",
  "password": "my_password_12345",
  "email": "steve@mail.com",
  "token": "fdsfdfsdfs9832r314gbt13..."
}
```
Answer:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78..."
}
```

**/login** (GET)

Request:
```json
{
  "email": "steve@mail.com",
  "password": "my_password_12345"
}
```
Answer:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78..."
}
```

**/logout** (POST)

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78..."
}
```

**/checktoken** (GET)
Just for another server, but may be used in frontend.

Request:
```json
{
  "userid": 1234,
  "token": "egy4983fbi3564342i78..."
}
```