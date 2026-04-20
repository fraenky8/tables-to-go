DROP TABLE IF EXISTS single_pk_implicit_autoincrement_table;
CREATE TABLE single_pk_implicit_autoincrement_table
(
    pk integer NOT NULL,
    name text,
    PRIMARY KEY (pk)
);

DROP TABLE IF EXISTS single_pk_explicit_autoincrement_table;
CREATE TABLE single_pk_explicit_autoincrement_table
(
    pk integer NOT NULL,
    name text,
    PRIMARY KEY (pk AUTOINCREMENT)
);

DROP TABLE IF EXISTS single_pk_text_table;
CREATE TABLE single_pk_text_table
(
    pk text NOT NULL,
    name text,
    PRIMARY KEY (pk)
);

DROP TABLE IF EXISTS multi_int_pk_table;
CREATE TABLE multi_int_pk_table
(
    pk_a integer NOT NULL,
    pk_b integer NOT NULL,
    name text,
    PRIMARY KEY (pk_a, pk_b)
);

DROP TABLE IF EXISTS multi_text_pk_table;
CREATE TABLE multi_text_pk_table
(
    pk_a text NOT NULL,
    pk_b text NOT NULL,
    name text,
    PRIMARY KEY (pk_a, pk_b)
);
