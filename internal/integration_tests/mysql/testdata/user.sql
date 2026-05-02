DROP TABLE IF EXISTS `user` CASCADE;
CREATE TABLE `user`
(
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id     INT UNSIGNED NOT NULL COMMENT 'This is the identifier of the user.',
    email       VARCHAR(300) NOT NULL COMMENT 'This is the email of the user.
Use it for notifications.',

    website_url VARCHAR(300) COMMENT 'This is the website URL of the user.
Optional field.',

    KEY user_email_idx (email),
    PRIMARY KEY (id)
) COMMENT = 'This is the user table.
Contains account information.';
