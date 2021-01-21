/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50649
Source Host           : localhost:3306
Source Database       : website

Target Server Type    : MYSQL
Target Server Version : 50649
File Encoding         : 65001

Date: 2020-12-02 14:33:28
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` varchar(255) DEFAULT NULL,
  `user_name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `gender` char(255) DEFAULT NULL,
  `addr` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of user
-- ----------------------------
