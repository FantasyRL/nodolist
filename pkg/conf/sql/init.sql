SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
SET global Time_Zone ='+8:00';

CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `user_name` varchar(255) NOT NULL COMMENT '用户名',
    `password` varchar(255) NOT NULL COMMENT '密码',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `user_name_password_idx` (`user_name`,`password`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `task` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '任务ID',
    `uid` bigint NOT NULL COMMENT '用户ID',
    `title` varchar(255) NOT NULL COMMENT '标题',
    `content` varchar(255) NOT NULL COMMENT '正文',
    `status` bigint NOT NULL DEFAULT 0 COMMENT '状态',
    `finished_at` varchar(255) NULL DEFAULT NULL COMMENT '完成时间',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    constraint `task_uid`
        foreign key (`uid`)
            references `user` (`id`)
            on delete cascade
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='任务表';