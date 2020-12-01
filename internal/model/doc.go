package model

import (
	"github.com/spf13/viper"
)

var APIdocZh Doc
var APIdocEn Doc
var APIdocCn Doc
var Cn *viper.Viper
var Zh *viper.Viper
var En *viper.Viper

type Doc struct {
	Title   string
	Version string
	Menu    []DocMenu
	Tag     []DocTag
}

type DocMenu struct {
	Name string
	Link string
}

type DocTag struct {
	Name  string
	Funcs []DocFunc
}

type DocFunc struct {
	FuncName    string
	Method      string
	Path        string
	Description string
	Explanation string
	Requests    []DocRequest
	Responses   []DocResponse
}

type DocRequest struct {
	Area string
	Body []DocBody
}

type DocBody struct {
	Name        string
	Type        string
	Required    bool
	Description string
}

type DocResponse struct {
	StateCode   string
	Description string
	Message     string
}
