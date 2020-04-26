package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Microposts_20200426_043110 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Microposts_20200426_043110{}
	m.Created = "20200426_043110"

	migration.Register("Microposts_20200426_043110", m)
}

// Run the migrations
func (m *Microposts_20200426_043110) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE microposts(`id` int(11) NOT NULL AUTO_INCREMENT,`title` varchar(256) NOT NULL,`content` longtext  NOT NULL,`user_id` int(11),`status` tinyint(1) NOT NULL,`comments_count` int(11) DEFAULT NULL,`followed_count` int(11) DEFAULT NULL,`updated_at` datetime NOT NULL,`created_at` datetime NOT NULL,PRIMARY KEY (`id`), FOREIGN KEY (`user_id`) REFERENCES users(id))")
}

// Reverse the migrations
func (m *Microposts_20200426_043110) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `microposts`")
	m.SQL("delete from migrations where name = \"Microposts_20200426_043110\"")
}
