# Redmine API with GoLang

## Запуск виртуальной машины

Для запуска виртуальной машины с помощью Oracle VirtualBox в фоновом режиме:

1. Запустите виртуальную машину через Oracle VirtualBox.
2. Откройте терминал в Windows и введите следующую команду для подключения к виртуальной машине:
    ```sh
    ssh user@<VM_HOST> -p 22
    ```
    Адрес и пароль лежат в локальном `.env` (`VM_HOST`, `VM_SSH_PASSWORD`), см. раздел «Конфигурация».
3. Закройте окно с виртуальной машиной, выбрав опцию "Продолжить работу в фоновом режиме".

## Запуск Redmine на виртуальной машине

Для запуска Redmine на виртуальной машине выполните следующие шаги:

1. Перейдите в директорию Redmine:
    ```sh
    cd redmine
    ```
2. Настройте окружение Ruby:
    ```sh
    export PATH="$HOME/.rbenv/bin:$PATH"
    eval "$(rbenv init -)"
    rbenv global 3.0.1
    ```
3. Запустите сервер Redmine с правами суперпользователя на основным портом:
    ```sh
    sudo /home/user/.rbenv/versions/3.0.1/bin/ruby bin/rails server -e production -u webrick -p 80 -b 0.0.0.0
    ```
    Или без них на любом другом порте:
    ```sh
    ruby bin/rails server -e production -u webrick -p 8080 -b 0.0.0.0
    ```

Для подключения к Redmine с другого устройства, откройте в браузере:

- `http://<VM_HOST>:80/` (если используется порт 80 с правами суперпользователя, иначе `http://<VM_HOST>:[port]/`)


## Работа с контейнером и пост-запросом 

Создание образа из Dockerfile:
```sh
docker buildx build --platform linux/amd64 -t stilusoff/redmine-api:latest .     
```

Пуш образа:
```sh
docker push stilusoff/redmine-api:latest
```

Пул образа:
```sh
docker pull stilusoff/redmine-api:latest
```

Запуск docker-контейнера (переменные `HOST` и `API_Key` берутся из `.env`):

```sh
docker run -p [порт на локальной машине]:[порт внутри контейнера для перенаправления] --env-file .env stilusoff/redmine-api:latest
```

Пример запуска:

```sh
docker run -p 8080:8080 --env-file .env stilusoff/redmine-api:latest
```

Примеры запросов (вместо localhost можно ввести ip-адресс; `API_Key` — из `.env`):

- Список проектов:
    ```sh
    source .env
    curl -X GET -H "API_KEY: $API_Key" http://localhost:8080/view_projects
    ```

- Создать задачу:
    ```sh
    source .env
    curl -X POST -H "API_KEY: $API_Key" "http://localhost:8080/create_issue?ProjectID=1&Subject=New%20Issue%20from%20Go&Description=This%20is%20a%20test%20issue%20created%20from%20Go&TrackerID=1"
    ```


## Взаимодействие с Redmine

Учётные данные админ-панели Redmine — в `.env` (`REDMINE_ADMIN_USER`, `REDMINE_ADMIN_PASSWORD`).

## Информация о базе данных

Для доступа к базе данных MySQL выполните следующие команды:

1. Откройте терминал и введите (пароль — `MYSQL_PASSWORD` из `.env`):
    ```sh
    mysql -u <MYSQL_USER> -p
    ```
2. Имя базы данных: `myredminedb`

## Конфигурация

Все адреса, ключи и пароли хранятся в файле `.env`, который не коммитится. Шаблон — `.env.example`:

```sh
cp .env.example .env
# заполните значения
```

## Все команды терминала
```sh
cd redmine
export PATH="$HOME/.rbenv/bin:$PATH"
eval "$(rbenv init -)"
rbenv global 3.0.1
sudo /home/user/.rbenv/versions/3.0.1/bin/ruby bin/rails server -e production -u webrick -p 80 -b 0.0.0.0
```
