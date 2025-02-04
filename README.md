# Сервис публикации постов

# Использование

**Сборка**

Из корневой директории проекта:

`$ docker build -t ozon-project .`

**Запуск**

Из корневой директории проекта:

`$ docker run -d -p 8080:8080 --name ozon-project_container ozon-project`

**Переменные окружения**

Все переменные окружения по умолчанию заданы в докерфайлах. Можно их изменить либо внутри файлов, либо при запуске докер контейнера, пример:

`$ docker run -d -p 8080:8080 -e STORAGE_TYPE=inmem --name ozon-project_container ozon-project`

 По умолчанию в качестве хранилища используется БД.