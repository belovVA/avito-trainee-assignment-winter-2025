FROM postgres:13

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /docker-entrypoint-initdb.d/

# Копируем SQL-скрипты и файлы миграций в контейнер
# COPY migrations/ /docker-entrypoint-initdb.d/

# Указываем стандартные переменные окружения (могут быть переопределены в docker-compose)
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=shop

# Открываем порт для подключения к базе данных
EXPOSE 5432
