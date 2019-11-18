--
-- Database: `unknown`
--
CREATE DATABASE IF NOT EXISTS `unknown`;
USE `unknown`;

--
-- user
--
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL DEFAULT '',
  `avatar` varchar(512) NOT NULL DEFAULT '',
  `unionid` varchar(64) NOT NULL DEFAULT '',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- wishing pool
--
CREATE TABLE `wishing_pool` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `name` varchar(64) NOT NULL DEFAULT '',
  `count` int(11) NOT NULL DEFAULT 0,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- options
--
CREATE TABLE `options` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL,
  `wid` int(11) NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`, `uid`, `wid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
