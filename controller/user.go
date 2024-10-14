package controller

import (
	"fmt"
	"im/model"
	"im/service"
	"im/util"
	"math/rand"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	passwd := r.PostForm.Get("passwd")

	loginOk := false
	if mobile == "18600000000" && passwd == "123456" {
		loginOk = true
	}

	if loginOk {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		util.RespOk(w, data, "")
	} else {
		util.RespFail(w, "密码不正确")
	}
}

var userService service.UserService

func UserRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	plainPwd := r.PostForm.Get("passwd")
	nickName := fmt.Sprintf("user%06d", rand.Int31())
	avater := ""
	sex := model.SEX_UNKNOW

	user, err := userService.Register(mobile, plainPwd, nickName, avater, sex)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "")

	}
}
