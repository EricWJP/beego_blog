package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Comments_20200426_043735 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Comments_20200426_043735{}
	m.Created = "20200426_043735"

	migration.Register("Comments_20200426_043735", m)
}

// Run the migrations
func (m *Comments_20200426_043735) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE comments(`id` int(11) NOT NULL AUTO_INCREMENT,`micropost_id` int(11),`user_id` int(11),`status` tinyint(1) NOT NULL,`like_count` int(11) DEFAULT NULL,`unlike_count` int(11) DEFAULT NULL,`comment` longtext  NOT NULL,`updated_at` datetime NOT NULL,`created_at` datetime NOT NULL,PRIMARY KEY (`id`), FOREIGN KEY (`micropost_id`) REFERENCES microposts(id), FOREIGN KEY (`user_id`) REFERENCES users(id))")
}

// Reverse the migrations
func (m *Comments_20200426_043735) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `comments`")
	m.SQL("delete from migrations where name = \"Comments_20200426_043735\"")
}
