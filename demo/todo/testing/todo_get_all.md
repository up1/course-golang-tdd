# Manage TODO

* Root: "http://localhost:8080/"

## `POST /todo/`

Create new todo

### Example request

```json
{
  "title":    "TODO 01"
}
```

* `Content-Type`: "application/json"

===

### Example response

* `Status`: `200`
* `Content-Type`: `"application/json; charset=utf-8"`

```json
{
  "id":      0,
  "title":    "TODO 01",
  "done": false
}
```

## `GET` `/todo/1`

Get todo by id

===

* `Status`: `200`
* `Content-Type`: `"application/json; charset=utf-8"`

```json
{
  "id":      1,
  "title":    "XXXX",
  "done": true
}
```

## `DELETE /todo/1`

Delete todo by id

===

* `Status`: `200` // OK