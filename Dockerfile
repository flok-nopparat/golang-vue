FROM node:14 AS build-node-env
WORKDIR /app

COPY ./client/line-interview/package*.json .
RUN yarn

COPY ./client/line-interview .

RUN yarn build


FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY --from=build-node-env /app/dist/ ./client/line-interview/dist/
COPY . .
RUN go build -v -o /docker-golang
EXPOSE ${RUN_PORT}

CMD ["/docker-golang"]