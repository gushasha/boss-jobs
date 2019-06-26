CREATE TABLE `spiders_boss_jobs` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
     `jid` varchar(255) NOT NULL DEFAULT '' COMMENT '第三方平台上的ID：jobID',
     `job_type` varchar(30) NOT NULL COMMENT '工作类型',
     `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
     `salary_range` varchar(255) NOT NULL DEFAULT '' COMMENT '薪资范围',
     `work_years` varchar(255) DEFAULT '' COMMENT '工作年限',
     `education` varchar(255) DEFAULT '' COMMENT '学历要求',
     `company_name` varchar(255) DEFAULT '' COMMENT '公司名称',
     `company_address` varchar(255) NOT NULL DEFAULT '' COMMENT '公司地址',
     `company_label` varchar(255) DEFAULT '' COMMENT '公司类型',
     `financing_stage` varchar(255) DEFAULT '' COMMENT '融资阶段',
     `company_employees_count` varchar(255) DEFAULT '' COMMENT '公司规模-员工人数',
     `create_time` int(10) unsigned DEFAULT '0',
     `update_time` int(10) unsigned DEFAULT NULL,
     PRIMARY KEY (`id`),
     UNIQUE KEY `uni_jid` (`jid`)
) ENGINE=InnoDB AUTO_INCREMENT=9167 DEFAULT CHARSET=utf8;