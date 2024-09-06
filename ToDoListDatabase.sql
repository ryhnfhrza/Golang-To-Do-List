create database todo_list;
use todo_list;

CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(254) NOT NULL UNIQUE,
    password VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    due_date TIMESTAMP,
    notified BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    # on delate cascade maksudnya jika table user dihapus maka akan otematis terhapus juga di task yang berhubungan dengan user yang dihapus tersebut
);


#validasi duedate agar tidak bisa memasukkan data yang tanggalnya sudah lewat
DELIMITER //

CREATE TRIGGER before_task_insert
BEFORE INSERT ON tasks
FOR EACH ROW
BEGIN
    IF NEW.due_date < NOW() THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Due date cannot be in the past';
    END IF;
END//

CREATE TRIGGER before_task_update
BEFORE UPDATE ON tasks
FOR EACH ROW
BEGIN
    IF NEW.due_date < NOW() THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Due date cannot be in the past';
    END IF;
END//

DELIMITER ;


select *from users;
desc tasks;
insert into tasks (id,user_id,title,description,completed,due_date,notified,created_at,updated_at) values ("jfoazjjaaadsdacaaofj","06a4745fa05f43739a573edb69cc84c4","Tugas Besar","ini adalah tugas besar dan harus dikumpulkan",0,"2024-08-12 14:29:00",0,now(),now());

show triggers;
select *from tasks;
select *from users where id = "55843c058c2d4d6099e8854f83ae122c";

show tables;
desc users;
desc tasks;

select *from users;

select *from tasks where user_id = "55843c058c2d4d6099e8854f83ae122c";

desc users;
select *from tasks;
update tasks set title = "oke 1", description = "okeoke", due_date = "2025-08-22" where id = "08701474074148a5adc5d6d8e739a626" and user_id = "2b0417185bab4d50beaa876b3a6f2677";  


 SET sql_safe_updates = 1;
 
 show triggers;
 #Testing
DELIMITER //

CREATE TRIGGER before_task_insert
BEFORE INSERT ON tasks
FOR EACH ROW
BEGIN
    IF NEW.due_date <= NOW() THEN
        INSERT INTO debug_log (log_message) VALUES (CONCAT('Validation failed: due_date=', NEW.due_date, ', now=', NOW()));
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Due date cannot be in the past or present time';
    ELSE
        INSERT INTO debug_log (log_message) VALUES (CONCAT('Validation passed: due_date=', NEW.due_date, ', now=', NOW()));
    END IF;
END//

DELIMITER ;

DELIMITER //

CREATE TRIGGER before_task_update
BEFORE UPDATE ON tasks
FOR EACH ROW
BEGIN
    IF NEW.due_date <= NOW() THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Due date cannot be in the past or present time';
    END IF;
END//



DELIMITER ;

show triggers;
select now();

-- Uji dengan due_date di masa depan
INSERT INTO tasks (id, user_id, title, description, completed, due_date, notified, created_at, updated_at) 
VALUES ("test1", "06a4745fa05f43739a573edb69cc84c4", "Test Future Date", "Testing future due_date", 0, "2024-08-14 15:00:00", 0, NOW(), NOW());

-- Uji dengan due_date di masa lalu
INSERT INTO tasks (id, user_id, title, description, completed, due_date, notified, created_at, updated_at) 
VALUES ("testNew02", "8947591baf134f069f15f6e8b2384c49", "Testhour", "Hour testing", 0, "2024-09-02 15:25:00", 1, "2024-08-23 14:27:00", NOW());

desc tasks;
select * from tasks;

select *from tasks where id = "96db4b11323b4f32a94acb01ffa7f3ee";

#end testing
select id,user_id,title,description,completed,due_date,notified,created_at,updated_at from tasks where id = ? and user_id = ?;

delete from tasks where id = "27fe54fb4cdb4ed18b8773a907540231" and user_id = "55843c058c2d4d6099e8854f83ae122c";

delete from tasks where id = ? and user_id = ?;

# atur created at menjadi wita

DELIMITER //

CREATE TRIGGER set_wita_timestamp
BEFORE INSERT ON tasks
FOR EACH ROW
BEGIN
    SET NEW.created_at = NOW();
    SET NEW.updated_at = NOW();
END//

DELIMITER ;

ALTER TABLE users MODIFY username VARCHAR(30) COLLATE utf8mb4_bin NOT NULL UNIQUE;


#end atur created at menjadi wita
SELECT @@global.time_zone, @@session.time_zone;

select *from users;
select *from tasks where user_id = "88681d19730949d8b10044ad48d75adc";
select title,description,due_date,completed,created_at from tasks where user_id = "45c46a35728f427f87f66ef915864996";

#searching
SELECT * FROM tasks 
WHERE user_id = "55843c058c2d4d6099e8854f83ae122c" 
AND (title LIKE "%date%" OR description LIKE "%no description%");

SELECT title, description, due_date, completed, created_at
FROM tasks
WHERE user_id = "88681d19730949d8b10044ad48d75adc"
AND (
  LEVENSHTEIN(title, "hari") < 3 
  OR LEVENSHTEIN(description, "hari") < 3
);

SELECT * FROM tasks WHERE due_date = DATE(NOW() + INTERVAL 1 DAY) AND notified = 0;
select *from users;

SELECT 
	u.id,
    u.email , 
    u.username , 
    t.id,
    t.title , 
    t.description , 
    t.due_date ,
    t.created_at,
    t.updated_at
FROM 
    tasks t 
JOIN 
    users u 
ON 
    t.user_id = u.id 
WHERE 
    DATE(t.due_date) = DATE(NOW() + INTERVAL 1 DAY) 
    AND t.notified = 0
    AND t.completed = 0;

select *from tasks where notified = 1;
select *from users;
delete from tasks where user_id = "8947591baf134f069f15f6e8b2384c49";

select *from tasks where user_id = "8947591baf134f069f15f6e8b2384c49";
update tasks set completed = 1 where id = "de37bc6c09164b55a0e87bb8cf8becdd";
desc tasks;

select* from users;



