-- Trigger untuk validasi due_date sebelum insert
CREATE TRIGGER `before_task_insert`
BEFORE INSERT ON `tasks`
FOR EACH ROW
BEGIN
    IF NEW.due_date IS NOT NULL AND NEW.due_date < NOW() THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Due date cannot be in the past';
    END IF;
END;


