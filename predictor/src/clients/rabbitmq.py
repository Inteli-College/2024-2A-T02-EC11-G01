import pika
from fastapi import Depends

class RabbitMQClient:
    def __init__(self, amqp_url: str):
        print(f"Connecting to RabbitMQ server with URL {amqp_url}")

        self.connection_params = pika.URLParameters(amqp_url)
        self.connection = pika.BlockingConnection(self.connection_params)
        self.channel = self.connection.channel()
        self.channel.queue_declare(queue='prediction', durable=True)

    def publish_message(self, message: str):
        self.channel.basic_publish(
            exchange='',
            routing_key='prediction',  # Use a fila apropriada
            body=message,
            properties=pika.BasicProperties(
                delivery_mode=2,  # Persistência da mensagem
            )
        )
        print(f"Message published: {message}")

    def close_connection(self):
        self.connection.close()

    def __del__(self):
        self.close_connection()