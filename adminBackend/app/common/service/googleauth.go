package service

import (
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

type GoogleAuth struct {
}

func NewGoogleAuth() *GoogleAuth {
	return &GoogleAuth{}
}

func (this *GoogleAuth) GenerateSecretAndQRCode(userName string) (string, []byte, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "KADSAEBRIDGE",
		AccountName: userName,
	})
	if err != nil {
		return "", nil, err
	}

	secret := key.Secret()
	otpURL := key.URL()
	png, _ := qrcode.Encode(otpURL, qrcode.Medium, 256)

	return secret, png, nil
}

func (this *GoogleAuth) VerifyCode(code string, secret string) bool {
	return totp.Validate(code, secret)
}
