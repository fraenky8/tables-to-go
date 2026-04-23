DROP TABLE IF EXISTS `user` CASCADE;
CREATE TABLE `user`
(
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id     INT UNSIGNED NOT NULL,
    email       VARCHAR(300) NOT NULL,

    website_url VARCHAR(300),

    KEY user_email_idx (email),
    PRIMARY KEY (id)
);
