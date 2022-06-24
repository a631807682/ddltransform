CREATE TABLE `products` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `code` longtext,
  `price` double DEFAULT NULL,
  `after_find_call_times` bigint DEFAULT NULL,
  `before_create_call_times` bigint DEFAULT NULL,
  `after_create_call_times` bigint DEFAULT NULL,
  `before_update_call_times` bigint DEFAULT NULL,
  `after_update_call_times` bigint DEFAULT NULL,
  `before_save_call_times` bigint DEFAULT NULL,
  `after_save_call_times` bigint DEFAULT NULL,
  `before_delete_call_times` bigint DEFAULT NULL,
  `after_delete_call_times` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_products_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;