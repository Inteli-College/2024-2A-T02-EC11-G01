FROM rabbitmq:3.12-management

ENV RABBITMQ_SERVER_ADDITIONAL_ERL_ARGS="-rabbitmq_management load_definitions \"/etc/rabbitmq/definitions.json\""

ADD --chown=rabbitmq definitions.json /etc/rabbitmq/
ADD --chown=rabbitmq rabbitmq.conf /etc/rabbitmq/

RUN rabbitmq-plugins enable rabbitmq_top

EXPOSE 5672 15672

CMD ["rabbitmq-server"]