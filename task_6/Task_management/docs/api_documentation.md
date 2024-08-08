# API Documentation

## Create Task

**Endpoint**: `POST /tasks/`

**Middleware**:
- **AuthMiddleware**: Validates the JWT token from the `Authorization` header.
- **AdminMiddleware**: Ensures the user has an admin role.

### Request

**Headers**:
- `Authorization`: Bearer token for authentication.

**Body** (JSON):
```json
{
  "id": 1,
  "username": "exampleUser",
  "title": "Task Title",
  "description": "Task Description",
  "completed": false
}
```

  ### Responses

   -`200 OK`: Task created successfully.

```json

{
  "message": "success"
}
```

-`400 Bad Request`: Invalid request payload.

```json

{
  "message": "status bad request"
}
```
  -`401 Unauthorized`: Unauthorized access or invalid token.

```json

{
  "message": "unauthorized access"
}
```

  -`404 Not Found`: User not found or task creation failed.

```json

{
  "message": "page not found"
}
```

## Get All Tasks

**Endpoint**: GET /tasks/getall

**Middleware**:

  **AuthMiddleware**: Validates the JWT token from the Authorization header.

### Request

**Headers**:

  **Authorization**: Bearer token for authentication.

**Responses**

  -`200 OK`: Successfully retrieved tasks.

```json

[
  {
    "id": 1,
    "username": "exampleUser",
    "title": "Task Title",
    "description": "Task Description",
    "completed": false
  },
  {
    "id": 2,
    "username": "anotherUser",
    "title": "Another Task",
    "description": "Another Description",
    "completed": true
  }
]
```
  -`401 Unauthorized`: Invalid token or unauthorized access.

```json

{
  "message": "unauthorized access"
}
```
  -`04 Not Found`: No tasks found or user not found.

```json

{
  "message": "not found"
}
```

## Update Task by ID

**Endpoint**: PATCH /tasks/update/:id

**Path Parameters**:

    id (int): The ID of the task to be updated.

Body (JSON):

```json

{
  "id": 1,
  "username": "exampleUser",
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "completed": true
}
```

### Responses

  -`200 OK`: Successfully updated the task.

```json

{
  "message": "successfully updated"
}
```

  -`400 Bad Request`: Invalid input or request payload.

```json

{
  "message": "invalid input"
}
```

  - `404 Not Found`: Task with the given ID not found or update failed.

```json

{
  "message": "task not found"
}
```

## Delete Task by ID

**Endpoint**: DELETE /tasks/delete/:id

**Path Parameters**:

    id (int): The ID of the task to be deleted.

### Responses

  -`200 OK`: Successfully deleted the task.

```json

{
  "message": "deleted"
}
```

  - `400 Bad Request`: Invalid ID or request payload.

```json

{
  "message": "invalid input"
}
```

  - `404 Not Found`: Task with the given ID not found or deletion failed.

```json

{
  "message": "task not found"
}
```

  - `500 Internal Server Error`: Server error during task deletion.

```json

{
  "message": "internal server error"
}
```


## Register User

**Endpoint**: `POST /users/register`

### Request

**Body** (JSON):
```json
{
  "username": "exampleUser",
  "password": "examplePassword",
  "role": "user"
}
```

  ### Responses

  -`200 OK`: User registered successfully.

```json

{
  "message": "registration successful"
}
```

  -`400 Bad Request`: Invalid request payload or missing fields.

```json

{
  "message": "bad request"
}
```
   -`500 Internal Server Error`: Server error during registration.

```json

{
  "message": "server error"
}
```
## Login

**Endpoint**: POST /users/login
## Request

Body (JSON):

```json

{
  "username": "exampleUser",
  "password": "examplePassword"
}
```
## Responses

  - `200 OK`: Successful login and returns a token.

```json

{
  "message": "login successful",
  "token": "your-jwt-token"
}
```
  - `400 Bad Request`: Invalid login credentials or request payload.

```json

{
  "message": "invalid credentials"
}
```

  - `500 Internal Server Error`: Server error during login.

