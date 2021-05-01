CREATE TABLE ATIVO (
    id bigint IDENTITY(1,1) PRIMARY KEY,
    codigo varchar(10) not null,
    nome varchar(60) not null,
)
INSERT INTO youasholding.dbo.ATIVO
    (codigo, nome)
VALUES('WEGE3', 'WEGE');

INSERT INTO youasholding.dbo.ATIVO
    (codigo, nome)
VALUES('ITUB3', 'ITAU');

INSERT INTO youasholding.dbo.ATIVO
    (codigo, nome)
VALUES('HYPE3', 'HYPERA');
