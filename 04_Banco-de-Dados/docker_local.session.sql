SELECT * FROM tipos;
SELECT * FROM servicos_tipos;
SELECT * FROM servicos ORDER BY nome;

/*Mysql*/
/*
--4 - Servico, Tipo
SELECT * FROM tipos;
SELECT * FROM servicos_tipos;
SELECT * FROM servicos ORDER BY nome;

--3 - Item, Categoria, Objeto, Serial
SELECT * FROM categoria;
SELECT * FROM items ORDER BY nome;
SELECT * FROM objetos ORDER BY nome;
SELECT * FROM numero_serials;

--2 - Pessoas
DROP TABLE pessoas;
DESC pessoas;
SELECT * FROM pessoas ORDER BY nome;
SELECT * FROM pessoas where deleted_at is null ORDER BY Nome;

--1 - Produtos
CREATE TABLE produtos (
    id VARCHAR(255) NOT NULL,
    nome VARCHAR(80) NOT NULL,
    preco DECIMAL(10,2) NOT NULL,
    PRIMARY KEY (id)
);
SELECT * FROM produtos ORDER BY nome;
SHOW tables;
*/