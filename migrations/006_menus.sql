CREATE DATABASE IF NOT EXISTS `cooking_server` CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE `cooking_server`;

CREATE TABLE IF NOT EXISTS  `menus`(
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `menu_name` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,
    `category_id` VARCHAR(10) NOT NULL ,
    `point` float(10)  NOT NULL ,
    `user_id` VARCHAR(10)  NOT NULL ,
    `admin_id` VARCHAR(10) NOT NULL ,
    `menu_calories` float(10) NULL,
    `method` TEXT NOT NULL,
    `created_date` VARCHAR(10) NULL
    ) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE IF NOT EXISTS `foods` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `menu_id` INT(10) NOT NULL ,
    `ingredients_id` INT(10) NOT NULL,
    `quantity` float(10) NOT NULL
    ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `ingredients` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `ingredients_name` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,
    `ingredients_calories` float(10) NULL ,
    `ingredients_type`VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `quantity` float(10) NULL,
    `type`VARCHAR(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL
    ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `cooking_server`.`foodsimg` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `menu_id` INT(10) NOT NULL ,
    `href` VARCHAR(1000) NOT NULL
     ) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;