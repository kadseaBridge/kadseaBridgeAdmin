/*
 Navicat Premium Dump SQL

 Source Server         : bridge
 Source Server Type    : MySQL
 Source Server Version : 80400 (8.4.0)
 Source Host           : ec2-16-162-106-158.ap-east-1.compute.amazonaws.com:7006
 Source Schema         : bridgedb

 Target Server Type    : MySQL
 Target Server Version : 80400 (8.4.0)
 File Encoding         : 65001

 Date: 24/06/2024 19:02:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;




-- ----------------------------
-- Table structure for daily_bridge_stats
-- ----------------------------
DROP TABLE IF EXISTS `daily_bridge_stats`;
CREATE TABLE `daily_bridge_stats`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `date` date NOT NULL,
  `coin_address` varchar(45) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `chain_id` varchar(20) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `transfer_in` decimal(40, 20) NOT NULL DEFAULT 0.00000000000000000000,
  `transfer_out` decimal(40, 20) NOT NULL DEFAULT 0.00000000000000000000,
  `transfer_difference` decimal(40, 20) NOT NULL DEFAULT 0.00000000000000000000,
  `fee` decimal(40, 20) NOT NULL DEFAULT 0.00000000000000000000,
  `platform_assets` decimal(40, 20) NOT NULL DEFAULT 0.00000000000000000000,
  `financial_balance` decimal(40, 20) NOT NULL DEFAULT 0.00000000000000000000,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of daily_bridge_stats
-- ----------------------------
INSERT INTO `daily_bridge_stats` VALUES (1, '2024-06-14', 'USDT', 'KAD', 100.00000000000000000000, 100.00000000000000000000, 0.00000000000000000000, 10.00000000000000000000, 1000.00000000000000000000, '2024-06-14 17:22:51', '2024-06-14 17:22:51');
INSERT INTO `daily_bridge_stats` VALUES (2, '2024-06-11', 'tBNB', 'BSC', 88.00000000000000000000, 88.00000000000000000000, 0.00000000000000000000, 1.00000000000000000000, 100.00000000000000000000, '2024-06-14 18:07:18', '2024-06-14 18:07:18');


-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('g', '22', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '22', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '31', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '31', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '15', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '16', '3', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '5', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '6', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '41', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '42', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '43', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '86', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '8', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '16', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '17', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '18', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '44', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '45', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '46', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '47', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '48', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '49', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '40', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '50', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '51', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '52', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '53', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '54', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '55', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '36', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '56', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '57', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '58', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '59', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '60', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '37', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '38', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '61', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '62', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '39', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '63', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '64', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '65', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '66', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '67', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '68', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '35', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '69', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '70', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '71', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '89', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '103', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '119', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '104', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '106', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '90', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '112', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '115', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '116', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '117', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '93', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '113', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '118', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '114', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '107', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '108', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '109', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '110', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '111', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '2', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '3', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '34', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '35', '3', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '35', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '5', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '14', '3', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '36', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '37', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '38', '5', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '38', '8', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'alice', 'data1', 'read', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'bob', 'data2', 'write', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'data2_admin', 'data2', 'read', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'data2_admin', 'data2', 'write', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'alice', 'data2_admin', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '1', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '1', '3', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '245', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '246', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '251', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '257', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '315', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '43', '9', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '42', '10', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', '44', '10', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '341', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '55', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '36', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '56', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '57', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '58', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '59', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '60', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '37', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '38', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '61', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '62', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '39', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '63', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '64', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '75', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '76', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '186', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '226', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '227', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '228', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '229', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '230', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '231', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '232', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '340', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '329', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '252', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '253', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '254', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '255', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '256', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '258', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '259', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '260', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '261', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '262', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '315', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '251', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '257', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '316', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '317', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '318', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '319', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '320', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '10', '321', 'All', '', '', '');



-- ----------------------------
-- Table structure for plugins_manage
-- ----------------------------
DROP TABLE IF EXISTS `plugins_manage`;
CREATE TABLE `plugins_manage`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `store_id` int NULL DEFAULT NULL COMMENT '插件在商城中的id',
  `p_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '插件名称英文',
  `p_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '插件名称',
  `p_description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '插件介绍',
  `p_auth` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '作者',
  `is_install` tinyint NOT NULL DEFAULT 0 COMMENT '是否安装',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态',
  `version` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '当前版本',
  `price` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '价格',
  `download_times` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '下载次数',
  `install_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '安装路径（用于卸载）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `storeIdUni`(`store_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 47 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '插件管理' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of plugins_manage
-- ----------------------------
INSERT INTO `plugins_manage` VALUES (34, 1, 'mail', '邮件发送', '邮件发送', 'demo01', 1, 1, '1.0.0', 1, 1248, 'plugins');
INSERT INTO `plugins_manage` VALUES (35, 2, 'weichat', '微信小程序', '微信小程序管理接口，包含关注回复、关键字回复、图文消息、素材管理、菜单管理等', 'demo01', 0, 0, '1.0.2', 1, 441, NULL);
INSERT INTO `plugins_manage` VALUES (36, 3, 'cms', 'cms管理系统-mysql版', 'cms管理系统', 'gfast01', 0, 0, '3.2.14', 29900, 635, NULL);
INSERT INTO `plugins_manage` VALUES (37, 4, 'demo2', '免费测试', '666', 'demo01', 0, 0, '1.2', 0, 2, NULL);
INSERT INTO `plugins_manage` VALUES (38, 12, 'gfast-v3.2-donate', 'gfast-v3.2捐赠版', '基于全新Go Frame 2.3+Vue3+Element Plus开发的全栈前后端分离的管理系统', 'gfast01', 0, 0, '3.2.1', 6800, 106, NULL);
INSERT INTO `plugins_manage` VALUES (39, 13, 'gfast-v3.2-mandatep', 'gfast-v3.2-个人授权版', '基于全新Go Frame 2.3+Vue3+Element Plus开发的全栈前后端分离的管理系统', 'gfast01', 0, 0, '3.2.26', 39800, 1407, NULL);
INSERT INTO `plugins_manage` VALUES (40, 14, 'gfast-v3.2-mandateu', 'gfast-v3.2-企业授权版', '基于全新Go Frame 2.3+Vue3+Element Plus开发的全栈前后端分离的管理系统', 'gfast01', 0, 0, '3.2.26', 58800, 752, NULL);
INSERT INTO `plugins_manage` VALUES (41, 5, 'mail', '邮件发送', 'qq smtp邮件发送功能', 'wilgx0', 0, 0, '3.2.1', 0, 167, NULL);
INSERT INTO `plugins_manage` VALUES (42, 6, 'wechat', '微信公众号', '微信小程序管理接口，包含关注回复、关键字回复、图文消息、素材管理、菜单管理等', 'gfast01', 0, 0, '1.0.1', 15900, 57, NULL);
INSERT INTO `plugins_manage` VALUES (43, 7, 'gencode', '代码生成', '代码生成功能完善，支持生成上传文件、图片及富文本编辑器功能', 'gfast01', 0, 0, '3.2.3', 19900, 422, NULL);
INSERT INTO `plugins_manage` VALUES (44, 8, 'im', '在线客服', '在线客服', 'wilgx0', 0, 0, '1.0.1', 9900, 22, NULL);
INSERT INTO `plugins_manage` VALUES (45, 9, 'doorkeeper', 'IP过滤', 'IP地址过滤，支持ipv4 ipv6 ,可设置访问ip白名单和黑名单', 'wilgx0', 0, 0, '1.0.1', 1990, 65, NULL);
INSERT INTO `plugins_manage` VALUES (46, 10, 'smsPlugin', '短信发送', '短信发送，支持网易云信、腾讯云、阿里云短信平台，可自定义配置。', 'wilgx0', 0, 0, 'v3.2.4', 1999, 66, NULL);

-- ----------------------------
-- Table structure for sys_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth_rule`;
CREATE TABLE `sys_auth_rule`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '条件',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 0目录 1菜单 2按钮',
  `weigh` int NOT NULL DEFAULT 0 COMMENT '权重',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `always_show` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示状态',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `jump_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '跳转路由',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_frame` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '所属模块',
  `model_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 342 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '菜单节点表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_auth_rule
-- ----------------------------
INSERT INTO `sys_auth_rule` VALUES (1, 341, 'system/config', '系统配置', 'system', '', '', 0, 8, 1, 1, 'config', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 15:46:11', NULL);
INSERT INTO `sys_auth_rule` VALUES (2, 341, 'system/auth', '权限管理', 'peoples', '', '', 0, 7, 1, 1, 'system/auth', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 15:46:21', NULL);
INSERT INTO `sys_auth_rule` VALUES (3, 341, 'system/monitor', '系统监控', 'monitor', '', '', 0, 6, 1, 1, 'monitor', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 15:46:38', NULL);
INSERT INTO `sys_auth_rule` VALUES (5, 1, 'system/config/dict/type/list', '字典管理', 'dict', '', '', 1, 0, 1, 1, 'dict/list', '', 'system/config/dict/list', 0, 'sys_admin', 0, NULL, '2021-07-20 08:47:37', NULL);
INSERT INTO `sys_auth_rule` VALUES (6, 1, 'system/config/sysConfig/list', '参数管理', 'date-range', '', '', 1, 0, 1, 1, 'params/list', '', 'system/config/params/list', 0, 'sys_admin', 0, NULL, '2021-07-20 08:58:57', NULL);
INSERT INTO `sys_auth_rule` VALUES (8, 2, 'system/auth/menuList', '菜单管理', 'nested', '', '', 1, 0, 1, 1, 'menuList', '', 'system/auth/menuList', 0, 'sys_admin', 0, NULL, '2021-07-20 09:01:49', NULL);
INSERT INTO `sys_auth_rule` VALUES (21, 5, 'system/config/dict/type/delete', '删除字典', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:56:49', NULL);
INSERT INTO `sys_auth_rule` VALUES (16, 8, 'system/auth/addMenu', '添加菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:02:20', NULL);
INSERT INTO `sys_auth_rule` VALUES (17, 8, 'system/auth/editMenu', '修改菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (18, 8, 'system/auth/deleteMenu', '删除菜单', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (19, 5, 'system/config/dict/type/add', '添加字典', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:56:29', NULL);
INSERT INTO `sys_auth_rule` VALUES (20, 5, 'system/config/dict/type/edit', '修改字典', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:56:39', NULL);
INSERT INTO `sys_auth_rule` VALUES (22, 5, 'system/config/dict/data/list', '字典数据管理', '', '', '', 2, 0, 1, 1, 'dataList', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:57:25', NULL);
INSERT INTO `sys_auth_rule` VALUES (23, 5, 'system/config/dict/data/add', '添加字典数据', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:58:06', NULL);
INSERT INTO `sys_auth_rule` VALUES (24, 5, 'system/config/dict/data/edit', '修改字典数据', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:58:14', NULL);
INSERT INTO `sys_auth_rule` VALUES (25, 5, 'system/config/dict/data/delete', '删除字典数据', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:58:22', NULL);
INSERT INTO `sys_auth_rule` VALUES (26, 2, 'system/auth/roleList', '角色管理', 'logininfor', '', '', 1, 0, 1, 1, 'role', '', 'system/auth/roleList', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (27, 26, 'system/auth/addRole', '添加角色', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (28, 2, 'system/auth/deptList', '部门管理', 'peoples', '', '', 1, 0, 1, 1, 'dept', '', 'system/auth/dept', 0, 'sys_admin', 0, NULL, '2021-07-20 09:03:46', NULL);
INSERT INTO `sys_auth_rule` VALUES (29, 26, 'system/auth/editRole', '修改角色', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (30, 26, 'system/auth/statusSetRole', '设置角色状态', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (31, 26, 'system/auth/deleteRole', '删除角色', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (33, 2, 'system/auth/postList', '岗位管理', 'tab', '', '', 1, 0, 1, 1, 'post', '', 'system/auth/post', 0, 'sys_admin', 0, NULL, '2021-07-20 09:04:51', NULL);
INSERT INTO `sys_auth_rule` VALUES (34, 3, 'system/monitor/online/list', '在线用户', 'cascader', '', '', 1, 0, 1, 1, 'online', '', 'system/monitor/online/list', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (36, 3, 'system/monitor/job', '定时任务', 'clipboard', '', '', 1, 0, 1, 1, 'job', '', 'system/monitor/job', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (37, 3, 'system/monitor/server/info', '服务监控', 'dict', '', '', 1, 0, 1, 1, 'server', '', 'system/monitor/server', 0, 'sys_admin', 0, NULL, '2021-07-19 16:07:23', NULL);
INSERT INTO `sys_auth_rule` VALUES (38, 3, 'system/monitor/loginLog', '登录日志', 'chart', '', '', 1, 0, 1, 1, 'logininfor', '', 'system/monitor/logininfor', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:00', NULL);
INSERT INTO `sys_auth_rule` VALUES (39, 3, 'system/monitor/operLog', '操作日志', 'dashboard', '', '', 1, 0, 1, 1, 'operlog', '', 'system/monitor/operlog', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:31', NULL);
INSERT INTO `sys_auth_rule` VALUES (40, 2, 'system/auth/userList', '用户管理', 'user', '', '', 1, 0, 1, 1, 'user', '', 'system/auth/userList', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (41, 6, 'system/config/sysConfig/add', '添加参数', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:59:09', NULL);
INSERT INTO `sys_auth_rule` VALUES (42, 6, 'system/config/sysConfig/edit', '修改参数', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:59:17', NULL);
INSERT INTO `sys_auth_rule` VALUES (43, 6, 'system/config/sysConfig/delete', '删除参数', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 08:59:25', NULL);
INSERT INTO `sys_auth_rule` VALUES (44, 28, 'system/auth/deptAdd', '添加部门', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:04:05', NULL);
INSERT INTO `sys_auth_rule` VALUES (45, 28, 'system/auth/deptEdit', '修改部门', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:04:19', NULL);
INSERT INTO `sys_auth_rule` VALUES (46, 28, 'system/auth/deptDelete', '删除部门', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:04:35', NULL);
INSERT INTO `sys_auth_rule` VALUES (47, 33, 'system/auth/postAdd', '添加岗位', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:05:04', NULL);
INSERT INTO `sys_auth_rule` VALUES (48, 33, 'system/auth/postEdit', '修改岗位', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:05:15', NULL);
INSERT INTO `sys_auth_rule` VALUES (49, 33, 'system/auth/postDelete', '删除岗位', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:05:25', NULL);
INSERT INTO `sys_auth_rule` VALUES (50, 40, 'system/auth/addUser', '添加用户', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (51, 40, 'system/auth/editUser', '修改用户', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (52, 40, 'system/auth/resetUserPwd', '密码重置', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (53, 40, 'system/auth/changeUserStatus', '状态设置', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (54, 40, 'system/auth/deleteUser', '删除用户', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:06:16', NULL);
INSERT INTO `sys_auth_rule` VALUES (55, 34, 'system/monitor/online/forceLogout', '强制退出', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (56, 36, 'system/monitor/job/add', '添加任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (57, 36, 'system/monitor/job/edit', '修改任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (58, 36, 'system/monitor/job/start', '开启任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (59, 36, 'system/monitor/job/stop', '停止任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (60, 36, 'system/monitor/job/delete', '删除任务', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (61, 38, 'system/monitor/loginLog/delete', '删除', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:10', NULL);
INSERT INTO `sys_auth_rule` VALUES (62, 38, 'system/monitor/loginLog/clear', '清空', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:15', NULL);
INSERT INTO `sys_auth_rule` VALUES (63, 39, 'system/monitor/operLog/delete', '删除', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:36', NULL);
INSERT INTO `sys_auth_rule` VALUES (64, 39, 'system/monitor/operLog/clear', '清空', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, '2021-07-20 09:08:41', NULL);
INSERT INTO `sys_auth_rule` VALUES (75, 341, 'system/tools', '系统工具', 'server', '', '', 0, 5, 1, 1, 'system/tools', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 15:46:48', NULL);
INSERT INTO `sys_auth_rule` VALUES (76, 75, 'system/tools/build', '表单构建', 'build', '', '', 1, 0, 1, 1, 'build', '', 'system/tools/build', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (86, 1, 'system/config/sysWebSet', '站点设置', 'system', '', '', 1, 0, 1, 1, '/webSet', '', 'system/config/webSet', 0, 'sys_admin', 0, NULL, '2021-07-20 08:59:41', NULL);
INSERT INTO `sys_auth_rule` VALUES (185, 86, 'system/config/sysWebSet/update', '站点配置更新', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-07-20 09:00:18', '2021-07-20 09:00:18', NULL);
INSERT INTO `sys_auth_rule` VALUES (186, 75, 'system/tools/gen/tableList', '代码生成', 'code', '', '', 1, 0, 1, 1, 'gen', '', 'system/tools/gen', 0, 'sys_admin', 0, '2021-08-16 10:36:54', '2021-08-16 12:12:49', NULL);
INSERT INTO `sys_auth_rule` VALUES (231, 227, 'system/pluginsManage/delete', '插件管理删除', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:12:09', '2021-12-22 10:12:09', NULL);
INSERT INTO `sys_auth_rule` VALUES (232, 227, 'system/pluginsManage/changeStatus', '插件管理状态修改', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:12:31', '2021-12-22 10:12:31', NULL);
INSERT INTO `sys_auth_rule` VALUES (230, 227, 'system/pluginsManage/edit', '插件管理修改', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:11:48', '2021-12-22 10:11:48', NULL);
INSERT INTO `sys_auth_rule` VALUES (229, 227, 'system/pluginsManage/add', '插件管理添加', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:11:29', '2021-12-22 10:11:29', NULL);
INSERT INTO `sys_auth_rule` VALUES (228, 227, 'system/pluginsManage/get', '插件管理查询', '', '', '', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, '2021-12-22 10:10:54', '2021-12-22 10:10:54', NULL);
INSERT INTO `sys_auth_rule` VALUES (227, 75, 'system/pluginsManage/list', '插件管理', 'list', '', '', 1, 0, 1, 1, 'pluginsManageList', '', 'system/pluginsManage/list', 0, 'sys_admin', 0, '2021-12-22 10:10:15', '2021-12-22 10:10:15', NULL);
INSERT INTO `sys_auth_rule` VALUES (226, 75, 'system/tools/api', '系统接口', 'guide', '', '', 1, 0, 1, 1, 'api', '', 'system/tools/api', 0, 'sys_admin', 0, '2021-12-22 09:42:42', '2021-12-22 09:42:42', NULL);
INSERT INTO `sys_auth_rule` VALUES (254, 252, 'admin/coin/add', '币种管理添加', '', '', '币种管理添加', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (253, 252, 'admin/coin/get', '币种管理查询', '', '', '币种管理查询', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (252, 340, 'admin/coin/list', '币种管理列表', 'list', '', '', 1, 3, 1, 1, 'coinList', '', 'admin/coin/list', 0, 'sys_admin', 0, NULL, '2024-06-17 16:41:16', NULL);
INSERT INTO `sys_auth_rule` VALUES (251, 340, 'admin/coin', '币种管理', 'form', '', '', 0, 0, 1, 0, 'coin', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 16:37:53', NULL);
INSERT INTO `sys_auth_rule` VALUES (328, 340, 'admin/chain', '链管理', 'form', '', '', 0, 1, 1, 0, 'chain', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 16:38:00', NULL);
INSERT INTO `sys_auth_rule` VALUES (329, 340, 'admin/chain/list', '链管理列表', 'list', '', '', 1, 4, 1, 1, 'chainList', '', 'admin/chain/list', 0, 'sys_admin', 0, NULL, '2024-06-17 16:41:03', NULL);
INSERT INTO `sys_auth_rule` VALUES (330, 329, 'admin/chain/get', '链管理查询', '', '', '链管理查询', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (331, 329, 'admin/chain/add', '链管理添加', '', '', '链管理添加', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (332, 329, 'admin/chain/edit', '链管理修改', '', '', '链管理修改', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (333, 329, 'admin/chain/delete', '链管理删除', '', '', '链管理删除', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (255, 252, 'admin/coin/edit', '币种管理修改', '', '', '币种管理修改', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (256, 252, 'admin/coin/delete', '币种管理删除', '', '', '币种管理删除', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (257, 340, 'admin/bridgeConfig', '跨链币对管理', 'form', '', '', 0, 0, 1, 0, 'bridgeConfig', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 16:38:09', NULL);
INSERT INTO `sys_auth_rule` VALUES (258, 340, 'admin/bridgeConfig/list', '跨链币对列表', 'list', '', '', 1, 2, 1, 1, 'bridgeConfigList', '', 'admin/bridgeConfig/list', 0, 'sys_admin', 0, NULL, '2024-06-17 16:41:36', NULL);
INSERT INTO `sys_auth_rule` VALUES (259, 258, 'admin/bridgeConfig/get', '可跨币对查询', '', '', '可跨币对查询', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (260, 258, 'admin/bridgeConfig/add', '可跨币对添加', '', '', '可跨币对添加', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (261, 258, 'admin/bridgeConfig/edit', '可跨币对修改', '', '', '可跨币对修改', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (262, 258, 'admin/bridgeConfig/delete', '可跨币对删除', '', '', '可跨币对删除', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (316, 340, 'admin/bridgeOrder/list', '跨链记录列表', 'list', '', '', 1, 0, 1, 1, 'bridgeOrderList', '', 'admin/bridgeOrder/list', 0, 'sys_admin', 0, NULL, '2024-06-17 16:36:47', NULL);
INSERT INTO `sys_auth_rule` VALUES (317, 316, 'admin/bridgeOrder/get', '跨链记录查询', '', '', '跨链记录查询', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (318, 316, 'admin/bridgeOrder/add', '跨链记录添加', '', '', '跨链记录添加', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (319, 316, 'admin/bridgeOrder/edit', '跨链记录修改', '', '', '跨链记录修改', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (320, 316, 'admin/bridgeOrder/delete', '跨链记录删除', '', '', '跨链记录删除', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (321, 316, 'admin/bridgeOrder/changeStatus', '跨链记录跨链记录状态修改', '', '', '跨链记录跨链记录状态修改', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (315, 340, 'admin/bridgeOrder', '跨链记录管理', 'form', '', '', 0, 1, 1, 0, 'bridgeOrder', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 16:41:50', NULL);
INSERT INTO `sys_auth_rule` VALUES (336, 335, 'admin/dailyBridgeStats/get', '跨链日统计查询', '', '', '跨链日统计查询', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (335, 340, 'admin/dailyBridgeStats/list', '跨链日统计列表', 'list', '', '', 1, 0, 0, 0, 'dailyBridgeStatsList', '', 'admin/dailyBridgeStats/list', 0, 'sys_admin', 0, NULL, '2024-06-17 16:37:10', NULL);
INSERT INTO `sys_auth_rule` VALUES (334, 340, 'admin/dailyBridgeStats', '跨链日统计管理', 'form', '', '', 0, 0, 1, 0, 'dailyBridgeStats', '', '', 0, 'sys_admin', 0, NULL, '2024-06-17 16:38:18', NULL);
INSERT INTO `sys_auth_rule` VALUES (337, 335, 'admin/dailyBridgeStats/add', '跨链日统计添加', '', '', '跨链日统计添加', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (338, 335, 'admin/dailyBridgeStats/edit', '跨链日统计修改', '', '', '跨链日统计修改', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (339, 335, 'admin/dailyBridgeStats/delete', '跨链日统计删除', '', '', '跨链日统计删除', 2, 0, 1, 1, '', '', '', 0, 'sys_admin', 0, NULL, NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (340, 0, 'admin', '业务管理', '', '', '', 0, 0, 1, 1, 'admin', '', '', 0, 'sys_admin', 0, '2024-06-17 15:41:37', '2024-06-17 16:34:32', NULL);
INSERT INTO `sys_auth_rule` VALUES (341, 0, 'sys', '系统管理', '', '', '', 0, 1, 1, 1, 'sys', '', '', 0, 'sys_admin', 0, '2024-06-17 15:45:29', '2024-06-17 16:42:41', NULL);

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` tinyint(1) NULL DEFAULT 0 COMMENT '系统内置（Y是 N否）',
  `create_by` int UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE INDEX `uni_config_key`(`config_key`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '文件上传-文件大小', 'sys.uploadFile.fileSize', '50M', 1, 31, 31, '文件上传大小限制', NULL, '2021-07-06 14:57:35', NULL);
INSERT INTO `sys_config` VALUES (2, '文件上传-文件类型', 'sys.uploadFile.fileType', 'doc,docx,zip,xls,xlsx,rar,jpg,jpeg,gif,npm,png', 1, 31, 31, '文件上传后缀类型限制', NULL, NULL, NULL);
INSERT INTO `sys_config` VALUES (3, '图片上传-图片类型', 'sys.uploadFile.imageType', 'jpg,jpeg,gif,npm,png', 1, 31, 0, '图片上传后缀类型限制', NULL, NULL, NULL);
INSERT INTO `sys_config` VALUES (4, '图片上传-图片大小', 'sys.uploadFile.imageSize', '50M', 1, 31, 31, '图片上传大小限制', NULL, NULL, NULL);
INSERT INTO `sys_config` VALUES (11, '静态资源', 'static.resource', '/', 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_config` VALUES (12, '啊实打实的22', 'aasdd22', 'asdasd22', 0, 31, 31, 'dasdasd22', '2021-07-06 12:11:37', '2021-07-06 14:58:27', '2021-07-06 15:21:48');

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `created_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint NULL DEFAULT NULL COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 208 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '部门表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', 'KAD跨链桥', 0, '若依', '15888888888', 'ry@qq.com', '1', 0, 1, '2021-07-13 15:56:52', '2024-06-01 13:35:50', NULL);
INSERT INTO `sys_dept` VALUES (101, 100, '0,100', '深圳总公司', 1, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:38:20');
INSERT INTO `sys_dept` VALUES (102, 100, '0,100', '长沙分公司', 2, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:25:07');
INSERT INTO `sys_dept` VALUES (103, 101, '0,100,101', '研发部门', 1, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:38:20');
INSERT INTO `sys_dept` VALUES (104, 101, '0,100,101', '市场部门', 2, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:38:20');
INSERT INTO `sys_dept` VALUES (105, 101, '0,100,101', '测试部门', 3, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:38:20');
INSERT INTO `sys_dept` VALUES (106, 101, '0,100,101', '财务部门', 4, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:38:20');
INSERT INTO `sys_dept` VALUES (107, 101, '0,100,101', '运维部门', 5, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:38:20');
INSERT INTO `sys_dept` VALUES (108, 102, '0,100,102', '市场部门', 1, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:25:07');
INSERT INTO `sys_dept` VALUES (109, 102, '0,100,102', '财务部门', 2, '若依', '15888888888', 'ry@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:25:07');
INSERT INTO `sys_dept` VALUES (200, 100, '', '研发部门', 1, 'jimmy', '18888888888', '888@qq.com', '1', 0, 1, '2021-07-13 15:56:52', '2024-06-01 13:32:07', NULL);
INSERT INTO `sys_dept` VALUES (201, 100, '', '测试部门', 3, '老李', '18888888888', 'li@qq.com', '1', 0, 1, '2021-07-13 15:56:52', '2024-06-01 13:39:07', NULL);
INSERT INTO `sys_dept` VALUES (202, 108, '', '外勤', 1, '小a', '18888888888', 'aa@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:25:07');
INSERT INTO `sys_dept` VALUES (203, 108, '', '行政', 0, 'aa', '18888888888', 'aa@qq.com', '0', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2024-06-01 13:25:07');
INSERT INTO `sys_dept` VALUES (204, 100, '', '测试一下', 0, '张三', '13688999888', '656@qq.com', '1', 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2021-07-12 10:12:00');
INSERT INTO `sys_dept` VALUES (205, 100, '', '测试222', 0, '李四22', '13788556554', '1546@2ad.com', '1', 31, 31, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2021-07-12 10:12:20');
INSERT INTO `sys_dept` VALUES (206, 205, '', 'asdasd', 0, 'dasd', '13699878797', '4654@aa.com', '1', 31, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', '2021-07-12 10:12:20');
INSERT INTO `sys_dept` VALUES (207, 100, '', '市场部门', 2, '张三', '18888888888', '888@qq.com', '0', 1, 1, '2024-06-01 13:33:37', '2024-06-01 13:39:17', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认（1是 0否）',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 125 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '字典数据表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '男', '1', 'sys_user_sex', '', '', 0, 1, 31, 2, '备注信息', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (2, 0, '女', '2', 'sys_user_sex', '', '', 0, 1, 31, 31, '备注信息', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (3, 0, '保密', '0', 'sys_user_sex', '', '', 1, 1, 31, 31, '备注信息', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (28, 0, '正常', '0', 'sys_job_status', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, '暂停', '1', 'sys_job_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (30, 0, '默认', 'DEFAULT', 'sys_job_group', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '系统', 'SYSTEM', 'sys_job_group', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (32, 0, '成功', '1', 'admin_login_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (33, 0, '失败', '0', 'admin_login_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (34, 0, '成功', '1', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (35, 0, '失败', '0', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (36, 0, '重复执行', '1', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (37, 0, '执行一次', '2', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (38, 0, '显示', '1', 'sys_show_hide', NULL, 'default', 1, 1, 31, 0, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (39, 0, '隐藏', '0', 'sys_show_hide', NULL, 'default', 0, 1, 31, 0, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (40, 0, '正常', '1', 'sys_normal_disable', '', 'default', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (41, 0, '停用', '0', 'sys_normal_disable', '', 'default', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (49, 0, '是', '1', 'sys_yes_no', '', '', 1, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (50, 0, '否', '0', 'sys_yes_no', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (61, 0, '政府工作目标', '1', 'gov_cate_models', '', '', 0, 1, 2, 0, '', NULL, NULL, '2024-06-05 01:34:08');
INSERT INTO `sys_dict_data` VALUES (62, 0, '系统后台', 'sys_admin', 'menu_module_type', '', '', 1, 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (63, 0, '政务工作', 'gov_work', 'menu_module_type', '', '', 0, 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (65, 0, '[work]测试业务表', 'wf_news', 'flow_type', '', '', 0, 1, 2, 2, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (66, 0, '回退修改', '-1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (67, 0, '保存中', '0', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (68, 0, '流程中', '1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (69, 0, '审批通过', '2', 'flow_status', '', '', 0, 1, 31, 2, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (70, 2, '发布栏目', '2', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (71, 3, '跳转栏目', '3', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (72, 4, '单页栏目', '4', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (73, 2, '置顶', '1', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, '2024-06-05 01:34:40');
INSERT INTO `sys_dict_data` VALUES (74, 3, '幻灯', '2', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, '2024-06-05 01:34:40');
INSERT INTO `sys_dict_data` VALUES (75, 4, '推荐', '3', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, '2024-06-05 01:34:40');
INSERT INTO `sys_dict_data` VALUES (76, 1, '一般', '0', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, '2024-06-05 01:34:40');
INSERT INTO `sys_dict_data` VALUES (77, 1, '频道页', '1', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (78, 0, '普通', '0', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20', NULL);
INSERT INTO `sys_dict_data` VALUES (79, 0, '加急', '1', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20', NULL);
INSERT INTO `sys_dict_data` VALUES (80, 0, '紧急', '2', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20', NULL);
INSERT INTO `sys_dict_data` VALUES (81, 0, '特急', '3', 'flow_level', '', '', 0, 1, 31, 31, '', NULL, '2021-07-20 08:55:25', NULL);
INSERT INTO `sys_dict_data` VALUES (82, 0, '频道页', '1', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL, '2024-06-05 01:34:24');
INSERT INTO `sys_dict_data` VALUES (83, 0, '发布栏目', '2', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL, '2024-06-05 01:34:24');
INSERT INTO `sys_dict_data` VALUES (84, 0, '跳转栏目', '3', 'sys_blog_type', '', '', 0, 1, 31, 31, '', NULL, NULL, '2024-06-05 01:34:24');
INSERT INTO `sys_dict_data` VALUES (85, 0, '单页栏目', '4', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL, '2024-06-05 01:34:24');
INSERT INTO `sys_dict_data` VALUES (87, 0, '[cms]文章表', 'cms_news', 'flow_type', '', '', 0, 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (88, 1, '下架', '0', 'is_enable', '', '', 0, 1, 1, 1, '', '2024-05-31 14:22:51', '2024-06-09 03:44:28', NULL);
INSERT INTO `sys_dict_data` VALUES (89, 0, '上架', '1', 'is_enable', '', '', 0, 1, 1, 1, '', '2024-05-31 14:25:06', '2024-06-09 03:44:38', NULL);
INSERT INTO `sys_dict_data` VALUES (90, 0, 'ethereum', '0', 'chain_type', '', '', 1, 1, 1, 0, '', '2024-05-31 15:16:16', '2024-05-31 15:16:16', NULL);
INSERT INTO `sys_dict_data` VALUES (91, 0, 'ethereum', '0', 'chain_type', '', '', 1, 1, 1, 0, '', '2024-05-31 15:16:34', '2024-05-31 15:16:34', '2024-05-31 15:19:13');
INSERT INTO `sys_dict_data` VALUES (92, 2, 'tron', '2', 'chain_type', '', '', 0, 1, 1, 1, '', '2024-05-31 15:17:44', '2024-05-31 15:20:20', NULL);
INSERT INTO `sys_dict_data` VALUES (93, 1, 'tron', '1', 'chain_type', '', '', 0, 1, 1, 0, '', '2024-05-31 15:17:58', '2024-05-31 15:17:58', '2024-05-31 15:19:18');
INSERT INTO `sys_dict_data` VALUES (94, 1, 'btc', '1', 'chain_type', '', '', 0, 1, 1, 0, '', '2024-05-31 15:18:36', '2024-05-31 15:18:36', NULL);
INSERT INTO `sys_dict_data` VALUES (95, 0, 'native', 'native', 'token_type', '', '', 0, 1, 1, 1, '', '2024-05-31 15:25:08', '2024-06-13 02:46:01', NULL);
INSERT INTO `sys_dict_data` VALUES (96, 1, 'erc20', 'erc20', 'token_type', '', '', 0, 1, 1, 1, '', '2024-05-31 15:26:08', '2024-06-13 02:46:08', NULL);
INSERT INTO `sys_dict_data` VALUES (97, 0, '等待自动审核 ', '0', 'bridge_order_status', '', '', 1, 1, 1, 1, '', '2024-05-31 15:29:46', '2024-06-14 19:08:21', '2024-06-14 19:29:17');
INSERT INTO `sys_dict_data` VALUES (98, 3, '审核失败', '7', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-05-31 15:32:08', '2024-06-14 19:42:06', '2024-06-14 19:51:44');
INSERT INTO `sys_dict_data` VALUES (99, 2, '审核通过', '3', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-05-31 15:33:28', '2024-06-14 19:42:06', '2024-06-14 19:51:14');
INSERT INTO `sys_dict_data` VALUES (100, 3, '付款处理中', '3', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-05-31 15:35:38', '2024-06-14 19:08:21', '2024-06-14 19:29:14');
INSERT INTO `sys_dict_data` VALUES (101, 4, '上链成功', '4', 'bridge_order_status', '', '', 0, 1, 1, 0, '', '2024-06-05 23:30:55', '2024-06-14 19:08:21', '2024-06-14 19:29:08');
INSERT INTO `sys_dict_data` VALUES (102, 5, '付款成功 ', '5', 'bridge_order_status', '', '', 0, 1, 1, 0, '', '2024-06-05 23:31:18', '2024-06-14 19:08:21', '2024-06-14 19:29:06');
INSERT INTO `sys_dict_data` VALUES (103, 6, '付款失败', '6', 'bridge_order_status', '', '', 0, 1, 1, 0, '', '2024-06-05 23:31:35', '2024-06-14 19:08:21', '2024-06-14 19:29:04');
INSERT INTO `sys_dict_data` VALUES (104, 0, '未删除', '0', 'isDelete', '', '', 1, 1, 1, 0, '', '2024-06-12 23:49:43', '2024-06-12 23:49:43', NULL);
INSERT INTO `sys_dict_data` VALUES (105, 1, '已删除', '1', 'isDelete', '', '', 0, 1, 1, 0, '已删除', '2024-06-12 23:50:20', '2024-06-12 23:50:20', NULL);
INSERT INTO `sys_dict_data` VALUES (106, 0, '0', 'ETH', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:00:59', '2024-06-13 02:00:59', NULL);
INSERT INTO `sys_dict_data` VALUES (107, 0, '1', 'tBNB', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:04:59', '2024-06-13 02:04:59', NULL);
INSERT INTO `sys_dict_data` VALUES (108, 0, '2', 'KAD', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:05:16', '2024-06-13 02:05:16', NULL);
INSERT INTO `sys_dict_data` VALUES (109, 0, '3', 'BSC test token', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:05:32', '2024-06-13 02:05:32', NULL);
INSERT INTO `sys_dict_data` VALUES (110, 0, '4', 'KAD test token', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:05:48', '2024-06-13 02:05:48', NULL);
INSERT INTO `sys_dict_data` VALUES (111, 0, '5', 'arbETH', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:06:12', '2024-06-13 02:06:12', NULL);
INSERT INTO `sys_dict_data` VALUES (112, 0, '6', 'arb test token', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:06:26', '2024-06-13 02:06:26', NULL);
INSERT INTO `sys_dict_data` VALUES (113, 0, '8', 'Sepolia eth', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:06:42', '2024-06-13 02:06:42', NULL);
INSERT INTO `sys_dict_data` VALUES (114, 0, '9', 'USDT', 'coinName', '', '', 0, 1, 1, 0, '', '2024-06-13 02:07:04', '2024-06-13 02:07:04', NULL);
INSERT INTO `sys_dict_data` VALUES (115, 0, '待审核', '0', 'bridge_order_status', '', '', 1, 1, 1, 1, '', '2024-06-14 19:44:18', '2024-06-15 15:00:07', '2024-06-15 15:15:49');
INSERT INTO `sys_dict_data` VALUES (116, 0, '待审核', '1', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-06-14 19:44:46', '2024-06-15 15:02:12', NULL);
INSERT INTO `sys_dict_data` VALUES (117, 0, '进行中', '2', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-06-14 19:48:35', '2024-06-15 15:04:30', '2024-06-15 15:16:01');
INSERT INTO `sys_dict_data` VALUES (118, 0, '进行中', '3', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-06-14 19:49:09', '2024-06-15 15:04:11', '2024-06-15 15:16:11');
INSERT INTO `sys_dict_data` VALUES (119, 0, '进行中', '4', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-06-14 19:49:27', '2024-06-15 15:03:58', NULL);
INSERT INTO `sys_dict_data` VALUES (120, 0, '已完成', '5', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-06-14 19:49:43', '2024-06-15 15:03:36', NULL);
INSERT INTO `sys_dict_data` VALUES (121, 0, '跨链失败', '6', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-06-14 19:49:57', '2024-06-15 15:03:03', '2024-06-15 15:16:18');
INSERT INTO `sys_dict_data` VALUES (122, 0, '跨链失败', '7', 'bridge_order_status', '', '', 0, 1, 1, 1, '', '2024-06-14 19:50:17', '2024-06-15 15:03:15', NULL);
INSERT INTO `sys_dict_data` VALUES (123, 0, '人工审核通过', '2', 'manual_review_status', '', '', 1, 1, 1, 0, '', '2024-06-14 19:57:21', '2024-06-14 19:57:21', NULL);
INSERT INTO `sys_dict_data` VALUES (124, 0, '人工审核失败', '7', 'manual_review_status', '', '', 0, 1, 1, 0, '', '2024-06-14 19:57:37', '2024-06-14 19:57:37', NULL);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint UNSIGNED NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` int UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '字典类型表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', 1, 31, 1, '用于选择用户性别', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (3, '任务状态', 'sys_job_status', 1, 31, 31, '任务状态列表', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (13, '任务分组', 'sys_job_group', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (14, '管理员登录状态', 'admin_login_status', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (15, '操作日志状态', 'sys_oper_log_status', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (16, '任务策略', 'sys_job_policy', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (17, '菜单状态', 'sys_show_hide', 1, 31, 0, '菜单状态', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (18, '系统开关', 'sys_normal_disable', 1, 31, 31, '系统开关', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (24, '系统内置', 'sys_yes_no', 1, 31, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (29, '政务工作模型分类', 'gov_cate_models', 1, 2, 0, '', NULL, NULL, '2024-06-05 01:34:08');
INSERT INTO `sys_dict_type` VALUES (30, '菜单模块类型', 'menu_module_type', 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (31, '工作流程类型', 'flow_type', 1, 2, 0, '', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (32, '工作流程审批状态', 'flow_status', 1, 31, 0, '工作流程审批状态', NULL, NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (33, '博客分类类型', 'sys_blog_type', 1, 31, 31, '博客分类中的标志', NULL, NULL, '2024-06-05 01:34:24');
INSERT INTO `sys_dict_type` VALUES (34, '博客日志标志', 'sys_log_sign', 1, 31, 0, '博客日志管理中的标志数据字典', NULL, NULL, '2024-06-05 01:34:39');
INSERT INTO `sys_dict_type` VALUES (35, '工作流紧急状态', 'flow_level', 1, 31, 31, '', NULL, '2021-07-20 08:55:20', NULL);
INSERT INTO `sys_dict_type` VALUES (36, '上架开关', 'is_enable', 1, 1, 1, '', '2024-05-30 20:00:41', '2024-05-31 15:41:13', NULL);
INSERT INTO `sys_dict_type` VALUES (37, '公链类型', 'chain_type', 1, 1, 0, '', '2024-05-31 14:58:28', '2024-05-31 14:58:28', NULL);
INSERT INTO `sys_dict_type` VALUES (38, '币种类型', 'token_type', 1, 1, 0, '', '2024-05-31 15:02:05', '2024-05-31 15:02:05', NULL);
INSERT INTO `sys_dict_type` VALUES (39, '跨链记录状态', 'bridge_order_status', 1, 1, 1, '', '2024-05-31 15:09:19', '2024-06-14 19:42:05', NULL);
INSERT INTO `sys_dict_type` VALUES (40, '软删除', 'isDelete', 1, 1, 0, '软删除状态', '2024-06-12 23:48:52', '2024-06-12 23:48:52', NULL);
INSERT INTO `sys_dict_type` VALUES (41, '链名称', 'chainName', 1, 1, 0, '链名称', '2024-06-13 00:03:00', '2024-06-13 00:03:00', NULL);
INSERT INTO `sys_dict_type` VALUES (42, '币种名', 'coinName', 1, 1, 0, '', '2024-06-13 02:00:13', '2024-06-13 02:00:13', NULL);
INSERT INTO `sys_dict_type` VALUES (43, '人工审核状态', 'manual_review_status', 1, 1, 0, '', '2024-06-14 19:55:40', '2024-06-14 19:55:40', NULL);
INSERT INTO `sys_dict_type` VALUES (44, '审核通过', '2', 1, 1, 0, '', '2024-06-14 19:56:35', '2024-06-14 19:56:35', '2024-06-14 19:56:45');

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` tinyint NULL DEFAULT 1 COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` tinyint NULL DEFAULT 1 COMMENT '是否并发执行（0允许 1禁止）',
  `status` tinyint NULL DEFAULT 0 COMMENT '状态（0正常 1暂停）',
  `create_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '备注信息',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`job_id`, `job_name`, `job_group`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '定时任务调度表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
INSERT INTO `sys_job` VALUES (1, '测试任务1', '', 'DEFAULT', 'test1', '* * * * * ?', 1, 0, 1, 1, 31, '', NULL, '2021-07-16 16:01:59', NULL);
INSERT INTO `sys_job` VALUES (2, '测试任务2', 'hello|gfast', 'DEFAULT', 'test2', '* * * * * ?', 1, 0, 1, 1, 31, '备注', NULL, '2021-07-16 17:15:09', NULL);
INSERT INTO `sys_job` VALUES (6, '测试任务3', 'hello|gfast', 'DEFAULT', 'test2', '* * * * * *', 1, 0, 1, 1, 31, '备注', NULL, NULL, '2021-07-16 16:51:20');
INSERT INTO `sys_job` VALUES (8, '在线用户定时更新', '', 'DEFAULT', 'checkUserOnline', '5 */10 * * * ?', 1, 0, 0, 2, 1, '', NULL, '2021-07-19 08:57:24', NULL);
INSERT INTO `sys_job` VALUES (9, '测试', 'arg1', 'DEFAULT', 'demo1', ' */5 * * * * ? ', 1, 1, 1, 31, 31, '', '2021-07-15 15:00:15', '2021-07-16 16:47:56', '2021-07-16 16:48:00');
INSERT INTO `sys_job` VALUES (10, '测试', 'arg1', 'DEFAULT', 'demo1', ' */5 * * * * ? ', 2, 1, 1, 31, 31, '', '2021-07-15 22:12:15', '2021-07-16 16:47:52', '2021-07-16 16:48:00');
INSERT INTO `sys_job` VALUES (11, '测试', 'arg1', 'DEFAULT', 'demo1', ' */5 * * * * ? ', 2, 1, 1, 31, 31, '', '2021-07-16 08:49:43', '2021-07-16 08:59:32', '2021-07-16 16:46:51');
INSERT INTO `sys_job` VALUES (12, '测试366', 'arg1', 'DEFAULT', 'demo1', ' */5 * * * * ? ', 2, 1, 1, 31, 31, '', '2021-07-16 08:50:10', '2021-07-16 08:59:26', '2021-07-16 16:46:51');

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` bigint NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` tinyint NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime NULL DEFAULT '0000-00-00 00:00:00' COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 86 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统访问记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
INSERT INTO `sys_login_log` VALUES (1, 'admin', '79.127.205.163', ' ', 'Chrome', 'Windows 10', 1, '登录成功', '2024-06-16 01:44:53', '系统后台');
INSERT INTO `sys_login_log` VALUES (2, 'amin', '149.102.253.233', ' ', 'Chrome', 'Windows 10', 0, '获取用户信息失败', '2024-06-16 01:48:24', '系统后台');
INSERT INTO `sys_login_log` VALUES (84, 'Jimmy', '[::1]', '内网IP', 'Chrome', 'Windows 10', 1, '登录成功', '2024-06-23 21:49:09', '系统后台');
INSERT INTO `sys_login_log` VALUES (85, 'admin', '89.187.163.233', ' ', 'Chrome', 'Windows 10', 1, '登录成功', '2024-06-23 14:07:40', '系统后台');

