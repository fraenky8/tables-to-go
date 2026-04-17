DROP TABLE IF EXISTS multi_pk_table;
CREATE TABLE multi_pk_table
(
    pk_a integer NOT NULL,
    pk_b integer NOT NULL,
    name text,
    PRIMARY KEY (pk_a, pk_b)
);
