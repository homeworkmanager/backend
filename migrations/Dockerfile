FROM gomicro/goose


COPY ./migrations/*.sql ./migrations/entrypoint.sh ./
RUN chmod a+x ./entrypoint.sh

RUN tr -d '\r' <./entrypoint.sh> ./entrypoint.tmp
RUN cp ./entrypoint.tmp ./entrypoint.sh


RUN echo "run migrations"
CMD ["bash", "./entrypoint.sh"]

