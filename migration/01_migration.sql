CREATE TABLE IF NOT EXISTS ATIVO (
id 			bigint 			PRIMARY KEY,
codigo 		varchar(10) 	UNIQUE NOT NULL,
nome 		varchar(60) 	NOT NULL
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
