package jwt

import jwt "github.com/dgrijalva/jwt-go"

// SigningMethod is the enum that lists the supported token signature hashing algorithms.
type SigningMethod int

const (
	_ = iota
	// RSA256 signing algorithm
	RSA256 SigningMethod = iota + 1
	// RSA384 signing algorithm
	RSA384
	// RSA512 signing algorithm
	RSA512
	// HMAC256 signing algorithm
	HMAC256
	// HMAC384 signing algorithm
	HMAC384
	// HMAC512 signing algorithm
	HMAC512
	// ECDSA256 signing algorithm
	ECDSA256
	// ECDSA384 signing algorithm
	ECDSA384
	// ECDSA512 signing algorithm
	ECDSA512
)

var signingmethods map[SigningMethod]string

func init() {
	signingmethods = make(map[SigningMethod]string)

	signingmethods[RSA256] = "RS256"
	signingmethods[RSA384] = "RS512"
	signingmethods[RSA512] = "RS512"
	signingmethods[HMAC256] = "HS256"
	signingmethods[HMAC384] = "HS384"
	signingmethods[HMAC512] = "HS512"
	signingmethods[ECDSA256] = "ES256"
	signingmethods[ECDSA384] = "ES384"
	signingmethods[ECDSA512] = "ES512"
}

// TokenManagerKey is the JWT middleware key used to store the token manager in the context.
const TokenManagerKey middlewareKey = 1

// ValidationKeyfunc is a function that takes a token and returns the key to validate that
// token, which allows it to use inforamtion from the key itself to choose the key
// to return.
type ValidationKeyfunc func(*jwt.Token) (interface{}, error)

func keyFuncWrapper(k ValidationKeyfunc) jwt.Keyfunc {
	return func(tok *jwt.Token) (interface{}, error) {
		return k(tok)
	}
}

// KeyFunc is a function that returns the key to sign a
// token.  It should return a []byte (for all)
// or a *rsa.PrivateKey or *ecdsa.PrivateKey
type KeyFunc func() (interface{}, error)

// Specification describes the JWT authorization properties.
// It is used to both instantiate a middleware and a token manager.
// The middleware uses the specification properties to authorize the incoming
// request while the token manager uses it to create authorization tokens.
type Specification struct {
	// TokenHeader is the HTTP header to search for the JWT Token
	// Defaults to "Authorization"
	TokenHeader string
	// TokenParam is the request parameter to parse for the JWT Token
	// Defaults to "token"
	TokenParam string
	// AllowParam is a flag that determines whether it is allowable
	// to parse tokens from the querystring
	// Defaults to false
	AllowParam bool
	// ValidationFunc is a function that returns the key to validate the JWT
	// Required, no default
	ValidationFunc ValidationKeyfunc
	// AuthOptions is a flag that determines whether a token is required on OPTIONS
	// requests
	AuthOptions bool
	// TTLMinutes is the TTL for tokens that are generated
	TTLMinutes int
	// RefreshTTLMinutes is the TTL for refresh tokens that are generated
	// and should generally be much longer than TTLMinutes
	RefreshTTLMinutes int
	// Issuer is the name of the issuer that will be inserted into the
	// generated token's claims
	Issuer string
	// KeySigningMethod determines the type of key that will be used to sign
	// Tokens.
	KeySigningMethod SigningMethod
	// SigningKeyFunc is a function that returns the key used to sign the token
	SigningKeyFunc KeyFunc
	// CommonClaims is a list of claims added to all tokens issued
	CommonClaims map[string]interface{}
}
