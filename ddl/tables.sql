CREATE DATABASE IF NOT EXISTS household_ledger;

CREATE TABLE IF NOT EXISTS user_account (
    user_id INT UNSIGNED AUTO_INCREMENT NOT NULL,
    username VARCHAR(64) NOT NULL,
    PRIMARY KEY (user_id)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS user_login (
    user_id INT UNSIGNED AUTO_INCREMENT NOT NULL,
    email VARCHAR(64) NOT NULL,
    pw_hash VARCHAR(250) NOT NULL,
    PRIMARY KEY (user_id)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS ledger (
    ledger_id INT UNSIGNED AUTO_INCREMENT NOT NULL,
    income INT NOT NULL,
    ledger_date DATE NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    PRIMARY KEY (ledger_id),
    FOREIGN KEY (user_id) REFERENCES user_account(user_id)
) ENGINE=INNODB;
