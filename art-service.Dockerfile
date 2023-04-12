FROM alpine:latest

RUN mkdir /app 

COPY /dist/art_service /app 

CMD [ "/app/art_service" ]