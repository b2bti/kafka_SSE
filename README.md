SSE
Artigo usado como base para o SSE em Go: https://thedevelopercafe.com/articles/server-sent-events-in-go-595ae2740c7a

Artigo usado como base para criar o cluster Kafka: https://betterprogramming.pub/a-simple-apache-kafka-cluster-with-docker-kafdrop-and-python-cf45ab99e2b9

iniciar container kqsl:  docker exec -it ksqldb-cli ksql http://ksqldb-server:8088
    criar stream para escutar os dados do topico :  CREATE STREAM prices (price INT) WITH (kafka_topic='odds', value_format='json', partitions=1);
        PRINT 'odds' FROM BEGINNING;

===========================================
docker exec -it ksqldb-cli /bin/sh

ksql http://ksqldb-server:8088

CREATE TABLE odds_table (
    ID INT PRIMARY KEY,
    PRICE INT
) WITH (
    KAFKA_TOPIC = 'odds',
    VALUE_FORMAT = 'JSON'
);

CREATE TABLE QUERYABLE_ODDS_TABLE AS SELECT * FROM ODDS_TABLE;
SELECT * FROM QUERYABLE_ODDS_TABLE;
DESCRIBE QUERYABLE_ODDS_TABLE;


SELECT * FROM ODDS_TABLE;


SHOW TABLES;

