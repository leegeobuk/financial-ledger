CREATE TABLE IF NOT EXISTS user_login (
    user_id VARCHAR(64) NOT NULL,
    passwd VARCHAR(250) NOT NULL,
    PRIMARY KEY (user_id)
) ENGINE=INNODB;