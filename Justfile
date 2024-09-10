[private]
default:
  @just --choose

infra option='':
  @if [ '{{option}}' = 'reset' ]; then \
    docker compose -f ./docker-compose.yml down -v postgres migrations rabbitmq ; \
  fi

  @docker compose -f ./docker-compose.yml up postgres migrations rabbitmq -d --build

clean:
  @docker compose -f ./docker-compose.yml down -v

app option='':
  @if [ '{{option}}' = 'local-infra' ]; then \
    docker compose -f ./docker-compose.yml up postgres migrations rabbitmq -d ; \
  fi

  @go run ./backend/cmd/app/main.go