```json

{
  "message": "server error"
}
```
## Promote User

**Endpoint**: PATCH /users/promote/:username

**Middleware**:

  **AuthMiddleware**: Validates the JWT token from the Authorization header.
  **AdminMiddleware**: Ensures the user has an admin role.

## Request

**Headers**:

  **Authorization**: Bearer token for authentication.

**Path Parameters**:

    username (string): The username of the user to be promoted.

## Responses

  - `200 OK`: User promoted successfully.

```json

{
  "message": "user promoted"
}
```
  - `400 Bad Request`: Invalid username or request payload.

```json

{
  "message": "invalid username"
}
```
  -`401 Unauthorized`: Unauthorized access or invalid token.

```json

{
  "message": "unauthorized access"
}
```
  -`404 Not Found`: User not found or promotion failed.

```json

{
  "message": "user not found"
}
```
### Create Task

**Endpoint**: POST /tasks/


## Request

**Headers**:

    *Authorization*: Bearer token for authentication.

Body (JSON):

```json

{
  "id": 1,
  "username": "exampleUser",
  "title": "Task Title",
  "description": "Task Description",
  "completed": false
}
```

### Responses

  -`200 OK`: Task created successfully.

```json

{
  "message": "success"
}
```
  -`400 Bad Request`: Invalid request payload.

```json

{
  "message": "status bad request"
}
```
  -`401 Unauthorized`: Unauthorized access or invalid token.

```json

{
  "message": "unauthorized access"
}
```
  - `404 Not Found`: User not found or task creation failed.

```json

{
  "message": "page not found"
}
```
## Get All Tasks

**Endpoint**: GET /tasks/getall

**Middleware**:

  **AuthMiddleware**: Validates the JWT token from the Authorization header.

### Request

**Headers**:

  **Authorization**: Bearer token for authentication.

### Responses

  -`200 OK`: Successfully retrieved tasks.

```json

[
  {
    "id": 1,
    "username": "exampleUser",
    "title": "Task Title",
    "description": "Task Description",
    "completed": false
  },
  {
    "id": 2,
    "username": "anotherUser",
    "title": "Another Task",
    "description": "Another Description",
    "completed": true
  }
]
```
  -`401 Unauthorized`: Invalid token or unauthorized access.

```json

{
  "message": "unauthorized access"
}
```
  - `404 Not Found`: No tasks found or user not found.

```json

{
  "message": "not found"
}
```

## Update Task by ID

**Endpoint**: PATCH /tasks/update/:id

**Middleware**:

  **AuthMiddleware**: Validates the JWT token from the Authorization header.
  **AdminMiddleware**: Ensures the user has an admin role.

### Request

**Headers**:

  **Authorization**: Bearer token for authentication.

**Path Parameters**:

    id (int): The ID of the task to be updated.

Body (JSON):

```json

{
  "id": 1,
  "username": "exampleUser",
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "completed": true
}
```
### Responses

    200 OK: Successfully updated the task.

```json

{
  "message": "successfully updated"
}
```

  -`400 Bad Request`: Invalid input or request payload.

```json

{
  "message": "invalid input"
}
```

  -`404 Not Found`: Task with the given ID not found or update failed.

```json

{
  "message": "task not found"
}
```


## Delete Task by ID

**Endpoint**: DELETE /tasks/delete/:id

**Middleware**:

  **AuthMiddleware**: Validates the JWT token from the Authorization header.
  **AdminMiddleware**: Ensures the user has an admin role.

### Request

**Headers**:

  **Authorization**: Bearer token for authentication.

**Path Parameters**:

    id (int): The ID of the task to be deleted.

### Responses

  -`200 OK`: Successfully deleted the task.

```json

{
  "message": "deleted"
}
```
  -`400 Bad Request`: Invalid ID or request payload.

```json

{
  "message": "invalid input"
}
```
  _`404 Not Found`: Task with the given ID not found or deletion failed.

```json

{
  "message": "task not found"
}
```
  - `500 Internal Server Error`: Server error during task deletion.

```json

{
  "message": "internal server error"
}
```
