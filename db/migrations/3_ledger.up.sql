CREATE TABLE IF NOT EXISTS ledger (
    ledger_id INT UNSIGNED AUTO_INCREMENT NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    ledger_desc VARCHAR(64) NOT NULL,
    income INT NOT NULL,
    ledger_date DATE NOT NULL,
    PRIMARY KEY (ledger_id)
) ENGINE=INNODB;