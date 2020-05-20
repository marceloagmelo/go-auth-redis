#!/usr/bin/env bash

# Tabela
echo "Criando a tabela mensagem..."
mysql -h localhost -u root -p -D goauthdb << EOF
use goauthdb;
CREATE TABLE usuario (
id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
nome VARCHAR(20), senha VARCHAR(100),
email VARCHAR(255), status INTEGER,
dtcriacao DATETIME, dtatualizacao DATETIME,
PRIMARY KEY (id)
);
EOF

