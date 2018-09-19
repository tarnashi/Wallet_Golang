# Wallet_Golang

Тестовое задание.

Надо написать web-приложение на go, которое будет имитировать простое банковское приложение. По сути там должны быть реализованы следующие операции:

1. Создать аккаунт с определенным балансом (значение баланса должно приходить в запросе).
2. Получить баланс аккаунта по его id.
3. Перевести средства с одного аккаунта на другой.

Хранить данные приложение должно в памяти, то есть не надо подключать внешние базы данных.

Примеры запросов:
curl.exe --header "Content-Type: application/json" --request POST --data '{\"balance\":2000}' http://localhost:8181/api/add

curl.exe --header "Content-Type: application/json" --request POST --data '{\"id\":2}' http://localhost:8181/api/getBalance

curl.exe --header "Content-Type: application/json" --request POST --data '{\"senderId\":2, \"recipientId\":1, \"sum\":1555}' http://localhost:8181/api/makePayment

curl.exe --header "Content-Type: application/json" --request GET http://localhost:8181/api/showAll
