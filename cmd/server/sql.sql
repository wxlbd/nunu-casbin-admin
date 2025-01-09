CREATE TABLE `menu_meta` (
    `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    `title` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '标题',
    `i18n` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '国际化',
    `badge` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '徽章',
    `icon` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '图标',
    `affix` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否固定（0:否, 1:是）',
    `hidden` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否隐藏（0:否, 1:是）',
    `type` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '类型',
    `cache` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否缓存（0:否, 1:是）',
    `copyright` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否显示版权（0:否, 1:是）',
    `breadcrumb_enable` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '是否显示面包屑（0:否, 1:是）',
    `component_path` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '视图前缀路径',
    `component_suffix` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '视图文件类型',
    `link` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '链接',
    `active_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '高亮菜单标识',
    `auth` JSON COMMENT '权限码（JSON数组）',
    `role` JSON COMMENT '角色码（JSON数组）',
    `user` JSON COMMENT '用户名（JSON数组）',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单元数据表';


CREATE TABLE `menu` (
    `id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    `parent_id` INT DEFAULT NULL COMMENT '父ID',
    `name` VARCHAR(255) NOT NULL COMMENT '菜单名称',
    `path` VARCHAR(255) NOT NULL COMMENT '路由地址',
    `component` VARCHAR(255) NOT NULL COMMENT '组件路径',
    `redirect` VARCHAR(255) DEFAULT NULL COMMENT '重定向地址',
    `type` VARCHAR(1) NOT NULL COMMENT '菜单类型 (M:菜单, B:按钮, L:链接, I:iframe)',
    `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 (1:正常, 2:停用)',
    `sort` INT NOT NULL DEFAULT 0 COMMENT '排序',
    `created_by` INT DEFAULT NULL COMMENT '创建者',
    `updated_by` INT DEFAULT NULL COMMENT '更新者',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注',
    `meta` JSON COMMENT '附加属性（JSON对象）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';