# Deep Work App

## Описание
Deep Work App - это приложение, предназначенное для повышения продуктивности пользователей путем управления сессиями глубокой работы. Оно позволяет пользователям сосредоточиться на задачах, минимизируя отвлекающие факторы и отслеживая их продуктивность.

## Структура проекта
Проект организован по принципам DDD (Domain-Driven Design) и разделен на несколько ключевых компонентов:

- **cmd/deepwork**: Точка входа в приложение, инициализирует приложение и запускает интерфейс командной строки.
- **internal/app**: Основная логика приложения, включая команды и запросы.
- **internal/config**: Управление конфигурацией приложения.
- **internal/domain**: Основная доменная логика, включая управление сессиями глубокой работы.
- **internal/infrastructure**: Реализация инфраструктурных компонентов, таких как хранилище данных и сервисы.
- **internal/interfaces**: Интерфейсы для взаимодействия с пользователем, включая CLI.
- **pkg/utils**: Утилиты для работы с файлами и временем.

## Установка
1. Клонируйте репозиторий:
   ```
   git clone <URL>
   cd deep-work-app
   ```

2. Установите зависимости:
   ```
   go mod tidy
   ```

## Использование
Запустите приложение с помощью команды:
```
go run cmd/deepwork/main.go start
```
Для остановки сессии используйте:
```
go run cmd/deepwork/main.go stop
```

## Логирование
Логи приложения сохраняются в файл `logs/deep_work.log`.

## Конфигурация
Конфигурация приложения хранится в файле `config/config.json`. Вы можете настроить параметры, такие как продолжительность сессий, настройки рабочего пространства и т.д.

## Вклад
Если вы хотите внести свой вклад в проект, пожалуйста, создайте запрос на слияние с описанием ваших изменений.

## Лицензия
Этот проект лицензирован под MIT License.