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


## Работа с контейнером и пост-запросом 

Создание образа из Dockerfile:
```sh
docker buildx build --platform linux/amd64 -t stilusoff/redmine-api:latest .     
```

Пуш образа:
```sh
docker push stilusoff/redmine-api:latest
```

Пул образа
```sh
docker pull stilusoff/redmine-api:latest
```

Запуск docker-контейнера:

```sh
docker run -p [порт на локальной машине]:[порт внутри контейнера для перенаправления] -e HOST=[хост, на котором контейнер будет слушать] -e API_KEY=[API-ключ redmine, находиться в личном кабинете пользователя в redmine] stilusoff/redmine-api:latest
```

Пример запуска:

```sh
docker run -p 8080:8080 -e HOST="http://<VM_HOST>:80" -e API_Key="<redmine-api-key>" stilusoff/redmine-api:latest
```

Или:

```sh
docker run -p 8080:8080 -e HOST="http://<local-ip>:80" -e API_Key="<redmine-api-key>" stilusoff/redmine-api:latest
```

Примеры запросов (вместо localhost можно ввести ip-адресс):

Список проектов:
```sh
curl -X GET -H "API_KEY: <redmine-api-key>" http://localhost:8080/view_projects

```

Создать задачу:
```sh
curl -X POST -H "API_KEY: <redmine-api-key>" "http://localhost:8080/create_issue?ProjectID=1&Subject=New%20Issue%20from%20Go&Description=This%20is%20a%20test%20issue%20created%20from%20Go&TrackerID=1"

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