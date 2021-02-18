# use ginblog
# source  + .sql文件绝对路径
DROP TABLE IF EXISTS user;
CREATE TABLE user (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
    email VARCHAR(100) NOT NULL,
    role INT DEFAULT 2
);
DROP TABLE IF EXISTS email;
CREATE TABLE email(
  id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  email_name VARCHAR(50) NOT NULL,
  active TINYINT(1) DEFAULT 0
);
INSERT INTO user VALUES (1,'Peterliang','M52d3F/7u05g+g==','ncuyanping666@126.com',1);
INSERT INTO email VALUES (1,'ncuyanping666@126.com',1);