DROP TABLE IF EXISTS serial_ref CASCADE;
CREATE TABLE serial_ref (
    serial_ref serial UNIQUE
);

DROP TABLE IF EXISTS serial;
CREATE TABLE serial (
    serial serial
);

DROP TABLE IF EXISTS serial_notnull;
CREATE TABLE serial_notnull (
    serial serial NOT NULL
);

DROP TABLE IF EXISTS serial_unique;
CREATE TABLE serial_unique (
    serial serial UNIQUE
);

DROP TABLE IF EXISTS serial_check;
CREATE TABLE serial_check (
    serial serial CHECK ( serial > 0 )
);

DROP TABLE IF EXISTS serial_pk;
CREATE TABLE serial_pk (
    serial serial PRIMARY KEY
);

DROP TABLE IF EXISTS serial_references;
CREATE TABLE serial_references (
    serial serial REFERENCES serial_ref (serial_ref)
);

DROP TABLE IF EXISTS serial_unique_check;
CREATE TABLE serial_unique_check (
    serial serial UNIQUE CHECK ( serial > 0 )
);

DROP TABLE IF EXISTS serial_unique_check_pk;
CREATE TABLE serial_unique_check_pk (
    serial serial PRIMARY KEY UNIQUE CHECK ( serial > 0 )
);

DROP TABLE IF EXISTS serial_unique_check_pk_ref;
CREATE TABLE serial_unique_check_pk_ref (
    serial serial PRIMARY KEY UNIQUE CHECK ( serial > 0 ) REFERENCES serial_ref (serial_ref)
);

DROP TABLE IF EXISTS serial_unique_check_ref;
CREATE TABLE serial_unique_check_ref (
    serial serial UNIQUE CHECK ( serial > 0 ) REFERENCES serial_ref (serial_ref)
);