-- Create "trx_transaction_details" table
CREATE TABLE `trx_transaction_details` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `transaction_id` varchar(100) NOT NULL DEFAULT "",
  `amount` double NOT NULL DEFAULT 0,
  `charge` double NOT NULL DEFAULT 0,
  `commission` double NOT NULL DEFAULT 0,
  `sender` varchar(255) NOT NULL DEFAULT "",
  `recipient` varchar(255) NOT NULL DEFAULT "",
  `status` int NOT NULL DEFAULT 0,
  `response_code` varchar(50) NOT NULL DEFAULT "",
  `recipient_name` varchar(255) NOT NULL DEFAULT "",
  `date_created` datetime NOT NULL,
  `date_modified` datetime NOT NULL,
  `created_by` int NOT NULL DEFAULT 0,
  `modified_by` int NOT NULL DEFAULT 0,
  `active` int NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
