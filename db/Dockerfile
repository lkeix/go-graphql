FROM postgres:13.6-alpine

COPY ./ddl /docker-entrypoint-initdb.d

ENV POSTGRES_USER user
ENV POSTGRES_PASSWORD password