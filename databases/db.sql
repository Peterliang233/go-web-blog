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
INSERT INTO user VALUES (1,'Peterliang','M52d3F/7u05g+g==','ncuyanping666@126.com',1);