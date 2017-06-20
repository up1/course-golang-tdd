# TODO

## GET /todo/
ดึงข้อมูลของ todo ทั้งหมด

===
### Response

* Status: 200
* Content-Type: "application/json; charset=utf-8"
* Data[0].id: 1
* Data[0].title: "Todo 1"

```json
[
{"id":1,"title":"Todo 1","done":false},
{"id":2,"title":"Todo 2","done":false},
{"id":3,"title":"Todo 3","done":false}]
```

## POST /todo/
ทำการสร้าง TODO ใหม่

### Request

```json
{
  "title":    "TODO 01"
}
```

* Content-Type: "application/json"

===
### Response

* Status: 200
* Content-Type: "application/json; charset=utf-8"

```json
{
  "id": 0,
  "title": "TODO 01",
  "done": false
}
```