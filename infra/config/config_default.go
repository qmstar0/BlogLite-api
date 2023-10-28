package config

func InitDefaultConfig() {
	Conf.Article = &Article{
		DefaultLimit:          "15",
		TagListKey:            "all:tags",
		CateListKey:           "all:cate",
		ArticleIndexKey:       "index:article",
		TagArticleIndexKey:    "index:tag:article",
		CateArticleIndexKey:   "index:cate:article",
		SystemIndexKey:        "index:system",
		ArticleDetailIndexKey: "index:articleDetail",
	}
	//Conf.User = &User{
	//	CaptchaLength:          6,
	//	HashCaptchaSalt:        "hash-captcha-salt-180",
	//	JwtAuthTokenLifeDay:    180,
	//	JwtCaptchaTokenLifeSec: 180,
	//}
	//Conf.System = &System{
	//	Theme:        "",
	//	Title:        "blog",
	//	Keywords:     "blog,yvye,于野,探索日志,博客",
	//	Description:  "于野的探索日志",
	//	RecordNumber: "",
	//}
	Conf.Jwt = &JWT{
		PrivateKeyPath: "system/private_key.pem",
		PublicKeyPath:  "system/public_key.pem",
	}
	Conf.Logger = &logger{
		Level:      "Info",
		OutputPath: "logs",
	}
}
