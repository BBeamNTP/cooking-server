-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: mariadb
-- Generation Time: Feb 26, 2020 at 04:56 PM
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
-- Table structure for table `foodsimg`
--

CREATE TABLE `foodsimg` (
  `id` int(10) NOT NULL,
  `menu_id` int(10) NOT NULL,
  `href` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `foodsimg`
--

INSERT INTO `foodsimg` (`id`, `menu_id`, `href`) VALUES
(1, 1, 'http://localhost:9000/img/imgfood/menuid1/เต้าเจียวหล่น.jpg'),
(2, 1, 'http://localhost:9000/img/imgfood/menuid1/เต้าเจียวหล่น1.jpg'),
(3, 2, 'http://localhost:9000/img/imgfood/menuid2/ไข่น้ำ.jpg'),
(4, 2, 'http://localhost:9000/img/imgfood/menuid2/ไข่น้ำ1.jpg'),
(5, 2, 'http://localhost:9000/img/imgfood/menuid2/ไข่น้ำผักชีโรย.jpg'),
(6, 3, 'http://localhost:9000/img/imgfood/menuid3/1382879944-1381419602-o.jpg'),
(7, 3, 'http://localhost:9000/img/imgfood/menuid3/1382880014-2553602995-o.jpg'),
(8, 4, 'http://localhost:9000/img/imgfood/menuid4/เตี๋ยวหมูเด้งต้มยำไข่ยางมะตูม.jpg'),
(9, 5, 'http://localhost:9000/img/imgfood/menuid5/กระเพรารวมมิตร กุ้ง หมึก ไข่ดาว.jpg'),
(12, 7, 'http://localhost:9000/img/imgfood/menuid7/กระดูกหมู.jpg'),
(22, 8, 'http://localhost:9000/img/imgfood/menuid8/เส้นเล็กต้มยำ.jpg'),
(23, 9, 'http://localhost:9000/img/imgfood/menuid9/กุ้งทอดกระเทียม.jpg'),
(24, 9, 'http://localhost:9000/img/imgfood/menuid9/กุ้งทอดกระเทียม1.jpg'),
(25, 10, 'http://localhost:9000/img/imgfood/menuid10/ข้าวไก่กระเทียม1.jpg'),
(26, 11, 'http://localhost:9000/img/imgfood/menuid11/EGG-AND-RICE.jpg'),
(27, 11, 'http://localhost:9000/img/imgfood/menuid11/ข้าวไข่ดาว2ฟอง.jpg'),
(28, 11, 'http://localhost:9000/img/imgfood/menuid11/รูป-หลัก-ของ-สูตร-ข้าวไข่ดาว-2-ฟองเมนูประหยัด.jpg'),
(29, 6, 'http://localhost:9000/img/imgfood/menuid6/101526.jpg'),
(30, 6, 'http://localhost:9000/img/imgfood/menuid6/DR-oxNyUEAAiTtk.jpg'),
(31, 12, 'http://localhost:9000/img/imgfood/menuid12/ข้าวผัดหมู1.jpg'),
(32, 13, 'http://localhost:9000/img/imgfood/menuid13/ข้าวผัดหมูไข่ดาว1.jpg'),
(33, 14, 'http://localhost:9000/img/imgfood/menuid14/ข้าวหมูผัดโทริยากิ.jpg'),
(34, 15, 'http://localhost:9000/img/imgfood/menuid15/ซุปเห็ด.jpg'),
(35, 16, 'http://localhost:9000/img/imgfood/menuid16/ต้มเลือดหมู.jpg'),
(36, 16, 'http://localhost:9000/img/imgfood/menuid16/ต้มเลือดหมู1.jpg'),
(37, 17, 'http://localhost:9000/img/imgfood/menuid17/ต้มข่าไก่.jpg'),
(38, 17, 'http://localhost:9000/img/imgfood/menuid17/ต้มข่าไก่1.jpg'),
(39, 18, 'http://localhost:9000/img/imgfood/menuid18/ต้มยำกุ้งน้ำข้น.jpg'),
(40, 18, 'http://localhost:9000/img/imgfood/menuid18/ต้มยำกุ้งน้ำข้น1.jpg'),
(41, 19, 'http://localhost:9000/img/imgfood/menuid19/หอยแมลงภู่.jpg'),
(42, 19, 'http://localhost:9000/img/imgfood/menuid19/หอยแมลงภู่1.jpg'),
(43, 20, 'http://localhost:9000/img/imgfood/menuid20/ทอดกระเทียม.jpg'),
(44, 20, 'http://localhost:9000/img/imgfood/menuid20/ปลากระพงทอดกระเทียม.jpg'),
(45, 21, 'http://localhost:9000/img/imgfood/menuid21/ปลากระพงทอดน้ำปลา.jpg'),
(46, 22, 'http://localhost:9000/img/imgfood/menuid22/ปลากระพงนึ่งมะนาว.jpg'),
(47, 22, 'http://localhost:9000/img/imgfood/menuid22/ปลากระพงนึ่งมะนาว1.jpg'),
(48, 23, 'http://localhost:9000/img/imgfood/menuid23/ผัดบุ้งไฟแดง.png'),
(49, 24, 'http://localhost:9000/img/imgfood/menuid24/ผัดผักกระเฉด.jpg'),
(50, 25, 'http://localhost:9000/img/imgfood/menuid25/ผัดสตอกุ้ง.png'),
(51, 25, 'http://localhost:9000/img/imgfood/menuid25/ผัดสตอกุ้ง1.jpg'),
(52, 25, 'http://localhost:9000/img/imgfood/menuid25/ผัดสะตอกุ้ง.png'),
(53, 26, 'http://localhost:9000/img/imgfood/menuid26/ยำปลากระป๋องโรซ่า1.jpg'),
(54, 27, 'http://localhost:9000/img/imgfood/menuid27/ยำขนมจีนหมูยอไข่แดง.jpg'),
(55, 27, 'http://localhost:9000/img/imgfood/menuid27/ยำขนมจีนหมูยอไข่แดง1.jpg'),
(56, 28, 'http://localhost:9000/img/imgfood/menuid28/ยำหมูยอไข่แดง.jpg'),
(57, 28, 'http://localhost:9000/img/imgfood/menuid28/ยำหมูยอไข่แดง1.jpg'),
(58, 29, 'http://localhost:9000/img/imgfood/menuid29/ยำปลาหมึก.jpg'),
(59, 29, 'http://localhost:9000/img/imgfood/menuid29/ยำปลาหมึก1.jpg'),
(60, 30, 'http://localhost:9000/img/imgfood/menuid30/ผัดมาม่ากุ้ง.png'),
(61, 31, 'http://localhost:9000/img/imgfood/menuid31/ยำมาม่าหมูยอ.jpg'),
(62, 32, 'http://localhost:9000/img/imgfood/menuid32/เส้นใหญ่ราดหน้า.jpg'),
(63, 32, 'http://localhost:9000/img/imgfood/menuid32/ราดหน้าเส้นใหญ่1.jpg'),
(64, 33, 'http://localhost:9000/img/imgfood/menuid33/สเต็กหมูสันคอ.jpg'),
(65, 34, 'http://localhost:9000/img/imgfood/menuid34/หอยนิวซีแลนด์อบชีส.jpg'),
(66, 34, 'http://localhost:9000/img/imgfood/menuid34/หอยนิวซีแลนด์อบชีส1.jpg'),
(67, 35, 'http://localhost:9000/img/imgfood/menuid35/สปาเก็ตตี้หอยลาย.jpg'),
(68, 35, 'http://localhost:9000/img/imgfood/menuid35/สปาหอยลาย1.jpg'),
(69, 36, 'http://localhost:9000/img/imgfood/menuid36/สปาเก็ตตี้คาโบนาล่า.jpg'),
(70, 36, 'http://localhost:9000/img/imgfood/menuid36/สปาเก็ตตี้คาโบนาล่า1.jpg'),
(71, 37, 'http://localhost:9000/img/imgfood/menuid37/สปาเก็ตตี้ขี้เมาทะเล.jpg'),
(72, 37, 'http://localhost:9000/img/imgfood/menuid37/สปาเก็ตตี้ขี้เมาทะเล1.jpg'),
(73, 38, 'http://localhost:9000/img/imgfood/menuid38/สปาเก็ตตี้เบ1.jpg'),
(74, 38, 'http://localhost:9000/img/imgfood/menuid38/สปาเก็ตตี้เบคอน.jpg'),
(75, 39, 'http://localhost:9000/img/imgfood/menuid39/สปาเก็ตตี้เขียวหวานไก่.jpg'),
(76, 39, 'http://localhost:9000/img/imgfood/menuid39/สปาเก็ตตี้เขียวหวานไก่1.jpg'),
(77, 40, 'http://localhost:9000/img/imgfood/menuid40/ชีสบอล.jpg'),
(78, 41, 'http://localhost:9000/img/imgfood/menuid41/บาบีคิวไก่.jpg'),
(79, 42, 'http://localhost:9000/img/imgfood/menuid42/ไก่ชุบแป้งทอด.jpg'),
(80, 43, 'http://localhost:9000/img/imgfood/menuid43/เมี่ยงหอยแครง.jpg'),
(81, 44, 'http://localhost:9000/img/imgfood/menuid44/มัสมั่นไก่.jpg'),
(82, 45, 'http://localhost:9000/img/imgfood/menuid45/1.jpg'),
(83, 46, 'http://localhost:9000/img/imgfood/menuid46/ไข่พะโล้.jpg'),
(84, 47, 'http://localhost:9000/img/imgfood/menuid47/แกงส้มสายบัว.jpg'),
(85, 48, 'http://localhost:9000/img/imgfood/menuid48/1.jpg'),
(86, 49, 'http://localhost:9000/img/imgfood/menuid49/เทโพหมู.jpg'),
(87, 50, 'http://localhost:9000/img/imgfood/menuid50/แกงเขียวหวานหมู.jpg'),
(88, 51, 'http://localhost:9000/img/imgfood/menuid51/แกงเขียวหวานทะเล.jpg'),
(89, 52, 'http://localhost:9000/img/imgfood/menuid52/1.jpg'),
(90, 53, 'http://localhost:9000/img/imgfood/menuid53/ลูกชิ้นกุ้งระเบิด.jpg'),
(91, 54, 'http://localhost:9000/img/imgfood/menuid54/ปีกไก่ทอดน้ำปลา.jpg'),
(92, 55, 'http://localhost:9000/img/imgfood/menuid55/ปีกไก่ทอดเกลือ.jpg'),
(93, 56, 'http://localhost:9000/img/imgfood/menuid56/ปีกไก่ชุบแป้งทอด.jpg'),
(94, 57, 'http://localhost:9000/img/imgfood/menuid57/บาบีคิวหมู.jpg'),
(95, 58, 'http://localhost:9000/img/imgfood/menuid58/ทอดมันหมู.jpg'),
(96, 59, 'http://localhost:9000/img/imgfood/menuid59/ข้าวตู.jpg'),
(97, 60, 'http://localhost:9000/img/imgfood/menuid60/คั่วไก่ใบชะพลู.jpg'),
(98, 61, 'http://localhost:9000/img/imgfood/menuid61/ซุปอุด้งน้ำใส.jpg'),
(99, 62, 'http://localhost:9000/img/imgfood/menuid62/ซุปฟักทอง.jpg'),
(100, 63, 'http://localhost:9000/img/imgfood/menuid63/ซุปครีมมันฝรั่งหอยลาย.jpg'),
(101, 64, 'http://localhost:9000/img/imgfood/menuid64/ซุปไก่ข้น.jpg'),
(102, 65, 'http://localhost:9000/img/imgfood/menuid65/ซุปไก่ข้น.jpg'),
(103, 66, 'http://localhost:9000/img/imgfood/menuid66/ไข่น้ำอนามัย.jpg'),
(104, 67, 'http://localhost:9000/img/imgfood/menuid67/โป๊แตกต้มยำทะเลน้ำใส.jpg'),
(105, 68, 'http://localhost:9000/img/imgfood/menuid68/ซุปหอมใหญ่.jpg'),
(106, 69, 'http://localhost:9000/img/imgfood/menuid69/ซุปไก่อิสลาม.jpg'),
(107, 70, 'http://localhost:9000/img/imgfood/menuid70/หอยตลับ.jpg'),
(108, 71, 'http://localhost:9000/img/imgfood/menuid71/ต้มส้มปลาหมึกยัดไส้.jpg'),
(109, 72, 'http://localhost:9000/img/imgfood/menuid72/ต้มยำปลานิล.jpg'),
(110, 73, 'http://localhost:9000/img/imgfood/menuid73/ต้มยำปลากระป๋อง.jpg'),
(111, 74, 'http://localhost:9000/img/imgfood/menuid74/ต้มผักกาดดองหมูสามชั้น.jpg'),
(112, 75, 'http://localhost:9000/img/imgfood/menuid75/ต้มส้มปลาทูสด.jpg'),
(113, 76, 'http://localhost:9000/img/imgfood/menuid76/ต้มแซ่บปลาทู.jpg'),
(114, 77, 'http://localhost:9000/img/imgfood/menuid77/ลาบปลาหมึก.jpg'),
(115, 78, 'http://localhost:9000/img/imgfood/menuid78/ยำผักบุ้งกุ้งสด.jpg'),
(116, 79, 'http://localhost:9000/img/imgfood/menuid79/ยำไข่แดง.jpg'),
(117, 80, 'http://localhost:9000/img/imgfood/menuid80/ก้อยกุ้ง.jpg'),
(118, 81, 'http://localhost:9000/img/imgfood/menuid81/101776.jpg'),
(119, 82, 'http://localhost:9000/img/imgfood/menuid82/101775.jpg');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `foodsimg`
--
ALTER TABLE `foodsimg`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `foodsimg`
--
ALTER TABLE `foodsimg`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=120;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
