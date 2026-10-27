# СРО: контейнеризация приложений (Docker)

Мукашев Асет, СИБ-34

Два консольных приложения на тему кибер-рисков, упакованные в Docker-образы
и опубликованные на Docker Hub под аккаунтом `mukashevasset`.

| Приложение | Язык | Docker Hub |
|---|---|---|
| `risk-calc` — калькулятор кибер-рисков (вероятность × ущерб, реестр и матрица 5×5) | Python | [mukashevasset/risk-calc](https://hub.docker.com/r/mukashevasset/risk-calc) |
| `passcheck` — проверка стойкости паролей (энтропия, словарь утёкших, время перебора) | Go | [mukashevasset/passcheck](https://hub.docker.com/r/mukashevasset/passcheck) |

## Запуск из Docker Hub

```bash
docker run --rm mukashevasset/risk-calc:1.0
docker run --rm mukashevasset/risk-calc:1.0 --min 15        # только критические

docker run --rm mukashevasset/passcheck:1.0 qwerty 'P@ssw0rd!2026'
printf 'admin\nTr0ub4dor&3\n' | docker run --rm -i mukashevasset/passcheck:1.0
```

Свой реестр угроз (CSV `asset,threat,likelihood,impact`):

```bash
docker run --rm -v "$PWD/my.csv:/data.csv:ro" mukashevasset/risk-calc:1.0 /data.csv
```

## Сборка локально

```bash
docker build -t mukashevasset/risk-calc:1.0 risk-calc
docker build -t mukashevasset/passcheck:1.0 passcheck
```

## Особенности образов

- `risk-calc` — на `python:3.13-alpine`, запуск от непривилегированного пользователя `app`.
- `passcheck` — multi-stage сборка: компиляция в `golang:1.23-alpine`, итоговый образ
  `scratch` (~2 МБ, только статический бинарник), запуск от `nobody` (UID 65534).
