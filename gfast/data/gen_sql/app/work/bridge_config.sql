/*
==========================================================================
GFast自动生成菜单SQL,只生成一次,按需修改.
生成日期：2024-05-27 18:48:21
生成路径: data/gen_sql/work/bridge_config_menu.sql
生成人：gfast
==========================================================================
*/
-- 当前日期
select @now := now();
-- 目录 SQL
INSERT INTO `sys_auth_rule` (`pid`,`name`,`title`,`icon`,`condition`,`remark`,`menu_type`,`weigh`,`status`,`always_show`,`path`,`jump_path`,`component`,`is_frame`,`module_type`,`model_id`,`created_at`,`updated_at`,`deleted_at` )
VALUES(0,'work/bridgeConfig','可跨币种设置管理','form','','可跨币种设置管理',0,0,1,1,'bridgeConfig','','',0,'sys_admin',0,@now,@now,NULL );
-- 菜单父目录ID
SELECT @parentId := LAST_INSERT_ID();
-- 菜单 SQL
INSERT INTO `sys_auth_rule` (`pid`,`name`,`title`,`icon`,`condition`,`remark`,`menu_type`,`weigh`,`status`,`always_show`,`path`,`jump_path`,`component`,`is_frame`,`module_type`,`model_id`,`created_at`,`updated_at`,`deleted_at` )
VALUES(@parentId,'work/bridgeConfig/list','可跨币种设置列表','list','','可跨币种设置列表',1,0,1,1,'bridgeConfigList','','work/bridgeConfig/list',0,'sys_admin',0,@now,@now,NULL );
-- 按钮父目录ID
SELECT @parentId := LAST_INSERT_ID();
-- 按钮 SQL
INSERT INTO `sys_auth_rule` (`pid`,`name`,`title`,`icon`,`condition`,`remark`,`menu_type`,`weigh`,`status`,`always_show`,`path`,`jump_path`,`component`,`is_frame`,`module_type`,`model_id`,`created_at`,`updated_at`,`deleted_at` )
VALUES(@parentId,'work/bridgeConfig/get','可跨币种设置查询','','','可跨币种设置查询',2,0,1,1,'','','',0,'sys_admin',0,@now,@now,NULL );
INSERT INTO `sys_auth_rule` (`pid`,`name`,`title`,`icon`,`condition`,`remark`,`menu_type`,`weigh`,`status`,`always_show`,`path`,`jump_path`,`component`,`is_frame`,`module_type`,`model_id`,`created_at`,`updated_at`,`deleted_at` )
VALUES(@parentId,'work/bridgeConfig/add','可跨币种设置添加','','','可跨币种设置添加',2,0,1,1,'','','',0,'sys_admin',0,@now,@now,NULL );
INSERT INTO `sys_auth_rule` (`pid`,`name`,`title`,`icon`,`condition`,`remark`,`menu_type`,`weigh`,`status`,`always_show`,`path`,`jump_path`,`component`,`is_frame`,`module_type`,`model_id`,`created_at`,`updated_at`,`deleted_at` )
VALUES(@parentId,'work/bridgeConfig/edit','可跨币种设置修改','','','可跨币种设置修改',2,0,1,1,'','','',0,'sys_admin',0,@now,@now,NULL );
INSERT INTO `sys_auth_rule` (`pid`,`name`,`title`,`icon`,`condition`,`remark`,`menu_type`,`weigh`,`status`,`always_show`,`path`,`jump_path`,`component`,`is_frame`,`module_type`,`model_id`,`created_at`,`updated_at`,`deleted_at` )
VALUES(@parentId,'work/bridgeConfig/delete','可跨币种设置删除','','','可跨币种设置删除',2,0,1,1,'','','',0,'sys_admin',0,@now,@now,NULL );
