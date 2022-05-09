FROM postgres:13.6-alpine

COPY ./ddl /docker-entrypoint-initdb.d

RUN rm -rf data && mkdir data

ENV POSTGRES_USER user
ENV POSTGRES_PASSWORD password