package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20200425_074627 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20200425_074627{}
	m.Created = "20200425_074627"

	migration.Register("Users_20200425_074627", m)
}

// Run the migrations
func (m *Users_20200425_074627) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE users(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`gender` varchar(128) NOT NULL,`phone` varchar(128) NOT NULL,`email` varchar(128) NOT NULL,`password` varchar(256) NOT NULL,`status` tinyint(1) NOT NULL,`is_admin` tinyint(1) NOT NULL,`followed_count` int(11) DEFAULT NULL,`microposts_count` int(11) DEFAULT NULL,`comment` longtext  NOT NULL,`updated_at` datetime NOT NULL,`created_at` datetime NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Users_20200425_074627) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `users`")
	m.SQL("delete from migrations where name = \"Users_20200425_074627\"")
}
