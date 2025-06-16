package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"webIM/internal/models"
	"webIM/internal/pkg"
	"webIM/internal/service"
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
	var loginReq models.UserLogin

	// 优先解析 JSON
	if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Content-Type") == "application/json; charset=utf-8" {
		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			ret := Resp{
				Code: -1,
				Msg:  "参数解析失败",
				Data: nil,
			}
			RespWithMsg(w, ret)
			return
		}
	} else {
		// 再解析 url 参数
		r.ParseForm()
		loginReq.Mobile = r.PostForm.Get("mobile")
		loginReq.PlainPwd = r.PostForm.Get("password")
	}

	userService := service.UserService{}
	fmt.Println(loginReq)
	user, err := userService.Login(&loginReq)
	if err != nil {
		ret := Resp{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}
		fmt.Println(ret)
		RespWithMsg(w, ret)
		return
	}

	ret := Resp{
		Code: 0,
		Msg:  "Success",
		Data: UserData{
			ID:    int(user.ID),
			Token: user.Token,
		},
	}
	fmt.Println(ret)
	RespWithMsg(w, ret)
}

func userRegister(w http.ResponseWriter, r *http.Request) {
	var registerReq models.UserRegister

	// 优先解析 JSON
	if r.Header.Get("Content-Type") == "application/json" || r.Header.Get("Content-Type") == "application/json; charset=utf-8" {
		err := json.NewDecoder(r.Body).Decode(&registerReq)
		if err != nil {
			ret := Resp{
				Code: -1,
				Msg:  "参数解析失败",
				Data: nil,
			}
			RespWithMsg(w, ret)
			return
		}
	} else {
		// 路径参数
		r.ParseForm()
		registerReq.Mobile = r.PostForm.Get("mobile")
		registerReq.PlainPwd = r.PostForm.Get("plainPwd")
		registerReq.RePwd = r.PostForm.Get("rePwd")
		registerReq.Nickname = r.PostForm.Get("nickname")
		registerReq.Avatar = r.PostForm.Get("avatar")
		registerReq.Sex = r.PostForm.Get("sex")
	}

	userService := service.UserService{}
	user, err := userService.Register(&registerReq)
	if err != nil {
		ret := Resp{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		}
		RespWithMsg(w, ret)
		return
	}

	ret := Resp{
		Code: 0,
		Msg:  "Success",
		Data: UserData{
			ID:    int(user.ID),
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

func registerView() {
	// 获取 view 路径下的所有模版html
	ts, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatalln(err)
		return
	}
	for _, t := range ts.Templates() {
		tn := t.Name()
		// 注册模版html为对应的请求地址
		http.HandleFunc("/"+tn, func(w http.ResponseWriter, r *http.Request) {
			err = ts.ExecuteTemplate(w, tn, nil)
			if err != nil {
				log.Fatalln(err)
			}
		})
	}

}

func main() {
	pkg.InitMysql()

	http.HandleFunc("/", home)
	http.HandleFunc("/user/login", userLogin)
	http.HandleFunc("/user/register", userRegister)

	// 设置静态文件路径
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// 注册前端页面
	registerView()

	// 启动服务器
	http.ListenAndServe(":9090", nil)
}
