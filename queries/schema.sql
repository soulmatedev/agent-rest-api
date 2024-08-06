CREATE TABLE Agent (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL
);

-- Создание таблицы для хранения информации о заданиях
CREATE TABLE Task (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    agent_id INT,
    FOREIGN KEY (agent_id) REFERENCES Agent(id) ON DELETE CASCADE
);
