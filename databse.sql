DROP TABLE IF EXISTS `rolls`;
CREATE TABLE `rolls` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `roll` text,
  `code` text,
  `link` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;
LOCK TABLES `rolls` WRITE;
UNLOCK TABLES;