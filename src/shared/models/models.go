package models

import "github.com/dgrijalva/jwt-go"

type ResponseError struct {
	Error string `json:"error"`
}

type ResponseList struct {
	Total int `json:"total"`
}

type PublicEndpoint struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type Error struct {
	Text string
	Core error
	Code int
}

type SwaggerJson struct {
	Swagger             any    `json:"swagger"`
	Host                string `json:"host"`
	Definitions         any    `json:"definitions"`
	ExternalDocs        any    `json:"externalDocs"`
	Info                any    `json:"info"`
	Paths               any    `json:"paths"`
	Schemes             any    `json:"schemes"`
	SecurityDefinitions any    `json:"securityDefinitions"`
	Tags                any    `json:"tags"`
}

type JwtPayload struct {
	UserID string `json:"user_id"`
	AppID  string `json:"app_id"`
	ACL    string `json:"acl"`
	jwt.StandardClaims
}
