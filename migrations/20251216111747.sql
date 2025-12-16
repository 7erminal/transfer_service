-- Modify "trx_transactions" table
ALTER TABLE `trx_transactions` MODIFY COLUMN `status_id` bigint NOT NULL, MODIFY COLUMN `date_created` datetime NOT NULL, MODIFY COLUMN `date_modified` datetime NOT NULL, MODIFY COLUMN `active` int NOT NULL DEFAULT 0, ADD COLUMN `total_debit_amount` double NOT NULL DEFAULT 0 AFTER `amount`, ADD COLUMN `description` varchar(255) NOT NULL DEFAULT "" AFTER `response_message`;
-- Create "status" table
CREATE TABLE `status` (
  `status_id` bigint NOT NULL AUTO_INCREMENT,
  `status` varchar(128) NOT NULL DEFAULT "",
  `status_code` varchar(128) NOT NULL DEFAULT "",
  `date_created` datetime NOT NULL,
  `date_modified` datetime NOT NULL,
  `created_by` int NOT NULL DEFAULT 0,
  `modified_by` int NOT NULL DEFAULT 0,
  `active` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`status_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "status_codes" table
CREATE TABLE `status_codes` (
  `status_id` bigint NOT NULL AUTO_INCREMENT,
  `status_code` varchar(50) NOT NULL DEFAULT "",
  `status_description` varchar(255) NOT NULL DEFAULT "",
  `active` int NOT NULL DEFAULT 0,
  `date_created` datetime NOT NULL,
  `date_modified` datetime NOT NULL,
  `created_by` int NOT NULL DEFAULT 0,
  `modified_by` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`status_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
