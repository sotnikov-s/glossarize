# select image
FROM golang:1.14.1-alpine

# install dependensies and migration tool
RUN apk add --no-cache build-base git && go get github.com/pressly/goose/cmd/goose

# copy migration files.
COPY versions /versions
COPY migrate.sh /usr/local/bin/migrate.sh

# run migration
WORKDIR /versions
CMD migrate.sh