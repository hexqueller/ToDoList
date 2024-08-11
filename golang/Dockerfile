FROM alpine:latest

WORKDIR /app

COPY server .

RUN chmod +x /app/server

EXPOSE 1234
CMD ["/app/server"]