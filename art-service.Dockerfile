FROM alpine:latest

RUN mkdir -p /app/img

COPY /static/*.*png /app/img
COPY /dist/art_service /app 

CMD [ "/app/art_service" ]