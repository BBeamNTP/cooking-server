CREATE DATABASE IF NOT EXISTS `cooking_server` CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `cooking_server`;

CREATE TABLE IF NOT EXISTS `genders` (
    `id` varchar(10) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `gender` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci

) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
