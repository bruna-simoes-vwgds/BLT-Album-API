FROM mongo:latest

ARG MONGO_URI
ARG USERNAME
ARG PASSWORD

ENV MONGO_URI ${MONGO_URI}
ENV USERNAME ${USERNAME}
ENV PASSWORD ${PASSWORD}

COPY public.albums.json /albums.json
COPY public.songs.json  /songs.json

CMD mongoimport --username ${USERNAME} --password ${PASSWORD} --uri ${MONGO_URI} --collection albums --drop --file /albums.json && \
    mongoimport --username ${USERNAME} --password ${PASSWORD} --uri ${MONGO_URI} --collection songs --drop --file /songs.json