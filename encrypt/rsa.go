package encrypt

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
)

/*
   RSA 加解密
*/

type RasEncrypt struct {
    privateKey    *rsa.PrivateKey
    privateKeyStr string
    publicKeyStr  string
}

func NewRasEncrypt() *RasEncrypt {
    return &RasEncrypt{}
}

func NewRasEncryptDefault(bits int) (*RasEncrypt, error) {
    rasEncrypt := &RasEncrypt{}

    return rasEncrypt, rasEncrypt.GenKey(bits)
}

func NewRasEncryptByPrivateKey(privateKey string) (*RasEncrypt, error) {
    if len(privateKey) == 0 {
        return nil, errors.New("private key is empty")
    }
    block, _ := pem.Decode([]byte(privateKey))
    if block == nil {
        return nil, errors.New("private key error")
    }

    parsePKCS1PrivateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    rasEncrypt := &RasEncrypt{privateKey: parsePKCS1PrivateKey}

    return rasEncrypt, rasEncrypt.init(parsePKCS1PrivateKey)
}

func (s *RasEncrypt) init(privateKey *rsa.PrivateKey) error {
    x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
    privateBlock := &pem.Block{Type: "RSA Private Key", Bytes: x509PrivateKey}

    x509PublicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
    if err != nil {
        return err
    }
    publicBlock := &pem.Block{Type: "RSA Public Key", Bytes: x509PublicKey}

    s.privateKeyStr = string(pem.EncodeToMemory(privateBlock))
    s.publicKeyStr = string(pem.EncodeToMemory(publicBlock))

    return nil
}

func (s *RasEncrypt) GetPublicKey() string {
    return s.publicKeyStr
}

func (s *RasEncrypt) GetPrivateKey() string {
    return s.privateKeyStr
}

// GenKey 生成密钥对
func (s *RasEncrypt) GenKey(bits int) (err error) {
    tmpPrivateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return err
    }
    s.privateKey = tmpPrivateKey

    return s.init(tmpPrivateKey)
}

// Encrypt 加密
func (s *RasEncrypt) Encrypt(plainText []byte) (cipherText []byte, err error) {
    // 3.使用公钥加密
    cipherTextBt, err := rsa.EncryptPKCS1v15(rand.Reader, &s.privateKey.PublicKey, plainText)
    if err != nil {
        return cipherText, err
    }

    cipherText = cipherTextBt
    return
}

// Decrypt 解密
func (s *RasEncrypt) Decrypt(cipherText string) (plainText []byte, err error) {
    cipherTextBt, err := base64.StdEncoding.DecodeString(cipherText)
    if err != nil {
        return plainText, err
    }

    // 3.解密数据
    plainTextBt, err := rsa.DecryptPKCS1v15(rand.Reader, s.privateKey, cipherTextBt)
    if err != nil {
        return plainText, err
    }

    return plainTextBt, nil
}
