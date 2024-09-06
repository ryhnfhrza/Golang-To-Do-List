-- Up Migration

-- Tabel users
CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) NOT NULL,
  `username` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `email` varchar(254) NOT NULL,
  `password` varchar(64) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Tabel tasks
CREATE TABLE IF NOT EXISTS `tasks` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` text,
  `completed` tinyint(1) NOT NULL DEFAULT '0',
  `due_date` timestamp NULL DEFAULT NULL,
  `notified` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Trigger untuk validasi due_date
CREATE TRIGGER `before_task_insert` BEFORE INSERT ON `tasks`
FOR EACH ROW
BEGIN
    IF NEW.due_date < NOW() THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Due date cannot be in the past';
    END IF;
END;

CREATE TRIGGER `before_task_update` BEFORE UPDATE ON `tasks`
FOR EACH ROW
BEGIN
    IF NEW.due_date < NOW() THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Due date cannot be in the past';
    END IF;
END;
