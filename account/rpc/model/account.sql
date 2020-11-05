CREATE TABLE IF NOT EXISTS account
(
    `id`                   VARCHAR(255),
    `name`                 VARCHAR(255) NOT NULL default '',
    `email`                VARCHAR(255) NOT NULL,
    `phone_number`         VARCHAR(255) NOT NULL,
    `confirmed_and_active` TINYINT(1)      NOT NULL DEFAULT 0,
    `member_since`         TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `password_hash`        VARCHAR(100)          DEFAULT '',
    `password_salt`        VARCHAR(100)          DEFAULT '',
    `photo_url`            VARCHAR(255) NOT NULL,
    `support`              TINYINT(1)      NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    key ix_account_email (email),
    key ix_account_phone_number (phone_number)
) ENGINE = InnoDB;