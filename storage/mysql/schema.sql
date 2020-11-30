CREATE TABLE `store-test`.`user`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL,
    `mobile` VARCHAR(11) NOT NULL,
    PRIMARY KEY(`id`),
    UNIQUE `uniq_mobile`(`mobile`)
) ENGINE = InnoDB;

CREATE TABLE `store-test`.`post`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `uid` INT UNSIGNED NOT NULL,
    `title` VARCHAR(64) NOT NULL,
    `content` TEXT NOT NULL,
    PRIMARY KEY(`id`)
) ENGINE = InnoDB;

CREATE TABLE `store-test`.`comment`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `pid` INT UNSIGNED NOT NULL,
    `cid` INT UNSIGNED NOT NULL DEFAULT 0,
    `uid` INT UNSIGNED NOT NULL,
    `content` VARCHAR(255) NOT NULL,
    PRIMARY KEY(`id`),
    KEY `idx_pid_cid`(`id`, `cid`)
) ENGINE = InnoDB;