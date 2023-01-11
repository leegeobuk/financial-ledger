CREATE TABLE IF NOT EXISTS ledger (
    ledger_id INT UNSIGNED AUTO_INCREMENT NOT NULL,
    income INT NOT NULL,
    ledger_date DATE NOT NULL,
    user_id INT UNSIGNED NOT NULL,
    PRIMARY KEY (ledger_id),
    FOREIGN KEY (user_id) REFERENCES user_account(user_id)
) ENGINE=INNODB;