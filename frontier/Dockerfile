FROM php:7.4-fpm-alpine

WORKDIR /var/www/html

RUN apk update \
    && apk add --no-cache postgresql-dev libzip-dev curl-dev \
    && docker-php-ext-configure pgsql -with-pgsql=/usr/local/pgsql \
    && docker-php-ext-install \
        json \
        curl \
        zip \
        opcache \
        pdo \
        pdo_pgsql \
        pgsql
