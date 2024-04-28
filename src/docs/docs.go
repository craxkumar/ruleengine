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
        "/management/health/liveness": {
            "get": {
                "description": "Get the liveness status of the service",
                "produces": [
                    "application/json"
                ],
                "summary": "Get the liveness status",
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/controllers.HealthResponse"
                        }
                    }
                }
            }
        },
        "/management/health/readiness": {
            "get": {
                "description": "Get the readiness status of the service",
                "produces": [
                    "application/json"
                ],
                "summary": "Get the readiness status",
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/controllers.HealthResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.HealthResponse": {
            "type": "object",
            "properties": {
                "components": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/controllers.HealthStatus"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "controllers.HealthStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
