// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://localhost:8080/swagger/index.html",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/address/geocode": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geocode"
                ],
                "summary": "Geocode an address",
                "operationId": "geocodeAddress",
                "parameters": [
                    {
                        "description": "Geocode request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.GeocodeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AddressGeo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/address/search": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "search"
                ],
                "summary": "Search for an address",
                "operationId": "searchAddress",
                "parameters": [
                    {
                        "description": "Search request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AddressSearch"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddressGeo": {
            "description": "AddressGeo represents the geocode result for an address.",
            "type": "object",
            "properties": {
                "suggestions": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "data": {
                                "type": "object",
                                "properties": {
                                    "country": {
                                        "type": "string"
                                    },
                                    "postal_code": {
                                        "type": "string"
                                    }
                                }
                            },
                            "unrestricted_value": {
                                "type": "string"
                            },
                            "value": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "models.AddressSearch": {
            "description": "AddressSearch represents the search result for an address.",
            "type": "object",
            "properties": {
                "metro": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "distance": {
                                "type": "number"
                            },
                            "line": {
                                "type": "string"
                            },
                            "name": {
                                "type": "string"
                            }
                        }
                    }
                },
                "result": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "models.GeocodeRequest": {
            "description": "GeocodeRequest represents the request body for address geocoding.",
            "type": "object",
            "properties": {
                "lat": {
                    "type": "string"
                },
                "lon": {
                    "type": "string"
                }
            }
        },
        "models.SearchRequest": {
            "description": "SearchRequest represents the request body for address search",
            "type": "object",
            "properties": {
                "query": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.1",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "My API",
	Description:      "This is a sample API for address searching and geocoding using Dadata API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
