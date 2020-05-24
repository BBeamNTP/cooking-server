CREATE DATABASE IF NOT EXISTS `cooking_server` CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `cooking_server`;

CREATE TABLE IF NOT EXISTS `otps` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `otp` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci,
    `start_time` DATETIME COLLATE utf8_unicode_ci,
    `end_time` DATETIME COLLATE utf8_unicode_ci
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
