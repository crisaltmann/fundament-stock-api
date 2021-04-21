CREATE TABLE IF NOT EXISTS ATIVO (
    id 			bigint 			PRIMARY KEY,
    codigo 		varchar(10) 	UNIQUE NOT NULL,
    nome 		varchar(60) 	NOT NULL
)

INSERT INTO ATIVO
    (codigo, nome)
VALUES('WEGE3', 'WEGE');

INSERT INTO ATIVO
    (codigo, nome)
VALUES('ITUB3', 'ITAU');

INSERT INTO ATIVO
    (codigo, nome)
VALUES('HYPE3', 'HYPERA');

CREATE TABLE IF NOT EXISTS USERS (
    id          bigint          PRIMARY KEY,
    username    varchar(10)     NOT NULL,
    password    varchar(60)     NOT NULL
)

-- CREATE TABLE IF NOT EXISTS PORTFOLIO (
--     id          bigint          PRIMARY KEY,
--     user_id     bigint          NOT NULL,
--     nome        varchar(100) NOT NULL,
--     CONSTRAINT PORTFOLIO_FK FOREIGN KEY (id) REFERENCES youasholding.dbo.USERS(id)
-- )
--
-- INSERT INTO USERS (username, password) VALUES ('teste', 'teste')
--
-- INSERT INTO PORTFOLIO (user_id, nome) VALUES ((SELECT id FROM USERS WHERE username = 'teste') ,'PADRÃO');
