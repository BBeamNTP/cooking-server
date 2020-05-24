CREATE DATABASE IF NOT EXISTS `cooking_server` CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `cooking_server`;

CREATE TABLE IF NOT EXISTS `tokens` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT(10)  NOT NULL ,
    `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `token` VARCHAR(1000) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `signin_method` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

