package jwt_test

import (
	"go/adv-demo/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@a.ru"
	jwtService := jwt.NewJWT("uuiFTsvOHu2zkLS6LIqbjz8WJE7hY1YH")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}
}
