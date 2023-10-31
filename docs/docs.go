// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://www.google.com",
        "contact": {
            "name": "Jason Yang",
            "url": "http://www.google.com",
            "email": "jjkk900925@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customer": {
            "post": {
                "description": "Get Customer by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer ID",
                        "name": "ID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Customer"
                        }
                    },
                    "500": {
                        "description": "{\"Message\": err.Error()}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customerList": {
            "get": {
                "description": "Get all Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Customer"
                        }
                    },
                    "500": {
                        "description": "{\"Message\": \"Internal Error!\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Customer": {
            "type": "object",
            "properties": {
                "Address": {
                    "type": "string"
                },
                "Birthday": {
                    "type": "string"
                },
                "CarNumber": {
                    "type": "string"
                },
                "Citizenship": {
                    "type": "string"
                },
                "CustomerId": {
                    "type": "integer"
                },
                "Gender": {
                    "type": "string"
                },
                "ID": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Note": {
                    "type": "string"
                },
                "PhoneNumber": {
                    "type": "string"
                },
                "history": {
                    "$ref": "#/definitions/model.History"
                }
            }
        },
        "model.History": {
            "type": "object",
            "properties": {
                "CustomerId": {
                    "type": "integer"
                },
                "Date": {
                    "type": "string"
                },
                "HistoryId": {
                    "type": "integer"
                },
                "Nofpeople": {
                    "type": "integer"
                },
                "Note": {
                    "type": "string"
                },
                "Price": {
                    "type": "integer"
                },
                "Room": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "CRMS_Swagger",
	Description:      "CRMS_Swagger information",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
