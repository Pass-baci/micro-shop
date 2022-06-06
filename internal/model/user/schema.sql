Create Table User (
    `id` varchar(50) PRIMARY KEY NOT NULL,
    `username` varchar(255) NOT NULL DEFAULT '',
    `password` varchar(500) NOT NULL DEFAULT '',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP
);