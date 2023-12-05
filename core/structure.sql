-- MySQL dump 10.13  Distrib 8.0.33, for macos13 (x86_64)
--
-- Host: 127.0.0.1    Database: coin
-- ------------------------------------------------------
-- Server version	5.7.41

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `account_types`
--

DROP TABLE IF EXISTS `account_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account_types` (
  `signatura` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Строковый идентификатор типа счета',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Название типа счета',
  PRIMARY KEY (`signatura`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Типы счетов';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `accounts`
--

DROP TABLE IF EXISTS `accounts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Идентификатор счета',
  `budget` double NOT NULL DEFAULT '0' COMMENT 'Бюджет на месяц',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Название',
  `icon_id` int(11) NOT NULL COMMENT 'Идентификатор иконки',
  `type_signatura` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Тип счета',
  `currency_signatura` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Строковый идентификатор валюты',
  `visible` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'Видимость счета',
  `account_group_id` int(11) NOT NULL COMMENT 'Идентификатор группы счетов',
  `accounting` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'Учет в общем балансе',
  `user_id` int(11) NOT NULL DEFAULT '1' COMMENT 'Идентфикатор пользователя',
  `parent_account_id` int(11) DEFAULT NULL COMMENT 'Идентификатор привязки к другому счету',
  `serial_number` int(10) unsigned NOT NULL,
  `is_parent` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `account_FK` (`icon_id`),
  KEY `accounts_FK_1` (`type_signatura`),
  KEY `accounts_FK_2` (`currency_signatura`),
  KEY `accounts_FK` (`account_group_id`),
  KEY `accounts_FK_3` (`user_id`),
  KEY `accounts_FK_5` (`parent_account_id`),
  CONSTRAINT `accounts_FK` FOREIGN KEY (`account_group_id`) REFERENCES `accounts_group` (`id`),
  CONSTRAINT `accounts_FK_1` FOREIGN KEY (`type_signatura`) REFERENCES `account_types` (`signatura`) ON UPDATE CASCADE,
  CONSTRAINT `accounts_FK_2` FOREIGN KEY (`currency_signatura`) REFERENCES `currencies` (`signatura`),
  CONSTRAINT `accounts_FK_3` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `accounts_FK_4` FOREIGN KEY (`icon_id`) REFERENCES `icons` (`id`),
  CONSTRAINT `accounts_FK_5` FOREIGN KEY (`parent_account_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8029 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Счета';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `accounts_group`
--

DROP TABLE IF EXISTS `accounts_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `accounts_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Идектификатор группы',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Название группы',
  `user_id` int(11) NOT NULL COMMENT 'Идентификатор пользователя',
  PRIMARY KEY (`id`),
  KEY `accounts_group_FK` (`user_id`),
  CONSTRAINT `accounts_group_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Группы аккаунтов';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `action_history`
--

DROP TABLE IF EXISTS `action_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `action_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Идентификатор действия',
  `action_type_signatura` varchar(100) NOT NULL COMMENT 'Тип действия пользователя',
  `user_id` int(11) NOT NULL COMMENT 'Пользователь, который произвел действие',
  `object_id` int(11) DEFAULT NULL COMMENT 'Идентификатор измененного объекта',
  `note` varchar(200) DEFAULT NULL COMMENT 'Заметка от администратора',
  `action_time` datetime NOT NULL COMMENT 'Время, когда совершилось действие',
  PRIMARY KEY (`id`),
  KEY `NewTable_FK` (`user_id`),
  KEY `action_history_FK` (`action_type_signatura`),
  CONSTRAINT `NewTable_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `action_history_FK` FOREIGN KEY (`action_type_signatura`) REFERENCES `action_types` (`type_signatura`)
) ENGINE=InnoDB AUTO_INCREMENT=1154 DEFAULT CHARSET=utf8mb4 COMMENT='История действий';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `action_types`
--

DROP TABLE IF EXISTS `action_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `action_types` (
  `type_signatura` varchar(100) NOT NULL COMMENT 'Название действия',
  `note` varchar(100) DEFAULT NULL COMMENT 'Заметка о типе действия',
  PRIMARY KEY (`type_signatura`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Типы действий';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `currencies`
--

DROP TABLE IF EXISTS `currencies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `currencies` (
  `signatura` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Строковый идентификатор',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Название валюты',
  `rate` decimal(15,8) NOT NULL COMMENT 'Курс валюты относительно доллара',
  PRIMARY KEY (`signatura`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Валюты';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `icons`
--

DROP TABLE IF EXISTS `icons`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `icons` (
  `id` int(11) NOT NULL COMMENT 'Строковый идентификатор',
  `img` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Ссылка на изображение',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Название изображения',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Иконки счетов';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `sessions`
--

DROP TABLE IF EXISTS `sessions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sessions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `refresh_token` varchar(200) NOT NULL,
  `expires_at` datetime NOT NULL,
  `device_id` varchar(100) NOT NULL,
  `user_id` int(11) NOT NULL,
  `last_sync` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `sessions_FK` (`user_id`),
  CONSTRAINT `sessions_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=314 DEFAULT CHARSET=utf8mb4 COMMENT='Сессии';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Идентификатор подкатегории',
  `name` varchar(100) DEFAULT NULL COMMENT 'Название подкатегории',
  `user_id` int(11) DEFAULT '1' COMMENT 'Идентификатор пользователя',
  PRIMARY KEY (`id`),
  KEY `tags_FK` (`user_id`),
  CONSTRAINT `tags_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=108 DEFAULT CHARSET=utf8mb4 COMMENT='Подкатегории';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tags_to_transaction`
--

DROP TABLE IF EXISTS `tags_to_transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tags_to_transaction` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Идентификатор привязки подкатегории',
  `transaction_id` int(11) NOT NULL COMMENT 'Идентификатор транзакции',
  `tag_id` int(11) DEFAULT NULL COMMENT 'Идентификатор подкатегории',
  PRIMARY KEY (`id`),
  KEY `tags_FK` (`transaction_id`),
  KEY `tags_to_order_FK` (`tag_id`),
  CONSTRAINT `tags_to_transaction_FK` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`),
  CONSTRAINT `tags_to_transaction_FK_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1935 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Подкатегории в транзакции';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `transaction_types`
--

DROP TABLE IF EXISTS `transaction_types`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction_types` (
  `signatura` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Строковый идентификатор типа операции',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Название типа операции',
  PRIMARY KEY (`signatura`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Типы счетов';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Идентификатор транзакции',
  `date_transaction` date NOT NULL COMMENT 'Дата совершения транзакции',
  `type_signatura` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Тип транзакции',
  `amount_from` decimal(17,7) NOT NULL COMMENT 'Сумма, ушедшая из первого счета',
  `amount_to` decimal(17,7) NOT NULL COMMENT 'Сумма, пришедшая во второй счет',
  `note` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'Заметка',
  `account_from_id` int(11) NOT NULL COMMENT 'Откуда взяли деньги',
  `account_to_id` int(11) NOT NULL COMMENT 'Куда положили деньги',
  `is_executed` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'Исполнена ли операция',
  `date_create` datetime DEFAULT NULL COMMENT 'Дата создания транзакции',
  `accounting` tinyint(1) NOT NULL DEFAULT '1',
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `orders_FK` (`account_from_id`),
  KEY `orders_FK_1` (`account_to_id`),
  KEY `orders_FK_3` (`type_signatura`),
  KEY `transactions_FK` (`user_id`),
  CONSTRAINT `orders_FK` FOREIGN KEY (`account_from_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `orders_FK_1` FOREIGN KEY (`account_to_id`) REFERENCES `accounts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `orders_FK_3` FOREIGN KEY (`type_signatura`) REFERENCES `transaction_types` (`signatura`),
  CONSTRAINT `transactions_FK` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4760 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Транзакции';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Идентификато пользователя',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Имя пользователя',
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Почтовый адрес пользователя',
  `password_hash` varchar(200) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Хэш пароля',
  `verification_email_code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Код верификации почты',
  `date_create` date NOT NULL COMMENT 'Дата создания аккаунта',
  `fcm_token` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Токен для push-уведомлений',
  `default_currency_signatura` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Валюта по умолчанию',
  `last_sync` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `members_FK_1` (`default_currency_signatura`),
  CONSTRAINT `members_FK_1` FOREIGN KEY (`default_currency_signatura`) REFERENCES `currencies` (`signatura`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Участники проекта';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping routines for database 'coin'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-07-26 22:25:19
