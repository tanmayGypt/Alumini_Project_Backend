package helper

import "fmt"

func GenerateOTPMailBody(otp string) string {
	emailBody := fmt.Sprintf(`<p>Dear User,</p>

    <p>Welcome to the BPIT Alumni Website!</p>

    <p>To complete your registration, please use the following One-Time Password (OTP):</p>

    <h2>%s</h2>

    <p>This OTP is valid for the next 5 minutes. Please do not share this code with anyone.</p>

    <p>If you did not request this registration, please ignore this email.</p>

    <p>Thank you for joining our community!</p>

    <p>Best regards,</p>
    <p>BPIT Alumni Team</p>

    <hr>
    <p>Bhagwan Parshuram Institute of Technology</p>
    <p>Alumni Association</p>
    <p><a href="https://alumni.bpitindia.com/">BPIT Alumni Website</a></p>`, otp)
	return emailBody
}
