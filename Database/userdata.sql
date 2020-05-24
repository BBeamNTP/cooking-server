-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Feb 26, 2020 at 04:55 PM
-- Server version: 10.4.12-MariaDB-1:10.4.12+maria~bionic-log
-- PHP Version: 7.4.1

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
-- Table structure for table `userdata`
--

CREATE TABLE `userdata` (
  `id` int(10) NOT NULL,
  `email` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `password` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `genderid` varchar(10) COLLATE utf8_unicode_ci DEFAULT NULL,
  `titleid` varchar(10) COLLATE utf8_unicode_ci DEFAULT NULL,
  `firstname` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `lastname` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL,
  `avatar` varchar(1000) COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_id` int(10) NOT NULL,
  `admin_id` int(10) NOT NULL,
  `signin_method` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created_date` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `userdata`
--

INSERT INTO `userdata` (`id`, `email`, `password`, `genderid`, `titleid`, `firstname`, `lastname`, `avatar`, `user_id`, `admin_id`, `signin_method`, `created_date`) VALUES
(1, 'admin@gmail.com', '$2a$14$0Y4PsiqqNtBhutTmTK3b8O/vS11HgYKgxqVMeoW80rlh2HW4HCEl.', '1', '1', 'Admin', 'Barn Arharn', 'http://localhost:9000/img/avatar/1.png', 0, 1, '', '2020-02-23 11:37:37'),
(4, 'fonarts05@gmail.com', '$2a$14$LMXjKUpGhoypvhqwxe2uMOWsvdw.zl4cMSvPuaWnQNHzdhZdN2R0i', '1', '1', 'Arram', 'Arm', 'http://localhost:9000/img/avatar/4.png', 2, 0, '', '2020-02-26 09:01:39');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `userdata`
--
ALTER TABLE `userdata`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `userdata`
--
ALTER TABLE `userdata`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
