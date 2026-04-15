DROP TABLE IF EXISTS numeric_ref;
CREATE TABLE numeric_ref
(
    numeric_ref numeric UNIQUE
);

DROP TABLE IF EXISTS numeric_table;
CREATE TABLE numeric_table
(
    n              numeric,
    numeric_nn     numeric NOT NULL,
    numeric_unique numeric UNIQUE,
    numeric_check  numeric CHECK (numeric_check > 0),
    numeric_ref    numeric REFERENCES numeric_ref (numeric_ref),
    numeric_def    numeric DEFAULT 123.456,
    numeric_pk     numeric PRIMARY KEY
);

DROP TABLE IF EXISTS numeric_nn_pk;
CREATE TABLE numeric_nn_pk
(
    numeric_nn_pk numeric NOT NULL PRIMARY KEY
);
