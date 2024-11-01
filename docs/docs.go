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
        "/trait-probability": {
            "post": {
                "description": "Calculate the probability of a trait based on the genotype of two plants",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "trait-probability"
                ],
                "summary": "Calculate Trait Probability",
                "parameters": [
                    {
                        "description": "Trait probability calculation request",
                        "name": "trait",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TraitRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TraitResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorTraitResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorTraitResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Plant": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "traits": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Trait"
                    }
                }
            }
        },
        "models.Trait": {
            "type": "object",
            "properties": {
                "dominant": {
                    "type": "boolean"
                },
                "geneLabel": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.TraitCombination": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "genotype": {
                    "type": "string"
                },
                "numberOfPlants": {
                    "type": "integer"
                },
                "percentage": {
                    "type": "number"
                }
            }
        },
        "models.TraitRequest": {
            "type": "object",
            "properties": {
                "fatherPlant": {
                    "$ref": "#/definitions/models.Plant"
                },
                "motherPlant": {
                    "$ref": "#/definitions/models.Plant"
                },
                "plantCount": {
                    "type": "integer"
                }
            }
        },
        "models.TraitResponse": {
            "type": "object",
            "properties": {
                "combinations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TraitCombination"
                    }
                },
                "fatherPlantName": {
                    "type": "string"
                },
                "motherPlantName": {
                    "type": "string"
                },
                "plantCount": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "Pepper Analytics API",
	Description:      "A Go REST API that calculates F2 generation probabilities in plant breeding genetics. Calculate inheritance patterns and expected phenotype distributions based on parent plant genotypes. Features Swagger documentation and handles multiple genetic traits with dominant/recessive expressions.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
