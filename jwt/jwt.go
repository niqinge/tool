package jwt

import (
    "crypto/rsa"
    "errors"
    "github.com/golang-jwt/jwt/v4"
)

var (
    ErrKeyIsNil         = errors.New("prvk or pk is nil")
    ErrTokenParseFailed = errors.New("token parse failed")
)

type Jwt struct {
    privateKey *rsa.PrivateKey
    publicKey  *rsa.PublicKey
}

func NewJwt(privatekeyPEM, publickeyPEM string) (*Jwt, error) {
    prvk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatekeyPEM))
    if err != nil {
        return nil, err
    }

    pk, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publickeyPEM))
    if err != nil {
        return nil, err
    }

    if prvk == nil || pk == nil {
        return nil, ErrKeyIsNil
    }

    srv := &Jwt{}
    srv.privateKey = prvk
    srv.publicKey = pk

    return srv, nil
}

func (receiver *Jwt) Decode(token string) (map[string]interface{}, error) {
    return receiver.claimsFromToken(token)
}

func (receiver *Jwt) Encode(param map[string]interface{}) (string, error) {
    token := receiver.encode(param)

    return token.SignedString(receiver.privateKey)
}

func (receiver *Jwt) encode(param map[string]interface{}) *jwt.Token {
    mapClaims := jwt.MapClaims{}
    for k, v := range param {
        mapClaims[k] = v
    }
    return jwt.NewWithClaims(jwt.SigningMethodRS256, mapClaims)
}

func (receiver *Jwt) claimsFromToken(str string) (jwt.MapClaims, error) {
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
        e = ErrTokenParseFailed
    }

    return claims, e
}
