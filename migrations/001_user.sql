CREATE DATABASE IF NOT EXISTS `cooking_server` CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `cooking_server`;

CREATE TABLE IF NOT EXISTS  `userdata`(
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `genderid` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `titleid` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `firstname` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `lastname` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `avatar` varchar(1000) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `user_id` INT(10)  NOT NULL ,
    `admin_id` INT(10) NOT NULL ,
    `signin_method` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `created_date` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
