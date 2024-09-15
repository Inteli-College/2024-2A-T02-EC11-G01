#!/bin/bash
until curl -u user:password -s http://rabbitmq:15672/api/overview; do
  echo "Esperando RabbitMQ..."
  sleep 5
done

curl -i -u user:password -H "content-type:application/json" \
    -XPUT -d'{"type":"topic","durable":true}' \
    http://rabbitmq:15672/api/exchanges/%2f/predictions

exec "$@"