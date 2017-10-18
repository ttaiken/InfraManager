CREATE TABLE IF NOT EXISTS `test`.`servers` (
`server_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
 `hostname` VARCHAR(45) NOT NULL DEFAULT 0,
 `ip1` int(10) UNSIGNED NOT NULL DEFAULT 0,
 `ip2` int(10) UNSIGNED NOT NULL DEFAULT 0,
 `ip3` int(10) UNSIGNED NOT NULL DEFAULT 0,   
`platform` VARCHAR(45) NOT NULL DEFAULT 0,
`os` VARCHAR(20) NOT NULL DEFAULT 0,
 PRIMARY KEY (`server_id`))
 AUTO_INCREMENT = 1
 DEFAULT CHARACTER SET = utf8
 COLLATE = utf8_general_ci;

INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('dns', inet_aton('8.8.8.8'), 'unknown', 'unknown');
INSERT INTO `test`.`servers` (`hostname`) VALUES ('www.baidu.com');
INSERT INTO `test`.`servers` (`hostname`) VALUES ('www.51cto.com');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('mysql', '3232237061', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p01', '3232237066', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p02', '3232237067', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p03', '3232237068', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p04', '3232237069', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p05', '3232237070', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p06', '3232237071', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p07', '3232237072', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p08', '3232237073', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p09', '3232237074', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p10', '3232237075', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p11', '3232237076', 'vmware', 'linux');
INSERT INTO `test`.`servers` (`hostname`, `ip1`, `platform`, `os`) VALUES ('p12', '3232237077', 'vmware', 'linux');
