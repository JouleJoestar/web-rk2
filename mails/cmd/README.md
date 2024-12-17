1. Отправка письма
Метод: POST
URL: http://<ваш_сервер>:<порт>/mails
Тело запроса (JSON):
{
    "theme": "Hello World",
    "text": "This is a test email.",
    "image": "http://example.com/image.png",
    "sender_id": 1,
    "receivers": [2, 3]
}

2. Получение списка писем для пользователя
Метод: GET
URL: http://localhost:8089/mails/<user_id>

3. Удаление письма
Метод: DELETE
URL: http://localhost:8089/mails/<mail_id>
