# Tarantool team intern task

---
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

### Storage
Realised with use of pure Tarantool with http module.

Notes:
- entity 'id' hasn't been specificated, so id == key;
- since a key gets sent through url it can cosists from [a-z, A-Z, -, _] letters only;
- log writes to '/opt/tarantool/storage.log' without stdout mirroring;

To build it (tests-storage:latest) run <code>./control.sh -b</code> command.

### Tets
Golang application connects to a local storage (configures) 
and executes scripted test cases.

To execute tests run <code>./control.sh -t</code> command.

The command will deploy a helm package with a pod consits off stateless storage (no mounts) and tester images. Tester's ouput will be placed to the console.

To remove the package call <code>./control.sh -r</code>.
