package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UsersFollows_20200426_044254 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UsersFollows_20200426_044254{}
	m.Created = "20200426_044254"

	migration.Register("UsersFollows_20200426_044254", m)
}

// Run the migrations
func (m *UsersFollows_20200426_044254) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users_follows(`id` int(11) NOT NULL AUTO_INCREMENT,`target_type` varchar(128) NOT NULL,`target_id` int(11) DEFAULT NULL,`user_id` int(11),`updated_at` datetime NOT NULL,`created_at` datetime NOT NULL,PRIMARY KEY (`id`), FOREIGN KEY (`user_id`) REFERENCES users(id))")
}

// Reverse the migrations
func (m *UsersFollows_20200426_044254) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users_follows`")
	m.SQL("delete from migrations where name = \"UsersFollows_20200426_044254\"")
}
