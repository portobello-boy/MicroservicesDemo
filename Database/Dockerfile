FROM mongo

RUN echo "Spinning up calendar database"

COPY populate-data.js /docker-entrypoint-initdb.d/

EXPOSE 27017