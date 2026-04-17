FROM mongo:latest

ARG MONGO_URI

ENV MONGO_URI ${MONGO_URI}

COPY public.albums.json /albums.json
COPY public.songs.json  /songs.json

CMD mongoimport --uri ${MONGO_URI} --collection albums --drop --file /albums.json && \
    mongoimport --uri ${MONGO_URI} --collection songs --drop --file /songs.json