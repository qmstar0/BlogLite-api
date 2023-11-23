package config

type User struct {
	CaptchaLength   int
	HashCaptchaSalt string

	JwtAuthTokenLifeDay    int64
	JwtCaptchaTokenLifeSec int64
}
