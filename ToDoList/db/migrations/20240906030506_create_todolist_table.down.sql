-- Down Migration

DROP TRIGGER IF EXISTS `before_task_update`;
DROP TRIGGER IF EXISTS `before_task_insert`;

DROP TABLE IF EXISTS `tasks`;
DROP TABLE IF EXISTS `users`;

