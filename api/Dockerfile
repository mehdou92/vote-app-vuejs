# FROM golang:alpine as build
# RUN apk add --no-cache ca-certificates
# WORKDIR /build
# ADD . .
# RUN CGO_ENABLED=0 GOOS=linux \
#   go build -ldflags '-extldflags "-static"' -o app

# FROM scratch
# COPY --from=build /etc/ssl/certs/ca-certificates.crt \
#   /etc/ssl/certs/ca-certificates.crt
# COPY --from=build /build/app /app
# ENTRYPOINT ["/app"]

FROM golang:alpine

# if left blank app will run with dev settings
# to build production image run:
# $ docker build ./api --build-args app_env=production
ARG app_env
ENV APP_ENV $app_env

# it is okay to leave user/GoDoRP as long as you do not want to share code with other libraries
COPY . /go/src/github.com/mehdou92/vote-app-vuejs/api
WORKDIR /go/src/github.com/mehdou92/vote-app-vuejs/api

RUN apk add --no-cache git mercurial
RUN ls
RUN go get ./
RUN go build

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
  then \
  Projet-GO; \
  else \
  sleep 10 && \
  go get github.com/pilu/fresh && \
  fresh; \
  fi

EXPOSE 3000
