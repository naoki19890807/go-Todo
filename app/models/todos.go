package models

import (
	"log"
	"time"
)

// todoの構造体
type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

// タスク追加
// usersのIDと紐づけるため、usersのメソッドして定義
func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content, 
		user_id, 
		created_at) values (?,?,?)`
	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// タスク取得(単一)
func GetTodo(id int) (todo Todo, err error) {
	//SELECT結果を格納する構造体
	todo = Todo{}
	cmd := `select * from todos where id = ?`
	row := Db.QueryRow(cmd, id)
	err = row.Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	return todo, err
}

// タスク取得(複数)
func GetTodos() (ts []Todo, err error) {
	cmd := `select * from todos`
	rows, _ := Db.Query(cmd)
	for rows.Next() {
		//１件データ取得
		t := Todo{}
		//複数データを格納できるようにスライスを定義する
		//ts = Todo{}

		err = rows.Scan(
			&t.ID,
			&t.Content,
			&t.UserID,
			&t.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		ts = append(ts, t)
	}
	rows.Close()

	return ts, err

}

// タスク取得(特定ユーザ)
func (u *User) GetTodosByUser() (u1 []Todo, err error) {
	cmd := `select * from todos where user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		//1件取得の構造体
		t := Todo{}
		err = rows.Scan(&t.ID, &t.Content, &t.UserID, &t.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		u1 = append(u1, t)
	}
	rows.Close()
	return u1, err
}

// タスク更新
func (t *Todo) UpdateTodo() (err error) {
	cmd := `update todos set content = ?, user_id = ? where id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// タスク削除
func (t *Todo) DeleteTodo() (err error) {
	cmd := `delete from todos where id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
