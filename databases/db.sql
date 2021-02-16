# use ginblog
# source  + .sql文件路径
DROP TABLE IF EXISTS user;
CREATE TABLE user (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(30) NOT NULL,
    role INT DEFAULT 2
);
INSERT INTO user VALUES (1,'Peterliang','M52d3F/7u05g+g==',1);