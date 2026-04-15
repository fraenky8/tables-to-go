DROP TABLE IF EXISTS blob_ref;
CREATE TABLE blob_ref
(
    blob_ref blob UNIQUE
);

DROP TABLE IF EXISTS blob_table;
CREATE TABLE blob_table
(
    b              blob,
    blob_nn        blob NOT NULL,
    blob_unique    blob UNIQUE,
    blob_check     blob CHECK (length(blob_check) > 0),
    blob_ref       blob REFERENCES blob_ref (blob_ref),
    blob_def_const blob DEFAULT X'00FF',
    blob_pk        blob PRIMARY KEY
);

DROP TABLE IF EXISTS blob_nn_pk;
CREATE TABLE blob_nn_pk
(
    blob_nn_pk blob NOT NULL PRIMARY KEY
);
