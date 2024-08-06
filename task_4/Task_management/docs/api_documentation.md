## POST /tasks

This endpoint is used to create a new task.

### Request Body

- **id** (number, required): The unique identifier for the task.
- **title** (string, required): The title of the task.
- **description** (string, required): The description of the task.
- **completed** (boolean, required): Indicates whether the task is completed or not.

### Response

The response is in JSON format with the following schema:

```json
{
  "type": "object",
  "properties": {
    "message": {
      "type": "string"
    }
  }
}
```

## Update Task

### Request

**PUT** `/tasks/:id`

**Path Parameters:**

- `id` (integer, required): The ID of the task to retrieve.

**Request Body:**

```json
{
  "title": "Updated Task 1",
  "done": true
}
  ```
## Delete Task

### Request

**DELETE** `/tasks/:id`

**Path Parameters:**

- `id` (integer, required): The ID of the task to retrieve.

### Response

**Status Code: 200 OK**

**Response Body:**

```json
{
  "message": "task deleted"
}
```
