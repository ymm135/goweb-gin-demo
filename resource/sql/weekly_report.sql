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

 Date: 01/11/2021 17:25:45
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

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
  `authority_id` varchar(90) COLLATE utf8_bin NOT NULL COMMENT '角色ID',
  `authority_name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '角色名',
  `parent_id` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) COLLATE utf8_bin DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
BEGIN;
INSERT INTO `sys_authorities` VALUES ('2021-11-01 10:57:52', '2021-11-01 10:57:55', NULL, '100', '管理员', '0', 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2021-11-01 11:02:49', '2021-11-01 11:02:51', NULL, '200', '普通用户', '0', 'dashboard');
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` varchar(90) COLLATE utf8_bin NOT NULL COMMENT '角色ID',
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
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menus` VALUES (1, '2021-11-01 11:14:12', '2021-11-01 11:14:14', NULL, NULL, '0', NULL, '周报', NULL, NULL, 1, NULL, NULL, '周报', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (2, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', NULL, '查看周报', NULL, NULL, 1, NULL, NULL, '查看周报', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (3, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', NULL, '写周报', NULL, NULL, 2, NULL, NULL, '写周报', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (4, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', NULL, '统计导出', NULL, NULL, 3, NULL, NULL, '统计导出', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (5, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', NULL, '模板编辑', NULL, NULL, 4, NULL, NULL, '模板编辑', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (6, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '0', NULL, '设置', NULL, NULL, 2, NULL, NULL, '设置', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (7, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', NULL, '用户密码', NULL, NULL, 1, NULL, NULL, '用户密码', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (8, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', NULL, '统计规则', NULL, NULL, 2, NULL, NULL, '统计规则', NULL, NULL);
INSERT INTO `sys_base_menus` VALUES (9, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', NULL, '用户管理', NULL, NULL, 3, NULL, NULL, '用户管理', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` varchar(90) COLLATE utf8_bin NOT NULL COMMENT '角色ID',
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
  `uuid` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) COLLATE utf8_bin DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) COLLATE utf8_bin DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) COLLATE utf8_bin DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) COLLATE utf8_bin DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) COLLATE utf8_bin DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` varchar(90) COLLATE utf8_bin DEFAULT '888' COMMENT '用户角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` VALUES (1, '2021-11-01 14:53:31', '2021-11-01 14:53:31', NULL, '8e600b7f-3297-4979-a445-d218205ef9a6', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'dark', '', '#fff', '#1890ff', '100');
INSERT INTO `sys_users` VALUES (2, '2021-11-01 14:53:31', '2021-11-01 14:53:31', NULL, '9d9ab82d-81a7-433b-9ea1-bde1c4f1b990', 'user', 'e10adc3949ba59abbe56e057f20f883e', '普通用户', 'dark', '', '#fff', '#1890ff', '200');
COMMIT;

-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
DROP VIEW IF EXISTS `authority_menu`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`close_tab` AS `close_tab`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
