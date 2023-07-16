-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Jul 16, 2023 at 02:17 PM
-- Server version: 5.7.39
-- PHP Version: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `sanbercode_final-project`
--

-- --------------------------------------------------------

--
-- Table structure for table `admins`
--

CREATE TABLE `admins` (
  `id` varchar(191) NOT NULL,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `admins`
--

INSERT INTO `admins` (`id`, `username`, `password`, `created_at`, `updated_at`) VALUES
('397ac10e-b41f-4afc-93b7-5079b8a48ef3', 'caesar', '$2a$10$X0JQuELlPoXIJtb183t0lO/Vs7RZHNBqD7.YEFDTesD28qd/AO3oq', '2023-07-16 06:52:51.269', '2023-07-16 06:53:24.727'),
('424416fa-311a-4ad7-96f2-34b95d6ba1c0', 'malta', '$2a$10$8T297h/r9nYiJj7co9UC3uCtJ9OM.f2V5xxUeG89.7.UqggGfoqTC', '2023-07-16 06:33:38.698', '2023-07-16 06:33:38.698');

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` varchar(191) NOT NULL,
  `name` varchar(191) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `created_at`, `updated_at`) VALUES
('65b4ab22-3b9c-40cc-b7e0-911dbce8e65b', 'fashion', '2023-07-16 07:05:00.938', '2023-07-16 07:05:00.938'),
('c960d7fc-b4cd-4eba-9d42-57ed15b048bf', 'furniture', '2023-07-16 07:07:46.405', '2023-07-16 07:07:46.405'),
('e1670d03-5c43-4c80-93ac-2dd850a25c9a', 'electronic', '2023-07-16 07:02:39.571', '2023-07-16 07:02:39.571');

-- --------------------------------------------------------

--
-- Table structure for table `invoices`
--

CREATE TABLE `invoices` (
  `id` varchar(191) NOT NULL,
  `user_id` varchar(191) DEFAULT NULL,
  `total_price` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `invoices`
--

INSERT INTO `invoices` (`id`, `user_id`, `total_price`, `created_at`, `updated_at`) VALUES
('70d11b73-622a-4d69-af0e-832d7adff96a', 'f7fed851-2505-41ba-ad51-217ccea62ea4', 7700000, '2023-07-16 09:26:05.566', '2023-07-16 09:26:05.566'),
('9fb8032f-52c3-46ec-9e15-f8d1aa808d51', 'f7fed851-2505-41ba-ad51-217ccea62ea4', 3000000, '2023-07-16 09:26:44.804', '2023-07-16 09:26:44.804'),
('d76ba9f9-4a85-4421-949a-d3775aaf5513', 'f7fed851-2505-41ba-ad51-217ccea62ea4', 600000, '2023-07-16 09:38:53.573', '2023-07-16 09:38:53.573');

-- --------------------------------------------------------

--
-- Table structure for table `invoice_items`
--

CREATE TABLE `invoice_items` (
  `id` varchar(191) NOT NULL,
  `invoice_id` varchar(191) DEFAULT NULL,
  `product_id` varchar(191) DEFAULT NULL,
  `qty` bigint(20) UNSIGNED NOT NULL,
  `price` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `invoice_items`
--

INSERT INTO `invoice_items` (`id`, `invoice_id`, `product_id`, `qty`, `price`, `created_at`, `updated_at`) VALUES
('3c7f4b3e-ee74-4879-b93d-0d12a182f5d4', '70d11b73-622a-4d69-af0e-832d7adff96a', 'cf08e430-9ad4-4e62-a686-aa9fb26a0df0', 1, 7000000, '2023-07-16 09:26:05.568', '2023-07-16 09:26:05.568'),
('3f464421-6938-412d-a4f3-142e7ae0ab53', 'd76ba9f9-4a85-4421-949a-d3775aaf5513', 'c6723e70-e522-4414-9d3d-ceb038f801a4', 3, 600000, '2023-07-16 09:38:53.575', '2023-07-16 09:38:53.575'),
('6abc74d7-8ec6-4424-b3db-8d6ee1416e36', '70d11b73-622a-4d69-af0e-832d7adff96a', '4991c9fd-d61d-4719-9ebb-553fe90526cc', 2, 700000, '2023-07-16 09:26:05.567', '2023-07-16 09:26:05.567'),
('c83232a4-3da6-4bce-95f2-d075225b04b4', '9fb8032f-52c3-46ec-9e15-f8d1aa808d51', '5cad48fe-75dd-4154-bbd6-0406578dfb0f', 1, 3000000, '2023-07-16 09:26:44.806', '2023-07-16 09:26:44.806');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` varchar(191) NOT NULL,
  `name` longtext NOT NULL,
  `price` bigint(20) UNSIGNED NOT NULL,
  `description` longtext NOT NULL,
  `stock` bigint(20) UNSIGNED NOT NULL,
  `category_id` varchar(191) DEFAULT NULL,
  `user_id` varchar(191) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `price`, `description`, `stock`, `category_id`, `user_id`, `created_at`, `updated_at`) VALUES
