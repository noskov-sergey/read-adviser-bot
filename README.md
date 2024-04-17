# Read-Adviser-Bot
Бот умеет сохранять ссылки, которые ему скидывают собеседники, и по запросу отправлять случайную ссылку из сохраненных.
Это полезно для тех людей, которые часто сохраняют много статей, но забывают их читать :)

## Комментарий
Код написан по обучающему видео Николая Тузова, для практики в понимании структурированноый архитектуры с разбиением на клиентов и возможностью их подмены через интерфейсы.
* использована стандарнтые пакеты библиотеки.

## Как развернуть проект

Клонируйте репозиторий:

```git clone git@github.com:noskov-sergey/read-adviser-bot```

Запустите проект, в тег -tg-token необходимо передать токен канала

`go run examples/backtesting/main.go -tg-token=X1X2X3X4X:AAAAAAAAA`
