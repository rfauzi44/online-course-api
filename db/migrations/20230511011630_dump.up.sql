-- MySQL dump 10.13  Distrib 8.0.32, for Win64 (x86_64)
--
-- Host: localhost    Database: online-course
-- ------------------------------------------------------
-- Server version	8.0.32

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
-- Table structure for table `courses`
--

DROP TABLE IF EXISTS `courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `courses` (
  `id` varchar(100) NOT NULL,
  `title` varchar(100) NOT NULL,
  `description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` int NOT NULL,
  `image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `category` varchar(100) NOT NULL,
  `author_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `image_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `courses_FK` (`author_id`),
  CONSTRAINT `courses_FK` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES ('002af416-72bb-49da-b570-13feed28a7f6','Japanese','Donec ut ante eleifend, consequat mi at, sodales augue. Mauris nisi arcu, venenatis eget nulla et, posuere elementum felis. Praesent eget elit sit amet enim dictum lobortis. Vivamus tincidunt, lorem sit amet consequat gravida, metus leo dapibus nunc, sed tempor libero nunc tempor arcu. In porttitor ante et nibh lobortis, id fringilla purus aliquet. Nullam faucibus blandit pretium. Suspendisse eu nibh eros. Sed porta rhoncus aliquet. Etiam eu vestibulum augue. Praesent vitae purus in dui auctor varius.',0,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683728002/cld-sample.jpg','language','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 04:00:30','2023-05-10 04:00:30',''),('043bb448-0e49-483d-acba-5eb5219dbb2c','SCRUM','Nullam egestas, elit id gravida pellentesque, massa neque ultricies magna, nec ultricies quam orci ac purus. Proin posuere nunc id erat scelerisque, id luctus arcu scelerisque. Donec gravida erat a purus ullamcorper pharetra. Proin lacinia ipsum malesuada dolor iaculis blandit. Vivamus volutpat in tortor sed ornare. Proin quis lorem auctor, pulvinar mi ac',25000,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683728002/cld-sample.jpg','technic','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 08:01:53','2023-05-10 08:01:53',''),('0be9c151-be2d-4d88-b8d4-edaaf17a7897','Echo','Nullam egestas, elit id gravida pellentesque, massa neque ultricies magna, nec ultricies quam orci ac purus. Proin posuere nunc id erat scelerisque, id luctus arcu scelerisque. Donec gravida erat a purus ullamcorper pharetra. Proin lacinia ipsum malesuada dolor iaculis blandit. Vivamus volutpat in tortor sed ornare. Proin quis lorem auctor, pulvinar mi ac',0,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683727990/samples/cloudinary-group.jpg','programming','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 03:59:45','2023-05-10 03:59:45',''),('0e30022f-932c-4f19-a43d-cb95edc1909b','Leadership','Nullam egestas, elit id gravida pellentesque, massa neque ultricies magna, nec ultricies quam orci ac purus. Proin posuere nunc id erat scelerisque, id luctus arcu scelerisque. Donec gravida erat a purus ullamcorper pharetra. Proin lacinia ipsum malesuada dolor iaculis blandit. Vivamus volutpat in tortor sed ornare. Proin quis lorem auctor, pulvinar mi ac',80000,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683727990/samples/cloudinary-group.jpg','soft skill','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 04:00:56','2023-05-10 04:00:56',''),('43ad4ddc-a819-444a-9f21-472bd22f8e48','SCRUM','Nullam egestas, elit id gravida pellentesque, massa neque ultricies magna, nec ultricies quam orci ac purus. Proin posuere nunc id erat scelerisque, id luctus arcu scelerisque. Donec gravida erat a purus ullamcorper pharetra. Proin lacinia ipsum malesuada dolor iaculis blandit. Vivamus volutpat in tortor sed ornare. Proin quis lorem auctor, pulvinar mi ac',25000,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683727990/samples/cloudinary-group.jpg','technic','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 04:01:18','2023-05-10 04:01:18',''),('5cceae0e-e226-4eea-8bc5-ccca82803c60','Golang','Donec pulvinar porttitor ante, nec dapibus purus convallis finibus. Vivamus pharetra massa ac neque vestibulum ultricies. Maecenas pellentesque euismod metus, id scelerisque nibh. Vivamus viverra tellus neque, a sagittis sem egestas at. Quisque faucibus eleifend lectus ut ullamcorper.',50000,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683727990/samples/cloudinary-group.jpg','programming','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 03:59:24','2023-05-10 03:59:24',''),('6a27387b-aa60-4de7-a058-f0255969e39c','English','Donec pulvinar porttitor ante, nec dapibus purus convallis finibus. Vivamus pharetra massa ac neque vestibulum ultricies. Maecenas pellentesque euismod metus, id scelerisque nibh. Vivamus viverra tellus neque, a sagittis sem egestas at. Quisque faucibus eleifend lectus ut ullamcorper.',20000,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683727990/samples/cloudinary-group.jpg','language','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 04:00:12','2023-05-10 04:00:12',''),('c650cab1-a6b9-449a-bd62-365393d26695','Gin','Donec pulvinar porttitor ante, nec dapibus purus convallis finibus. Vivamus pharetra massa ac neque vestibulum ultricies. Maecenas pellentesque euismod metus, id scelerisque nibh. Vivamus viverra tellus neque, a sagittis sem egestas at. Quisque faucibus eleifend lectus ut ullamcorper.',0,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683728003/cld-sample-5.jpg','programming','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 03:59:56','2023-05-10 03:59:56',''),('e82f2705-9ed5-40bd-be0f-5d37569179d2','Communication','Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.',0,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683728003/cld-sample-5.jpg','soft skill','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 04:00:43','2023-05-10 04:00:43',''),('f8ca9305-f9e5-42c9-99b2-ac2bf0b8e446','Indonesia','Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.',70000,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683728003/cld-sample-5.jpg','language','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 04:00:22','2023-05-10 04:00:22',''),('f9453760-4a69-40aa-b913-73249360e3bf','Javascript','Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.',20000,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683728003/cld-sample-5.jpg','programming','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 03:59:12','2023-05-10 03:59:12',''),('fa004718-2b00-42cd-a71b-9d67228275d5','React','Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry\'s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.',100000,'https://res.cloudinary.com/dxbe9tibu/image/upload/v1683728003/cld-sample-5.jpg','programming','62019004-3ca3-415c-97e5-3106fe37f7d1','2023-05-10 03:59:34','2023-05-10 03:59:34','');
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `role` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('3accf472-02dc-4857-b9be-06a281a4de57','user@user1.com','$2a$10$1tzDLd3etZtjzByoR6k.Leos6penyrc/9yt4eue13TEcCjdHvAvo6','user','2023-05-10 17:50:54','2023-05-10 17:50:54',1,'user'),('50a7c009-8a75-4543-958b-a4f7fb10e60b','cristiano@cristiano.com','$2a$10$cvXNrbfcOfNMCZQHZHPCN.4C/43./MQ8U2dkTQoKUoecQbM4RMvhC','admin','2023-05-10 17:38:59','2023-05-10 17:38:59',0,'Christiano Ronaldo'),('62019004-3ca3-415c-97e5-3106fe37f7d1','admin@admin.com','$2a$10$Ds9hPjV8Yn6L12/T5cNsG.b4EmZhgBFO3oQmUJxXyMnw3n3haUSUG','admin','2023-05-10 03:56:11','2023-05-10 03:56:11',0,'John Doe'),('c10c0af6-368a-48ff-bc6e-2f33f030a2b9','user@user.com','$2a$10$60toR4SmncmWlC1bZ6VMU.umgejZISwqII2urtGHGywwi2GOwALIy','admin','2023-05-10 14:21:13','2023-05-10 14:21:13',0,'John Travolta');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'online-course'
--

--
-- Dumping routines for database 'online-course'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-05-11  8:12:06