('4991c9fd-d61d-4719-9ebb-553fe90526cc', 'Celana chinos coklat', 350000, 'Celana chinos coklat formal', 1, '65b4ab22-3b9c-40cc-b7e0-911dbce8e65b', 'f7fed851-2505-41ba-ad51-217ccea62ea4', '2023-07-16 08:20:35.749', '2023-07-16 09:26:05.561'),
('5cad48fe-75dd-4154-bbd6-0406578dfb0f', 'Meja belajar informa', 3000000, 'Meja belajar besi kuat', 9, 'c960d7fc-b4cd-4eba-9d42-57ed15b048bf', 'f7fed851-2505-41ba-ad51-217ccea62ea4', '2023-07-16 08:16:54.584', '2023-07-16 09:26:44.802'),
('c6723e70-e522-4414-9d3d-ceb038f801a4', 'Kemeja kantor formal', 200000, 'Kemeja putih bahan combed 30s', 47, '65b4ab22-3b9c-40cc-b7e0-911dbce8e65b', 'f7fed851-2505-41ba-ad51-217ccea62ea4', '2023-07-16 08:18:34.706', '2023-07-16 09:38:53.570'),
('cf08e430-9ad4-4e62-a686-aa9fb26a0df0', 'Monitor Toshiba', 7000000, '45inch monitor', 3, 'e1670d03-5c43-4c80-93ac-2dd850a25c9a', 'f7fed851-2505-41ba-ad51-217ccea62ea4', '2023-07-16 08:15:08.519', '2023-07-16 09:26:05.565');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` varchar(191) NOT NULL,
  `name` longtext NOT NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `address` longtext NOT NULL,
  `is_seller` tinyint(1) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `address`, `is_seller`, `created_at`, `updated_at`) VALUES
('017838a3-335e-4ea0-9b01-fbfa19f78987', 'Septian', 'septian@gmail.com', '$2a$10$f5WYcQ3FWqr9x6q0UN/5AepqVbVTuuYl2dyz931WCuNrNYAI0LIfm', 'Jakarta', 0, '2023-07-16 07:27:13.601', '2023-07-16 07:27:13.601'),
('f7fed851-2505-41ba-ad51-217ccea62ea4', 'Malta', 'arta@gmail.com', '$2a$10$QeVswueMb/3WGUVVvoWK5Ow/jw5aobK3qdszFPjF0mQxQ.BHxIxQC', 'Bandung', 1, '2023-07-16 07:08:40.117', '2023-07-16 07:26:13.330');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admins`
--
ALTER TABLE `admins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD UNIQUE KEY `username_2` (`username`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`),
  ADD UNIQUE KEY `name_2` (`name`);

--
-- Indexes for table `invoices`
--
ALTER TABLE `invoices`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_invoices_user_id` (`user_id`);

--
-- Indexes for table `invoice_items`
--
ALTER TABLE `invoice_items`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_invoice_items_invoice_id` (`invoice_id`),
  ADD KEY `idx_invoice_items_product_id` (`product_id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_products_category_id` (`category_id`),
  ADD KEY `idx_products_user_id` (`user_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD UNIQUE KEY `email_2` (`email`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `invoices`
--
ALTER TABLE `invoices`
  ADD CONSTRAINT `fk_invoices_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `invoice_items`
--
ALTER TABLE `invoice_items`
  ADD CONSTRAINT `fk_invoice_items_invoice` FOREIGN KEY (`invoice_id`) REFERENCES `invoices` (`id`),
  ADD CONSTRAINT `fk_invoice_items_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `fk_products_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`),
  ADD CONSTRAINT `fk_products_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
