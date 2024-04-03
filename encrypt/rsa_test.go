package encrypt

import (
    "encoding/base64"
    "github.com/stretchr/testify/suite"
    "testing"
)

type RsaTestSuite struct {
    suite.Suite
    encrypt *RasEncrypt
    text    string
    data    string
}

// suite.Run
func TestOrderTestSuite(t *testing.T) {
    t.Log("TestOrderTestSuite")
    suite.Run(t, new(RsaTestSuite))
}

// before test
func (r *RsaTestSuite) SetupSuite() {
    r.T().Log("SetupSuite")
    r.encrypt = NewRasEncrypt()
    err := r.encrypt.GenKey(2048)
    r.Require().NoError(err)
    r.text = "zhou ge ge is a big pig"
    r.T().Logf(r.encrypt.privateKeyStr)
}

// after test
func (r *RsaTestSuite) TearDownSuite() {
    r.T().Log("TearDownSuite")
}

// before each test
func (r *RsaTestSuite) SetupTest() {
    r.T().Log("SetupTest")

    text, err := r.encrypt.Encrypt([]byte(r.text))
    r.Require().NoError(err)
    r.Require().NotNil(text)
    r.data = base64.StdEncoding.EncodeToString(text)
}

// after each test
func (r *RsaTestSuite) TearDownTest() {
    r.T().Log("TearDownTest")
}

func (r *RsaTestSuite) TestDecrypt() {
    r.T().Log("TestDecrypt")
    text, err := r.encrypt.Decrypt(r.data)
    r.Require().NoError(err)
    r.T().Log(string(text))
}

func (r *RsaTestSuite) TestEncrypt() {
    r.T().Log("TestEncrypt")
}
