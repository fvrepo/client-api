// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Client API",
    "version": "1.0"
  },
  "basePath": "/api/v1.0",
  "paths": {
    "/ports": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "get ports",
        "operationId": "getAllPorts",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "additionalProperties": {
                "type": "object",
                "$ref": "#/definitions/Port"
              }
            }
          },
          "500": {
            "description": "General server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "post ports",
        "operationId": "postPorts",
        "parameters": [
          {
            "type": "file",
            "description": "The file to upload",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "500": {
            "description": "General server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "internal status code",
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Port": {
      "description": "port object",
      "type": "object",
      "properties": {
        "alias": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "city": {
          "description": "port city",
          "type": "string"
        },
        "code": {
          "description": "port country",
          "type": "string"
        },
        "coordinates": {
          "description": "port coordinates",
          "type": "array",
          "items": {
            "type": "number",
            "format": "float64"
          }
        },
        "country": {
          "description": "port country",
          "type": "string"
        },
        "name": {
          "description": "port name",
          "type": "string"
        },
        "pinCode": {
          "description": "pin code",
          "type": "string"
        },
        "province": {
          "description": "port province",
          "type": "string"
        },
        "regions": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "timezone": {
          "description": "port timezone",
          "type": "string"
        },
        "unlocs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Client API",
    "version": "1.0"
  },
  "basePath": "/api/v1.0",
  "paths": {
    "/ports": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "get ports",
        "operationId": "getAllPorts",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "additionalProperties": {
                "type": "object",
                "$ref": "#/definitions/Port"
              }
            }
          },
          "500": {
            "description": "General server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "summary": "post ports",
        "operationId": "postPorts",
        "parameters": [
          {
            "type": "file",
            "description": "The file to upload",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "500": {
            "description": "General server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "internal status code",
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Port": {
      "description": "port object",
      "type": "object",
      "properties": {
        "alias": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "city": {
          "description": "port city",
          "type": "string"
        },
        "code": {
          "description": "port country",
          "type": "string"
        },
        "coordinates": {
          "description": "port coordinates",
          "type": "array",
          "items": {
            "type": "number",
            "format": "float64"
          }
        },
        "country": {
          "description": "port country",
          "type": "string"
        },
        "name": {
          "description": "port name",
          "type": "string"
        },
        "pinCode": {
          "description": "pin code",
          "type": "string"
        },
        "province": {
          "description": "port province",
          "type": "string"
        },
        "regions": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "timezone": {
          "description": "port timezone",
          "type": "string"
        },
        "unlocs": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    }
  }
}`))
}
