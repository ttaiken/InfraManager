CREATE TABLE IF NOT EXISTS `test`.`user` (
 `userid` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
 `username` VARCHAR(45) NOT NULL COMMENT '用户名称',
`userpw` VARCHAR(45) NOT NULL COMMENT '用户名称',
  `role` VARCHAR(45) NOT NULL COMMENT '权限',
`phone` VARCHAR(45) NOT NULL COMMENT '电话',
`maill` VARCHAR(45) NOT NULL COMMENT '邮件',
`qq` VARCHAR(45) NOT NULL COMMENT 'QQ',
`微信` VARCHAR(45) NOT NULL COMMENT '微信', `
userage` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户年龄', 
 `usersex` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别', 
 PRIMARY KEY (`userid`))
 AUTO_INCREMENT = 1 
 DEFAULT CHARACTER SET = utf8 
 COLLATE = utf8_general_ci 
 COMMENT = '用户表';

INSERT INTO `user` (username, userpw) VALUES ('admin', 'admin');
INSERT INTO `user` (username, userpw) VALUES ('byron', 'pa55word');
INSERT INTO `user` (username, userpw) VALUES ('tom', '123456');
