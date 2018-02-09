// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server Petstore server.  You can find out more about     Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key ` + "`" + `special-key` + "`" + ` to test the authorization     filters.",
    "title": "Swagger Petstore",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@swagger.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "petstore.swagger.io",
  "basePath": "/v2",
  "paths": {
    "/pet": {
      "put": {
        "consumes": [
          "application/json",
          "application/xml"
        ],
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Update an existing pet",
        "operationId": "updatePet",
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          },
          "405": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json",
          "application/xml"
        ],
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Add a new pet to the store",
        "operationId": "addPet",
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "parameters": [
          {
            "description": "Pet object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/pet/findByStatus": {
      "get": {
        "description": "Multiple status values can be provided with comma separated strings",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Finds Pets by status",
        "operationId": "findPetsByStatus",
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "array",
            "items": {
              "enum": [
                "available",
                "pending",
                "sold"
              ],
              "type": "string",
              "default": "available"
            },
            "collectionFormat": "multi",
            "description": "Status values that need to be considered for filter",
            "name": "status",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid status value"
          }
        }
      }
    },
    "/pet/findByTags": {
      "get": {
        "description": "Muliple tags can be provided with comma separated strings. Use         tag1, tag2, tag3 for testing.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Finds Pets by tags",
        "operationId": "findPetsByTags",
        "deprecated": true,
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Tags to filter by",
            "name": "tags",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Pet"
              }
            }
          },
          "400": {
            "description": "Invalid tag value"
          }
        }
      }
    },
    "/pet/{petId}": {
      "get": {
        "description": "Returns a single pet",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Find pet by ID",
        "operationId": "getPetById",
        "security": [
          {
            "api_key": []
          }
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet to return",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Pet"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      },
      "post": {
        "consumes": [
          "application/x-www-form-urlencoded"
        ],
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Updates a pet in the store with form data",
        "operationId": "updatePetWithForm",
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet that needs to be updated",
            "name": "petId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Updated name of the pet",
            "name": "name",
            "in": "formData"
          },
          {
            "type": "string",
            "description": "Updated status of the pet",
            "name": "status",
            "in": "formData"
          }
        ],
        "responses": {
          "405": {
            "description": "Invalid input"
          }
        }
      },
      "delete": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "Deletes a pet",
        "operationId": "deletePet",
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "string",
            "name": "api_key",
            "in": "header"
          },
          {
            "type": "integer",
            "format": "int64",
            "description": "Pet id to delete",
            "name": "petId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Pet not found"
          }
        }
      }
    },
    "/pet/{petId}/uploadImage": {
      "post": {
        "consumes": [
          "multipart/form-data"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "pet"
        ],
        "summary": "uploads an image",
        "operationId": "uploadFile",
        "security": [
          {
            "petstore_auth": [
              "write:pets",
              "read:pets"
            ]
          }
        ],
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of pet to update",
            "name": "petId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "Additional data to pass to server",
            "name": "additionalMetadata",
            "in": "formData"
          },
          {
            "type": "file",
            "description": "file to upload",
            "name": "file",
            "in": "formData"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/ApiResponse"
            }
          }
        }
      }
    },
    "/store/inventory": {
      "get": {
        "description": "Returns a map of status codes to quantities",
        "produces": [
          "application/json"
        ],
        "tags": [
          "store"
        ],
        "summary": "Returns pet inventories by status",
        "operationId": "getInventory",
        "security": [
          {
            "api_key": []
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "object",
              "additionalProperties": {
                "type": "integer",
                "format": "int32"
              }
            }
          }
        }
      }
    },
    "/store/order": {
      "post": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "store"
        ],
        "summary": "Place an order for a pet",
        "operationId": "placeOrder",
        "parameters": [
          {
            "description": "order placed for purchasing the pet",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Order"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          },
          "400": {
            "description": "Invalid Order"
          }
        }
      }
    },
    "/store/order/{orderId}": {
      "get": {
        "description": "For valid response try integer IDs with value \u003e= 1 and \u003c= 10.         Other values will generated exceptions",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "store"
        ],
        "summary": "Find purchase order by ID",
        "operationId": "getOrderById",
        "parameters": [
          {
            "maximum": 10,
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "ID of pet that needs to be fetched",
            "name": "orderId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Order"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Order not found"
          }
        }
      },
      "delete": {
        "description": "For valid response try integer IDs with positive integer value.         Negative or non-integer values will generate API errors",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "store"
        ],
        "summary": "Delete purchase order by ID",
        "operationId": "deleteOrder",
        "parameters": [
          {
            "minimum": 1,
            "type": "integer",
            "format": "int64",
            "description": "ID of the order that needs to be deleted",
            "name": "orderId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Order not found"
          }
        }
      }
    },
    "/user": {
      "post": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Create user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/createWithArray": {
      "post": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Creates list of users with given input array",
        "operationId": "createUsersWithArrayInput",
        "parameters": [
          {
            "description": "List of user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          }
        ],
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/createWithList": {
      "post": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Creates list of users with given input array",
        "operationId": "createUsersWithListInput",
        "parameters": [
          {
            "description": "List of user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          }
        ],
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/login": {
      "get": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "type": "string",
            "description": "The user name for login",
            "name": "username",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "The password for login in clear text",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "string"
            },
            "headers": {
              "X-Expires-After": {
                "type": "string",
                "format": "date-time",
                "description": "date in UTC when token expires"
              },
              "X-Rate-Limit": {
                "type": "integer",
                "format": "int32",
                "description": "calls per hour allowed by the user"
              }
            }
          },
          "400": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/{username}": {
      "get": {
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched. Use user1 for testing. ",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "put": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Updated user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "name that need to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid user supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "description": "This can only be done by the logged in user.",
        "produces": [
          "application/xml",
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    }
  },
  "definitions": {
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "Category": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Category"
      }
    },
    "Order": {
      "type": "object",
      "properties": {
        "complete": {
          "type": "boolean",
          "default": false
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "petId": {
          "type": "integer",
          "format": "int64"
        },
        "quantity": {
          "type": "integer",
          "format": "int32"
        },
        "shipDate": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "description": "Order Status",
          "type": "string",
          "enum": [
            "placed",
            "approved",
            "delivered"
          ]
        }
      },
      "xml": {
        "name": "Order"
      }
    },
    "Pet": {
      "type": "object",
      "required": [
        "name",
        "photoUrls"
      ],
      "properties": {
        "category": {
          "$ref": "#/definitions/Category"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string",
          "example": "doggie"
        },
        "photoUrls": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "xml": {
            "name": "photoUrl",
            "wrapped": true
          }
        },
        "status": {
          "description": "pet status in the store",
          "type": "string",
          "enum": [
            "available",
            "pending",
            "sold"
          ]
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          },
          "xml": {
            "name": "tag",
            "wrapped": true
          }
        }
      },
      "xml": {
        "name": "Pet"
      }
    },
    "Tag": {
      "type": "object",
      "properties": {
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Tag"
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "lastName": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phone": {
          "type": "string"
        },
        "userStatus": {
          "description": "User Status",
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      },
      "xml": {
        "name": "User"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "petstore_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://petstore.swagger.io/oauth/dialog",
      "scopes": {
        "read:pets": "read your pets",
        "write:pets": "modify pets in your account"
      }
    }
  },
  "tags": [
    {
      "description": "Everything about your Pets",
      "name": "pet",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://swagger.io"
      }
    },
    {
      "description": "Access to Petstore orders",
      "name": "store"
    },
    {
      "description": "Operations about user",
      "name": "user",
      "externalDocs": {
        "description": "Find out more about our store",
        "url": "http://swagger.io"
      }
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}
