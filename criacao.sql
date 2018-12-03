create schema formularioEletronico;

use formularioEletronico;

select database();

create table Patologia
(
  id_patologia int not null,
  nome         varchar(100),
  primary key (id_patologia)
);

create table ClassificacaoEtaria
(
  id_classificacao int not null,
  faixa_etaria     varchar(100),
  classificacao    varchar(50),
  primary key (id_classificacao)
);

create table UF
(
  id_uf int not null,
  sigla char(2),
  nome  varchar(50),
  primary key (id_uf)
);

create table Zona
(
  id_zona      int not null,
  nome         varchar(50),
  primary key (id_zona)
);

create table Paciente
(
  id_paciente        int not null,
  id_classificacao   int not null,
  sexo               varchar(8),
  endereco           varchar(150),
  id_patologia       int not null,
  presenca_patologia boolean,
  id_uf     int not null,
  id_zona            int not null,
  primary key (id_paciente),
  foreign key (id_classificacao) references ClassificacaoEtaria (id_classificacao) on delete cascade,
  foreign key (id_patologia) references Patologia (id_patologia) on delete restrict,
  foreign key (id_uf) references UF (id_uf) on delete restrict,
  foreign key (id_zona) references Zona (id_zona) on delete restrict
);
