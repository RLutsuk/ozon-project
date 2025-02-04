FROM golang:alpine AS build

COPY . /server/

WORKDIR /server/

RUN go build cmd/main.go

FROM ubuntu:20.04
COPY . .

RUN apt-get -y update && apt-get install -y tzdata
RUN ln -snf /usr/share/zoneinfo/Russia/Moscow /etc/localtime && echo Russia/Moscow > /etc/timezone

ENV PGVER 12
RUN apt-get -y update && apt-get install -y postgresql-$PGVER
USER postgres

RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER db_pg WITH SUPERUSER PASSWORD 'db_postgres';" &&\
    createdb -O db_pg db_post &&\
    psql -f db/migrations/000001_enable_uuid_ossp.up.sql -d db_post &&\
    psql -f db/migrations/000002_create_users_table.up.sql -d db_post &&\
    psql -f db/migrations/000003_create_posts_table.up.sql -d db_post &&\
    psql -f db/migrations/000004_create_comments_table.up.sql -d db_post &&\
    psql -f db/migrations/000005_test_data.up.sql -d db_post &&\
    /etc/init.d/postgresql stop

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/$PGVER/main/postgresql.conf

RUN /etc/init.d/postgresql restart

USER root
COPY --from=build /server/main .

EXPOSE 8080

ENV STORAGE_TYPE=postgres
ENV POSTGRES_HOST=localhost
ENV POSTGRES_PORT=5432
ENV POSTGRES_USERNAME=db_pg
ENV POSTGRES_PASSWORD=db_postgres
ENV POSTGRES_DATABASE=db_post
ENV SERVER_PORT=8080

CMD service postgresql start && ./main