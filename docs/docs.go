// Code generated by swaggo/swag. DO NOT EDIT
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
        "/group/groupr": {
            "post": {
                "description": "token 为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "创建群聊接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": 200,\"msg\": \"创建群聊成功,群号为:8660920\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/group/join": {
            "post": {
                "description": "room_id token 为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "加入群聊接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "群号",
                        "name": "room_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": 200, \"msg\": \"加入群聊成功！\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "get": {
                "description": "用户名 token 邮箱为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "删除好友接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"msg\":\"删除成功！\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "上传文件\nfile 为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "上传文件接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "表单name",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":1,\"msg\":\"\\u0001个文件上传成功\",\"url\":\"127.0.0.1:8080/img/12.png\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/join": {
            "get": {
                "description": "用户名 token 邮箱为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "添加好友接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "账号",
                        "name": "account",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\":200,\"msg\":\"添加好友成功！\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户名 密码为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "登录接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": 200, \"msg\": \"登陆成功\", \"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpbmRlbnRseSI6IjZhMmE0NjJjLWExMDctNDhlYS04MmU1LTc0ZTMwODMyN2U2ZiIsInVzZXJuYW1lIjoiYWRtaW4iLCJpc3MiOiJ0ZXN0IiwiZXhwIjoxNjc4Nzg2NTM1fQ.P4dJ_f2UGhKbpiIqHxTxghRKwKIlCpF2XjryHCSnKKk\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户名 密码 邮箱为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "注册接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "account\": \"3169387148\", \"code\": 200, \"msg\": \"注册成功\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/send_code": {
            "get": {
                "description": "用户名 为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "验证码接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\": \"530757\", \"msg\": \"获取验证码成功！\" }",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/websocket": {
            "get": {
                "description": "token 邮箱为必填",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共方法"
                ],
                "summary": "websocket连接接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
