# Redmine API with GoLang

## Запуск виртуальной машины

Для запуска виртуальной машины с помощью Oracle VirtualBox в фоновом режиме:

1. Запустите виртуальную машину через Oracle VirtualBox.
2. Откройте терминал в Windows и введите следующую команду для подключения к виртуальной машине:
    ```sh
    ssh user@<VM_HOST> -p 22
    ```
    Пароль: `<ssh-password>`
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

- [http://<VM_HOST>:80/](http://<VM_HOST>:80/) (если используется порт 80 с правами суперпользователя, иначе <http://<VM_HOST>:[port]/>)


## Взаимодействие с API

Для взаимодействия с Redmine API введите в терминал следующее:

```sh
go run main.go -HOST "http://<VM_HOST>" -API_Key "<redmine-api-key>"
```
Или:
```sh
go run main.go -HOST "http://localhost:80" -API_Key "<redmine-api-key>"
```
Также можно ввести свои значения:
```sh
go run main.go -HOST "YOUR_URL" -API_Key "YOUR_API_KEY"
```

## Работа с контейнером и пост-запросом 

Создание образа из Dockerfile:
```sh
docker build -t redmine-api .
```

Запуск docker-контейнера:

```sh
docker run -p [порт на локальной машине]:[порт внутри контейнера для перенаправления] -e HOST=[хост, на котором контейнер будет слушать] -e API_KEY=[API-ключ redmine, находиться в личном кабинете пользователя в redmine] stilusoff/redmine-api
```

Пример запуска:

```sh
docker run -p 8080:8080 -e HOST="http://<VM_HOST>" -e API_Key="<redmine-api-key>" redmine-api
```

Примеры запросов:

```sh
curl -X POST http://localhost:8080/viwe_tasks
```

```sh
curl -X POST http://localhost:8080/viwe_tasks
```


## Взаимодействие с Redmine

Для входа в админ-панель Redmine используйте следующие учетные данные:

- Пользователь: `<admin-user>`
- Пароль: `<admin-password>`

## Информация о базе данных

Для доступа к базе данных MySQL выполните следующие команды:

1. Откройте терминал и введите:
    ```sh
    mysql -u <mysql-user> -p
    ```
    Пароль: `<mysql-password>`
2. Имя базы данных: `myredminedb`

## Все команды терминала
```sh
cd redmine
export PATH="$HOME/.rbenv/bin:$PATH"
eval "$(rbenv init -)"
rbenv global 3.0.1
sudo /home/user/.rbenv/versions/3.0.1/bin/ruby bin/rails server -e production -u webrick -p 80 -b 0.0.0.0
```