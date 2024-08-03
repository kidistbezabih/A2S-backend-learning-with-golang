POST /tasks
This endpoint is used to create a new task.
Request Body
id (number, required): The unique identifier for the task.
title (string, required): The title of the task.
description (string, required): The description of the task.
completed (boolean, required): Indicates whether the task is completed or not.

Response
The response is in JSON format with the following schema:


JSON


{
  "type": "object",
  "properties": {
    "message": {
      "type": "string"
    }
  }
}


message (string): A message indicating the status of the request.

Example
Request Body:


JSON

{
  "id": 0,
  "title": "",
  "description": "",
  "completed": true
}


Response:


JSON

{
  "message": ""
}


Status Code: 201
Content-Type: application/json


GET /tasks
This endpoint retrieves a list of tasks.
Request
The request should be sent to localhost:8080/tasks using an HTTP GET method.
Request Body
id (number, required): The ID of the task.
title (string, required): The title of the task.
description (string, required): The description of the task.
completed (boolean, required): Indicates whether the task is completed.

Response
Upon successful execution, the endpoint returns a status code of 200 and a JSON array with the following schema:


JSON

[
    {
        "id": "number",
        "title": "string",
        "description": "string",
        "completed": "boolean"
    }
]




