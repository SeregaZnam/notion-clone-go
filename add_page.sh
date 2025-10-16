#!/bin/bash

# Скрипт для добавления страницы в Notion Clone API
# Использование: ./add_page.sh [title] [iconSrc] [iconClass] [coverSrc]

# Проверяем количество аргументов
if [ $# -lt 1 ]; then
    echo "Использование: $0 <title> [iconSrc] [iconClass] [coverSrc]"
    echo "Пример: $0 \"Моя страница\" \"https://example.com/icon.png\" \"fas fa-file\" \"https://example.com/cover.jpg\""
    exit 1
fi

# Параметры
TITLE=$1
ICON_SRC=${2:-""}
ICON_CLASS=${3:-""}
COVER_SRC=${4:-""}

# URL API (можно изменить на нужный)
API_URL="http://localhost:8080/pages"

# Создаем JSON payload
JSON_PAYLOAD=$(cat <<EOF
{
  "title": "$TITLE",
  "iconSrc": "$ICON_SRC",
  "iconClass": "$ICON_CLASS",
  "coverSrc": "$COVER_SRC"
}
EOF
)

echo "Отправляем запрос на добавление страницы..."
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
# Минимальный запрос
# ./add_page.sh 1 "Тестовая страница"

# С иконкой
# ./add_page.sh 2 "Страница с иконкой" "https://cdn-icons-png.flaticon.com/512/25/25694.png"

# Полный запрос
# ./add_page.sh 3 "Полная страница" "https://example.com/icon.png" "fas fa-file" "https://example.com/cover.jpg"