package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenPair(user *UserData, secret string) (*Tokens, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = user.ID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  t,
		RefreshToken: rt,
	}, nil
}

func ParseToken(token string, secret string) (*UserData, error) {
	t, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	switch {
	case t == nil:
		return nil, fmt.Errorf("invalid token")
	case t.Valid:
		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			return nil, fmt.Errorf("invalid token")
		}

		id, err := claims.GetSubject()
		if err != nil {
			return nil, fmt.Errorf("invalid token")
		}
		email, ok := claims["email"].(string)

		return &UserData{
			ID:    id,
			Email: email,
		}, nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		return nil, fmt.Errorf("invalid token")
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		// Invalid signature
		return nil, fmt.Errorf("Invalid signature")
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		// Token is either expired or not active yet
		return nil, err
	default:
		return nil, fmt.Errorf("invalid token")
	}
}
