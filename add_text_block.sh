#!/bin/bash

# Скрипт для добавления текстового блока в Notion Clone API
# Использование: ./add_text_block.sh [text] [page_id] [order] [type]

# Проверяем количество аргументов
if [ $# -lt 4 ]; then
    echo "Использование: $0 <text> <page_id> <order> <type>"
    echo "Пример: $0 \"Привет, мир!\" 1 1 \"paragraph\""
    echo "Типы: paragraph, heading, list_item, code_block"
    exit 1
fi

# Параметры
TEXT=$1
PAGE_ID=$2
ORDER=$3
TYPE=$4

# URL API (можно изменить на нужный)
API_URL="http://localhost:8080/text-blocks"

# Создаем JSON payload
JSON_PAYLOAD=$(cat <<EOF
{
  "text": "$TEXT",
  "page_id": $PAGE_ID,
  "order": $ORDER,
  "type": "$TYPE"
}
EOF
)

echo "Отправляем запрос на добавление текстового блока..."
echo "URL: $API_URL"
echo "Данные: $JSON_PAYLOAD"
echo ""

# Выполняем curl запрос
curl -X POST "$API_URL" \
  -H "Content-Type: application/json" \
  -d "$JSON_PAYLOAD" \
  -w "\n\nHTTP Status: %{http_code}\n" \
  -s

echo ""

# Примеры вызова:
# Добавляем заголовок
# ./add_text_block.sh "Моя страница" 1 1 "heading"

# Добавляем параграф
# ./add_text_block.sh "Это содержимое страницы" 1 2 "paragraph"

# Добавляем элемент списка
# ./add_text_block.sh "Первый пункт" 1 3 "list_item"