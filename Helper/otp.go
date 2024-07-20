package helper

import (
	"crypto/rand"
	"errors"
	"math/big"
	models "my-go-backend/Models"
	database "my-go-backend/config"
	"time"
)

func GenerateOTP(length int) (string, error) {
	const charset = "0123456789"
	otp := make([]byte, length)
	for i := range otp {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		otp[i] = charset[num.Int64()]
	}
	return string(otp), nil
}

func VerifyOTP(email string, inputOTP string) error {
	var otp models.OTP
	if err := database.DB.Where("email = ? AND code = ?", email, inputOTP).First(&otp).Error; err != nil {
		return errors.New("invalid OTP")
	}

	if time.Now().After(otp.ExpiresAt) {
		database.DB.Delete(&otp)
		return errors.New("OTP has expired")
	}

	database.DB.Delete(&otp)
	return nil
}
