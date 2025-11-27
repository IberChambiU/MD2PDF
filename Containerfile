
# Etapa de build con Go y Chromium
FROM golang:1.25.1 AS build-back
RUN apt-get update && \
    # Solo instalar herramientas necesarias para compilar en la etapa build
    apt-get install -y --no-install-recommends make binutils ca-certificates wget && \
    rm -rf /var/lib/apt/lists/*
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN make build
RUN strip /app/md2pdf

# Imagen final basada en Debian slim con Chromium y binario
FROM debian:bullseye-slim AS base
LABEL name="md2pdf" version="0.0.1"
WORKDIR /
RUN apt-get update && \
    # Instalar runtime deps (Chromium + fuentes emoji) solamente en la imagen final
    apt-get install -y --no-install-recommends \
        chromium \
        fonts-liberation \
        fonts-noto-color-emoji \
        fonts-noto-core \
        fonts-symbola \
        fontconfig \
        libappindicator3-1 libasound2 libatk-bridge2.0-0 libatk1.0-0 libcups2 \
        libdbus-1-3 libgdk-pixbuf-xlib-2.0-0 libnspr4 libnss3 libx11-xcb1 \
        libxcomposite1 libxdamage1 libxrandr2 xdg-utils wget && \
    rm -rf /var/lib/apt/lists/* && \
    fc-cache -f -v
COPY --from=build-back /app/md2pdf ./
COPY static/ ./static/
ENV URL_SERVIDOR=http://localhost:3030 PORT=3030 CHROMEDP_PATH=/usr/bin/chromium
CMD ["/md2pdf"]

FROM base AS cert
ENV URL_SERVIDOR=http://localhost:3030

FROM base AS prod
ENV URL_SERVIDOR=http://localhost:3030