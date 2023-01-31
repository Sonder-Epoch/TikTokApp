
-- ---
-- Table 'Video'
-- 
-- ---

DROP TABLE IF EXISTS `Video`;
		
CREATE TABLE `Video` (
  `id` BIGINT NOT NULL,
  `author_id` BIGINT COMMENT '作者id',
  `play_url` VARCHAR(255) COMMENT '视频播放地址',
  `cover_url` VARCHAR(255) COMMENT '视频封面地址',
  `favorite_count` BIGINT COMMENT '视频点赞总数',
  `comment_count` BIGINT COMMENT '评论总数',
  `is_favorite` CHAR(2) COMMENT '是否点赞',
  `title` VARCHAR(255) COMMENT '视频标题',
  `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

-- ---
-- Table 'User'
-- 
-- ---

DROP TABLE IF EXISTS `User`;
		
CREATE TABLE `User` (
  `id` BIGINT NOT NULL,
  `name` VARCHAR(255) COMMENT '用户名',
  `password` VARCHAR(255) COMMENT '密码',
  `follow_count` BIGINT COMMENT '关注总数',
  `follower_count` BIGINT COMMENT '粉丝总数',
  `is_follow` CHAR(2) COMMENT '是否关注',
  PRIMARY KEY (`id`)
);

-- ---
-- Table 'favorite'
-- 
-- ---

DROP TABLE IF EXISTS `favorite`;
		
CREATE TABLE `favorite` (
  `video_id` BIGINT NOT NULL COMMENT '视频id',
  `user_id` BIGINT NOT NULL COMMENT '用户id',
  PRIMARY KEY (`video_id`, `user_id`)
);

-- ---
-- Table 'comment'
-- 
-- ---

DROP TABLE IF EXISTS `comment`;
		
CREATE TABLE `comment` (
  `id` BIGINT NOT NULL,
  `user_id` BIGINT COMMENT '用户id',
  `video_id` BIGINT COMMENT '视频id',
  `content` TEXT COMMENT '内容',
  `create_date` DATE NOT NULL,
  PRIMARY KEY (`id`)
);

-- ---
-- Table 'relation'
-- 
-- ---

DROP TABLE IF EXISTS `relation`;
		
CREATE TABLE `relation` (
  `user_id` BIGINT NOT NULL COMMENT '用户id',
  `follow_id` BIGINT NOT NULL COMMENT '关注id'
);
