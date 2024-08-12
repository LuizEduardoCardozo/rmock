# Rmock - A no-fancy tool for mocking your API

Rmock is a simple tool, with almost nothing to setup, designed for mocking your API. It is designed to be simple and easy to use. It is not a full-fledged API mocking tool, but it is a good starting point for simple projects.


Perfect for when you want to talk less, and show them more code.

## Installation

TODO

## Usage

First of all, make sure that you have the configuration file in the root of your project. The configuration file is a simple JSON file that contains the routes and the responses that you want to mock.

Eg:
```json
{
  "routes": [
    {
      "path": "/api/v1/users",
      "method": "GET",
      "response": {
        "status": 200,
        "body": {
          "users": [
            {
              "id": 1,
              "name": "John Doe"
            },
            {
              "id": 2,
              "name": "Jane Doe"
            }
          ]
        }
      }
    }
  ]
}
```

Then, you can start the server by running the following command:

```bash
$ rmock --port <port> --json <config file>
```

The server will start, then you can access the routes that you have defined in the configuration file.

It's not fancy, but it gets the job done.