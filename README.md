# About

> Тестовое задание

    Web-server, по обращении (GET-запрос) отдает данные (формируются на основе данных в БД) в измеённом формате. 
    Данные берутся из "внешнего источника" и сохраняются в БД раз в 30 секунд,
    В качестве СУБД используется Postgres.

# Usage

- Создать файл `.env` в корне проекта 
- Прописать константы (см. пример в [.exapmle_env](https://github.com/bbt-t/tech-ex-server/blob/master/.example_env))
- Запустить postrges
- Запуск через `make run` || `make && ./apiserver` ... etc
