package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TokenManager creates tokens given a JWT specification and claims.
const (
	ttldefault        = 5
	refreshttldefault = 1440 // default to one day
)

// TokenManager provides for the creation of access and refresh JWT Tokens
type TokenManager struct {
	spec *Specification
}

// NewTokenManager returns a TokenManager.  If TTLMinutes isn't specified
// it will default to 5 minutes.  Use the same Specification as you use for
// Middleware() to ensure your tokens are compatible.
func NewTokenManager(spec *Specification) *TokenManager {
	if spec.TTLMinutes == 0 {
		spec.TTLMinutes = ttldefault
	}
	if spec.RefreshTTLMinutes == 0 {
		spec.RefreshTTLMinutes = refreshttldefault
	}
	return &TokenManager{spec: spec}
}

// Create makes a new token, adding the claims provided.  It returns
// a token as a string.
func (tm *TokenManager) Create(claims map[string]interface{}) (string, error) {

	t := jwt.New(jwt.GetSigningMethod(signingmethods[tm.spec.KeySigningMethod]))
	for k, v := range claims {
		t.Claims[k] = v
	}

	if tm.spec.Issuer != "" {
		t.Claims["iss"] = tm.spec.Issuer
	}

	for k, v := range tm.spec.CommonClaims {
		t.Claims[k] = v
	}
	// set issued at time
	t.Claims["iat"] = time.Now().Unix()
	// set the expire time
	t.Claims["exp"] = time.Now().Add(time.Minute * time.Duration(tm.spec.TTLMinutes)).Unix()
	bytes, err := tm.spec.SigningKeyFunc()
	if err != nil {
		return "", fmt.Errorf("Error retrieving Signing Key: %v", err)
	}
	return t.SignedString(bytes)

}
