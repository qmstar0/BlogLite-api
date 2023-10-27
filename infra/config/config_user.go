package config

type User struct {
	CaptchaLength   int
	HashCaptchaSalt string

	JwtAuthTokenLifeDay    uint
	JwtCaptchaTokenLifeSec uint
}
