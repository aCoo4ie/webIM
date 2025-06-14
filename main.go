package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 定义公共的返回结构体
type Resp struct {
	Code int         `json:"id"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 定义用户数据结构
type UserData struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

func RespWithMsg(w http.ResponseWriter, r Resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ret, err := json.Marshal(r)
	if err != nil {
		log.Println(err)
	}
	w.Write(ret)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	// 获取路径参数
	r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	token := r.PostForm.Get("token")

	// 校验
	if mobile != "19912345678" || token != "test" {
		// 失败
		ret := Resp{
			Code: -1,
			Msg:  "Error",
			Data: nil,
		}
		RespWithMsg(w, ret)
		return
	}

	// 成功
	ret := Resp{
		Code: 0,
		Msg:  "Success",
		Data: UserData{
			ID:    1234,
			Token: "test",
		},
	}
	RespWithMsg(w, ret)
}

func home(w http.ResponseWriter, r *http.Request) {
	ret := Resp{
		Code: 0,
		Msg:  "Success",
		Data: "Hello",
	}
	RespWithMsg(w, ret)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user/login", userLogin)

	// 启动服务器
	http.ListenAndServe(":9090", nil)
}
