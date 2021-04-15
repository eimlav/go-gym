# go-gym

Basic Gym class booking API built in Go.

GORM and a SQLite database (`go-gym.db`) is used by the application.

## Setup

Run `make build` to build the application.

## Running

Run `make run` to run the application.

## Testing

Run `make test` to run tests.

## Config

You can configure go-gym using the config.yaml file:

```
server:
    address: <IP address of API server>
    port: <Port to expose API server on>
```

This file contains preconfigured default values to get up and running quickly.

## Endpoints

### `/api/v1/classes` POST

Create a class and associated class events.

#### Body Parameters

- `name` | string | required,min=2,max=128
- `start_date` | time.Time | required time_format:"2006-01-02T15:04:05Z07:00"
- `name` | time.Time | required time_format:"2006-01-02T15:04:05Z07:00"
- `capacity` | int | required gt=0

#### Responses

- `201 Created` Class was created successfully
- `400 Bad Request` Paramters supplied were invalid.
- `500 Internal Server Error` Unexpected server error.

### `/api/v1/bookings` POST

Create a booking record.

#### Body Parameters

- `member_name` | string | required,min=2,max=128
- `class_event_id` | int | required

#### Responses

- `201 Created` Booking was created successfully
- `400 Bad Request` Paramters supplied were invalid.
- `500 Internal Server Error` Unexpected server error.

## Database Schema

The schema for the database can be found in `db/models/models.go`.

## Examples

### Create a class

```
⇒  curl -i -X POST -H "Content-Type:application/json"  -d '{"name":"Yoga","start_date":"2021-04-02T15:00:00+00:00","end_date":"2021-04-04T15:00:00+00:00","capacity":0}' 0.0.0.0:8080/api/v1/classes
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Thu, 14 Apr 2021 20:46:15 GMT
Content-Length: 8

{"id":9}
```

### Create a booking

```
⇒  curl -i -X POST -H "Content-Type:application/json"  -d '{"member_name": "bob", "class_event_id":17}' 0.0.0.0:8080/api/v1/bookings
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Thu, 15 Apr 2021 17:52:24 GMT
Content-Length: 8

{"id":6}
```
