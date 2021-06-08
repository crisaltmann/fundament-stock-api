CREATE TABLE IF NOT EXISTS ATIVO (
    id 			serial 			PRIMARY KEY,
    codigo 		varchar(10) 	UNIQUE NOT NULL,
    nome 		varchar(60) 	NOT NULL,
    logo        text            NULL,
    cotacao     decimal(10,2)   NULL
);

INSERT INTO ATIVO
    (codigo, nome, logo)
VALUES('WEGE3', 'WEGE',
       'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAYAAABccqhmAAADvElEQVR4nO3dwY3bMBBAUTpwQ+nHqmWdWqx60pJy3AC5OPYS1OK/V4A00OGDPHh8OY5jAE0/Vg8ArCMAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAEHYZ236sHgJYwwkAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwgQAwq7T3/C4Xaa/A3iJEwCECQCECQCECQCECQCECQCECQCECQCECQCECQCECQCECQCECQCECQCECQCEzd8H8Kpt/xhj3Cc9/T4et1+Tng3fxnkDcGbbfkx79jsLVGbO9V28+v2i384VAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMIEAMKsBec576wr/9u2/x5j/PySZ/E2AYDZTvxfD64AECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAEGYjEM+ZvJmGNZwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIEwAIMxWYJ7zuF1Wj8DXcwKAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAsPk/B972Y/o7/t99bPt99RCwmn0APOecIf/XGfcWnPjbuQJAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABAmABA2Pk2qAKfJm8UdgKAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAMAGAsOvY9o/VQwBrXMcY99VDAGu4AkCYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAECYAEDY5TiO1TMAizgBQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQJgAQNgfipQtDghRvOoAAAAASUVORK5CYII=');

INSERT INTO ATIVO
    (codigo, nome, logo)
VALUES('ITUB3', 'ITAU',
       'data:image/png;base64,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');

INSERT INTO ATIVO
    (codigo, nome, logo)
VALUES('HYPE3', 'HYPERA',
       'data:image/png;base64,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');

CREATE TABLE IF NOT EXISTS MOVIMENTACAO (
    id          serial          PRIMARY KEY,
    id_ativo    bigint          NOT NULL,
    quantidade  int             NOT NULL,
    valor       decimal(10,2)   NOT NULL,
    data        date            NOT NULL,
    id_usuario  bigint          NOT NULL
)

CREATE TABLE IF NOT EXISTS TRIMESTRE (
    id              serial          PRIMARY KEY,
    codigo          varchar(20)     UNIQUE NOT NULL,
    ano             int             NOT NULL,
    trimestre       int             NOT NULL,
    data_inicio     date            NOT NULL,
    data_fim        date            NOT NULL
);

INSERT INTO TRIMESTRE (codigo, ano, trimestre, data_inicio, data_fim) VALUES
    ('2020_01', 2020, 1, '2020-01-01', '2020-03-31'),
    ('2020_02', 2020, 2, '2020-04-01', '2020-06-30'),
    ('2020_03', 2020, 3, '2020-07-01', '2020-09-30'),
    ('2020_04', 2020, 4, '2020-10-01', '2020-12-31'),
    ('2021_01', 2021, 1, '2021-01-01', '2020-03-31'),
    ('2021_02', 2021, 2, '2021-04-01', '2020-06-30');

CREATE TABLE IF NOT EXISTS COTACAO_ATIVO (
    id              serial          PRIMARY KEY,
    id_ativo        bigint          UNIQUE NOT NULL,
    data            date            NOT NULL,
    cotacao         decimal(10,2)   NOT NULL,
    total           bigint          NOT NULL
);

CREATE TABLE IF NOT EXISTS RESULTADO_TRIMESTRE (
       id                  serial          PRIMARY KEY,
       id_trimestre        bigint          NOT NULL,
       id_ativo            bigint          NOT NULL,
       receita_liquida     bigint          NOT NULL,
       ebitda              bigint          NOT NULL,
       lucro_liquido       bigint          NOT NULL,
       divida_liquida      bigint          NOT NULL
);

CREATE TABLE IF NOT EXISTS PORTFOLIO_TRIMESTRE (
    id                  serial          PRIMARY KEY,
    id_trimestre        bigint          NOT NULL,
    id_usuario          bigint          NOT NULL,
    id_ativo            bigint          NOT NULL,
    quantidade          int             NOT NULL,
    receita_liquida     bigint          NOT NULL,
    ebitda              bigint          NOT NULL,
    lucro_liquido       bigint          NOT NULL,
    divida_liquida      bigint          NOT NULL
);