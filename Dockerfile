# Use uma imagem base do ksqlDB
FROM confluentinc/cp-ksqldb-server:5.3.0

# Copie os arquivos de configuração personalizados, se necessário
# COPY ./path/to/ksqldb/config/ksqldb.properties /etc/ksqldb/ksqldb.properties

# Defina quaisquer variáveis de ambiente necessárias
# ENV VARIABLE_NAME=value

# Exponha a porta em que o ksqlDB vai rodar (por padrão é 8088)
EXPOSE 8088

# Comando para iniciar o ksqlDB
# CMD ["ksql-server-start", "/etc/ksqldb/ksqldb.properties"]
