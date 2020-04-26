package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UsersComments_20200426_044112 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UsersComments_20200426_044112{}
	m.Created = "20200426_044112"

	migration.Register("UsersComments_20200426_044112", m)
}

// Run the migrations
func (m *UsersComments_20200426_044112) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE `users_comments`(`id` int(11) NOT NULL AUTO_INCREMENT,`comment_id` int(11),`user_id` int(11),`flag` tinyint(1) NOT NULL,`updated_at` datetime NOT NULL,`created_at` datetime NOT NULL,PRIMARY KEY (`id`), FOREIGN KEY (`comment_id`) REFERENCES comments(id), FOREIGN KEY (`user_id`) REFERENCES users(id))")
}

// Reverse the migrations
func (m *UsersComments_20200426_044112) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users_comments`")
	m.SQL("delete from migrations where name = \"UsersComments_20200426_044112\"")
}
