/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : 65001

 Date: 22/03/2019 10:02:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`
(
  `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id`      int(10) unsigned    DEFAULT '0' COMMENT '标签ID',
  `title`       varchar(100)        DEFAULT '' COMMENT '文章标题',
  `desc`        varchar(255)        DEFAULT '' COMMENT '简述',
  `content`     text,
  `created_on`  int(11)             DEFAULT NULL,
  `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255)        DEFAULT '' COMMENT '修改人',
  `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='文章管理';

-- ----------------------------
-- Table structure for blog_auth
-- ----------------------------
DROP TABLE IF EXISTS `blog_auth`;
CREATE TABLE `blog_auth`
(
  `id`       int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8 COMMENT ='认证表';

-- ----------------------------
-- Records of blog_auth
-- ----------------------------
BEGIN;
INSERT INTO `blog_auth`
VALUES (1, 'test', 'test123456');
COMMIT;

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`
(
  `id`          int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name`        varchar(100)        DEFAULT '' COMMENT '标签名称',
  `created_on`  int(10) unsigned    DEFAULT '0' COMMENT '创建时间',
  `created_by`  varchar(100)        DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned    DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100)        DEFAULT '' COMMENT '修改人',
  `state`       tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='文章标签管理';

SET FOREIGN_KEY_CHECKS = 1;
