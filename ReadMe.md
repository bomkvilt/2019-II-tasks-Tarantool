# Tarantool team intern task

## Task

Requirements:
- download/build Tarantool
- launch test application
- realise http-accesable key-value storage
- publish the solution to the Github
- deploy the solution in a public cloud

Additional requirements:
- all actions must be logged

Storage's API:

| method | uri      | body |
| ---    | ---      | ---  |
| POST   | /kv      | {"key": "test", "value": {non-specified json}} |
| PUT    | /kv/{id} | {"value": {non-specified json}} |
| GET    | /kv/{id} | - |
| DELETE | /kv/{id} | - |

- POST: 409 if the key already exists
- POST, PUT: 400 if a body is incorrect
- PUT, GET, DELET: 404 if the key doesn't exist

---
## Solution

Solution consists of two docker containers:
- KV-storage (/storage)
- API test   (/test)

The containers could be built with <code>./control.sh -b</code> command.

# TODO::

### Storage
Realised with use of NGinx + Tarantool upstream module + Tarantool.


### Tets
Scripted and randomised api tests. 
