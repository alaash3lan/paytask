## paytask

- Build and run:

```bash
$ go mod download
$ go run .
```

## APIs

```http
GET localhost:8000/accounts/
```
List all account data
## Responses

```javascript
{
  "status" : int,
  "success" : bool,
  "message"    :
    {
        "id": string,
        "name": string,
        "balance": string
    }
}
```

```http
GET localhost:8000/accounts/{id}
```
Get account data

| Parameter | Type | Description |
|:----------| :--- |:------------|
| `id`      | `string` | Account id  |

## Responses

```javascript
{
    "status" : int,
        "success" : bool,
        "message"    :
    {
        "id": string,
        "name": string,
        "balance": string
    }
}
```

```http
POST localhost:8000/transaction}
```
Transfer balance

| Parameter | Type     | Description              |
|:----------|:---------|:-------------------------|
| `from`    | `string` | Account id               |
| `to`      | `string` | Account id               |
| `amount` | `float`  | blance will be transferd |

## Responses

```javascript
{
    "status" : int,
    "success" : bool,
    "error" : string
}
```