DROP TABLE IF EXISTS real_ref;
CREATE TABLE real_ref
(
    real_ref real UNIQUE
);

DROP TABLE IF EXISTS real_table;
CREATE TABLE real_table
(
    r              real,
    real_nn        real NOT NULL,
    real_unique    real UNIQUE,
    real_check     real CHECK (real_check > 0),
    real_ref       real REFERENCES real_ref (real_ref),
    real_def_const real DEFAULT 42.5,
    real_pk        real PRIMARY KEY
);

DROP TABLE IF EXISTS real_nn_pk;
CREATE TABLE real_nn_pk
(
    real_nn_pk real NOT NULL PRIMARY KEY
);
