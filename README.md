# Homework 2. Market bot

## Основная идея
Бот с реализацией мониторинга стоимости из разных финансовых инструментовты: валюты, акции, фонды и криптовалюты с пространством для дальнейшего подключения торговых ботов:
- API handler будет собирать данные о состояннии настоящей стоимости портфеля
- *Коммуникатор будет обрабатывать эти данные*
- Вьюшка будет собирать данные (с апи или коммуникатора) с коммуникатора и присылать данные о портфеле потребителю

## Вью
В качестве интерфейса для конечного потребителя будет использоваться телеграм бот и/или небольшой сайт с информацией и графиками.


## Реализация
В реализации планируется использовать несколько различных модулей. Из-за малого опыта разработки больших проектов я решил представить на обсуждение две различных архитектуры.


### 1. Делигируем по максимуму
``` mermaid
graph TD;
    TinkoffAPI --gRPC--> TinkoffAPI_Hand  
    TinkoffAPI_Hand --HTTP---> TelegramBot 
    TinkoffAPI_Hand ---> any_view 

    
    any_api --HTTP--> any_api_Hand  
    any_api_Hand --HTTP---> TelegramBot 
    any_api_Hand ---> any_view 
```
<!-- BinanceAPI --HTTP-- BinanceAPI_Hand  
BinanceAPI_Hand --HTTP--- TelegramBot 
BinanceAPI_Hand --gRPC--- WEBSite 
BinanceAPI_Hand --- any_view -->
<!-- TinkoffAPI_Hand --gRPC--- WEBSite  -->
<!-- any_api_Hand --gRPC--- WEBSite  -->
#### **Плюсы**
+ Можно находить индивидуальных подход к каждому сервису.
+ В реализации MVP будет меньше работы
#### **Минусы**
Напрашиваются сами собой
- Очень сложная структура после потенциального масштабирования
- На каждой вьюшке для каждого сервиса после масштабирования нужно будет делать отдельный бработчик данных

### 2. Связной между view и api хендлерами
``` mermaid
graph TD;
    TinkoffAPI --gRPC--> TinkoffAPI_Hand
    TinkoffAPI_Hand --gRPC--> Communicator


    any_api --> any_api_hand
    any_api_hand --> Communicator

    Communicator --gRPC--> TelegramBot_Hand
    Communicator --> any_view
```
<!-- BinanceAPI --HTTP-- BinanceAPI_Hand
BinanceAPI_Hand --gRPC-- Communicator -->
<!-- Communicator --gRPC-- WEBSite -->
#### **Плюсы**
+ Меньше отдельных модулей после потенциального масштабирования
+ В будущем нужно будет придумывать отдельный хендлер для каждой апишки на вьюшках
#### **Минусы**
- В реализации MVP нужно будет писать больше модулей
- В будущем нужно унифицировать подход к каждому сервису. На мой взгляд, с достаточно разными по структуре сервисами, это будет сделать сложно

## Дальнейшее масштабирование
### Финансовые площадки

- Binance 
- всё, что угодно

### Вьюшки

- Сайт
- Приложение
- Боты в других мессенджерах

### Торговые боты

Боты, которые будут торговать (спекулировать) на биржах без или с минимальным вмешательством пользователя.
## Заключение

MVP может быть реализован с помощью всего двух модулей: TinkofAPI_Hand и TelegramBot_Hand, но с потенциалом масштабирования такая модель вряд-ли будет работать, поэтому, на мой взгляд, лучше сразу реализовать модуль коммуникатора, чтобы оставить пространство для добавления других финансовых инструментов, вьюшек и торговых ботов без необходимости переписывать весь проект. 


