package main

type JWK struct {
	Kty   string `json:"kty"`
	D     string `json:"d"`
	E     string `json:"e"`
	Use   string `json:"use"`
	Kid   string `json:"kid"`
	Alias string `json:"alias"`
	Alg   string `json:"alg"`
	N     string `json:"n"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}