-- ----------------------------
-- Table structure for sys_model_info
-- ----------------------------
DROP TABLE IF EXISTS `sys_model_info`;
CREATE TABLE `sys_model_info`  (
  `model_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '模型ID',
  `model_category_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '模板分类id',
  `model_name` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '模型标识',
  `model_title` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '模型名称',
  `model_pk` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '主键字段',
  `model_order` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '默认排序字段',
  `model_sort` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '表单字段排序',
  `model_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '列表显示字段，为空显示全部',
  `model_edit` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '可编辑字段，为空则除主键外均可以编辑',
  `model_indexes` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '索引字段',
  `search_list` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '高级搜索的字段',
  `create_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `model_status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态',
  `model_engine` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'MyISAM' COMMENT '数据库引擎',
  `create_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `update_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  PRIMARY KEY (`model_id`) USING BTREE,
  UNIQUE INDEX `name_uni`(`model_name` ASC) USING BTREE COMMENT '模型名唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '文档模型表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_model_info
-- ----------------------------
INSERT INTO `sys_model_info` VALUES (1, 5, 'house_info', '房屋信息', 'house_id', 'hi_order', 'house_id', 'house_id,belong_houses,belong_house_type,door_number', 'belong_houses,belong_house_type,door_number', 'door_number', 'belong_houses,belong_house_type,door_number', 1498101800, 1605838046, 1, 'MyISAM', 0, 1);
INSERT INTO `sys_model_info` VALUES (9, 6, 'cms_news_attr', '文章模型', 'art_id', '', 'art_id', 'news_info,text_demo', 'news_content,thumbnail,news_files,news_file,news_imgs,news_title,news_info,text_demo,time_demo', '', 'news_info,text_demo', 1595316217, 1596095352, 1, 'InnoDB', 1, 31);
INSERT INTO `sys_model_info` VALUES (10, 1, 'gov_work_report', '政府工作报告', 'matter_id', 'matter_id', '', 'title,content,sponsor,organizer', 'title,content,sponsor,organizer', 'matter_id', 'title', 1596528880, 1599701551, 1, 'InnoDB', 2, 31);
INSERT INTO `sys_model_info` VALUES (12, 1, 'gov_work_major_project', '重大项目', 'matter_id', 'matter_id', '', 'project_no,title,content,sponsor,total_investment,annual_investment,progress_status', 'project_no,title,content,sponsor,total_investment,annual_investment,progress_status', 'matter_id', 'title', 1599702575, 1605512186, 1, 'InnoDB', 31, 1);
INSERT INTO `sys_model_info` VALUES (13, 1, 'gov_work_important_documents', '重要文件办理', 'matter_id', 'matter_id', '', 'title,original_unit,document_number,receive_time,sponsor,content', 'title,original_unit,document_number,receive_time,sponsor,content', 'matter_id', 'title', 1599710390, 1599710390, 1, 'InnoDB', 31, 0);
INSERT INTO `sys_model_info` VALUES (14, 1, 'gov_work_livelihood', '民生目标', 'matter_id', 'matter_id', '', 'title,type,content,sponsor,organizer', 'title,type,content,sponsor,organizer', 'matter_id', 'title', 1599721999, 1599721999, 1, 'InnoDB', 31, 0);
INSERT INTO `sys_model_info` VALUES (15, 1, 'gov_work_meeting', '会议议定', 'matter_id', 'matter_id', '', 'title,content,sponsor,organizer', 'title,content,sponsor,organizer', 'matter_id', 'title', 1599722042, 1605512125, 1, 'InnoDB', 31, 1);
INSERT INTO `sys_model_info` VALUES (16, 1, 'gov_work_assigned_leader', '领导交办', 'matter_id', 'matter_id', '', 'title,number,content,sponsor,organizer,instruction_time,instruction_leader', 'title,number,content,sponsor,organizer,instruction_time,instruction_leader', 'matter_id', 'title', 1599722097, 1599722097, 1, 'InnoDB', 31, 0);
INSERT INTO `sys_model_info` VALUES (17, 1, 'gov_work_other', '其它任务', 'matter_id', 'matter_id', '', 'title,content,sponsor,organizer', 'title,content,sponsor,organizer', 'matter_id', 'title', 1599722127, 1599722127, 1, 'InnoDB', 31, 0);
INSERT INTO `sys_model_info` VALUES (18, 1, 'gov_work_investment', '招商引资', 'matter_id', 'matter_id', '', 'project_type,title,company_name,investment_volume,paid_in_investment,content,address,project_progress,project_target,manage_leader,responsible_department,sponsor,organizer,remark', 'project_type,title,company_name,investment_volume,paid_in_investment,content,address,project_progress,project_target,manage_leader,responsible_department,sponsor,organizer,remark', '', 'title', 1611542517, 1611542540, 1, 'InnoDB', 1, 1);

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '请求参数',
  `json_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '返回参数',
  `status` int NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 10449 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '操作日志记录' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
INSERT INTO `sys_oper_log` VALUES (1030, '', 0, '/system/user/getInfo', 'GET', 1, 'admin', '深圳总公司', '/system/user/getInfo', '[::1]', '内网IP', '{}', '{\"code\":-1,\"msg\":\"open ./config/casbin_conf/rbac_model.conf: The system cannot find the path specified.\"}', 0, '', '2024-05-30 17:48:31');
INSERT INTO `sys_oper_log` VALUES (10448, '', 0, '/system/auth/unBind', 'POST', 1, '', '', '/system/auth/unBind', '171.22.78.221', ' ', '{\"userName\":\"admin\"}', '{\"code\":0,\"msg\":\"success\"}', 1, '', '2024-06-24 09:15:12');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int NOT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `created_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '岗位信息表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES (1, 'ceo', '董事长', 1, '1', '', 0, 0, '2021-07-11 11:32:58', NULL, NULL);
INSERT INTO `sys_post` VALUES (2, 'se', '项目经理', 2, '1', '', 0, 0, '2021-07-12 11:01:26', NULL, NULL);
INSERT INTO `sys_post` VALUES (3, 'hr', '人力资源', 3, '1', '', 0, 0, '2021-07-12 11:01:30', NULL, NULL);
INSERT INTO `sys_post` VALUES (4, 'user', '普通员工', 4, '1', '', 0, 0, '2021-07-12 11:01:33', NULL, NULL);
INSERT INTO `sys_post` VALUES (5, 'it', 'IT部', 5, '1', '信息部', 31, 31, '2021-07-12 11:09:42', '2021-07-12 11:21:00', NULL);
INSERT INTO `sys_post` VALUES (6, 'asd', 'aaa', 100, '0', '', 31, 0, '2021-07-12 11:21:12', '2021-07-12 11:21:12', '2021-07-12 11:24:49');
INSERT INTO `sys_post` VALUES (7, 'asda', 'ddd', 222, '0', 'asdas', 31, 31, '2021-07-12 11:21:18', '2021-07-12 11:21:27', '2021-07-12 11:24:49');
INSERT INTO `sys_post` VALUES (8, 'dasd', 'dasd', 3444, '0', 'asdas', 31, 31, '2021-07-12 11:21:22', '2021-07-12 11:21:31', '2021-07-12 11:24:45');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;0:禁用;1:正常',
  `list_order` float NOT NULL DEFAULT 0 COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint UNSIGNED NOT NULL DEFAULT 3 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 1, 0, '超级管理员', '备注', 3);
INSERT INTO `sys_role` VALUES (3, 1, 0, '站点管理员', '站点管理人员', 3);
INSERT INTO `sys_role` VALUES (4, 1, 0, '初级管理员', '初级管理员', 3);
INSERT INTO `sys_role` VALUES (5, 1, 0, '高级管理员', '高级管理员', 2);
INSERT INTO `sys_role` VALUES (9, 1, 5, '测试人员', '', 3);
INSERT INTO `sys_role` VALUES (10, 1, 6, '开发人员', '', 1);

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `dept_id` bigint NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色和部门关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
INSERT INTO `sys_role_dept` VALUES (5, 103);
INSERT INTO `sys_role_dept` VALUES (5, 104);
INSERT INTO `sys_role_dept` VALUES (5, 105);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int NOT NULL DEFAULT 0 COMMENT '生日',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `googleauth` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'google验证码',
  `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint NOT NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint NOT NULL DEFAULT 1 COMMENT '是否后台管理员 1 是  0   否',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `phone_num` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系电话',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_login`(`user_name` ASC) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile` ASC) USING BTREE,
  INDEX `user_nickname`(`user_nickname` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '13578342363', '超级管理员', 0, '542a6e44dbac171f260fc4a032cd5522', '', 'dlqVVBTADg', 1, 'yxh669@qq.com', 1, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 100, '', 1, 'asdasfdsaf大发放打发士大夫发按时', '描述信息', '18611111111', '89.187.163.233', '2024-06-23 14:07:39', '2021-06-22 17:58:00', '2024-06-24 09:15:11', NULL);
INSERT INTO `sys_user` VALUES (42, 'kk', '18888888888', 'KK', 0, '69d0e24453d13cf487c18f47fd6d703a', '', 'LhNox462he', 1, '88@qq.com', 1, '', 200, '', 1, '', '', '', '79.127.205.163', '2024-06-10 18:17:31', '2024-06-01 14:32:29', '2024-06-10 18:24:18', NULL);
INSERT INTO `sys_user` VALUES (43, 'testa', '13888888888', '测试a', 0, 'b86860028dccd78838b5385c4017943b', '', 'XxqSg6Bgfh', 1, 'test@qq.com', 1, '', 201, '', 1, '', '', '', '120.229.82.196', '2024-06-10 18:10:19', '2024-06-10 18:08:13', '2024-06-10 18:08:13', NULL);
INSERT INTO `sys_user` VALUES (44, 'Jimmy', '18378266988', 'Jimmy', 0, 'ada30cfdf98da656a48bc22eca8f1a98', '', 'f4GeHLhslw', 1, '8888@qq.com', 0, '', 200, '', 1, '', '', '', '[::1]', '2024-06-23 21:49:08', '2024-06-19 17:39:46', '2024-06-23 21:48:07', NULL);

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `create_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 17122 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户在线状态表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `post_id` bigint NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES (1, 2);
INSERT INTO `sys_user_post` VALUES (1, 3);
INSERT INTO `sys_user_post` VALUES (2, 1);
INSERT INTO `sys_user_post` VALUES (2, 2);
INSERT INTO `sys_user_post` VALUES (3, 2);
INSERT INTO `sys_user_post` VALUES (5, 2);
INSERT INTO `sys_user_post` VALUES (14, 1);
INSERT INTO `sys_user_post` VALUES (15, 4);
INSERT INTO `sys_user_post` VALUES (16, 2);
INSERT INTO `sys_user_post` VALUES (22, 1);
INSERT INTO `sys_user_post` VALUES (22, 2);
INSERT INTO `sys_user_post` VALUES (31, 2);
INSERT INTO `sys_user_post` VALUES (34, 1);
INSERT INTO `sys_user_post` VALUES (35, 2);
INSERT INTO `sys_user_post` VALUES (35, 3);
INSERT INTO `sys_user_post` VALUES (36, 1);
INSERT INTO `sys_user_post` VALUES (37, 3);
INSERT INTO `sys_user_post` VALUES (38, 2);
INSERT INTO `sys_user_post` VALUES (38, 3);
INSERT INTO `sys_user_post` VALUES (42, 4);
INSERT INTO `sys_user_post` VALUES (42, 5);
INSERT INTO `sys_user_post` VALUES (43, 4);

-- ----------------------------
-- Table structure for sys_web_set
-- ----------------------------
DROP TABLE IF EXISTS `sys_web_set`;
CREATE TABLE `sys_web_set`  (
  `web_id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `web_content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '站点信息',
  PRIMARY KEY (`web_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_web_set
-- ----------------------------
INSERT INTO `sys_web_set` VALUES (1, '{\"CopyrightInfo\":\"11122\",\"recordInfo\":\"111222\",\"statisticsCode\":\"11122\",\"webId\":1,\"webLogo\":\"https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-13/ccroz2q3sptczqwchk.jpg\",\"webName\":\"gfast\",\"webSite\":\"http://localhost/index#/webSet\"}');

-- ----------------------------
-- Table structure for tools_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table`;
CREATE TABLE `tools_gen_table`  (
  `table_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '表描述',
  `class_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '实体类名称',
  `tpl_category` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `package_name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成包路径',
  `module_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成模块名',
  `business_name` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成业务名',
  `function_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成功能名',
  `function_author` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '生成功能作者',
  `options` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '其它生成选项',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 87 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of tools_gen_table
-- ----------------------------
INSERT INTO `tools_gen_table` VALUES (65, 'demo_gen_class', '代码生成关联测试表', 'DemoGenClass', 'crud', 'gfast/app/system', 'system', 'demo_gen_class', '代码生成关联测试', 'gfast', '', '2021-07-26 11:07:30', '2021-08-16 15:53:39', '');
INSERT INTO `tools_gen_table` VALUES (68, 'demo_gen_tree', '代码生成树形结构测试表', 'DemoGenTree', 'tree', 'gfast/app/system', 'system', 'demo_gen_tree', '代码生成树形结构测试', 'gfast', '{\"treeCode\":\"id\",\"treeName\":\"demoName\",\"treeParentCode\":\"parentId\"}', '2021-08-02 11:55:43', '2021-08-16 15:53:47', '');
INSERT INTO `tools_gen_table` VALUES (69, 'demo_gen', '代码生成测试表', 'DemoGen', 'crud', 'gfast/app/system', 'system', 'demo_gen', '代码生成测试', 'gfast', '', '2021-08-03 18:14:06', '2021-08-16 15:53:56', '');
INSERT INTO `tools_gen_table` VALUES (70, 'bridge_config', '跨链币对', 'BridgeConfig', 'crud', 'gfast/app/admin', 'admin', 'bridge_config', '可跨币对', 'jimmy', '', '2024-05-30 18:31:01', '2024-05-31 19:51:46', '');
INSERT INTO `tools_gen_table` VALUES (71, 'bridge_contract_address', '', 'BridgeContractAddress', 'crud', 'gfast/app/system', 'system', 'bridge_contract_address', '', 'gfast', '', '2024-05-30 18:31:04', '2024-05-30 18:31:04', '');
INSERT INTO `tools_gen_table` VALUES (83, 'bridge_order', '跨链记录', 'BridgeOrder', 'crud', 'gfast/app/admin', 'admin', 'bridge_order', '跨链记录', 'jimmy', '', '2024-06-03 19:16:53', '2024-06-06 15:59:36', '');
INSERT INTO `tools_gen_table` VALUES (84, 'pay_detail', '支付详情', 'PayDetail', 'crud', 'gfast/app/admin', 'admin', 'pay_detail', '支付详情', 'jimmy', '', '2024-06-08 23:23:41', '2024-06-08 23:28:43', '');
INSERT INTO `tools_gen_table` VALUES (85, 'chain', '链管理', 'Chain', 'crud', 'gfast/app/admin', 'admin', 'chain', '链管理', 'jimmy', '', '2024-06-12 23:29:47', '2024-06-13 00:03:41', '');
INSERT INTO `tools_gen_table` VALUES (86, 'daily_bridge_stats', '跨链日统计', 'DailyBridgeStats', 'crud', 'gfast/app/admin', 'admin', 'daily_bridge_stats', '跨链日统计', 'jimmy', '', '2024-06-14 17:06:57', '2024-06-14 17:16:50', '');

-- ----------------------------
-- Table structure for tools_gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table_column`;
CREATE TABLE `tools_gen_table_column`  (
  `column_id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint NULL DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '列类型',
  `go_type` varchar(500) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'Go类型',
  `go_field` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'Go字段名',
  `html_field` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'html字段名',
  `is_pk` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `sort` int NULL DEFAULT NULL COMMENT '排序',
  `link_table_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表名',
  `link_table_class` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表类名',
  `link_table_package` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表包名',
  `link_label_id` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表键名',
  `link_label_name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '关联表字段值',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 923 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of tools_gen_table_column
-- ----------------------------
INSERT INTO `tools_gen_table_column` VALUES (666, 65, 'id', '分类id', 'int(10) unsigned', 'uint', 'Id', 'id', '1', '1', '0', '0', '0', '1', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (667, 65, 'class_name', '分类名', 'varchar(30)', 'string', 'ClassName', 'className', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (693, 68, 'id', 'ID', 'int(11) unsigned', 'uint', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (694, 68, 'parent_id', '父级ID', 'int(10) unsigned', 'uint', 'ParentId', 'parentId', '0', '0', '1', '1', '1', '0', '0', 'EQ', 'select', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (695, 68, 'demo_name', '姓名', 'varchar(20)', 'string', 'DemoName', 'demoName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (696, 68, 'demo_age', '年龄', 'int(10) unsigned', 'uint', 'DemoAge', 'demoAge', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (697, 68, 'classes', '班级', 'varchar(30)', 'string', 'Classes', 'classes', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', '', 5, 'demo_gen_class', 'DemoGenClass', 'gfast/app/system', 'id', 'className');
INSERT INTO `tools_gen_table_column` VALUES (698, 68, 'demo_born', '出生年月', 'datetime', 'Time', 'DemoBorn', 'demoBorn', '0', '0', '', '1', '1', '1', '', 'EQ', 'datetime', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (699, 68, 'demo_gender', '性别', 'tinyint(3) unsigned', 'uint', 'DemoGender', 'demoGender', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_user_sex', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (700, 68, 'created_at', '创建日期', 'datetime', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '1', '1', 'EQ', 'datetime', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (701, 68, 'updated_at', '修改日期', 'datetime', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (702, 68, 'deleted_at', '删除日期', 'datetime', 'Time', 'DeletedAt', 'deletedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (703, 68, 'created_by', '创建人', 'bigint(20) unsigned', 'uint64', 'CreatedBy', 'createdBy', '0', '0', '0', '0', '0', '1', '0', 'EQ', 'input', '', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (704, 68, 'updated_by', '修改人', 'bigint(20) unsigned', 'uint64', 'UpdatedBy', 'updatedBy', '0', '0', '0', '0', '0', '1', '0', 'EQ', 'input', '', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (705, 68, 'demo_status', '状态', 'tinyint(4)', 'int', 'DemoStatus', 'demoStatus', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', 'sys_normal_disable', 13, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (706, 68, 'demo_cate', '分类', 'varchar(30)', 'string', 'DemoCate', 'demoCate', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'checkbox', 'cms_news_type', 14, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (707, 69, 'id', '', 'int(11) unsigned', 'uint', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (708, 69, 'demo_name', '姓名', 'varchar(20)', 'string', 'DemoName', 'demoName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (709, 69, 'demo_age', '年龄', 'int(10) unsigned', 'uint', 'DemoAge', 'demoAge', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (710, 69, 'classes', '班级', 'varchar(30)', 'string', 'Classes', 'classes', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', '', 4, 'demo_gen_class', 'DemoGenClass', 'gfast/app/system', 'id', 'className');
INSERT INTO `tools_gen_table_column` VALUES (711, 69, 'demo_born', '出生年月', 'datetime', 'Time', 'DemoBorn', 'demoBorn', '0', '0', '', '1', '1', '1', '', 'EQ', 'datetime', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (712, 69, 'demo_gender', '性别', 'tinyint(3) unsigned', 'uint', 'DemoGender', 'demoGender', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_user_sex', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (713, 69, 'created_at', '创建日期', 'datetime', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (714, 69, 'updated_at', '修改日期', 'datetime', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (715, 69, 'deleted_at', '删除日期', 'datetime', 'Time', 'DeletedAt', 'deletedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (716, 69, 'created_by', '创建人', 'bigint(20) unsigned', 'uint64', 'CreatedBy', 'createdBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (717, 69, 'updated_by', '修改人', 'bigint(20) unsigned', 'uint64', 'UpdatedBy', 'updatedBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (718, 69, 'demo_status', '状态', 'tinyint(4)', 'int', 'DemoStatus', 'demoStatus', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', 'sys_normal_disable', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (719, 69, 'demo_cate', '分类', 'varchar(30)', 'string', 'DemoCate', 'demoCate', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'checkbox', 'cms_news_type', 13, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (720, 70, 'id', '', 'bigint', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (721, 70, 'source_chainId', '当前链', 'varchar(30)', 'string', 'SourceChainId', 'sourceChainId', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', '', 2, 'chain', 'Chain', 'gfast/app/admin', 'chainId', 'name');
INSERT INTO `tools_gen_table_column` VALUES (722, 70, 'target_chainId', '对手链', 'varchar(30)', 'string', 'TargetChainId', 'targetChainId', '0', '0', '1', '1', '1', '1', '', 'EQ', 'select', '', 3, 'chain', 'Chain', 'gfast/app/admin', 'chainId', 'name');
INSERT INTO `tools_gen_table_column` VALUES (723, 70, 'source_coin_address', '币种', 'varchar(45)', 'string', 'SourceCoinAddress', 'sourceCoinAddress', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', '', 4, 'coin', 'Coin', 'gfast/app/admin', 'address', 'symbol');
INSERT INTO `tools_gen_table_column` VALUES (724, 70, 'fee_fixed', '固定金额手续费，', 'double', 'float64', 'FeeFixed', 'feeFixed', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (725, 70, 'fee_percent', '手续费百分比，如果不需要则设置为0', 'double', 'float64', 'FeePercent', 'feePercent', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (726, 70, 'day_total', '跨入每日审核数量', 'decimal(20,10)', 'float64', 'DayTotal', 'dayTotal', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (727, 70, 'once_total', '跨入单次审核数量', 'decimal(20,10)', 'float64', 'OnceTotal', 'onceTotal', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (728, 70, 'isEnable', '状态', 'tinyint(1)', 'int', 'IsEnable', 'isEnable', '0', '0', '1', '1', '1', '1', '', 'EQ', 'radio', 'is_enable', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (729, 70, 'target_coin_address', '目标链合约地址', 'varchar(45)', 'string', 'TargetCoinAddress', 'targetCoinAddress', '0', '0', '1', '1', '1', '', '', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (730, 70, 'update_at', '更新时间', 'datetime', 'Time', 'UpdateAt', 'updateAt', '0', '0', '', '', '', '', '', 'EQ', 'datetime', '', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (731, 70, 'create_at', '创建时间', 'datetime', 'Time', 'CreateAt', 'createAt', '0', '0', '', '', '', '', '', 'EQ', 'datetime', '', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (732, 71, 'id', '', 'bigint', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (733, 71, 'chainId', '', 'varchar(30)', 'string', 'ChainId', 'chainId', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (734, 71, 'contract_address', '', 'varchar(45)', 'string', 'ContractAddress', 'contractAddress', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (735, 71, 'create_at', '', 'datetime', 'Time', 'CreateAt', 'createAt', '0', '0', '', '1', '1', '1', '1', 'EQ', 'datetime', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (736, 71, 'update_at', '', 'datetime', 'Time', 'UpdateAt', 'updateAt', '0', '0', '', '1', '1', '1', '1', 'EQ', 'datetime', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (869, 83, 'id', '', 'bigint', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (870, 83, 'source_address', '转出币种', 'varchar(45)', 'string', 'SourceAddress', 'sourceAddress', '0', '0', '1', '1', '', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (871, 83, 'target_address', '转入币种', 'varchar(45)', 'string', 'TargetAddress', 'targetAddress', '0', '0', '1', '1', '', '1', '1', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (872, 83, 'source_coin_address', '转出钱包地址', 'varchar(45)', 'string', 'SourceCoinAddress', 'sourceCoinAddress', '0', '0', '1', '1', '', '1', '1', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (873, 83, 'target_coin_address', '转入钱包地址', 'varchar(45)', 'string', 'TargetCoinAddress', 'targetCoinAddress', '0', '0', '1', '1', '', '1', '1', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (874, 83, 'source_chainId', '转出链', 'varchar(30)', 'string', 'SourceChainId', 'sourceChainId', '0', '0', '1', '1', '', '1', '1', 'EQ', 'select', '', 6, 'chain', 'Chain', 'gfast/app/admin', 'chainId', 'name');
INSERT INTO `tools_gen_table_column` VALUES (875, 83, 'target_chainId', '转入链', 'varchar(30)', 'string', 'TargetChainId', 'targetChainId', '0', '0', '1', '1', '', '1', '1', 'EQ', 'select', '', 7, 'chain', 'Chain', 'gfast/app/admin', 'chainId', 'name');
INSERT INTO `tools_gen_table_column` VALUES (876, 83, 'transaction_hash', '交易哈希', 'varchar(70)', 'string', 'TransactionHash', 'transactionHash', '0', '0', '1', '1', '', '1', '1', 'EQ', 'input', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (877, 83, 'orderId', '跨链订单id', 'varchar(70)', 'string', 'OrderId', 'orderId', '0', '0', '1', '1', '', '1', '', 'EQ', 'input', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (878, 83, 'amount', '数量', 'decimal(40,20)', 'float64', 'Amount', 'amount', '0', '0', '1', '1', '', '1', '1', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (879, 83, 'status', '跨链记录状态', 'int', 'int', 'Status', 'status', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', 'bridge_order_status', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (880, 83, 'create_at', '发起时间', 'datetime', 'Time', 'CreateAt', 'createAt', '0', '0', '', '1', '', '1', '1', 'EQ', 'datetime', '', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (881, 83, 'update_at', '结束时间', 'datetime', 'Time', 'UpdateAt', 'updateAt', '0', '0', '', '1', '', '1', '1', 'EQ', 'datetime', '', 13, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (882, 83, 'block_number', '区块高度', 'bigint', 'int64', 'BlockNumber', 'blockNumber', '0', '0', '1', '1', '', '', '', 'EQ', 'input', '', 14, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (883, 83, 'fee', '手续费', 'decimal(40,20)', 'float64', 'Fee', 'fee', '0', '0', '1', '1', '', '1', '1', 'EQ', 'input', '', 15, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (884, 83, 'remark', '评论', 'blob', 'string', 'Remark', 'remark', '0', '0', '', '1', '', '', '0', 'EQ', '', '', 16, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (885, 84, 'id', '', 'bigint', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (886, 84, 'order_id', '', 'varchar(70)', 'string', 'OrderId', 'orderId', '0', '0', '1', '1', '', '1', '1', 'EQ', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (887, 84, 'pay_chainId', '', 'varchar(30)', 'string', 'PayChainId', 'payChainId', '0', '0', '1', '1', '', '1', '', 'EQ', 'select', '', 3, 'chain', 'Chain', 'gfast/app/admin', 'chainId', 'name');
INSERT INTO `tools_gen_table_column` VALUES (888, 84, 'pay_coin_address', '', 'varchar(45)', 'string', 'PayCoinAddress', 'payCoinAddress', '0', '0', '1', '1', '', '1', '', 'EQ', 'select', '', 4, 'coin', 'Coin', 'gfast/app/admin', 'address', 'name');
INSERT INTO `tools_gen_table_column` VALUES (889, 84, 'transaction_hash', '', 'varchar(70)', 'string', 'TransactionHash', 'transactionHash', '0', '0', '1', '1', '', '1', '', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (890, 84, 'pay_amount', '', 'decimal(40,20)', 'float64', 'PayAmount', 'payAmount', '0', '0', '1', '1', '', '1', '', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (891, 84, 'pay_time', '', 'datetime', 'Time', 'PayTime', 'payTime', '0', '0', '1', '1', '', '1', '', 'EQ', 'datetime', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (892, 84, 'pay_address', '', 'varchar(45)', 'string', 'PayAddress', 'payAddress', '0', '0', '1', '1', '', '1', '', 'EQ', 'input', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (893, 84, 'create_at', '', 'datetime', 'Time', 'CreateAt', 'createAt', '0', '0', '', '1', '', '1', '', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (894, 84, 'update_at', '', 'datetime', 'Time', 'UpdateAt', 'updateAt', '0', '0', '', '1', '', '1', '', 'EQ', 'datetime', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (895, 84, 'remark', '', 'text', 'string', 'Remark', 'remark', '0', '0', '', '1', '', '', '0', 'EQ', 'input', '', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (896, 84, 'block_number', '', 'bigint', 'int64', 'BlockNumber', 'blockNumber', '0', '0', '', '1', '', '1', '', 'EQ', 'input', '', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (897, 84, 'pay_gas_fee', '', 'decimal(40,20)', 'float64', 'PayGasFee', 'payGasFee', '0', '0', '1', '1', '', '1', '', 'EQ', 'input', '', 13, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (898, 85, 'id', '', 'bigint', 'int64', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (899, 85, 'name', '链名', 'varchar(100)', 'string', 'Name', 'name', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (900, 85, 'chainId', '链Id', 'varchar(30)', 'string', 'ChainId', 'chainId', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (901, 85, 'rpc', '链rpc', 'varchar(128)', 'string', 'Rpc', 'rpc', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (902, 85, 'explorer', '链浏览器地址', 'varchar(128)', 'string', 'Explorer', 'explorer', '0', '0', '', '1', '1', '', '', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (903, 85, 'isEnable', '状态', 'tinyint(1)', 'int', 'IsEnable', 'isEnable', '0', '0', '1', '1', '1', '1', '', 'EQ', 'select', 'is_enable', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (904, 85, 'create_at', '创建时间', 'datetime', 'Time', 'CreateAt', 'createAt', '0', '0', '', '1', '1', '', '', 'EQ', 'datetime', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (905, 85, 'update_at', '更新时间', 'datetime', 'Time', 'UpdateAt', 'updateAt', '0', '0', '', '1', '1', '', '', 'EQ', 'datetime', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (906, 85, 'type', '0 ethereum 1 tron  2 btc', 'int unsigned', 'int', 'ChainType', 'type', '0', '0', '1', '1', '1', '1', '', 'EQ', 'select', 'chain_type', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (907, 85, 'icon', '链图标', 'varchar(100)', 'string', 'Icon', 'icon', '0', '0', '', '1', '1', '1', '', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (908, 85, 'isDelete', '软删除', 'tinyint(1)', 'int', 'IsDelete', 'isDelete', '0', '0', '', '1', '1', '', '', 'EQ', 'select', 'isDelete', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (909, 85, 'max_gas_price', '支付时，该公链最大接受的gasprice,如果超过则不进行目标链支付,为0时， 为0时，所有订单都不进行支付， 如果想要所有的订单都进行支付，不管gasprice 则设置很大   单位：wei', 'decimal(20,0)', 'float64', 'MaxGasPrice', 'maxGasPrice', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (910, 85, 'enablePay', '该公链作为目标时，是否允许支付', 'tinyint(1)', 'int', 'EnablePay', 'enablePay', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 13, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (911, 85, 'bridge_contract_address', '跨链桥合约地址', 'varchar(100)', 'string', 'BridgeContractAddress', 'bridgeContractAddress', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 14, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (912, 86, 'id', '序号', 'int', 'int', 'Id', 'id', '1', '1', '0', '0', '0', '1', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (913, 86, 'date', '时间', 'date', 'Time', 'Date', 'date', '0', '0', '', '', '', '1', '1', 'BETWEEN', 'datetime', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (914, 86, 'coin', '币种', 'varchar(10)', 'string', 'Coin', 'coin', '0', '0', '', '', '', '1', '1', 'LIKE', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (915, 86, 'chain_type', '链类型', 'varchar(10)', 'string', 'ChainType', 'chainType', '0', '0', '', '', '', '1', '1', 'LIKE', 'select', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (916, 86, 'transfer_in', '100', 'decimal(40,20)', 'float64', 'TransferIn', 'transferIn', '0', '0', '', '', '', '1', '', 'EQ', 'input', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (917, 86, 'transfer_out', '100', 'decimal(40,20)', 'float64', 'TransferOut', 'transferOut', '0', '0', '', '', '', '1', '', 'EQ', 'input', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (918, 86, 'transfer_difference', '跨链差额', 'decimal(40,20)', 'float64', 'TransferDifference', 'transferDifference', '0', '0', '', '', '', '1', '', 'EQ', 'input', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (919, 86, 'transfer_fee', '跨链手续费', 'decimal(40,20)', 'float64', 'TransferFee', 'transferFee', '0', '0', '', '', '', '1', '', 'EQ', 'input', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (920, 86, 'platform_assets', '平台资产', 'decimal(40,20)', 'float64', 'PlatformAssets', 'platformAssets', '0', '0', '', '', '', '1', '', 'EQ', 'input', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (921, 86, 'created_at', '', 'timestamp', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (922, 86, 'updated_at', '', 'timestamp', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 11, '', '', '', '', '');

SET FOREIGN_KEY_CHECKS = 1;
