-- Modify "trx_transactions" table
ALTER TABLE `trx_transactions` ADD COLUMN `charge` double NOT NULL DEFAULT 0 AFTER `total_debit_amount`, ADD COLUMN `commission` double NOT NULL DEFAULT 0 AFTER `charge`;
