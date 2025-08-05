CREATE TABLE usuario (
  id SERIAL PRIMARY KEY,
  senha VARCHAR,
  nome VARCHAR,
  email VARCHAR UNIQUE
);

CREATE TABLE area (
  id SERIAL PRIMARY KEY,
  consultor VARCHAR,
  propriedade VARCHAR,
  area VARCHAR,
  id_usuario INTEGER
);

CREATE TABLE analise (
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
  cao NUMERIC,
  mgo NUMERIC,
  teor_mg NUMERIC,
  mg_desejada NUMERIC,
  ca_desejada NUMERIC,
  id_area INTEGER
);

CREATE TABLE calculo (
  id SERIAL PRIMARY KEY,
  id_analise INTEGER,
  resultado NUMERIC,
  sat_extra NUMERIC,
  relacao_ca_mg NUMERIC,
  data_calculo DATE,
  id_metodo INTEGER
);

ALTER TABLE area
  ADD CONSTRAINT area_id_usuario_fkey FOREIGN KEY (id_usuario) REFERENCES usuario(id);

ALTER TABLE analise
  ADD CONSTRAINT analise_id_usuario_fkey FOREIGN KEY (id_usuario) REFERENCES usuario(id);

ALTER TABLE analise
  ADD CONSTRAINT fk_analise_area FOREIGN KEY (id_area) REFERENCES area(id);

ALTER TABLE calculo
  ADD CONSTRAINT Calculo_id_analise_fkey FOREIGN KEY (id_analise) REFERENCES analise(id);
