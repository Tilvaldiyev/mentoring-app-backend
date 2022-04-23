// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/article": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Create article",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/article/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get article info by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID article",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Article"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Update article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID article",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateArticleInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Delete article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID article",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/mentors": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get mentors list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mentor"
                ],
                "summary": "Get mentors list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search by expertise",
                        "name": "expertise",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by country",
                        "name": "country",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search by language",
                        "name": "language",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.MentorResponse"
                            }
                        }
                    },
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/post": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Create post",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Posts"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/post/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get post info by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Get post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID article",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Posts"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Delete post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID article",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password": {
            "post": {
                "description": "Send secret code to email when forgot password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Send secret code to email",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PasswordRecoveryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/auth/recover": {
            "put": {
                "description": "Update password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Update password",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/auth/sign-in": {
            "post": {
                "description": "User authorization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authorization",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "description": "User registration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User registration",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Users"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/auth/verify-secret-code": {
            "post": {
                "description": "Verify sent secret code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verify sent secret code",
                "parameters": [
                    {
                        "description": "request body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PasswordRecoveryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/countries": {
            "get": {
                "description": "Get list of countries",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    ""
                ],
                "summary": "Get list of countries",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Country"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/expertises": {
            "get": {
                "description": "Get list of user levels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    ""
                ],
                "summary": "Get list of user levels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Level"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/languages": {
            "get": {
                "description": "Get list of languages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    ""
                ],
                "summary": "Get list of languages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Language"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        },
        "/user-types": {
            "get": {
                "description": "Get list of user types",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    ""
                ],
                "summary": "Get list of user types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.UserType"
                            }
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Article": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "expertise_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "img_url": {
                    "type": "string"
                },
                "level_id": {
                    "type": "integer"
                },
                "likes": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "views": {
                    "type": "integer"
                }
            }
        },
        "model.Country": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Expertise": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Language": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Level": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.MentorResponse": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "expertise": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "level": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.PasswordRecoveryInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "secret_code": {
                    "type": "string"
                }
            }
        },
        "model.Posts": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "isSuccess": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.SignInInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.UpdateArticleInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.UserType": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.Users": {
            "type": "object",
            "properties": {
                "country_id": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "expertise_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "info": {
                    "type": "string"
                },
                "language_id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "level_id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user_type_id": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Mentoring App API",
	Description:      "API for virtual mentoring application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
