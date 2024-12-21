Скрипт создания бд

CREATE TABLE "tasks" (
    id SERIAL PRIMARY KEY,
    author_name VARCHAR(100) NOT NULL,
    assignee_name VARCHAR(100),
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    resolved_date TIMESTAMP,
    status VARCHAR(20) CHECK (status IN ('new', 'in progress', 'done')) NOT NULL
);

1. Создание задачи
Метод: POST
URL: http://localhost:8081/tasks
Тело запроса (JSON):
{
    "author_name": "John Doe",
    "assignee_name": "Jane Smith",
    "status": "new"
}

2. Получение списка задач
Метод: GET
URL: http://localhost:8081/tasks

3. Обновление статуса задачи
Метод: PUT
URL: http://localhost:8081/tasks/status?id=1
Тело запроса (JSON):
{
    "status": "in progress"
}