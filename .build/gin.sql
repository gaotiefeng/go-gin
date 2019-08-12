/*
Navicat MySQL Data Transfer

Source Server         : 192.168.41.99
Source Server Version : 50723
Source Host           : 192.168.41.99:3306
Source Database       : gin

Target Server Type    : MYSQL
Target Server Version : 50723
File Encoding         : 65001

Date: 2019-08-12 17:44:14
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mobile` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;
