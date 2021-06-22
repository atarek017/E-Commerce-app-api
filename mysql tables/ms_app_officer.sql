-- MySQL dump 10.13  Distrib 5.6.24, for Win32 (x86)
--
-- Host: localhost    Database: ms_app
-- ------------------------------------------------------
-- Server version	5.7.16-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `officer`
--

DROP TABLE IF EXISTS `officer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `officer` (
  `id` varchar(45) NOT NULL,
  `name` varchar(45) NOT NULL,
  `email` varchar(45) DEFAULT NULL,
  `password` varchar(45) DEFAULT NULL,
  `photo` text,
  `phone` varchar(11) DEFAULT NULL,
  `money_founder` int(11) DEFAULT NULL,
  `admin` int(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `officer`
--

LOCK TABLES `officer` WRITE;
/*!40000 ALTER TABLE `officer` DISABLE KEYS */;
INSERT INTO `officer` VALUES ('1','Ahmed','atarek017@gmail.com','Xaha4006','https://avatars3.githubusercontent.com/u/26245770?s=400&u=027902ef5b39db7e6df2576245c8e88343f3e4ac&v=4','01120530142',1000,1),('1hvLe2fhVmJNDyjQz7XczKefvcx','ali','ali@gmail.com','Xaha4006','','',350,0),('2','mohamed','ahmed150236@fci.bu.edu.eg','123','https://avatars3.githubusercontent.com/u/26245770?s=400&u=027902ef5b39db7e6df2576245c8e88343f3e4ac&v=4','214132',600,1),('3','ebrahime','eb12@gmail.com','123','https://avatars3.githubusercontent.com/u/26245770?s=400&u=027902ef5b39db7e6df2576245c8e88343f3e4ac&v=4','21431',400,0),('4','Ayman','mesan@gmail.com','Xaha4006','https://avatars3.githubusercontent.com/u/26245770?s=400&u=027902ef5b39db7e6df2576245c8e88343f3e4ac&v=4','1120530142',20,1);
/*!40000 ALTER TABLE `officer` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-10-28 22:09:54
