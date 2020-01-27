FROM mongo
WORKDIR /data
ARG MONGO_HOST
ENV MONGO=$MONGO_HOST
COPY products.json products.json
CMD mongoimport --host $MONGO --db shop --collection products --type json --file products.json --jsonArray
# RUN ["chmod", "+x", "import.sh"]
# CMD ["./import.sh"]