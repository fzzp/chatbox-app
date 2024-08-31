package config

import (
	"chatbox-app/lib/token"

	ut "github.com/go-playground/universal-translator"
)

const (
	Dev = "development"
	Pro = "production"
)

// App 提供给全局使用
var App app

type app struct {
	Env      Settings
	Trans    ut.Translator
	JwtMaker token.Maker
}

func SetupGlobalApp(env Settings, jwtMaker token.Maker, trans ut.Translator) {
	App = app{
		Env:      env,
		Trans:    trans,
		JwtMaker: jwtMaker,
	}
}
