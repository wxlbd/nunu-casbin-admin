/*
 Navicat Premium Dump SQL

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 50744 (5.7.44)
 Source Host           : localhost:3306
 Source Schema         : mineadmin

 Target Server Type    : MySQL
 Target Server Version : 50744 (5.7.44)
 File Encoding         : 65001

 Date: 06/02/2025 09:56:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for attachment
-- ----------------------------
DROP TABLE IF EXISTS `attachment`;
CREATE TABLE `attachment` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `storage_mode` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'local' COMMENT '存储模式:local=本地,oss=阿里云,qiniu=七牛云,cos=腾讯云',
  `origin_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '原文件名',
  `object_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '新文件名',
  `hash` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件hash',
  `mime_type` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '资源类型',
  `storage_path` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '存储目录',
  `suffix` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件后缀',
  `size_byte` bigint(20) DEFAULT NULL COMMENT '字节数',
  `size_info` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '文件大小',
  `url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'url地址',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `attachment_hash_unique` (`hash`),
  KEY `attachment_storage_path_index` (`storage_path`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='上传文件信息表';

-- ----------------------------
-- Records of attachment
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=245 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (235, 'p', 'testrole', '/api/system/menu', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (236, 'p', 'testrole', '/api/system/menu', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (237, 'p', 'testrole', '/api/system/menu/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (238, 'p', 'testrole', '/api/system/menu/:ids', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (239, 'p', 'testrole', '/api/system/role', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (240, 'p', 'testrole', '/api/system/role', 'POST', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (241, 'p', 'testrole', '/api/system/role/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (243, 'p', 'testrole', '/api/system/role/:id/menus', 'GET', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (244, 'p', 'testrole', '/api/system/role/:id/menus', 'PUT', '', '', '');
INSERT INTO `casbin_rule` (`id`, `ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES (242, 'p', 'testrole', '/api/system/role/:ids', 'DELETE', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for dict_data
-- ----------------------------
DROP TABLE IF EXISTS `dict_data`;
CREATE TABLE `dict_data` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type_code` varchar(30) NOT NULL COMMENT '字典类型编码',
  `label` varchar(20) NOT NULL COMMENT '字典key',
  `value` varchar(255) NOT NULL COMMENT '字典值',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '字典状态:1-正常,2-禁用',
  `sort` int(11) NOT NULL COMMENT '排序',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of dict_data
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for dict_types
-- ----------------------------
DROP TABLE IF EXISTS `dict_types`;
CREATE TABLE `dict_types` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `code` varchar(30) NOT NULL COMMENT '字典类型编码',
  `name` varchar(20) NOT NULL COMMENT '字典key',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '字典状态:1-正常,2-禁用',
  `sort` int(11) NOT NULL COMMENT '排序',
  `remark` varchar(255) NOT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of dict_types
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` bigint(20) unsigned NOT NULL COMMENT '父ID',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `meta` json DEFAULT NULL COMMENT '附加属性',
  `path` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路径',
  `component` varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `redirect` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '重定向地址',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:1=正常,2=停用',
  `sort` smallint(6) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(60) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `type` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `menu_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单信息表';

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (1, 0, 'permission', '{\"i18n\": \"baseMenu.permission.index\", \"icon\": \"ri:git-repository-private-line\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"权限管理\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/permission', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', 'menu');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (2, 1, 'permission:user', '{\"i18n\": \"baseMenu.permission.user\", \"icon\": \"material-symbols:manage-accounts-outline\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"用户管理\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/permission/user', 'base/views/permission/user/index', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:45:57', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (3, 2, 'permission:user:index', '{\"i18n\": \"baseMenu.permission.userList\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"用户列表\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:45:57', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (4, 2, 'permission:user:save', '{\"i18n\": \"baseMenu.permission.userSave\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"用户保存\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:45:57', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (5, 2, 'permission:user:update', '{\"i18n\": \"baseMenu.permission.userUpdate\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"用户更新\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:45:57', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (6, 2, 'permission:user:delete', '{\"i18n\": \"baseMenu.permission.userDelete\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"用户删除\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:45:57', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (7, 2, 'permission:user:set:password', '{\"i18n\": \"baseMenu.permission.userPassword\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"用户初始化密码\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:45:57', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (8, 2, 'permission:user:get:roles', '{\"i18n\": \"baseMenu.permission.getUserRole\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"获取用户角色\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:45:57', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (9, 2, 'permission:user:set:roles', '{\"i18n\": \"baseMenu.permission.setUserRole\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"用户角色赋予\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:45:57', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (10, 1, 'permission:menu', '{\"i18n\": \"baseMenu.permission.menu\", \"icon\": \"ph:list-bold\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"菜单管理\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/permission/menu', 'base/views/permission/menu/index', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 16:21:17', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (11, 10, 'permission:menu:tree', '{\"i18n\": \"baseMenu.permission.menuList\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"菜单列表\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-16 10:16:54', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (12, 10, 'permission:menu:create', '{\"i18n\": \"baseMenu.permission.menuSave\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"菜单保存\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 16:21:17', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (13, 10, 'permission:menu:update', '{\"i18n\": \"baseMenu.permission.menuUpdate\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"菜单更新\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 16:21:17', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (14, 10, 'permission:menu:delete', '{\"i18n\": \"baseMenu.permission.menuDelete\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"菜单删除\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 16:21:17', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (15, 1, 'permission:role', '{\"i18n\": \"baseMenu.permission.role\", \"icon\": \"material-symbols:supervisor-account-outline-rounded\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"角色管理\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/permission/role', 'base/views/permission/role/index', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 13:13:04', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (16, 15, 'permission:role:index', '{\"i18n\": \"baseMenu.permission.roleList\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"角色列表\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 13:13:04', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (17, 15, 'permission:role:create', '{\"i18n\": \"baseMenu.permission.roleSave\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"角色创建\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 13:17:24', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (18, 15, 'permission:role:update', '{\"i18n\": \"baseMenu.permission.roleUpdate\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"角色更新\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 13:13:04', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (19, 15, 'permission:role:delete', '{\"i18n\": \"baseMenu.permission.roleDelete\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"角色删除\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 13:13:04', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (20, 15, 'permission:role:get:menus', '{\"i18n\": \"baseMenu.permission.getRolePermission\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"获取角色权限\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 13:13:04', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (21, 15, 'permission:role:set:menus', '{\"i18n\": \"baseMenu.permission.setRolePermission\", \"icon\": \"\", \"type\": \"B\", \"affix\": false, \"cache\": false, \"title\": \"赋予角色权限\", \"hidden\": false, \"copyright\": false, \"componentPath\": \"\", \"componentSuffix\": \"\", \"breadcrumbEnable\": false}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 13:13:04', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (22, 0, 'log', '{\"i18n\": \"baseMenu.log.index\", \"icon\": \"ph:instagram-logo\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"日志管理\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/log', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (23, 22, 'log:userLogin', '{\"i18n\": \"baseMenu.log.userLoginLog\", \"icon\": \"ph:user-list\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"用户登录日志管理\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/log/userLoginLog', 'base/views/log/userLogin', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (24, 23, 'log:userLogin:list', '{\"i18n\": \"baseMenu.log.userLoginLogList\", \"type\": \"B\", \"title\": \"用户登录日志列表\"}', '/log/userLoginLog', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (25, 23, 'log:userLogin:delete', '{\"i18n\": \"baseMenu.log.userLoginLogDelete\", \"type\": \"B\", \"title\": \"删除用户登录日志\"}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (26, 22, 'log:userOperation', '{\"i18n\": \"baseMenu.log.operationLog\", \"icon\": \"ph:list-magnifying-glass\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"操作日志管理\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/log/operationLog', 'base/views/log/userOperation', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (27, 26, 'log:userOperation:list', '{\"i18n\": \"baseMenu.log.userOperationLog\", \"type\": \"B\", \"title\": \"用户操作日志列表\"}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (28, 26, 'log:userOperation:delete', '{\"i18n\": \"baseMenu.log.userOperationLogDelete\", \"type\": \"B\", \"title\": \"删除用户操作日志\"}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (29, 0, 'dataCenter', '{\"i18n\": \"baseMenu.dataCenter.index\", \"icon\": \"ri:database-line\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"数据中心\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/dataCenter', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (30, 29, 'dataCenter:attachment', '{\"i18n\": \"baseMenu.dataCenter.attachment\", \"icon\": \"ri:attachment-line\", \"type\": \"M\", \"affix\": false, \"cache\": true, \"title\": \"附件管理\", \"hidden\": false, \"copyright\": true, \"componentPath\": \"modules/\", \"componentSuffix\": \".vue\", \"breadcrumbEnable\": true}', '/dataCenter/attachment', 'base/views/dataCenter/attachment/index', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (31, 30, 'dataCenter:attachment:list', '{\"i18n\": \"baseMenu.dataCenter.attachmentList\", \"type\": \"B\", \"title\": \"附件列表\"}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (32, 30, 'dataCenter:attachment:upload', '{\"i18n\": \"baseMenu.dataCenter.attachmentUpload\", \"type\": \"B\", \"title\": \"上传附件\"}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (33, 30, 'dataCenter:attachment:delete', '{\"i18n\": \"baseMenu.dataCenter.attachmentDelete\", \"type\": \"B\", \"title\": \"删除附件\"}', '', '', '', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '', '');
INSERT INTO `menu` (`id`, `parent_id`, `name`, `meta`, `path`, `component`, `redirect`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`, `type`) VALUES (34, 0, 'PurePermission', NULL, '/permission', '', '', 1, 0, 0, 0, NULL, NULL, '', 'menu');
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色名称',
  `code` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色代码',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:1=正常,2=停用',
  `sort` smallint(6) NOT NULL DEFAULT '0' COMMENT '排序',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_code_unique` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色信息表';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (1, '超级管理员', 'SuperAdmin', 1, 0, 0, 0, '2025-01-15 11:22:58', '2025-01-15 11:22:58', '');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (2, '测试角色', 'testrole', 1, 1, 0, 0, '2025-01-15 13:14:23', '2025-02-05 14:45:58', '备注');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (6, 'Hui Chi Ming', 'crYgsHy02S', 1, 776, 499, 161, '2018-12-14 08:24:41', '2007-05-07 12:09:42', 'LNm8rSMG74');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (7, 'Sugiyama Nanami', 'iFM0eBwqoq', 1, 186, 138, 965, '2002-01-27 17:48:50', '2001-04-21 05:40:01', 'zwurjHaVfM');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (8, 'Joseph Rodriguez', '8ORx1RvYvB', 1, 163, 981, 474, '2014-07-24 07:33:02', '2015-05-16 16:02:49', 'cR1M3bOIW8');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (9, 'Margaret Fisher', 'BoB6x0v2EB', 1, 797, 279, 333, '2004-07-01 22:28:05', '2010-10-27 06:29:41', 'wfbyHQTSdE');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (10, 'Florence Morris', 'E5C3CSNNxU', 1, 53, 684, 97, '2005-02-03 00:41:03', '2004-08-21 21:07:44', 'UcZuZP5LTU');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (11, 'Chen Zhiyuan', 'K98HkFevAg', 1, 123, 446, 893, '2020-06-03 16:49:31', '2009-02-11 18:33:26', '9lbiHUrUTN');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (12, 'Mo Yuning', 'OxQFuuFLVx', 1, 912, 826, 709, '2006-04-15 16:18:29', '2003-05-07 22:59:36', 'b9dstXOS5i');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (14, 'Lau Chi Ming', 'yDI12iSZCG', 1, 485, 842, 266, '2012-02-29 08:23:51', '2022-05-15 01:54:08', 'ccUzr9gsTf');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (15, 'Donna Myers', 'tN8GUHMndE', 1, 681, 38, 330, '2018-09-25 03:15:36', '2014-11-10 08:59:40', '8DY1422AlO');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (16, 'Donald Henderson', 'euwmVRsndg', 1, 17, 558, 733, '2017-03-19 07:37:58', '2007-08-19 07:29:43', 'zBVstsxu3r');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (17, 'Sasaki Nanami', 'CN2ehoO1jN', 1, 712, 81, 916, '2024-06-02 11:14:43', '2016-08-07 18:54:27', 'hFy5qZqaRB');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (18, 'Patrick Gordon', '2LGygYXqhv', 1, 659, 832, 584, '2013-11-19 05:49:24', '2009-04-08 09:45:45', '2Q4hZqpfoJ');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (19, 'Choi Ming', 'XVmv6zMr1L', 1, 831, 573, 865, '2014-11-12 03:35:30', '2022-01-04 16:51:32', 'Jc1uqAGZM8');
INSERT INTO `role` (`id`, `name`, `code`, `status`, `sort`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (20, 'Nakano Ryota', 'zWCqcpMIut', 1, 599, 59, 778, '2023-10-22 03:21:14', '2018-07-16 21:51:51', 'Omto7GZLnR');
COMMIT;

-- ----------------------------
-- Table structure for role_menus
-- ----------------------------
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=302 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色菜单映射表';

-- ----------------------------
-- Records of role_menus
-- ----------------------------
BEGIN;
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (289, 2, 2, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (290, 2, 5, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (291, 2, 6, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (292, 2, 7, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (293, 2, 8, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (294, 2, 3, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (295, 2, 9, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (296, 2, 10, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (297, 2, 11, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (298, 2, 12, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (299, 2, 17, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (300, 2, 18, NULL, NULL);
INSERT INTO `role_menus` (`id`, `role_id`, `menu_id`, `created_at`, `updated_at`) VALUES (301, 2, 1, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint(20) DEFAULT '0' COMMENT '父菜单ID',
  `menu_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '菜单类型（1代表菜单、2代表iframe、3代表外链、4代表按钮）',
  `title` varchar(50) NOT NULL COMMENT '菜单名称',
  `name` varchar(50) DEFAULT NULL COMMENT '路由名称',
  `path` varchar(200) DEFAULT '' COMMENT '路由路径',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `rank` int(11) DEFAULT '99' COMMENT '显示排序',
  `redirect` varchar(255) DEFAULT NULL COMMENT '重定向地址',
  `icon` varchar(100) DEFAULT NULL COMMENT '菜单图标',
  `extra_icon` varchar(100) DEFAULT NULL COMMENT '右侧图标',
  `enter_transition` varchar(50) DEFAULT NULL COMMENT '进场动画',
  `leave_transition` varchar(50) DEFAULT NULL COMMENT '离场动画',
  `active_path` varchar(255) DEFAULT NULL COMMENT '激活路由路径',
  `auths` varchar(500) DEFAULT NULL COMMENT '权限标识',
  `frame_src` varchar(500) DEFAULT NULL COMMENT '外链地址',
  `frame_loading` tinyint(1) DEFAULT '1' COMMENT '是否显示加载动画',
  `keep_alive` tinyint(1) DEFAULT '0' COMMENT '是否缓存',
  `hidden_tag` tinyint(1) DEFAULT '0' COMMENT '是否隐藏标签',
  `fixed_tag` tinyint(1) DEFAULT '0' COMMENT '是否固定标签',
  `show_link` tinyint(1) DEFAULT '1' COMMENT '是否显示',
  `show_parent` tinyint(1) DEFAULT '0' COMMENT '是否显示父级菜单',
  `status` tinyint(4) DEFAULT '1' COMMENT '菜单状态（0停用 1正常）',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COMMENT='菜单权限表';

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (1, 0, 1, 'menus.pureSysManagement', 'PureSystem', '/system', '', 13, '', 'ri:settings-3-line', '', '', '', '', '', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-24 18:49:41', '2025-02-05 14:38:47');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (2, 1, 1, 'menus.pureSystemMenu', 'SystemMenu', '/system/menu/index', '', 99, '', 'ep:menu', '', '', '', '', '', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-24 18:50:12', '2025-01-24 18:50:12');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (3, 1, 1, 'menus.pureRole', 'SystemRole', '/system/role/index', '', 99, '', 'ri:admin-fill', '', '', '', '', '', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-24 18:50:40', '2025-01-24 18:50:40');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (4, 1, 1, 'menus.pureUser', 'SystemUser', '/system/user/index', '', 99, '', 'ri:admin-line', '', '', '', '', '', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-24 18:51:03', '2025-01-24 18:51:03');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (5, 2, 4, '菜单列表', '', '', '', 99, '', '', '', '', '', '', 'system:menu:list', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 10:50:14', '2025-01-25 10:50:14');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (6, 2, 4, '菜单添加', NULL, '', NULL, 99, NULL, NULL, NULL, NULL, NULL, NULL, 'system:menu:create', NULL, 1, 0, 0, 0, 1, 0, 1, '2025-01-25 10:56:06', '2025-01-25 10:56:10');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (7, 2, 4, '菜单修改', NULL, '', NULL, 99, NULL, NULL, NULL, NULL, NULL, NULL, 'system:menu:update', NULL, 1, 0, 0, 0, 1, 0, 1, '2025-01-25 10:57:14', NULL);
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (8, 2, 4, '菜单删除', '', '', '', 99, '', '', '', '', '', '', 'system:menu:delete', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 10:58:12', '2025-01-25 10:58:12');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (9, 3, 4, '角色列表', '', '', '', 99, '', '', '', '', '', '', 'system:role:list', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 10:58:57', '2025-01-25 10:58:57');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (10, 3, 4, '角色添加', '', '', '', 99, '', '', '', '', '', '', 'system:role:create', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 11:00:29', '2025-01-25 11:00:29');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (11, 3, 4, '角色修改', '', '', '', 99, '', '', '', '', '', '', 'system:role:update', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 11:00:59', '2025-01-25 11:00:59');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (12, 3, 4, '角色删除', '', '', '', 99, '', '', '', '', '', '', 'system:role:delete', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 11:09:18', '2025-01-25 11:09:18');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (13, 4, 4, '用户列表', '', '', '', 99, '', '', '', '', '', '', 'system:user:list', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 11:09:56', '2025-01-25 11:09:56');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (14, 4, 4, '用户添加', '', '', '', 99, '', '', '', '', '', '', 'system:user:create', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 11:12:23', '2025-01-25 11:12:23');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (15, 4, 4, '用户修改', '', '', '', 99, '', '', '', '', '', '', 'system:user:update', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 11:12:50', '2025-01-25 11:12:50');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (16, 4, 4, '用户删除', '', '', '', 99, '', '', '', '', '', '', 'system:user:delete', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 11:13:25', '2025-01-25 11:13:25');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (17, 3, 4, '获取角色权限', '', '', '', 99, '', '', '', '', '', '', 'system:role:get:menus', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 15:13:28', '2025-01-25 15:13:28');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (18, 3, 4, '赋予角色权限', '', '', '', 99, '', '', '', '', '', '', 'system:role:set:menus', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 15:14:01', '2025-01-25 15:14:01');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (19, 4, 4, '获取用户角色', '', '', '', 99, '', '', '', '', '', '', 'system:user:get:roles', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 15:15:20', '2025-01-25 15:15:20');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (20, 4, 4, '赋予用户角色', '', '', '', 99, '', '', '', '', '', '', 'system:user:set:roles', '', 1, 0, 0, 0, 1, 0, 1, '2025-01-25 15:15:45', '2025-01-25 15:15:45');
INSERT INTO `sys_menus` (`id`, `parent_id`, `menu_type`, `title`, `name`, `path`, `component`, `rank`, `redirect`, `icon`, `extra_icon`, `enter_transition`, `leave_transition`, `active_path`, `auths`, `frame_src`, `frame_loading`, `keep_alive`, `hidden_tag`, `fixed_tag`, `show_link`, `show_parent`, `status`, `created_at`, `updated_at`) VALUES (21, 0, 1, 'menus.pureHome', 'Home', '/welcome', '', 1, '', 'ep:home-filled', '', '', '', '', '', '', 1, 0, 0, 0, 1, 0, 1, '2025-02-06 09:54:19', '2025-02-06 09:55:24');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID,主键',
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `user_type` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '100' COMMENT '用户类型:100=系统用户',
  `nickname` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `phone` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机',
  `email` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `signed` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '个人签名',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态:1=正常,2=停用',
  `login_ip` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '127.0.0.1' COMMENT '最后登陆IP',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后登陆时间',
  `backend_setting` json DEFAULT NULL COMMENT '后台设置数据',
  `created_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
  `updated_by` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新者',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户信息表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `username`, `password`, `user_type`, `nickname`, `phone`, `email`, `avatar`, `signed`, `status`, `login_ip`, `login_time`, `backend_setting`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (1, 'admin', '$2y$10$T3Po5Ufu1pKiKczWqp.dbOOjmeZ4H3Oj0daATqlqXsZOvrRW2s2IS', '100', '创始人', '16858888988', 'admin@adminmine.com', '', '广阔天地，大有所为', 1, '127.0.0.1', '2025-02-06 08:42:37', NULL, 0, 0, '2025-01-15 11:22:58', '2025-02-06 08:42:37', '');
INSERT INTO `user` (`id`, `username`, `password`, `user_type`, `nickname`, `phone`, `email`, `avatar`, `signed`, `status`, `login_ip`, `login_time`, `backend_setting`, `created_by`, `updated_by`, `created_at`, `updated_at`, `remark`) VALUES (2, 'test', '$2a$10$Tl8cyMEFtXp7mBmG3KZ49.q0CiCRejuxw6cjNwr/CRiwuWBbNJk2a', '100', '测试用户', '16711411400', '', '', '', 1, '', '2025-02-05 14:23:20', '{\"app\": {\"layout\": \"\", \"asideDark\": false, \"colorMode\": \"\", \"useLocale\": \"\", \"whiteRoute\": null, \"pageAnimate\": \"\", \"primaryColor\": \"\", \"watermarkText\": \"\", \"showBreadcrumb\": false, \"enableWatermark\": false, \"loadUserSetting\": false}, \"tabbar\": {\"mode\": \"\", \"enable\": false}, \"subAside\": {\"showIcon\": false, \"showTitle\": false, \"fixedAsideState\": false, \"showCollapseButton\": false}, \"copyright\": {\"dates\": \"\", \"enable\": false, \"company\": \"\", \"website\": \"\", \"putOnRecord\": \"\"}, \"mainAside\": {\"showIcon\": false, \"showTitle\": false, \"enableOpenFirstRoute\": false}, \"welcomePage\": {\"icon\": \"\", \"name\": \"\", \"path\": \"\", \"title\": \"\"}}', 0, 0, '2025-01-15 13:20:34', '2025-02-06 09:29:59', '');
COMMIT;

-- ----------------------------
-- Table structure for user_login_log
-- ----------------------------
DROP TABLE IF EXISTS `user_login_log`;
CREATE TABLE `user_login_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录IP地址',
  `os` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '操作系统',
  `browser` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '浏览器',
  `status` smallint(6) NOT NULL DEFAULT '1' COMMENT '登录状态 (1成功 2失败)',
  `message` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '提示消息',
  `login_time` datetime NOT NULL COMMENT '登录时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `user_login_log_username_index` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='登录日志表';

-- ----------------------------
-- Records of user_login_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_operation_log
-- ----------------------------
DROP TABLE IF EXISTS `user_operation_log`;
CREATE TABLE `user_operation_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户名',
  `method` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求方式',
  `router` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求路由',
  `service_name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务名称',
  `ip` varchar(45) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求IP地址',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `user_operation_log_username_index` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- ----------------------------
-- Records of user_operation_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色映射表';

-- ----------------------------
-- Records of user_roles
-- ----------------------------
BEGIN;
INSERT INTO `user_roles` (`id`, `user_id`, `role_id`, `created_at`, `updated_at`) VALUES (1, 1, 1, NULL, NULL);
INSERT INTO `user_roles` (`id`, `user_id`, `role_id`, `created_at`, `updated_at`) VALUES (11, 2, 2, NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
