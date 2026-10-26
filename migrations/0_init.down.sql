-- Откатываем индексы (необязательно, если удаляется вся таблица, но полезно для явности)
DROP INDEX idx_tasks_team_status ON tasks;

-- 1. Удаляем таблицы с внешними ключами, которые ни от кого больше не зависят
DROP TABLE IF EXISTS task_comments;
DROP TABLE IF EXISTS task_history;

-- 2. Удаляем таблицу задач (она зависела от команд и участников)
DROP TABLE IF EXISTS tasks;

-- 3. Удаляем таблицу участников команд (зависела от команд и пользователей)
DROP TABLE IF EXISTS team_members;

-- 4. Удаляем таблицу команд (зависела от пользователей)
DROP TABLE IF EXISTS teams;

-- 5. В последнюю очередь удаляем пользователей (главный справочник)
DROP TABLE IF EXISTS users;