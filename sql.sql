create table xcx_aircraft
(
    `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `title`         varchar(50) NOT NULL DEFAULT '',
    `cover_picture` varchar(50) NOT NULL DEFAULT '',
    `type`          tinyint(4) NOT NULL default 1,
    `tag`           varchar(50) NOT NULL DEFAULT '',
    `content`       text,
    `create_time`   int         NOT NULL default 0,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table xcx_admin_user
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        varchar(50)  NOT NULL DEFAULT '',
    `email`       varchar(100) not null default '',
    `password`    varchar(100) not null default '',
    `create_time` int          not null default 0,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table xcx_admin_user_cookies
(

)

create table xcx_game
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name`        varchar(50) NOT NULL DEFAULT '',
    `create_time` INT(11) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table xcx_area
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `game_id`     INT UNSIGNED NOT NULL DEFAULT 0,
    `name`        varchar(50) NOT NULL DEFAULT '',
    `create_time` INT(11) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    KEY           `idx_game_id` (`game_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table xcx_props
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `game_id`     INT UNSIGNED NOT NULL DEFAULT 0,
    `name`        VARCHAR(255) NOT null DEFAULT "",
    `create_time` INT(11) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    KEY           `idx_game_id` (`game_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table xcx_price
(
    `id`           INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `props_id`     INT UNSIGNED NOT NULL DEFAULT 0,
    `area_id`      INT UNSIGNED NOT NULL DEFAULT 0,
    `unit_price`   int not null default 0,
    `quantity`     int not      default 0,
    `total_prices` int not      default 0,
    `create_time`  int not      default 0,
    PRIMARY KEY (`id`),
    KEY            `idx_props_id` (`props_id`),
    KEY            `idx_area_id` (`area_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;