

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


CREATE DATABASE IF NOT EXISTS weekly_report;

-- ----------------------------
-- Table structure for auto_code_examples
-- ----------------------------
DROP TABLE IF EXISTS `auto_code_examples`;
CREATE TABLE `auto_code_examples` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `auto_code_example_field` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '仅作示例条目无实际作用',
  PRIMARY KEY (`id`),
  KEY `idx_auto_code_examples_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of auto_code_examples
-- ----------------------------
BEGIN;
COMMIT;

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
INSERT INTO `casbin_rule` VALUES ('p', '100', 'user/getUserList', 'POST', NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `customer_name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks` (
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
-- Records of exa_file_chunks
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
BEGIN;
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '2021-10-25 14:53:31', '2021-10-25 14:53:31', NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (2, '2021-10-25 14:53:31', '2021-10-25 14:53:31', NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');
COMMIT;

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files` (
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
-- Records of exa_files
-- ----------------------------
BEGIN;
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
  `name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=142 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of file_upload_and_downloads
-- ----------------------------
BEGIN;
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
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `path` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) COLLATE utf8_bin DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
BEGIN;
INSERT INTO `sys_apis` VALUES (1, '2021-10-25 14:53:31', '2021-10-25 14:53:31', NULL, '/base/login', '用户登录（必选）', 'base', 'POST');
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
INSERT INTO `sys_authority_menus` VALUES (7, '200');
INSERT INTO `sys_authority_menus` VALUES (8, '100');
INSERT INTO `sys_authority_menus` VALUES (9, '100');
COMMIT;

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `table_name` varchar(191) COLLATE utf8_bin DEFAULT NULL,
  `request_meta` text COLLATE utf8_bin,
  `auto_code_path` text COLLATE utf8_bin,
  `injection_meta` text COLLATE utf8_bin,
  `struct_name` varchar(191) COLLATE utf8_bin DEFAULT NULL,
  `struct_cn_name` varchar(191) COLLATE utf8_bin DEFAULT NULL,
  `api_ids` varchar(191) COLLATE utf8_bin DEFAULT NULL,
  `flag` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------
BEGIN;
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
  `type` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '地址栏携带参数的值',
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menus` VALUES (1, '2021-11-01 11:14:12', '2021-11-01 11:14:14', NULL, NULL, '0', NULL, '周报', NULL, NULL, 1, NULL, NULL, '周报', 'file-word', NULL);
INSERT INTO `sys_base_menus` VALUES (2, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', '/viewWeeklyReport', '查看周报', NULL, NULL, 1, NULL, NULL, '查看周报', 'eye', NULL);
INSERT INTO `sys_base_menus` VALUES (3, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', '/writeWeeklyReport', '写周报', NULL, NULL, 2, NULL, NULL, '写周报', 'edit', NULL);
INSERT INTO `sys_base_menus` VALUES (4, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', '/statisticalExport', '统计导出', NULL, NULL, 3, NULL, NULL, '统计导出', 'line-chart', NULL);
INSERT INTO `sys_base_menus` VALUES (5, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '1', '/templateEditing', '模板编辑', NULL, NULL, 4, NULL, NULL, '模板编辑', 'align-left', NULL);
INSERT INTO `sys_base_menus` VALUES (6, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '0', NULL, '设置', NULL, NULL, 2, NULL, NULL, '设置', 'setting', NULL);
INSERT INTO `sys_base_menus` VALUES (7, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', '/userPassword', '用户密码', NULL, NULL, 1, NULL, NULL, '用户密码', 'user', NULL);
INSERT INTO `sys_base_menus` VALUES (8, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', '/statisticalRules', '统计规则', NULL, NULL, 2, NULL, NULL, '统计规则', 'calculator', NULL);
INSERT INTO `sys_base_menus` VALUES (9, '2021-11-01 11:16:40', '2021-11-01 11:16:43', NULL, NULL, '6', '/userManagement', '用户管理', NULL, NULL, 3, NULL, NULL, '用户管理', 'user', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` varchar(90) COLLATE utf8_bin NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` varchar(90) COLLATE utf8_bin NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
BEGIN;
INSERT INTO `sys_dictionaries` VALUES (1, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, '性别', 'sex', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `label` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
BEGIN;
INSERT INTO `sys_dictionary_details` VALUES (1, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (2, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (3, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'date', 0, 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (6, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (7, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (8, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (9, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'timestamp', 5, 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (10, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (11, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'double', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (12, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'decimal', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (13, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (14, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'varchar', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (15, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'tinyblob', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (16, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'tinytext', 3, 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (17, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'text', 4, 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (18, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'blob', 5, 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (19, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'mediumblob', 6, 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (20, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'mediumtext', 7, 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (21, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'longblob', 8, 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (22, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'longtext', 9, 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2021-10-25 15:11:46', '2021-10-25 15:11:46', NULL, 'tinyint', 0, 1, 0, 6);
COMMIT;

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `ip` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) COLLATE utf8_bin DEFAULT NULL COMMENT '错误信息',
  `body` longtext COLLATE utf8_bin COMMENT '请求Body',
  `resp` longtext COLLATE utf8_bin COMMENT '响应Body',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=269 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
BEGIN;
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
INSERT INTO `sys_user_authority` VALUES (5, '100');
INSERT INTO `sys_user_authority` VALUES (10, '200');
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` VALUES (1, '2021-11-01 14:53:31', '2021-11-02 10:30:56', NULL, '8e600b7f-3297-4979-a445-d218205ef9a6', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '100');
INSERT INTO `sys_users` VALUES (2, '2021-11-01 14:53:31', '2021-11-15 04:10:26', NULL, '9d9ab82d-81a7-433b-9ea1-bde1c4f1b990', 'user', 'e10adc3949ba59abbe56e057f20f883e', '普通用户', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', '200');
COMMIT;

-- ----------------------------
-- Table structure for wt_comments
-- ----------------------------
DROP TABLE IF EXISTS `wt_comments`;
CREATE TABLE `wt_comments` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `report_id` bigint(20) NOT NULL COMMENT '报告id',
  `user_name` varchar(100) COLLATE utf8_bin NOT NULL COMMENT '评论的用户',
  `comment` text COLLATE utf8_bin NOT NULL COMMENT '评论',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of wt_comments
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wt_reports
-- ----------------------------
DROP TABLE IF EXISTS `wt_reports`;
CREATE TABLE `wt_reports` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `user_id` int(11) DEFAULT NULL COMMENT '用户id',
  `send_to` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '发送给谁',
  `header` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '报告标题名',
  `contents` mediumtext COLLATE utf8_bin COMMENT '报告内容',
  `pictures` text COLLATE utf8_bin COMMENT '图片',
  `attachments` text COLLATE utf8_bin COMMENT '附件',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=157 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of wt_reports
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for wt_rules
-- ----------------------------
DROP TABLE IF EXISTS `wt_rules`;
CREATE TABLE `wt_rules` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `reporters` varchar(500) COLLATE utf8_bin DEFAULT NULL COMMENT '需要提交报告的用户名ID',
  `start_week` int(10) DEFAULT NULL COMMENT '提交报告的开始在周几',
  `start_hour` int(10) DEFAULT NULL COMMENT '提交周报的开始时间小时',
  `end_week` int(10) DEFAULT NULL COMMENT '提交报告的结束在周几',
  `end_hour` int(10) DEFAULT NULL COMMENT '提交报告的结束时间小时',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of wt_rules
-- ----------------------------
BEGIN;
INSERT INTO `wt_rules` VALUES (1, 1, '[{\"id\":2,\"name\":\"user\"}]', 1, 2, 7, 23, '2021-11-08 06:47:34', '2021-11-15 09:16:40', NULL);
COMMIT;

-- ----------------------------
-- Table structure for wt_templates
-- ----------------------------
DROP TABLE IF EXISTS `wt_templates`;
CREATE TABLE `wt_templates` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '用户名',
  `header` varchar(100) COLLATE utf8_bin DEFAULT NULL COMMENT '报告标题名',
  `contents` mediumtext COLLATE utf8_bin COMMENT '报告内容',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=104 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of wt_templates
-- ----------------------------
BEGIN;
INSERT INTO `wt_templates` VALUES (103, 'admin', '网藤周报', '[{\"title\":\"本周工作\",\"content\":\"\"},{\"title\":\"下周计划\",\"content\":\"\"}]', '2021-11-05 10:11:16', '2021-11-15 07:58:24', NULL);
COMMIT;

-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
DROP VIEW IF EXISTS `authority_menu`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`close_tab` AS `close_tab`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;