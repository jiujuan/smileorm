#
# Structure for table "follower"
#

DROP TABLE IF EXISTS `follower`;
CREATE TABLE `follower` (
  `user_id` int(11) NOT NULL DEFAULT '0',
  `follower_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`,`follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

#
# Data for table "follower"
#


#
# Structure for table "post"
#

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `body` varchar(180) DEFAULT NULL,
  `timestamp` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

#
# Data for table "post"
#

INSERT INTO `post` VALUES (1,1,'beautiful day in borlan!',NULL),(2,2,'the aven mov wo so col',NULL),(3,2,'sun shine is beautiful',NULL);

#
# Structure for table "user"
#

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `email` varchar(120) DEFAULT NULL,
  `password_hash` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

#
# Data for table "user"
#

INSERT INTO `user` VALUES (1,'bon','','e99a18c428cb38d5f260853678922e03'),(2,'frn','rene@test.com','e99a18c428cb38d5f260853678922e03'),(3,'jimmy','jimmy007@test.com','11223311'),(4,'tom','tom@test.com','tom11223311'),(5,'jimmy','jimmy@gogo.me','password123123123'),(6,'jimmy','jimmy@gogo.me','password123123123'),(7,'jimmy','jimmy@gogo.me','password123123123');
