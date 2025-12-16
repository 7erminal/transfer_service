-- Create "trx_transactions" table
CREATE TABLE `trx_transactions` (
  `transaction_id` varchar(255) NOT NULL,
  `amount` double NOT NULL DEFAULT 0,
  `sender_account_number` varchar(255) NOT NULL DEFAULT "",
  `recipient_account_number` varchar(255) NOT NULL DEFAULT "",
  `transfer_code` varchar(150) NOT NULL DEFAULT "",
  `status_id` int NOT NULL,
  `response_code` varchar(50) NOT NULL DEFAULT "",
  `response_message` varchar(255) NOT NULL DEFAULT "",
  `date_created` datetime DEFAULT CURRENT_TIMESTAMP,
  `date_modified` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `created_by` int NOT NULL DEFAULT 0,
  `modified_by` int NOT NULL DEFAULT 0,
  `active` int NOT NULL DEFAULT 1,
  PRIMARY KEY (`transaction_id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
