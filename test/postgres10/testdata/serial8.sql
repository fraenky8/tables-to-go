DROP TABLE IF EXISTS serial8_ref CASCADE;
CREATE TABLE serial8_ref (
    serial8_ref serial8 UNIQUE
);

DROP TABLE IF EXISTS serial8;
CREATE TABLE serial8 (
    serial8 serial8
);

DROP TABLE IF EXISTS serial8_notnull;
CREATE TABLE serial8_notnull (
    serial8 serial8 NOT NULL
);

DROP TABLE IF EXISTS serial8_unique;
CREATE TABLE serial8_unique (
    serial8 serial8 UNIQUE
);

DROP TABLE IF EXISTS serial8_check;
CREATE TABLE serial8_check (
    serial8 serial8 CHECK ( serial8 > 0 )
);

DROP TABLE IF EXISTS serial8_pk;
CREATE TABLE serial8_pk (
    serial8 serial8 PRIMARY KEY
);

DROP TABLE IF EXISTS serial8_references;
CREATE TABLE serial8_references (
    serial8 serial8 REFERENCES serial8_ref (serial8_ref)
);

DROP TABLE IF EXISTS serial8_unique_check;
CREATE TABLE serial8_unique_check (
    serial8 serial8 UNIQUE CHECK ( serial8 > 0 )
);

DROP TABLE IF EXISTS serial8_unique_check_pk;
CREATE TABLE serial8_unique_check_pk (
    serial8 serial8 PRIMARY KEY UNIQUE CHECK ( serial8 > 0 )
);

DROP TABLE IF EXISTS serial8_unique_check_pk_ref;
CREATE TABLE serial8_unique_check_pk_ref (
    serial8 serial8 PRIMARY KEY UNIQUE CHECK ( serial8 > 0 ) REFERENCES serial8_ref (serial8_ref)
);

DROP TABLE IF EXISTS serial8_unique_check_ref;
CREATE TABLE serial8_unique_check_ref (
    serial8 serial8 UNIQUE CHECK ( serial8 > 0 ) REFERENCES serial8_ref (serial8_ref)
);
