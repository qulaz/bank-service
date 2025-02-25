#!/bin/bash

# ID счета отправителя
FROM_ACCOUNT_ID=1
# ID счета получателя
TO_ACCOUNT_ID=2
# Сумма перевода
AMOUNT=10000
# Количество параллельных запросов
NUM_REQUESTS=10

# Проверка баланса до переводов
curl -s http://localhost:8080/api/accounts/$FROM_ACCOUNT_ID
echo "--------------------------------"

# Запускаем параллельные запросы
seq 1 $NUM_REQUESTS | xargs -P $NUM_REQUESTS -I {} curl -s -X POST http://localhost:8080/api/transactions/transfer \
    -H "Content-Type: application/json" \
    -d "{
      \"from_account_id\": $FROM_ACCOUNT_ID,
      \"to_account_id\": $TO_ACCOUNT_ID,
      \"amount\": $AMOUNT,
      \"description\": \"Test transfer {}\"
    }"

echo "All transfers completed"
echo "--------------------------------"
# Проверка баланса после переводов
curl -s http://localhost:8080/api/accounts/$FROM_ACCOUNT_ID
