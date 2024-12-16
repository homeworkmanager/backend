FROM gomicro/goose


COPY ./migrations/*.sql ./migrations/entrypoint.sh ./


RUN echo "run migrations"
CMD ["./entrypoint.sh"]

