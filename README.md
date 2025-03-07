# NFT Processing Service

Этот проект предоставляет API для обработки транзакций NFT, включая получение и отправку NFT через блокчейн. 

## Переменные окружения

Для работы сервиса необходимо задать следующие обязательные переменные окружения:

- `PORT`: Номер порта, который сервер будет слушать. По умолчанию `18300`, если не задано.
  
- `HOST`: Хост или IP-адрес, который сервер будет использовать. По умолчанию `0.0.0.0`.

- `SECRET`: Секретный ключ, используемый для аутентификации. Значение должно быть задано в окружении.

- `CALLBACK_URL`: URL-адрес для отправки обратных вызовов. Если не задан, по умолчанию пустая строка.

- `BLOCKCHAIN_WALLET`: Адрес блокчейн-кошелька который отслеживается и на который возвращается газ. По умолчанию пустая строка.

- `BLOCKCHAIN_NETWORK`: URL для конфигурации блокчейн-сети. Значение по умолчанию: `https://ton.org/global.config.json`.

- `BLOCKCHAIN_WORDS`: Массив с seed-словами, используемыми для восстановления или генерации кошелька. Значение по умолчанию: пустой срез. Выделено из переменной окружения, разделенной запятыми.

## Схема событий

### Поступление NFT на кошелек

Пример события при поступлении NFT на кошелек:

```json
{
  "address": "EQCXZ2WvHnv8hYZC5buYrGZrE9Ii51SMP-RL1gVYOv2pPKwj",
  "collection": "EQB6AtBPOuTtQml8oSA7X8ZqJ5QmcOYYqoz92sQYXGUQrxyB",
  "message": "",
  "meta_data": {
    "attributes": [
      {
        "trait_type": "Model",
        "value": "Red Mirage"
      },
      {
        "trait_type": "Backdrop",
        "value": "Pacific Green"
      },
      {
        "trait_type": "Symbol",
        "value": "Octopus"
      }
    ],
    "description": "An exclusive Hex Pot with the appearance Red Mirage on a Pacific Green background with Octopus icons.",
    "image": "https://nft.fragment.com/gift/hexpot-14619.webp",
    "lottie": "https://nft.fragment.com/gift/hexpot-14619.lottie.json",
    "name": "Hex Pot #14619"
  },
  "owner": "EQAno5PEMnsMt26bPgnXeFMOBVzSNHor2ctgyALrg3oD5fy_",
  "tx_hash": "VicLITs1cAY8JaEGuFKJx2svhRHkxFRZXVou6FPROUU=",
  "type": "received_nft"
}
```

`meta_data` может содержать разные данные и иметь иную структуру.

### Отчет об успешном переводе NFT

Пример события при успешном переводе NFT:

```json
{
  "address": "EQCXZ2WvHnv8hYZC5buYrGZrE9Ii51SMP-RL1gVYOv2pPKwj",
  "tx_hash": "cgG7tDykSVeevWFhfDKGgtl3oXGHmG3So+Ig+f56DcY=",
  "type": "success_send_nft"
}
```

## Использование API

### Аутентификация

Все API-запросы должны включать `SECRET` для проверки подлинности. Это можно сделать следующими способами:

- **Заголовок авторизации:**

  Включите `SECRET` в заголовок авторизации HTTP:

  ```http
  Authorization: your_secret_key
  ```

- **Параметр запроса:**

  Включите `SECRET` как параметр запроса в URL:

  ```http
  http://<HOST>:<PORT>/withdraw?secret=your_secret_key
  ```

### Запрос на вывод NFT

- **Маршрут:** `POST /nft.send`
- **Тело запроса:**

  В теле запроса должны быть указаны следующие поля JSON:

  ```json
  {
    "address": "EQCXZ2WvHnv8hYZC5buYrGZrE9Ii51SMP-RL1gVYOv2pPKwj",
    "owner": "UQB13vfb4jgMP7XLmu7FXOPhKYKlNi0qHeAQgsPB5w9LOxxB",
    "message": "Transaction message" // необязательное сообщение для транзакции
  }
  ```

### Успешный ответ

Изменен формат успешного ответа от сервера на запрос вывода средств. Теперь ответ будет выглядеть так:

```json
{
  "response": {
    "result": true
  }
}
```

### Ответ с ошибкой

Когда возникает ошибка, ответ будет содержать объект `error` в следующем формате:

- **Формат ошибки:**

  ```json
  {
    "error": {
      "code": 123,
      "message": "Error message",
      "critical": true
    }
  }
  ```
```