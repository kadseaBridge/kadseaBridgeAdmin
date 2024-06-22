package service

import (
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

type GoogleAuth2 struct {
}

func NewGoogleAuth2() *GoogleAuth2 {
	return &GoogleAuth2{}
}

func (this *GoogleAuth2) GenerateSecretAndQRCode(userName string) (string, []byte, error) {
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

func (this *GoogleAuth2) VerifyCode(code string, secret string) bool {
	return totp.Validate(code, secret)
}
