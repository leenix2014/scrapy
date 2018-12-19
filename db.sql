DROP TABLE IF EXISTS `t_pdf`;
CREATE TABLE `t_pdf`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_mail` VARCHAR(64) NULL DEFAULT NULL COMMENT '用户邮箱',
  `root` VARCHAR(512) NULL DEFAULT NULL COMMENT 'pdf根url',
  `url` VARCHAR(512) NULL DEFAULT NULL COMMENT 'pdf下载url',
  `visited` BOOLEAN NOT NULL DEFAULT 0 COMMENT '是否已访问',
  `createTime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updateTime` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX(`user_mail`,`url`)
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci;