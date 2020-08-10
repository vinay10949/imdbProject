-- MySQL dump 10.13  Distrib 8.0.21, for Linux (x86_64)
--
-- Host: localhost    Database: imdb
-- ------------------------------------------------------
-- Server version	8.0.21-0ubuntu0.20.04.4

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `language`
--

DROP TABLE IF EXISTS `language`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `language` (
  `ID` int NOT NULL,
  `LANGUAGE` varchar(100) NOT NULL,
  `CREATED_BY` varchar(45) NOT NULL,
  `CREATED_DATE` datetime NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `language`
--

LOCK TABLES `language` WRITE;
/*!40000 ALTER TABLE `language` DISABLE KEYS */;
INSERT INTO `language` VALUES (1,'HINDI','admin','2020-08-03 11:00:00'),(2,'ENGLISH','admin','2020-08-03 11:00:00');
/*!40000 ALTER TABLE `language` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movie`
--

DROP TABLE IF EXISTS `movie`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `movie` (
  `ID` bigint NOT NULL AUTO_INCREMENT,
  `TITLE` varchar(100) NOT NULL,
  `DESCRIPTION` longtext NOT NULL,
  `RATING` int NOT NULL DEFAULT '0',
  `CREATED_BY` varchar(100) NOT NULL,
  `CREATED_DATE` datetime NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie`
--

LOCK TABLES `movie` WRITE;
/*!40000 ALTER TABLE `movie` DISABLE KEYS */;
INSERT INTO `movie` VALUES (6,'KAAL','Et sit suscipit dolorum odio repellendus.',7,'ADMIN','2020-08-04 16:16:52'),(7,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:16:54'),(8,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:17:29'),(9,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',8,'ADMIN','2020-08-04 16:18:59'),(10,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:20:08'),(11,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:24:23'),(12,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:25:17'),(13,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:34:57'),(14,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:34:58'),(15,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:35:34'),(16,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:36:06'),(17,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:36:08'),(18,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:39:10'),(19,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:40:18'),(20,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:40:49'),(21,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:42:10'),(22,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:43:19'),(23,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:46:46'),(24,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:47:05'),(25,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:47:59'),(26,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:48:01'),(27,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:48:20'),(28,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:48:22'),(29,'Odit dolores hic.','Et sit suscipit dolorum odio repellendus.',0,'ADMIN','2020-08-04 16:53:17'),(30,'AAA','aaaa',0,'admin','2020-08-10 12:11:49'),(31,'Matrix','matrix movie',0,'admin','2020-08-10 12:13:05');
/*!40000 ALTER TABLE `movie` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_rating_comments`
--

DROP TABLE IF EXISTS `user_rating_comments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_rating_comments` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `MOVIE_ID` bigint NOT NULL,
  `USER_ID` bigint NOT NULL,
  `COMMENT` longtext,
  `RATING` int DEFAULT NULL,
  `CREATED_BY` varchar(100) DEFAULT NULL,
  `CREATED_DATE` datetime DEFAULT NULL,
  `MODIFIED_BY` varchar(100) DEFAULT NULL,
  `MODIFIED_DATE` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_rating_comments`
--

LOCK TABLES `user_rating_comments` WRITE;
/*!40000 ALTER TABLE `user_rating_comments` DISABLE KEYS */;
INSERT INTO `user_rating_comments` VALUES (1,6,1,NULL,5,'admin','2020-08-10 13:10:37',NULL,NULL),(3,6,2,'HELLO',8,'admin','2020-08-10 13:33:04',NULL,NULL),(5,6,3,'Good movie',8,'admin','2020-08-10 13:49:43',NULL,NULL),(12,9,2,NULL,8,'admin','2020-08-10 14:11:16',NULL,NULL);
/*!40000 ALTER TABLE `user_rating_comments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `ID` int NOT NULL AUTO_INCREMENT,
  `NAME` varchar(100) DEFAULT NULL,
  `LOGIN_ID` varchar(50) NOT NULL,
  `PASSWORD` varchar(100) NOT NULL,
  `ROLE` varchar(45) NOT NULL,
  `CREATED_BY` varchar(100) NOT NULL,
  `CREATED_DATE` datetime NOT NULL,
  `MODIFIED_BY` varchar(100) DEFAULT NULL,
  `MODIFIED_DATE` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'vinay','vinay10949','123456','admin','admin','2020-08-25 11:00:00',NULL,NULL),(2,'abhi','abhi','123456','admin','admin','2020-08-25 11:00:00',NULL,NULL),(3,'asd','adasd','123456','admin','admin','2020-08-25 11:00:00',NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-10 15:04:14
