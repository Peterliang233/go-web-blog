CREATE DATABASE IF NOT EXISTS ginblog;
use ginblog;
# source  + .sql文件绝对路径
DROP TABLE IF EXISTS user;
CREATE TABLE user (
    id  int primary key auto_increment,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL
)engine = InnoDB
auto_increment = 1
default charset = utf8;
INSERT INTO user VALUES (1,'Peterliang','M52d3F/7u05g+g==');

DROP TABlE IF EXISTS tag;
CREATE TABLE tag (
    id int primary key auto_increment,
    name varchar(30)
)engine = InnoDB
auto_increment = 1
default charset =utf8;

DROP TABlE IF EXISTS category;
CREATE TABLE category (
    id int primary key auto_increment,
    name varchar(30)
)engine = InnoDB
 auto_increment = 1
 default charset =utf8;

DROP TABLE IF EXISTS article;
CREATE TABLE article (
    id int primary key auto_increment,
    tag_id int,
    category_id int,
    `desc` text,
    title varchar(66),
    content text,
    create_at datetime not null,
    delete_at datetime,
    update_at datetime,
    foreign key (tag_id) references tag (id),
    foreign key (category_id) references category (id)
)engine = InnoDB
 auto_increment = 1
 default charset = utf8;

DROP TABLE IF EXISTS comment;
CREATE TABLE comment(
    id int primary key auto_increment,
    article_id int,
    content text,
    create_at datetime not null,
    delete_at datetime,
    foreign key (article_id) references article (id)
)engine = InnoDB
 auto_increment = 1
 default charset = utf8;

DROP TABlE IF EXISTS `like`;
CREATE TABLE `like` (
    id int primary key auto_increment,
    article_id int,
    foreign key (article_id) references article (id)
)engine = InnoDB
 auto_increment = 1
 default charset =utf8;
