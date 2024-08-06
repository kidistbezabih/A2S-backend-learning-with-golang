### Request

This is an HTTP POST request used to create a new task. The request body is in JSON format and includes the following parameters:

- `id` (number): The unique identifier for the task.
- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `completed` (boolean): Indicates whether the task is completed or not.

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


### Get All Tasks

This endpoint retrieves all tasks.

#### Request

- Method: GET
- URL: `http://localhost:9090/v1/tasks/getall`
- Headers: N/A
- Body:
    - id (number): The ID of the task.
    - title (string): The title of the task.
    - description (string): The description of the task.
    - completed (boolean): Indicates whether the task is completed.

#### Response

The response is in JSON format and returns an array of task objects with the following schema:

``` json
[
  {
    "id": 0,
    "title": "",
    "description": "",
    "completed": true
  }
]

 ```

 ### Get Task by ID

This endpoint retrieves a task by its ID.

#### Request

- Method: GET
- URL: `http://localhost:9090/v1/tasks/get/1`
- Headers:
    - Content-Type: application/json
- Body:
    
    ``` json
      {
          "id": 0,
          "title": "",
          "description": "",
          "completed": true
      }
    
     ```
    

#### Response

- Status: 200
- Content-Type: application/json
- Body:
    
    ``` json
      {
          "id": 0,
          "title": "",
          "description": "",
          "completed": true
      }
    
     ```
    

#### Response Schema

``` json
{
    "type": "object",
    "properties": {
        "id": {
            "type": "number"
        },
        "title": {
            "type": "string"
        },
        "description": {
            "type": "string"
        },
        "completed": {
            "type": "boolean"
        }
    }
}

 ```

 ### Delete Task

This endpoint is used to delete a specific task.

#### Request

- Method: DELETE
- URL: `http://localhost:9090/v1/tasks/delete/2`
- Body:
    - id (number, required): The ID of the task to be deleted.
    - title (string): The title of the task.
    - description (string): The description of the task.
    - completed (boolean): Indicates whether the task is completed.

#### Response

The response for this request is in JSON format and follows the schema below:

``` json
{
  "type": "object",
  "properties": {
    "message": {
      "type": "string"
    }
  }
}

 ```

- Status: 200 OK
- Content-Type: application/json

### Update Task Details

This endpoint is used to update the details of a specific task.

#### Request Body

- `id` (number): The unique identifier of the task.
- `title` (string): The title of the task.
- `description` (string): The description of the task.
- `completed` (boolean): Indicates whether the task is completed or not.
    

#### Response

The response is a JSON object with the following schema:

``` json
{
    "message": "string"
}

 ```

- `message` (string): A message indicating the status of the update operation.