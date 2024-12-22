FROM gomicro/goose


COPY ./migrations/*.sql ./migrations/entrypoint.sh ./
RUN chmod a+x ./entrypoint.sh


RUN echo "run migrations"
CMD ["bash", "./entrypoint.sh"]

