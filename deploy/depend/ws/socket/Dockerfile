FROM debian:stretch-slim
WORKDIR /app
COPY ./ws/socket/socket /app/socket
COPY ./ws/socket/config.toml /app/config.toml
RUN chmod +x /app/socket
