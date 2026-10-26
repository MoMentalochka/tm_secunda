-- Таблица пользователей
CREATE TABLE IF NOT EXISTS users (
 id INT AUTO_INCREMENT PRIMARY KEY,
 email VARCHAR(255) NOT NULL UNIQUE,
 password_hash VARCHAR(255) NOT NULL,
 name VARCHAR(255) NOT NULL,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Таблица команд
CREATE TABLE IF NOT EXISTS teams (
 id INT AUTO_INCREMENT PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 created_by INT NOT NULL,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Таблица участников команд (Связь многие-ко-многим)
CREATE TABLE IF NOT EXISTS team_members (
 team_id INT NOT NULL,
 user_id INT NOT NULL,
 role ENUM('owner', 'admin', 'member') NOT NULL DEFAULT 'member',
 PRIMARY KEY (team_id, user_id), -- Гарантирует уникальность пары
 FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
 FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Таблица задач
CREATE TABLE IF NOT EXISTS tasks (
 id INT AUTO_INCREMENT PRIMARY KEY,
 team_id INT NOT NULL,
 title VARCHAR(255) NOT NULL,
 description TEXT,
 status ENUM('todo', 'in_progress', 'done') NOT NULL DEFAULT 'todo',
 created_by INT NOT NULL,
 assignee_id INT DEFAULT NULL,
 version INT NOT NULL DEFAULT 1, -- Для защиты от конкурентного обновления
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
 closed_at TIMESTAMP NULL DEFAULT NULL,
 FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
 FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE RESTRICT,
 FOREIGN KEY (assignee_id) REFERENCES team_members(user_id) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Индекс для ускорения фильтрации и кеширования списков задач
CREATE INDEX idx_tasks_team_status ON tasks(team_id, status);

-- Таблица истории изменений
CREATE TABLE IF NOT EXISTS task_history (
 id INT AUTO_INCREMENT PRIMARY KEY,
 task_id INT NOT NULL,
 changed_by INT NOT NULL,
 changes JSON NOT NULL, -- Храним старые и новые значения в формате {"title": {"old": "X", "new": "Y"}}
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
 FOREIGN KEY (changed_by) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Таблица комментариев
CREATE TABLE IF NOT EXISTS task_comments (
 id INT AUTO_INCREMENT PRIMARY KEY,
 task_id INT NOT NULL,
 user_id INT NOT NULL,
 content TEXT NOT NULL,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
 FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;