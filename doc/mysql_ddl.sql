CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL comment='用户名',
  `mobile` varchar(11) NOT NULL comment='手机号码',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment='用户表';

CREATE TABLE `post` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uid` int unsigned NOT NULL comment='用户id',
  `title` varchar(64) NOT NULL comment='文章标题',
  `content` text NOT NULL comment='文章内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment='文章表';

CREATE TABLE `comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(10) unsigned NOT NULL comment='文章id',
  `cid` int(10) unsigned NOT NULL DEFAULT '0' comment='回复的评论id',
  `uid` int(10) unsigned NOT NULL comment='评论者id',
  `content` varchar(255) NOT NULL comment='评论内容',
  PRIMARY KEY (`id`),
  KEY `idx_pid_cid` (`id`,`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment='评论表';