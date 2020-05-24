CREATE DATABASE IF NOT EXISTS `cooking_server` CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `cooking_server`;

CREATE TABLE IF NOT EXISTS `cooking_server`.`points` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `menu_id` INT(10) NOT NULL ,
    `user_id` INT(10) NOT NULL ,
    `point` float(10) NOT NULL
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
