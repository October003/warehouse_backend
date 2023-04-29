
INSERT ignore INTO sys_api VALUES (11, 'warehouse/apis.SysUser.GetPage-fm', '管理员列表', '/api/sys-user/get', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (12, 'warehouse/apis.SysUser.Get-fm', '管理员通过id获取', '/api/sys-user/get/:id', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (13, 'warehouse/apis.SysUser.Insert-fm', '管理员创建', '/api/sys-user/create', 'BUS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (14, 'warehouse/apis.SysUser.Update-fm', '管理员编辑', '/api/sys-user/update', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (15, 'warehouse/apis.SysUser.Delete-fm', '管理员删除', '/api/sys-user/delete', 'BUS', 'DELETE', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (16, 'warehouse/apis.SysUser.ResetPwd-fm', '重置密码', '/api/user/pwd/reset', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);

INSERT ignore INTO sys_api VALUES (21, 'warehouse/apis.SysRole.GetPage-fm', '角色列表', '/api/role/get', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (22, 'warehouse/apis.SysRole.Get-fm', '角色通过id获取', '/api/role/get/:id', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (23, 'warehouse/apis.SysRole.Insert-fm', '角色创建', '/api/role/create', 'BUS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (24, 'warehouse/apis.SysRole.Update-fm', '角色编辑', '/api/role/update/:id', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (25, 'warehouse/apis.SysRole.Delete-fm', '角色删除', '/api/role/delete', 'BUS', 'DELETE', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (26, 'warehouse/apis.SysRole.Update2DataScope-fm', '角色数据权限修改', '/api/roledatascope', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (27, 'warehouse/apis.SysRole.Update2Status-fm', '角色状态', '/api/role-status', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);

INSERT ignore INTO sys_api VALUES (31, 'warehouse/apis.SysDept.GetPage-fm', '部门列表', '/api/dept/get', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (32, 'warehouse/apis.SysDept.Get-fm', '部门通过id获取', '/api/dept/get/:id', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (33, 'warehouse/apis.SysDept.Insert-fm', '部门创建', '/api/dept/create', 'BUS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (34, 'warehouse/apis.SysDept.Update-fm', '部门编辑', '/api/dept/update/:id', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (35, 'warehouse/apis.SysDept.Delete-fm', '部门删除', '/api/dept/delete/:id', 'BUS', 'DELETE', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (36, 'warehouse/apis.SysDept.Get2Tree-fm', '查询部门下拉树【角色权限-自定权限】', '/api/deptTree', 'SYS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);

INSERT ignore INTO sys_api VALUES (41, 'warehouse/apis.SysMenu.GetPage-fm', '菜单列表', '/api/menu/get', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (42, 'warehouse/apis.SysMenu.Get-fm', '菜单通过id获取', '/api/menu/get/:id', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (43, 'warehouse/apis.SysMenu.Insert-fm', '菜单创建', '/api/menu/create', 'BUS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (44, 'warehouse/apis.SysMenu.Update-fm', '编辑菜单', '/api/menu/update/:id', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (45, 'warehouse/apis.SysMenu.Delete-fm', '删除菜单', '/api/menu/delete', 'BUS', 'DELETE', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (46, 'warehouse/apis.SysMenu.GetMenuRole-fm', '角色菜单【顶部左侧菜单】', '/api/menurole', 'SYS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);

INSERT ignore INTO sys_api VALUES (47, 'warehouse/apis.SysMenu.GetMenuTreeSelect-fm', '菜单权限列表【角色配菜单使用】', '/api/roleMenuTreeselect/:roleId', 'SYS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (48, 'warehouse/apis.SysDept.GetDeptTreeRoleSelect-fm', '角色部门结构树【自定义数据权限】', '/api/roleDeptTreeselect/:roleId', 'SYS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);


INSERT ignore INTO sys_api VALUES (51, 'warehouse/apis.SysApi.GetPage-fm', '接口列表', '/api/sys-api/get', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (52, 'warehouse/apis.SysApi.Get-fm', '接口通过id获取', '/api/sys-api/get/:id', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (53, 'warehouse/apis.SysApi.Update-fm', '接口编辑', '/api/sys-api/update/:id', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);

INSERT ignore INTO sys_api VALUES (61, 'warehouse/apis.Admin.Get-fm', '获取系统设置', '/api/config/get', 'SYS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (62, 'warehouse/apis.Admin.Add-fm', '添加系统设置', '/api/config/add', 'SYS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (63, 'warehouse/apis.Admin.Delete-fm', '删除系统设置', '/api/config/delete', 'SYS', 'DELETE', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);

INSERT ignore INTO sys_api VALUES (71, 'warehouse/apis.Details.CreateInbound-fm', '创建入库明细', '/api/inbound/create', 'BUS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (72, 'warehouse/apis.Details.DeleteInbound-fm', '删除入库明细', '/api/inbound/delete', 'BUS', 'DELETE', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (73, 'warehouse/apis.Details.QueryInbound-fm', '查询入库明细', '/api/inbound/get', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (74, 'warehouse/apis.Details.CreateOutbound-fm', '创建出库明细', '/api/outbound/create', 'BUS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (75, 'warehouse/apis.Details.DeleteOutbound-fm', '删除出库明细', '/api/outbound/delete', 'BUS', 'DELETE', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (76, 'warehouse/apis.Details.QueryOutbound-fm', '查询出库明细', '/api/outbound/get', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (77, 'warehouse/apis.Details.QueryByTimestamp-fm', '时间查询', '/api/query_by_timestamp', 'BUS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);

INSERT ignore INTO sys_api VALUES (81, 'warehouse/apis.Material.Create-fm', '添加物料信息', '/api/item/create', 'BUS', 'POST', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (82, 'warehouse/apis.Material.Query-fm', '查询物料信息', '/api/item/get', 'BUS', 'GET', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (83, 'warehouse/apis.Material.Delete-fm', '删除物料信息', '/api/item/delete', 'BUS', 'DELETE', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);
INSERT ignore INTO sys_api VALUES (84, 'warehouse/apis.Material.Update-fm', '更新物料信息', '/api/item/update', 'BUS', 'PUT', '2023-03-24 14:20:15', '2023-03-24 14:20:15', NULL, 0, 0);

INSERT ignore INTO sys_menu VALUES (2, 'Admin', '系统管理', 'api-server', '/admin', '/0/2', 'M', '无', '', 0, 1, '', 'Layout', 10, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (3, 'SysUserManage', '用户管理', 'user', '/admin/sys-user', '/0/2/3', 'C', '无', 'admin:sysUser:list', 2, 0, '', '/sys-user/index', 10, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (11, '', '新增用户', 'app-group-fill', '', '/0/2/3/11', 'F', 'POST', 'admin:sysUser:add', 3, 0, '', '', 10, '0', '1', 0, 1, '2023-03-24 15:38:16.000','2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (12, '', '查询用户', 'app-group-fill', '', '/0/2/3/12', 'F', 'GET', 'admin:sysUser:query', 3, 0, '', '', 40, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (13, '', '修改用户', 'app-group-fill', '', '/0/2/3/13', 'F', 'PUT', 'admin:sysUser:edit', 3, 0, '', '', 30, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (14, '', '删除用户', 'app-group-fill', '', '/0/2/3/14', 'F', 'DELETE', 'admin:sysUser:remove', 3, 0, '', '', 20, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);

INSERT ignore INTO sys_menu VALUES (4, 'SysRoleManage', '角色管理', 'peoples', '/admin/sys-role', '/0/2/4', 'C', '无', 'admin:sysRole:list', 2, 1, '', '/sys-role/index', 20, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (21, '', '新增角色', 'app-group-fill', '', '/0/2/4/21', 'F', '', 'admin:sysRole:add', 4, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (22, '', '查询角色', 'app-group-fill', '', '/0/2/4/22', 'F', '', 'admin:sysRole:query', 4, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (23, '', '修改角色', 'app-group-fill', '', '/0/2/4/23', 'F', '', 'admin:sysRole:update', 4, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (24, '', '删除角色', 'app-group-fill', '', '/0/2/4/24', 'F', '', 'admin:sysRole:remove', 4, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);

INSERT ignore INTO sys_menu VALUES (5, 'SysDeptManage', '部门管理', 'tree', '/admin/sys-dept', '/0/2/5', 'C', '无', 'admin:sysDept:list', 2, 0, '', '/sys-dept/index', 40, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (31, '', '查询部门', 'app-group-fill', '', '/0/2/5/31', 'F', '', 'admin:sysDept:query', 5, 0, '', '', 40, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (32, '', '新增部门', 'app-group-fill', '', '/0/2/5/32', 'F', '', 'admin:sysDept:add', 5, 0, '', '', 10, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (33, '', '修改部门', 'app-group-fill', '', '/0/2/5/33', 'F', '', 'admin:sysDept:edit', 5, 0, '', '', 30, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (34, '', '删除部门', 'app-group-fill', '', '/0/2/5/34', 'F', '', 'admin:sysDept:remove',5, 0, '', '', 20, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);

INSERT ignore INTO sys_menu VALUES (6, 'SysMenuManage', '菜单管理', 'tree-table', '/admin/sys-menu', '/0/2/6', 'C', '无', 'admin:sysMenu:list', 2, 1, '', '/sys-menu/index', 30, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (41, '', '新增菜单', 'app-group-fill', '', '/0/2/6/41', 'F', '', 'admin:sysMenu:add', 6, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (42, '', '修改菜单', 'app-group-fill', '', '/0/2/6/42', 'F', '', 'admin:sysMenu:edit', 6, 0, '', '', 1, '0', '1', 1, 1,'2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (43, '', '查询菜单', 'app-group-fill', '', '/0/2/6/43', 'F', '', 'admin:sysMenu:query', 6, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (44, '', '删除菜单', 'app-group-fill', '', '/0/2/6/44', 'F', '', 'admin:sysMenu:remove', 6, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);

INSERT ignore INTO sys_menu VALUES (7, 'ConfigManage', '设置管理', 'tree-table', '/config', '/0/2/7', 'C', '无', 'admin:config:list', 2, 1, '', '/config/index', 30, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (51, '', '新增系统设置', 'app-group-fill', '', '/0/2/7/51', 'F', '', 'admin:config:add', 7, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (52, '', '查询系统设置', 'app-group-fill', '', '/0/2/7/52', 'F', '', 'admin:config:query', 7, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (53, '', '删除系统设置', 'app-group-fill', '', '/0/2/7/53', 'F', '', 'admin:config:remove', 7, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);

INSERT ignore INTO sys_menu VALUES (8, 'DetailsManage', '出入库管理', 'tree-table', '/details', '/0/2/8', 'C', '无', 'admin:detail:list', 2, 1, '', '/detail/index', 30, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (61, '', '新增入库明细', 'app-group-fill', '', '/0/2/8/61', 'F', '', 'admin:detail:add', 8, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (62, '', '查询入库明细', 'app-group-fill', '', '/0/2/8/62', 'F', '', 'admin:detail:query', 8, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (63, '', '删除入库明细', 'app-group-fill', '', '/0/2/8/63', 'F', '', 'admin:detail:remove', 8, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (64, '', '新增出库明细', 'app-group-fill', '', '/0/2/8/64', 'F', '', 'admin:detail:add', 8, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (65, '', '查询出库明细', 'app-group-fill', '', '/0/2/8/65', 'F', '', 'admin:detail:query', 8, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (66, '', '删除出库明细', 'app-group-fill', '', '/0/2/8/66', 'F', '', 'admin:detail:remove', 8, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (67, '', '时间查询', 'app-group-fill', '', '/0/2/8/67', 'F', '', 'admin:time:query', 8, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);

INSERT ignore INTO sys_menu VALUES (9, 'MaterialManage', '物料管理', 'tree-table', '/material', '/0/2/9', 'C', '无', 'admin:material:list', 2, 1, '', '/material/index', 30, '0', '1', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (71, '', '新增物料', 'app-group-fill', '', '/0/2/9/71', 'F', '', 'admin:material:add', 9, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (72, '', '查询物料', 'app-group-fill', '', '/0/2/9/72', 'F', '', 'admin:material:query',9, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (73, '', '删除物料', 'app-group-fill', '', '/0/2/9/73', 'F', '', 'admin:material:remove',9, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (74, '', '更新物料', 'app-group-fill', '', '/0/2/9/74', 'F', '', 'admin:material:edit',9, 0, '', '', 1, '0', '1', 1, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);

INSERT ignore INTO sys_menu VALUES (10, 'SysApiManage', '接口管理', 'api-doc', '/sys-api', '/0/2/10', 'C', '无', 'admin:sysApi:list', 2, 0, '', '/sys-api/index', 0, '0', '0', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (81, '', '查询接口', 'app-group-fill', '', '/0/2/10/81', 'F', '无', 'admin:sysApi:query', 10, 0, '', '', 40, '0', '0', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);
INSERT ignore INTO sys_menu VALUES (82, '', '修改接口', 'app-group-fill', '', '/0/2/10/82', 'F', '无', 'admin:sysApi:edit', 10, 0, '', '', 30, '0', '0', 0, 1, '2023-03-24 15:38:16.000', '2023-03-24 15:38:16.000', NULL);

INSERT ignore INTO sys_menu_api_rule VALUES (3,11);
INSERT ignore INTO sys_menu_api_rule VALUES (4,21);
INSERT ignore INTO sys_menu_api_rule VALUES (5,31);
INSERT ignore INTO sys_menu_api_rule VALUES (6,41);
INSERT ignore INTO sys_menu_api_rule VALUES (7,61);
INSERT ignore INTO sys_menu_api_rule VALUES (8,73);
INSERT ignore INTO sys_menu_api_rule VALUES (8,76);
INSERT ignore INTO sys_menu_api_rule VALUES (9,82);
INSERT ignore INTO sys_menu_api_rule VALUES (10,51);

INSERT ignore INTO sys_menu_api_rule VALUES (11, 13);
INSERT ignore INTO sys_menu_api_rule VALUES (12, 12);
INSERT ignore INTO sys_menu_api_rule VALUES (13,14);
INSERT ignore INTO sys_menu_api_rule VALUES (14,15);

INSERT ignore INTO sys_menu_api_rule VALUES (21, 23);
INSERT ignore INTO sys_menu_api_rule VALUES (22, 21);
INSERT ignore INTO sys_menu_api_rule VALUES (23, 24);
INSERT ignore INTO sys_menu_api_rule VALUES (24,25);

INSERT ignore INTO sys_menu_api_rule VALUES (31,31);
INSERT ignore INTO sys_menu_api_rule VALUES (32, 33);
INSERT ignore INTO sys_menu_api_rule VALUES (33, 34);
INSERT ignore INTO sys_menu_api_rule VALUES (34, 35);

INSERT ignore INTO sys_menu_api_rule VALUES (41,43);
INSERT ignore INTO sys_menu_api_rule VALUES (42, 44);
INSERT ignore INTO sys_menu_api_rule VALUES (43, 41);
INSERT ignore INTO sys_menu_api_rule VALUES (44, 45);

INSERT ignore INTO sys_menu_api_rule VALUES (51,62);
INSERT ignore INTO sys_menu_api_rule VALUES (52, 61);
INSERT ignore INTO sys_menu_api_rule VALUES (53, 63);

INSERT ignore INTO sys_menu_api_rule VALUES (61,71);
INSERT ignore INTO sys_menu_api_rule VALUES (62, 73);
INSERT ignore INTO sys_menu_api_rule VALUES (63, 72);
INSERT ignore INTO sys_menu_api_rule VALUES (64, 74);
INSERT ignore INTO sys_menu_api_rule VALUES (65,76);
INSERT ignore INTO sys_menu_api_rule VALUES (66, 75);
INSERT ignore INTO sys_menu_api_rule VALUES (67, 77);

INSERT ignore INTO sys_menu_api_rule VALUES (71,81);
INSERT ignore INTO sys_menu_api_rule VALUES (72, 82);
INSERT ignore INTO sys_menu_api_rule VALUES (73, 83);
INSERT ignore INTO sys_menu_api_rule VALUES (74, 84);

INSERT ignore INTO sys_menu_api_rule VALUES (81,51);
INSERT ignore INTO sys_menu_api_rule VALUES (82,53);

INSERT ignore INTO `sys_user` VALUES(1, 'admin', '$2a$10$.Dj5uoorxg3P9Byosj7hbuOFDNM/TVfvwlDHOHq6F1CCWFMHWktaK','系统管理员','电话','',1,1,'2',0,0,'2023-03-24 13:19:21','2023-03-24 13:19:21',NULL);

INSERT ignore INTO sys_role VALUES (1, '系统管理员', '2', 'admin', 1, '', '', 1, '', 1, 1, '2023-03-24 14:20:15.000', '2023-03-24 14:20:15.000', NULL);
INSERT ignore INTO sys_role VALUES (2, '普通用户', '2', '普通权限', 1, '', '', 1, '', 1, 1, '2023-03-24 14:20:15.000', '2023-03-24 14:20:15.000', NULL);

INSERT ignore INTO sys_dept VALUES (1, 0, '/0/1/', 'XX省', 0, '', '', '2', 1, 1, '2023-03-20 14:44:50.681', '2023-03-20 14:44:50.681', NULL);
INSERT ignore INTO sys_dept VALUES (11, 1, '/0/1/11/', 'XX市', 1, '', '', '2', 1, 1, '2023-03-20 14:44:50.681', '2023-03-20 14:44:50.681', NULL);
INSERT ignore INTO sys_dept VALUES (21, 11, '/0/1/11/21/', 'XX区', 1, '', '',  '2', 1, 1, '2023-03-20 14:44:50.681', '2023-03-20 14:44:50.681', NULL);
