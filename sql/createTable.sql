CREATE TABLE `user`(  
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id` BIGINT(20) NOT NULL COMMENT '用户id',
    `username` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
    `password` VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
    `email` VARCHAR(64) COLLATE utf8mb4_general_ci COMMENT '邮箱',
    `gender` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '性别',
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;