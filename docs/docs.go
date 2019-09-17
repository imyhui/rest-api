// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-09-17 20:38:07.808871 +0800 CST m=+0.035076311

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Login generates the authentication token",
                "parameters": [
                    {
                        "description": "login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UserModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ\"}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sd/cpu": {
            "get": {
                "description": "Checks the cpu usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the cpu usage",
                "responses": {
                    "200": {
                        "description": "CRITICAL - Load average: 1.78, 1.99, 2.02 | Cores: 2",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sd/disk": {
            "get": {
                "description": "Checks the disk usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the disk usage",
                "responses": {
                    "200": {
                        "description": "OK - Free space: 17233MB (16GB) / 51200MB (50GB) | Used: 33%",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sd/health": {
            "get": {
                "description": "Shows OK as the ping-pong result",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Shows OK as the ping-pong result",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sd/ram": {
            "get": {
                "description": "Checks the ram usage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sd"
                ],
                "summary": "Checks the ram usage",
                "responses": {
                    "200": {
                        "description": "OK - Free space: 402MB (0GB) / 8192MB (8GB) | Used: 4%",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/user": {
            "get": {
                "description": "List users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "List the users in the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "List users",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.ListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"totalCount\":1,\"userList\":[{\"id\":0,\"username\":\"admin\",\"random\":\"user 'admin' get random string 'EnqntiSig'\",\"password\":\"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG\",\"createdAt\":\"2018-05-28 00:25:33\",\"updatedAt\":\"2018-05-28 00:25:33\"}]}}",
                        "schema": {
                            "$ref": "#/definitions/user.ListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Add new user to the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"username\":\"kong\"}}",
                        "schema": {
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    }
                }
            }
        },
        "/v1/user/{id}": {
            "put": {
                "description": "Update a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a user info by the user identifier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The user's database id index num",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UserModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":null}",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete an user by the user identifier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The user's database id index num",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":null}",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/v1/user/{username}": {
            "get": {
                "description": "Get an user by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get an user by the user identifier",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"username\":\"kong\",\"password\":\"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS\"}}",
                        "schema": {
                            "$ref": "#/definitions/model.UserModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Token": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "sayHello": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.UserModel": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "user.ListRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.ListResponse": {
            "type": "object",
            "properties": {
                "totalCount": {
                    "type": "integer"
                },
                "userList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UserInfo"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
