// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "/createuser": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "新创建用户信息",
                "parameters": [
                    {
                        "description": "请求参数结构体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseData"
                        }
                    }
                }
            }
        },
        "/deleteuser/{userId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "删除指定用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseData"
                        }
                    }
                }
            }
        },
        "/getuser/{userId}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "获取单个用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登录信息",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseData"
                        }
                    }
                }
            }
        },
        "/getusers": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "获取用户列表（分页）",
                "parameters": [
                    {
                        "description": "请求参数结构体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserSearchDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseData"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "请求参数结构体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseData"
                        }
                    }
                }
            }
        },
        "/updateuser": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "请求参数结构体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.ResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateUserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginDto": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "dto.UserSearchDto": {
            "type": "object",
            "required": [
                "page",
                "pageSize"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "utils.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "object"
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
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/api/v1/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server.",
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
