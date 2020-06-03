SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for follower
-- ----------------------------
DROP TABLE IF EXISTS `follower`;
CREATE TABLE `follower` (
  `user_id` int(11) NOT NULL DEFAULT '0',
  `follower_id` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`,`follower_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of follower
-- ----------------------------

-- ----------------------------
-- Table structure for post
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `body` varchar(180) DEFAULT NULL,
  `timestamp` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of post
-- ----------------------------
INSERT INTO `post` VALUES ('1', '1', 'beautiful day in borlan!', null);
INSERT INTO `post` VALUES ('2', '2', 'the aven mov wo so col', null);
INSERT INTO `post` VALUES ('3', '2', 'sun shine is beautiful', null);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) DEFAULT NULL,
  `email` varchar(120) DEFAULT NULL,
  `password_hash` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'bon', '', 'e99a18c428cb38d5f260853678922e03');
INSERT INTO `user` VALUES ('2', 'frn', 'rene@test.com', 'e99a18c428cb38d5f260853678922e03');
INSERT INTO `user` VALUES ('3', 'jimmy', 'jimmy007@test.com', '11223311');
INSERT INTO `user` VALUES ('4', 'tom', 'tom@test.com', 'tom11223311');
INSERT INTO `user` VALUES ('5', 'jimmy', 'jimmy@gogo.me', 'password123123123');
INSERT INTO `user` VALUES ('6', 'jimmy', 'jimmy@gogo.me', 'password123123123');
INSERT INTO `user` VALUES ('7', 'jimmy', 'jimmy@gogo.me', 'password123123123');
