package utilities

import (
	"database/sql"
	"fmt"
	"log"
)

func UsersAdd(db *sql.DB, email string, username string, password string) error {
	_, err := db.Exec(`Insert INTO users (email,username,password) VALUES (?,?,?)`, email, username, password)
	return err

}

func UsersGet(db *sql.DB, id int) {
	rows, err := db.Query(`SELECT users.id,users.username,users.email,users.password,users.register_date,users.birth_date FROM users WHERE users.id = ?`, id)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var users []GetUser
	for rows.Next() {
		var u GetUser

		err := rows.Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.Register_date, &u.Birth_date)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}

func UsersGetByEmail(db *sql.DB, email string) GetUser {
	rows, err := db.Query(`SELECT users.id,users.username FROM users WHERE users.email = ?`, email)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var users GetUser
	for rows.Next() {
		var u GetUser

		err := rows.Scan(&u.Id, &u.Username)
		if err != nil {
			log.Fatal(err)
		}
		users = u

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func UsersGetAll(db *sql.DB) []GetUsersAll {
	rows, err := db.Query(`SELECT users.id,users.username,users.email,users.password,users.register_date FROM users`)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var users []GetUsersAll
	for rows.Next() {
		var u GetUsersAll

		err := rows.Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.Register_date)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users
}

func UsersUpdate(db *sql.DB, id int, email string, username string, password string, description string, birth_date string) {
	if description == "" {
		description = "Not set"
	}
	_, err := db.Exec(`UPDATE users SET email = ?, username = ?, password = ?,description =?,birth_date=? WHERE id = ?`, email, username, password, description, birth_date, id)
	if err != nil {
		panic(err.Error())
	}
}

func UsersHash(db *sql.DB, id int, email string, username string, password string, description string, birth_date string) {

}

func UsersDelete(db *sql.DB, id int) {
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)

	if err != nil {
		panic(err.Error())
	}
}

func UserRoles(db *sql.DB, id int) {
	rows, err := db.Query(`SELECT roles.name,roles.color FROM users_roles INNER JOIN roles ON users_roles.role_id = roles.id WHERE users_roles.user_id = ?`, id)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var roles []RoleUser
	for rows.Next() {
		var r RoleUser

		err := rows.Scan(&r.Name, &r.Color)
		if err != nil {
			log.Fatal(err)
		}
		roles = append(roles, r)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(roles)
}

func UserFollowedTopics(db *sql.DB, id int) {
	rows, err := db.Query(`SELECT topic.name FROM users_followed_topics INNER JOIN topic ON users_followed_topics.topic_id = topic.id WHERE users_followed_topics.user_id = ?`, id)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var topics []string
	for rows.Next() {
		var t string

		err := rows.Scan(&t)
		if err != nil {
			log.Fatal(err)
		}
		topics = append(topics, t)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(topics)
}

func UserMessagesInteractions(db *sql.DB, id int) {
	rows, err := db.Query(`SELECT messages.id,messages.content,messages.creation_date,messages.user_id,messages.topic_id,users_messages_interactions.status FROM users_messages_interactions INNER JOIN messages ON users_messages_interactions.message_id = messages.id WHERE users_messages_interactions.user_id = ?`, id)

	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var messages []MessageUser
	for rows.Next() {
		var m MessageUser

		err := rows.Scan(&m.Id, &m.Content, &m.Creation_date, &m.User_id, &m.Topic_id, &m.Status)
		if err != nil {
			log.Fatal(err)
		}
		messages = append(messages, m)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
