# Redmine API with GoLang

## Запуск виртуальной машины

Для запуска виртуальной машины с помощью Oracle VirtualBox в фоновом режиме:

1. Запустите виртуальную машину через Oracle VirtualBox.
2. Откройте терминал в Windows и введите следующую команду для подключения к виртуальной машине:
    ```sh
    ssh user@<local-ip>
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

- [http://<local-ip>:80/](http://<local-ip>:80/) (если используется порт 80 с правами суперпользователя, иначе <http://<local-ip>:[port]/>)


## Взаимодействие с API

Для взаимодействия с Redmine API введите в терминал следующее:

```sh
go run main.go -Host "http://<local-ip>" -API_Key "<redmine-api-key>"
```
Или введите свои значения:
```sh
go run main.go -Host "YOUR_URL" -API_Key "YOUR_API_KEY"
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