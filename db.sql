DROP TABLE IF EXISTS `t_pdf`;
CREATE TABLE `t_pdf`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_mail` VARCHAR(256) NULL DEFAULT NULL COMMENT '用户邮箱',
  `root` VARCHAR(1024) NULL DEFAULT NULL COMMENT 'pdf根url',
  `url` VARCHAR(1024) NULL DEFAULT NULL COMMENT 'pdf下载url',
  `visited` BOOLEAN NULL DEFAULT 0 COMMENT '是否已访问',
  `createTime` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updateTime` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci;