package controllers

import (
	"log"
	"net/http"
	"text/template"
	"todo_app/app/models"
)

// ハンドラー処理----------------------------------------

// 登録処理
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			//解析、実行
			t, err := template.ParseFiles("app/views/templates/signup.html")
			if err != nil {
				log.Println(err)
			}
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		//入力フォームの解析
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//Formより受け取った名前、Email、パスワードをuserの構造体に格納
		user := models.User{
			//name属性のnameより取得
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		//usersテーブルに格納
		if err = user.CreateUser(); err != nil {
			log.Panicln(err)
		}
		//完了後、Topページにリダイレクトする
		http.Redirect(w, r, "/", 302)
	}
}

// ログイン処理

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			//解析、実行
			t, err := template.ParseFiles("app/views/templates/login.html")
			if err != nil {
				log.Println(err)
			}
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		//入力フォームの解析
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		//Formで入力したメールアドレスを取得
		user, err := models.GetUserByEmail(r.PostFormValue("email"))
		if err != nil {
			log.Println(err)
			//ログイン不可の場合、ログインページにリダイレクトする
			http.Redirect(w, r, "/login", 302)
		}
		if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession()
			if err != nil {
				log.Println(err)
			}
			//cookieに保存
			cookie := http.Cookie{
				Name:     "_cookie",
				Value:    session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", 302)
		} else { //パスワード不一致
			http.Redirect(w, r, "/login", 302)
		}
	}
}

// ログアウト処理
func logout(w http.ResponseWriter, r *http.Request) {
	//ブラウザからcookie取得
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}
