CREATE TABLE Usuario (
  id SERIAL PRIMARY KEY,
  senha VARCHAR,
  nome VARCHAR,
  email VARCHAR UNIQUE
);


CREATE TABLE Area (
  id SERIAL PRIMARY KEY,
  consultor VARCHAR,
  propriedade VARCHAR,
  area VARCHAR,
  id_usuario INTEGER
);

CREATE TABLE Analise (
  id SERIAL PRIMARY KEY,
  id_usuario INTEGER,
  magnesio NUMERIC,
  aluminio NUMERIC,
  calcio NUMERIC,
  sat_desejada NUMERIC,
  prnt NUMERIC,
  ctc NUMERIC,
  argila NUMERIC,
  sat_atual NUMERIC,
  teor_ca NUMERIC,
  caO NUMERIC,
  mgO NUMERIC,
  teor_mg NUMERIC,
  mg_desejada NUMERIC,
  ca_desejada NUMERIC,
  id_area INTEGER
);


CREATE TABLE Calculo (
  id SERIAL PRIMARY KEY,
  id_analise INTEGER,
  resultado NUMERIC,
  sat_extra NUMERIC,
  relacao_ca_mg NUMERIC,
  data_calculo DATE,
  id_metodo INTEGER
);

ALTER TABLE Area
  ADD CONSTRAINT Area_id_usuario_fkey FOREIGN KEY (id_usuario) REFERENCES Usuario(id);

ALTER TABLE Analise
  ADD CONSTRAINT Analise_id_usuario_fkey FOREIGN KEY (id_usuario) REFERENCES Usuario(id);

ALTER TABLE Analise
  ADD CONSTRAINT fk_analise_area FOREIGN KEY (id_area) REFERENCES Area(id);

ALTER TABLE Calculo
  ADD CONSTRAINT Calculo_id_analise_fkey FOREIGN KEY (id_analise) REFERENCES Analise(id);
