package utils

import (
	"encoding/json"

	"errors"

	conf "../configuration"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//TokenGenerator provides basic methods to generate and parse jwt token
type TokenGenerator struct {
	TokenPassword string `json:"tokenPassword"`
	SigningMethod string `json:"signingMethod"`
}

//Token represents jwt token data
type Token struct {
	SessionID uuid.UUID

	jwt.StandardClaims
}

//InitTokenGenerator is used to read and parse token config and initialize token generator
func InitTokenGenerator(env conf.EnvironmentKey) (*TokenGenerator, error) {
	config, err := conf.ReadTokenConfig(env)
	if err != nil {
		return nil, err
	}

	tokenGenerator, err := parseTokenConfig(config)

	if err != nil {
		return nil, err
	}

	return tokenGenerator, nil
}

// ParseTokenConfig parses json config for token generator
func parseTokenConfig(config []byte) (*TokenGenerator, error) {
	var parsedConfig *TokenGenerator

	err := json.Unmarshal(config, &parsedConfig)

	return parsedConfig, err
}

//CreateNewToken returns new jwt token with user session UUID
func (t *TokenGenerator) CreateNewToken(uuid uuid.UUID) (string, error) {
	tk := &Token{SessionID: uuid}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(t.SigningMethod), tk)
	tokenString, err := token.SignedString([]byte(t.TokenPassword))

	return tokenString, err
}

//ParseToken parses jwt token
func (t *TokenGenerator) ParseToken(tkn string) (*Token, error) {
	// splitted := strings.Split(tkn, " ") // check `Bearer {token-body}` format
	// if len(splitted) != 2 {
	// 	return nil, errors.New("Invalid token format for token: " + tkn)
	// }

	// tokenPart := splitted[1] //getting the first part of the token
	tk := &Token{}

	token, err := jwt.ParseWithClaims(tkn, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.TokenPassword), nil
	})

	if err != nil {
		return nil, errors.New("Error parsing token: " + tkn)
	}

	if !token.Valid {
		return nil, errors.New("Invalid token: " + tkn)
	}

	return tk, nil
}
