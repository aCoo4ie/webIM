CREATE TABLE `user` (
	`id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT 'Primary Key, Auto Increment',
	`username` VARCHAR(50) NOT NULL UNIQUE COMMENT 'Username, unique and cannot be null',
	`password` VARCHAR(255) NOT NULL COMMENT 'Password, cannot be null',
	`mobile` VARCHAR(20) UNIQUE COMMENT 'Mobile number, unique',
	`avatar` VARCHAR(255) COMMENT 'Avatar URL',
	`sex` VARCHAR(10) COMMENT 'Gender',
	`nickname` VARCHAR(50) COMMENT 'Nickname',
	`salt` VARCHAR(50) COMMENT 'Salt for password hashing',
	`token` VARCHAR(255) COMMENT 'Authentication token',
	`online` INT DEFAULT 0 COMMENT 'Online status, 0 for offline, 1 for online',
	`memo` TEXT COMMENT 'Additional notes or remarks',
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Creation time',
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Last update time',
	`deleted_at` DATETIME DEFAULT NULL COMMENT 'Deletion time for soft delete'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='User table';

-- Insert some common content
INSERT INTO `user` (`username`, `password`, `mobile`, `avatar`, `sex`, `nickname`, `salt`, `token`, `online`, `memo`, `created_at`, `updated_at`) VALUES
('john_doe', 'hashed_password_1', '1234567890', 'https://example.com/avatar1.png', '男', 'John', 'random_salt_1', 'token_1', 1, 'Test user 1', NOW(), NOW()),
('jane_doe', 'hashed_password_2', '0987654321', 'https://example.com/avatar2.png', '女', 'Jane', 'random_salt_2', 'token_2', 0, 'Test user 2', NOW(), NOW()),
('alex_smith', 'hashed_password_3', '1122334455', 'https://example.com/avatar3.png', '未知', 'Alex', 'random_salt_3', 'token_3', 1, 'Test user 3', NOW(), NOW());