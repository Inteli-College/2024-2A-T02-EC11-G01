from pydantic import BaseModel, field_validator, model_validator
import json
from pydantic_settings import BaseSettings
from dotenv import load_dotenv

load_dotenv()


class RabbitMQConfig(BaseModel):
    host: str
    username: str
    password: str
    queue: str


class Settings(BaseSettings):
    rabbitmq_config: RabbitMQConfig | None = None

    @field_validator("rabbitmq_config", mode="before")
    def load_from_json(cls, values):
        env_json = values.get("RABBITMQ_CONFIG")
        if env_json:
            values["rabbitmq_config"] = RabbitMQConfig(**json.loads(env_json))
        return values

    class Config:
        env_file = ".env"


settings = Settings()
