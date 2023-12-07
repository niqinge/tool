package jwtimport

import (
    "crypto/rsa"
    "errors"
    "github.com/golang-jwt/jwt/v4"
)

type Jwt struct {
    privateKey *rsa.PrivateKey
    publicKey  *rsa.PublicKey
}

func NewJwt(privatekeyPEM, publickeyPEM string) *Jwt {
    prvk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatekeyPEM))
    if err != nil {
        panic(err)
    }

    pk, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publickeyPEM))
    if err != nil {
        panic(err)
    }

    if prvk == nil || pk == nil {
        panic("prvk or pk is nil")
    }

    srv := &Jwt{}
    srv.privateKey = prvk
    srv.publicKey = pk

    return srv
}

func (receiver *Jwt) Decode(token string) (string, error) {
    return receiver.claimsFromToken(token)
}

func (receiver *Jwt) Encode(userId string) (string, error) {
    token := receiver.encode(userId)

    return token.SignedString(receiver.privateKey)
}

func (receiver *Jwt) encode(userId string) *jwt.Token {
    return jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{"user_id": userId})
}

func (receiver *Jwt) claimsFromToken(str string) (string, error) {
    resultToken, e := jwt.ParseWithClaims(str, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
            // logger.Error("解析token出错:" + str)
            return nil, errors.New("parse token failed:" + str)
        }
        return receiver.publicKey, nil
    })
    var claims jwt.MapClaims
    if resultToken != nil && resultToken.Claims != nil {
        claims = resultToken.Claims.(jwt.MapClaims)
    } else {
        e = errors.New("无法解析token")
    }
    if userId, ok := claims["user_id"].(string); ok {
        return userId, e
    }

    return "", e
}
