/*
 Navicat Premium Data Transfer

 Source Server         : g1_mysql
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : g1:3306
 Source Schema         : weekly_report

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 02/11/2021 20:13:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `ptype` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES ('p', '100', '/menu/getMenu', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '100', '/jwt/jsonInBlacklist', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '100', '/base/login', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '100', '/user/register', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '100', '/user/changePassword', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '100', '/user/setUserAuthority', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '100', '/user/setUserInfo', 'PUT', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '100', '/user/getUserInfo', 'GET', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '200', '/menu/getMenu', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '200', '/jwt/jsonInBlacklist', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '200', '/base/login', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '200', '/user/register', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '200', '/user/changePassword', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '200', '/user/setUserAuthority', 'POST', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '200', '/user/setUserInfo', 'PUT', '', '', '', NULL);
INSERT INTO `casbin_rule` VALUES ('p', '200', '/user/getUserInfo', 'GET', '', '', '', NULL);
COMMIT;

-- ----------------------------
-- Table structure for file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `file_chunks`;
CREATE TABLE `file_chunks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `exa_file_id` bigint(20) unsigned DEFAULT NULL,
  `file_chunk_number` bigint(20) DEFAULT NULL,
  `file_chunk_path` varchar(191) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of file_chunks
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `file_upload_and_downloads`;
CREATE TABLE `file_upload_and_downloads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '?????????',
  `url` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '????????????',
  `tag` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '????????????',
  `key` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '??????',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of file_upload_and_downloads
-- ----------------------------
BEGIN;
INSERT INTO `file_upload_and_downloads` VALUES (3, '2021-11-02 09:04:27', '2021-11-02 09:04:27', NULL, 'demo.c', '/Users/zero/Documents/uploads/file/fe01ce2a7fbac8fafaed7c982a04e229_20211102170427.c', 'c', 'fe01ce2a7fbac8fafaed7c982a04e229_20211102170427.c');
INSERT INTO `file_upload_and_downloads` VALUES (4, '2021-11-02 12:10:17', '2021-11-02 12:10:17', NULL, 'mac-lan.png', '/Users/zero/Documents/uploads/file/2db0df442ea6d92d75657712a29e5604_20211102201017.png', 'png', '2db0df442ea6d92d75657712a29e5604_20211102201017.png');
COMMIT;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `file_name` varchar(191) COLLATE utf8_bin DEFAULT NULL,
  `file_md5` varchar(191) COLLATE utf8_bin DEFAULT NULL,
  `file_path` varchar(191) COLLATE utf8_bin DEFAULT NULL,
  `chunk_total` bigint(20) DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `jwt` text COLLATE utf8_bin COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `authority_id` varchar(90) COLLATE utf8_bin NOT NULL COMMENT '??????ID',
  `authority_name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '?????????',
  `parent_id` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '?????????ID',
  `default_router` varchar(191) COLLATE utf8_bin DEFAULT 'dashboard' COMMENT '????????????',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
BEGIN;
INSERT INTO `sys_authorities` VALUES ('2021-11-01 10:57:52', '2021-11-01 10:57:55', NULL, '100', '?????????', '0', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2021-11-01 11:02:49', '2021-11-01 11:02:51', NULL, '200', '????????????', '0', 'dashboard');
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` varchar(90) COLLATE utf8_bin NOT NULL COMMENT '??????ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_authority_menus` VALUES (1, '100');
INSERT INTO `sys_authority_menus` VALUES (1, '200');
INSERT INTO `sys_authority_menus` VALUES (2, '100');
INSERT INTO `sys_authority_menus` VALUES (2, '200');
INSERT INTO `sys_authority_menus` VALUES (3, '100');
INSERT INTO `sys_authority_menus` VALUES (3, '200');
INSERT INTO `sys_authority_menus` VALUES (4, '100');
INSERT INTO `sys_authority_menus` VALUES (5, '100');
INSERT INTO `sys_authority_menus` VALUES (6, '100');
INSERT INTO `sys_authority_menus` VALUES (6, '200');
INSERT INTO `sys_authority_menus` VALUES (7, '100');
INSERT INTO `sys_authority_menus` VALUES (7, '200');
INSERT INTO `sys_authority_menus` VALUES (8, '100');
INSERT INTO `sys_authority_menus` VALUES (9, '100');
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '????????????????????????params??????query',
  `key` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '????????????????????????key',
  `value` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '???????????????????????????',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '?????????ID',
  `path` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '??????path',
  `name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '??????name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '?????????????????????',
  `component` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '????????????????????????',
  `sort` bigint(20) DEFAULT NULL COMMENT '????????????',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '????????????',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '????????????',
  `title` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '????????????',
  `icon` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '????????????',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '????????????',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menus` VALUES (1, '2021-11-01 11:14:12', '2021-11-01 11:14:14', NULL, NULL, '0', NULL, '??????', NULL, NULL, 1, NULL, NULL, '??????', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (2, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', '/viewWeeklyReport', '????????????', NULL, NULL, 1, NULL, NULL, '????????????', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (3, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', '/writeWeeklyReport', '?????????', NULL, NULL, 2, NULL, NULL, '?????????', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (4, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', '/statisticalExport', '????????????', NULL, NULL, 3, NULL, NULL, '????????????', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (5, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', '/templateEditing', '????????????', NULL, NULL, 4, NULL, NULL, '????????????', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (6, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '0', NULL, '??????', NULL, NULL, 2, NULL, NULL, '??????', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (7, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', '/userPassword', '????????????', NULL, NULL, 1, NULL, NULL, '????????????', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (8, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', '/statisticalRules', '????????????', NULL, NULL, 2, NULL, NULL, '????????????', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (9, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', '/userManagement', '????????????', NULL, NULL, 3, NULL, NULL, '????????????', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` varchar(90) COLLATE utf8_bin NOT NULL COMMENT '??????ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_authority` VALUES (1, '100');
INSERT INTO `sys_user_authority` VALUES (2, '200');
COMMIT;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uuid` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '??????UUID',
  `username` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '???????????????',
  `password` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '??????????????????',
  `nick_name` varchar(191) COLLATE utf8_bin DEFAULT '????????????' COMMENT '????????????',
  `side_mode` varchar(191) COLLATE utf8_bin DEFAULT 'dark' COMMENT '??????????????????',
  `header_img` varchar(191) COLLATE utf8_bin DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '????????????',
  `base_color` varchar(191) COLLATE utf8_bin DEFAULT '#fff' COMMENT '????????????',
  `active_color` varchar(191) COLLATE utf8_bin DEFAULT '#1890ff' COMMENT '????????????',
  `authority_id` varchar(90) COLLATE utf8_bin DEFAULT '888' COMMENT '????????????ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` VALUES (1, '2021-11-01 14:53:31', '2021-11-02 10:30:56', NULL, '8e600b7f-3297-4979-a445-d218205ef9a6', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '???????????????', 'dark', '', '#fff', '#1890ff', '100');
INSERT INTO `sys_users` VALUES (2, '2021-11-01 14:53:31', '2021-11-01 14:53:31', NULL, '9d9ab82d-81a7-433b-9ea1-bde1c4f1b990', 'user', 'e10adc3949ba59abbe56e057f20f883e', '????????????', 'dark', '', '#fff', '#1890ff', '200');
COMMIT;

-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
DROP VIEW IF EXISTS `authority_menu`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`close_tab` AS `close_tab`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
