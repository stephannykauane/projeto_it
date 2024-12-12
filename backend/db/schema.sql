CREATE TABLE Usuario (
    id_usuario SERIAL PRIMARY KEY,
    email VARCHAR NOT NULL,
    senha VARCHAR NOT NULL,
    nome VARCHAR
);

CREATE TABLE Metodo (
    id_metodo SERIAL PRIMARY KEY,
    nome_metodo VARCHAR NOT NULL
);

CREATE TABLE Calculo (
    id_calculo SERIAL PRIMARY KEY,
    resultado NUMERIC,
    data_calculo DATE
);

CREATE TABLE Analise (
    id_analise SERIAL PRIMARY KEY,
    hidrogenio NUMERIC,
    potassio NUMERIC,
    magnesio NUMERIC,
    aluminio NUMERIC,
    calcio NUMERIC,
    sat_desejada NUMERIC,
    prnt NUMERIC,
    ctc NUMERIC
);
