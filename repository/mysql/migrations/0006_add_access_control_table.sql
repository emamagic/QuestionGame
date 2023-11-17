-- +migrate Up
CREATE TABLE `access_controls` (
                                `id` int primary key AUTO_INCREMENT,
                                 int NOT N`actor_id`ULL,
                                `actor_type` ENUM('role', 'user') NOT NULL,
                                `permission_id` INT NOT NULL,
                                `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                FOREIGN KEY (`permission_id`) REFERENCES `permissions`(`id`)
);

-- +migrate Down 
DROP TABLE access_controls;