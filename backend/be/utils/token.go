package utils

import (
	"encoding/json"

	conf "../configuration"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

//TokenGenerator provides basic methods to generate and parse jwt token
type TokenGenerator struct {
	tokenPassword string `json:"tokenPassword"`
	signingMethod string `json:"signingMethod"`
}

//Token represents jwt token data
type Token struct {
	sessionID uuid.UUID
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
	tk := &Token{sessionID: uuid}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(t.signingMethod), tk)
	tokenString, err := token.SignedString([]byte(t.tokenPassword))

	return tokenString, err
}

func (t *TokenGenerator) ParseToken() {
	
}
