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
-- Table structure for table `foods`
--

CREATE TABLE IF NOT EXISTS  `foods` (
  `id` int(10) NOT NULL,
  `menu_id` int(10) NOT NULL,
  `ingredients_id` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `foods`
--

INSERT INTO `foods` (`id`, `menu_id`, `ingredients_id`) VALUES
(1, 1, 1),
(2, 1, 2),
(3, 1, 6),
(4, 1, 11),
(5, 1, 13),
(6, 2, 1),
(7, 2, 3),
(8, 2, 5),
(9, 2, 6),
(10, 2, 7),
(11, 2, 8),
(12, 2, 10),
(13, 2, 11),
(14, 2, 12),
(15, 2, 13),
(16, 3, 4),
(17, 3, 9),
(18, 3, 10),
(19, 3, 11),
(20, 3, 12),
(21, 3, 13),
(22, 3, 14),
(23, 3, 15),
(24, 3, 16),
(25, 3, 17),
(26, 3, 18),
(27, 4, 1),
(28, 4, 2),
(29, 4, 6),
(30, 4, 11),
(31, 4, 13),
(32, 4, 17);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `foods`
--
ALTER TABLE `foods`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `foods`
--
ALTER TABLE `foods`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=33;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
