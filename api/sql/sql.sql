CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY (usuario_id, seguidor_id)
) ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    
    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp
) ENGINE=INNODB;

/* ---------------------------------------------------------------------------------- */

-- Cria o banco de dados (se não existir)
CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

-- Remove as tabelas na ordem correta (primeiro as que dependem de outras)
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS usuarios;

-- Cria a tabela 'usuarios'
CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

-- Cria a tabela 'seguidores'
CREATE TABLE seguidores(
    usuario_id int not null,
    seguidor_id int not null,
    PRIMARY KEY (usuario_id, seguidor_id),
    FOREIGN KEY (usuario_id) REFERENCES usuarios(id) ON DELETE CASCADE,
    FOREIGN KEY (seguidor_id) REFERENCES usuarios(id) ON DELETE CASCADE
) ENGINE=INNODB;

-- Cria a tabela 'publicacoes'
CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id int not null,
    curtidas int default 0,
    criadaEm timestamp default current_timestamp,
    FOREIGN KEY (autor_id) REFERENCES usuarios(id) ON DELETE CASCADE
) ENGINE=INNODB;