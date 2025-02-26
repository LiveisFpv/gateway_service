// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/country": {
            "get": {
                "description": "Возвращает страны",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "country"
                ],
                "summary": "Получение стран",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/httpgin.Get_CountryById_Response"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Редактирует страну",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "country"
                ],
                "summary": "Редактирование страны",
                "parameters": [
                    {
                        "description": "Данные страны",
                        "name": "country",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpgin.Update_CountryById_Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpgin.Update_CountryById_Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет страну",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "country"
                ],
                "summary": "Добавление страну",
                "parameters": [
                    {
                        "description": "Данные страны",
                        "name": "country",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpgin.Create_Country_Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpgin.Create_Country_Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет страну",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "country"
                ],
                "summary": "Удаление страны",
                "parameters": [
                    {
                        "description": "ID страны",
                        "name": "country",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/httpgin.Delete_CountryById_Request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpgin.Delete_CountryById_Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpgin.Create_Country_Request": {
            "type": "object",
            "properties": {
                "country_area": {
                    "type": "string"
                },
                "country_capital": {
                    "type": "string"
                },
                "country_title": {
                    "type": "string"
                }
            }
        },
        "httpgin.Create_Country_Response": {
            "type": "object",
            "properties": {
                "country_id": {
                    "type": "integer"
                }
            }
        },
        "httpgin.Delete_CountryById_Request": {
            "type": "object",
            "properties": {
                "country_id": {
                    "type": "integer"
                }
            }
        },
        "httpgin.Delete_CountryById_Response": {
            "type": "object",
            "properties": {
                "country_title": {
                    "type": "string"
                }
            }
        },
        "httpgin.Get_CountryById_Response": {
            "type": "object",
            "properties": {
                "country_area": {
                    "type": "string"
                },
                "country_capital": {
                    "type": "string"
                },
                "country_id": {
                    "type": "integer"
                },
                "country_title": {
                    "type": "string"
                }
            }
        },
        "httpgin.Update_CountryById_Request": {
            "type": "object",
            "properties": {
                "country_area": {
                    "type": "string"
                },
                "country_capital": {
                    "type": "string"
                },
                "country_id": {
                    "type": "integer"
                },
                "country_title": {
                    "type": "string"
                }
            }
        },
        "httpgin.Update_CountryById_Response": {
            "type": "object",
            "properties": {
                "country_area": {
                    "type": "string"
                },
                "country_capital": {
                    "type": "string"
                },
                "country_title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:18000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "country API",
	Description:      "API с документацией Swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
