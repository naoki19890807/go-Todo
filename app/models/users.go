package models

import (
	"log"
	"time"
)

// 構造体生成(User)
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos     []Todo
}

// 構造体生成(Session)
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    string
	CreatedAt time.Time
}

// ユーザー作成----------------------------
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ユーザー取得-----------------------
func GetUser(id int) (u User, err error) {
	//取得したデータを格納する構造体
	u = User{}
	cmd := `select * from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&u.ID,
		&u.UUID,
		&u.Name,
		&u.Email,
		&u.PassWord,
		&u.CreatedAt)

	if err != nil {
		log.Fatalln(err)
	}
	return u, err
}

// ユーザー更新
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ユーザー削除
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// session関連---------------------------------------------------------
// Formに入力されたemailアドレスに合致するデータを取得する
func GetUserByEmail(email string) (user User, err error) {
	cmd := `select * from users where email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

// Formに紐づくデータをsessionsテーブルに格納
func (u *User) CreateSession() (session Session, err error) {
	cmd1 := `insert into sessions (
		uuid, 
		email, 
		user_id, 
		created_at) values (?,?,?,?)`
	_, err = Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}
	//上記で格納したデータを取得
	cmd2 := `select * from sessions where email = ? and user_id = ?`
	err = Db.QueryRow(cmd2, u.Email, u.ID).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserID,
		&session.CreatedAt,
	)
	return session, err

}

// sessionDB確認(登録しているか否か)
func (sess *Session) CheckSession() (valid bool, err error) {
	valid = true
	cmd := `select * from sessions where uuid = ?`
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Email,
		&sess.UserID,
		&sess.CreatedAt,
	)
	if err != nil {
		valid = false
		return
	}
	return valid, err
}

// ログアウト処理
func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = ?`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (sess *Session) GetUserBySession() (user User, err error) {
	cmd := `select id, uuid, name, email, created_at from users where id = ?`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)
	return user, err

}
