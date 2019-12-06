--
-- Database: `unknown`
--
CREATE DATABASE IF NOT EXISTS `unknown`;
USE `unknown`;

-- --
-- -- user
-- --
-- CREATE TABLE `user` (
--   `id` int(11) NOT NULL AUTO_INCREMENT,
--   `name` varchar(64) NOT NULL DEFAULT '',
--   `avatar` varchar(512) NOT NULL DEFAULT '',
--   `unionid` varchar(64) NOT NULL DEFAULT '',
--   `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
--   `created_at` timestamp NOT NULL DEFAULT current_timestamp,
--   PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- shops
--
CREATE TABLE `shops` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` varchar(32) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `distance` int(11) NOT NULL,
  `star` float(5,2) NOT NULL DEFAULT 0,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "一口轻食", 588, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "麦当劳", 477, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "手工粉", 606, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "食其家", 800, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "一秀寿司", 1024, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "桃屋料理", 998, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "丸龟制面", 332, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "过桥米线", 755, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "一乐拉面", 666, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "KFC", 233, 4.8);
INSERT INTO `unknown`.`shops` (`uid`, `name`, `distance`, `star`) VALUES ("admin", "乐凯撒", 888, 4.8);

--
-- votes
--
CREATE TABLE `votes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` varchar(32) NOT NULL,
  `subject` varchar(64) NOT NULL DEFAULT '',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- vote options
--
CREATE TABLE `vote_options` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vid` int(11) NOT NULL,
  `sid` int(11) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`vid`, `sid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- vote results
--
CREATE TABLE `vote_results` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vid` int(11) NOT NULL,
  `sid` int(11) NOT NULL,
  `uid` varchar(32) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`vid`, `uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
