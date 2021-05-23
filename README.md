# Клонер для gitlab 

## Описание
Сервис предназначен для клонирования репозитория и поддержания его 
актуальности (методом вызова fetch) при пуше в репозиторий

## Сборка

1. Создать файл VERSION с версией приложения (например 0.1.0)
2. Создать файл APP_NAME с названием образа (например krisengine/cloner)
3. Собрать и залить образ командой make push

## Запуск

Создать файл конфигурации в формате json (например config.json)
```json
{
  "gitlab_host": "gitlab.mycompany.ru",
  "server_host": "0.0.0.0",
  "server_port": "8000",
  "gitlab_user": "<user>",
  "gitlab_token": "<token>",
  "repository_dir": "/app/rep",
  "projects": [
    "<group>/<project>",
    "<group>/<project>"
  ]
}
```

Запустить контейнер 
```
docker run -p 8000:8000 -v "${PWD}/rep:/app/rep" -v "${PWD}/config.json:/app/config.json" $(APP_NAME):latest
```

смотировав в контейнер конфиг /app/config.json и папку для репозиториев /app/rep (указана в конфиге)