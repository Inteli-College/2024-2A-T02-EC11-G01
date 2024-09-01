import pika
from fastapi import Depends

class RabbitMQClient:
    def __init__(self, host: str, queue: str, username: str, password: str):
        print(f"Connecting to RabbitMQ server at {host}")

        self.host = host
        self.queue = queue
        self.credentials = pika.PlainCredentials(username, password)
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(host=self.host, credentials=self.credentials)
        )
        self.channel = self.connection.channel()
        self.channel.queue_declare(queue=self.queue, durable=True)

    def publish_message(self, message: str):
        self.channel.basic_publish(
            exchange='',
            routing_key=self.queue,
            body=message,
            properties=pika.BasicProperties(
                delivery_mode=2,
            )
        )
        print(f"Message published: {message}")

    def close_connection(self):
        self.connection.close()

    def __del__(self):
        self.close_connection()
