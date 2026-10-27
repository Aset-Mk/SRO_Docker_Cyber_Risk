# Этап 1: сборка статического бинарника
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /passcheck .

# Этап 2: пустой образ только с бинарником
FROM scratch
LABEL maintainer="mukashevasset" description="CLI-проверка стойкости паролей"
COPY --from=build /passcheck /passcheck
USER 65534
ENTRYPOINT ["/passcheck"]
