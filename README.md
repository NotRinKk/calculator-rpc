# Микросервисный калькулятор на Golang с RPC

[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://golang.org/)
[![Docker](https://img.shields.io/badge/docker-ready-blue.svg)](https://www.docker.com/)


## Содержание

- [Описание](#описание)
- [Архитектура](#архитектура)
- [Запуск проекта](#запуск-проекта)
- [Использование](#использование)
- [API](#api)

## Описание

Микросервисный калькулятор, где каждая арифметическая операция выполняется отдельным сервисом. Основные компоненты:

- **Основной сервис** (`calculator-service`): Принимает запросы и распределяет их между специализированными сервисами
- **Сервисы операций**: Отдельные микросервисы для каждой арифметической операции

## Архитектура
calculator-rpc/  
├── calculator-service/&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Основной сервис-агрегатор  
├── addition-service/&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Сервис сложения  
├── subtraction-service/&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Сервис вычитания  
├── multiplication-service/&nbsp;# Сервис умножения  
├── division-service/&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Сервис деления  
├── square-root-service/&nbsp;&nbsp;&nbsp;&nbsp;# Сервис вычисления корня  
├── percentage-service/&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Сервис вычисления процентов  
├── rounding-service/&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Сервис округления  
└── power-service/&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;# Сервис возведения в степень  

## Запуск проекта

### Требования

- Docker и Docker Compose
- Go 1.22+

### Инструкция

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/NotRinKk/calculator-rpc.git
   cd calculator-rpc
   ```
2. Запустите сервисы:
   ```bash
   docker-compose up --build
   ```
3. Сервис будет доступен на http://localhost:8080

## Использование
Отправьте POST-запрос на /calculate с JSON-телом:
```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{
    "operations": [
        {"type": "addition", "a": 5, "b": 3},
        {"type": "subtraction", "a": 10, "b": 4},
        {"type": "multiplication", "a": 6, "b": 7},
        {"type": "division", "a": 20, "b": 5},
        {"type": "square_root", "value": 25},
        {"type": "percentage", "value": 200, "percent": 10},
        {"type": "rounding", "value": 3.14159, "places": 2},
        {"type": "power", "value": 2, "power": 3}
    ]
}'
```
Пример ответа:
```bash
{
  "results": [8, 32, 16],
  "finalResult": 16,
  "error": null
}
```
## API
Эндпоинт
POST /calculate

Формат запроса
```typescript
interface Operation {
  type: 'addition' | 'subtraction' | 'multiplication' | 'division' 
       | 'square_root' | 'percentage' | 'rounding' | 'power';
  a?: number;       // Первый операнд
  b?: number;       // Второй операнд
  value?: number;   // Значение (для унарных операций)
  places?: number;  // Знаков после запятой (для округления)
  percent?: number; // Процент
  power?: number;   // Степень
}

interface Request {
  operations: Operation[];
}
```

Формат ответа
```typescript
interface Response {
  results: number[];    // Результаты каждой операции
  finalResult: number;  // Итоговый результат
  error?: string;       // Сообщение об ошибке
}
```
