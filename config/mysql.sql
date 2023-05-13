DROP DATABASE IF EXISTS  tiktok;
CREATE DATABASE tiktok;
USE tiktok;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int NOT NULL,
  `user_id` int NULL DEFAULT NULL,
  `video_id` int NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `commit_time` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`  (
  `id` int NOT NULL,
  `verion_id` int NULL DEFAULT NULL,
  `user_id` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`  (
   `id` int NOT NULL AUTO_INCREMENT,
   `following_id` int NOT NULL COMMENT '被关注用户 id',
   `follower_id` int NOT NULL COMMENT '粉丝用户 id',
   PRIMARY KEY (`id`) USING BTREE,
   UNIQUE INDEX `follow_index`(`follower_id`, `following_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 133 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
    `id` int NOT NULL AUTO_INCREMENT,
    `from_user_id` int NULL DEFAULT NULL,
    `to_user_id` int NULL DEFAULT NULL,
    `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
    `create_time` int NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------
INSERT INTO `message` VALUES (5, 63, 62, '1', 1683961615);
INSERT INTO `message` VALUES (6, 62, 63, '2', 1683961630);
INSERT INTO `message` VALUES (7, 62, 63, '3', 1683961633);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name_index`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


INSERT INTO `tiktok`.`users` (`id`, `name`, `password`, `avatar`, `background_image`, `signature`)
VALUES (1, 'mike', 'password', 'asd','asd','asd');
INSERT INTO `tiktok`.`users` (`id`, `name`, `password`)
VALUES (-1, 'null', 'null');
-- ----------------------------
-- Table structure for video
-- ----------------------------
CREATE TABLE videos (
  ID int UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  author_id int,
  title varchar(255),
  play_url varchar(255),
  cover_url varchar(255),
  publish_time datetime
);

INSERT INTO `tiktok`.`videos` (`id`, `author_id`, `play_url`, `cover_url`, `title`, `publish_time`)
VALUES (1, 1, 'https://mini-tiktok-bytedance.oss-cn-beijing.aliyuncs.com/The%20Long%20Season.mp4',
        'https://mini-tiktok-bytedance.oss-cn-beijing.aliyuncs.com/The%20Long%20Season.mp4', 'test',
        '2023-02-02 19:49:36');