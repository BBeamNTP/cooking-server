-- phpMyAdmin SQL Dump
-- version 4.9.1
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Nov 18, 2019 at 05:05 PM
-- Server version: 10.4.8-MariaDB-1:10.4.8+maria~bionic-log
-- PHP Version: 7.2.23

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cooking_server`
--

-- --------------------------------------------------------

--
-- Table structure for table `ingredients`
--

CREATE TABLE `ingredients` (
  `id` int(10) NOT NULL,
  `ingredients_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ingredients_calories` int(10) DEFAULT NULL,
  `ingredients_type` int(10) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `ingredients`
--

INSERT INTO `ingredients` (`id`, `ingredients_name`, `ingredients_calories`, `ingredients_type`) VALUES
(1, 'ไข่ไก่', NULL, 1),
(2, 'หมูสับ', NULL, 1),
(3, 'หมูชิ้น', NULL, 1),
(4, 'กุ้งขาวสด', NULL, 1),
(5, 'ข้าวเปล่า', NULL, 1),
(6, 'ต้นหอม', NULL, 2),
(7, 'น้ำมันพืช', NULL, 3),
(8, 'ผมปรุงรส', NULL, 3),
(9, 'น้ำปลา', NULL, 3),
(10, 'มะนาว', NULL, 2),
(11, 'ซอสปรุงรส', NULL, 3),
(12, 'น้ำตาล', NULL, 3),
(13, 'เกลือ', NULL, 3),
(14, 'ข่า', NULL, 2),
(15, 'ตะไค้', NULL, 2),
(16, 'ใบมะกรูด', NULL, 2),
(17, 'น้ำเปล่า', NULL, 3),
(18, 'พริกเผา', NULL, 3);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `ingredients`
--
ALTER TABLE `ingredients`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `ingredients`
--
ALTER TABLE `ingredients`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
