## Проект «Телеграм бот для поздравлений с днем рождения»

## Описание проекта 

Проект представляет собой сервис,
который взаимодействует с Telegram API
и реализуя backend бота предоставляет ответы
на далее перечисленные команды. Структурно проект разделен на 3 слоя:  
  - Слой работы с базой данных (директория repository)
  - Сервисный слой (директория service)
  - Слой взаимодействия с API Telegram (директория app)  

Для возможного в будущем изменения реализации в service и repository есть по одному интерфейсу в котором указанны методы которые должен реализовывать сервисный и бд слой соответственно.
В слое работы с API определяется реализация и создаются экземпляры сервисного и бд слоя, а так же выбирается логгер для сервисного слоя.  
Модели объектов для взаимодействия с базой данных хранятся в директории domain.  
С учетом того что бот не рассчитан на широкое (более 100 человек) использование
,а так же в целях упрощения разработки и тестирования, было принято решение использовать в качестве базы данных SQLite. Файл базы данных храниться в папке storage. Юнит тесты для базы данных хранятся в папке test.  
Так же в папке migrations хранятся файлы для миграций. Сами миграции применяются, либо с помощью запуска через консоль IDE, либо через Taskfile.yaml, командой - task.
## Запуск приложения

Запуск приложения осуществляется при помощи запуска консольной программы - go run ./cmd/main.go. Перед этим стоит применить миграции одним из вышеописанных способов.

## Команды для взаимодействия с ботом

/add - Добавить нового пользователя не заристрированного в боте ранее  
/all - Посмотреть список всех пользователей а так же дни их рождения  
/sub - Подписаться на уведомления о дне рождения пользователя  
/tg_id - Посмотреть когда день рождения у пользователя с указанным тг_юзернеймом  
/subs - Посмотреть на чьи дни рождения ты подписан  
/unsub - Отписаться от уведомления о дне рождения пользователя   

@Username бота - @Happy_Birthday_m_Bot

