package jwt

import (
    "github.com/niqinge/tool/encrypt"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func Test_gen(t *testing.T) {

    rasEncrypt := encrypt.NewRasEncrypt()
    err := rasEncrypt.GenKey(2048)
    assert.NoError(t, err)

    jt, err := NewJwt(rasEncrypt.GetPrivateKey(), rasEncrypt.GetPublicKey())
    assert.NoError(t, err)

    param := map[string]interface{}{"user_id": time.Now().UnixNano()}
    token, err := jt.Encode(param)
    assert.NoError(t, err)

    decodeResult, err := jt.Decode(token)
    assert.NoError(t, err)
    for k, v := range param {
        decodeResult[k] = v
    }
}
