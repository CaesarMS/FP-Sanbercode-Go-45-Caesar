// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/admin": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get all admins, by id, or by username (select one)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get admins list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "insert admin id in UUID format",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "insert admin username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Admin Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.AdminOutput"
                        }
                    }
                }
            }
        },
        "/admin/": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete admin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Delete admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Admin Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.AdminOutput"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "Authenticate admin username \u0026 password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Login admin authenticate",
                "parameters": [
                    {
                        "description": "the body to login as admin",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.AdminAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.AdminOutput"
                        }
                    }
                }
            }
        },
        "/admin/password": {
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update admin password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update admin password",
                "parameters": [
                    {
                        "description": "the body to login as admin",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.AdminUpdate"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Admin Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.AdminOutput"
                        }
                    }
                }
            }
        },
        "/admin/register": {
            "post": {
                "security": [
                    {
                        "x-api-key": []
                    }
                ],
                "description": "Register by username \u0026 password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Register a new admin",
                "parameters": [
                    {
                        "description": "the body to register as admin",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.AdminAuth"
                        }
                    },
                    {
                        "type": "string",
                        "description": "API Key: 9c6f9769-6d5b-493d-ae2e-4fad70711564",
                        "name": "x-api-key",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.AdminOutput"
                        }
                    }
                }
            }
        },
        "/admin/seller": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get all seller",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get seller list by admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Admin Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.UserOutput"
                        }
                    }
                }
            }
        },
        "/admin/user": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get all user, by id, or by email (select one)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get user list by admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "insert user id in UUID format",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "insert user email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Admin Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.UserOutput"
                        }
                    }
                }
            }
        },
        "/category": {
            "get": {
                "description": "Get all categories, by id, or by name (select one)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get categories list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "insert category id in UUID format",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "insert category name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.CategoryOutput"
                        }
                    }
                }
            }
        },
        "/category/": {
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Insert a new category by admin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Insert a new category by admin",
                "parameters": [
                    {
                        "description": "the body to insert a new category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.CategoryInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Admin Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.CategoryOutput"
                        }
                    }
                }
            }
        },
        "/category/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a category by admin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete a category by admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category id in UUID format",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Admin Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.CategoryOutput"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update a category by admin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update a category by admin",
                "parameters": [
                    {
                        "description": "the body to update a category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.CategoryInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Category id in UUID format",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Admin Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.CategoryOutput"
                        }
                    }
                }
            }
        },
        "/invoice": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Get all invoice or by id (select one)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Get invoice list by buyer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "insert invoice id in UUID format",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.InvoiceOutput"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Create a new invoice as a buyer",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Create a new invoice by buyer",
                "parameters": [
                    {
                        "description": "the body to create a new invoice",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.InvoiceInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.InvoiceOutput"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "description": "Get all product, by id, or by category name (select one)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get product list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "insert product id in UUID format",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "insert product category name",
                        "name": "category",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ProductOutput"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Create a new product as a seller",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create a new product by seller",
                "parameters": [
                    {
                        "description": "the body to create a new product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.ProductInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Seller Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ProductOutput"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete product",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id in UUID format",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Seller Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ProductOutput"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update product data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update product data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id in UUID format",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update a product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.ProductInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Seller Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.UserOutput"
                        }
                    }
                }
            }
        },
        "/user/": {
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.AdminOutput"
                        }
                    }
                }
            }
        },
        "/user/data": {
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update user data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user data",
                "parameters": [
                    {
                        "description": "the body to login as user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.UserUpdateData"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.UserOutput"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Authenticate user email \u0026 password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login user authenticate",
                "parameters": [
                    {
                        "description": "the body to login as user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.UserAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.UserOutput"
                        }
                    }
                }
            }
        },
        "/user/password": {
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update user password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user password",
                "parameters": [
                    {
                        "description": "the body to login as user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.UserUpdatePassword"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.UserOutput"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Register as a new user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "the body to register as a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/config.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.UserOutput"
                        }
                    }
                }
            }
        },
        "/user/seller": {
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update user to seller. Must re-login to get a new access token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user to seller",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.UserOutput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.AdminAuth": {
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
        "config.AdminOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "config.AdminUpdate": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "config.CategoryInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "config.CategoryOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "config.InvoiceInput": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.InvoiceItemInput"
                    }
                }
            }
        },
        "config.InvoiceItemInput": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "qty": {
                    "type": "integer"
                }
            }
        },
        "config.InvoiceItemOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                },
                "qty": {
                    "type": "integer"
                }
            }
        },
        "config.InvoiceOutput": {
            "type": "object",
            "properties": {
                "bought_at": {
                    "type": "string"
                },
                "buyer_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.InvoiceItemOutput"
                    }
                },
                "total_price": {
                    "type": "integer"
                }
            }
        },
        "config.ProductInput": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "config.ProductOutput": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "seller": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "config.UserAuth": {
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
        "config.UserCreate": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "config.UserOutput": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "config.UserUpdateData": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "config.UserUpdatePassword": {
            "type": "object",
            "properties": {
                "password": {
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
