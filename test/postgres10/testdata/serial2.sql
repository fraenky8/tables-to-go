DROP TABLE IF EXISTS serial2_ref CASCADE;
CREATE TABLE serial2_ref (
    serial2_ref serial2 UNIQUE
);

DROP TABLE IF EXISTS serial2;
CREATE TABLE serial2 (
    serial2 serial2
);

DROP TABLE IF EXISTS serial2_notnull;
CREATE TABLE serial2_notnull (
    serial2 serial2 NOT NULL
);

DROP TABLE IF EXISTS serial2_unique;
CREATE TABLE serial2_unique (
    serial2 serial2 UNIQUE
);

DROP TABLE IF EXISTS serial2_check;
CREATE TABLE serial2_check (
    serial2 serial2 CHECK ( serial2 > 0 )
);

DROP TABLE IF EXISTS serial2_pk;
CREATE TABLE serial2_pk (
    serial2 serial2 PRIMARY KEY
);

DROP TABLE IF EXISTS serial2_references;
CREATE TABLE serial2_references (
    serial2 serial2 REFERENCES serial2_ref (serial2_ref)
);

DROP TABLE IF EXISTS serial2_unique_check;
CREATE TABLE serial2_unique_check (
    serial2 serial2 UNIQUE CHECK ( serial2 > 0 )
);

DROP TABLE IF EXISTS serial2_unique_check_pk;
CREATE TABLE serial2_unique_check_pk (
    serial2 serial2 PRIMARY KEY UNIQUE CHECK ( serial2 > 0 )
);

DROP TABLE IF EXISTS serial2_unique_check_pk_ref;
CREATE TABLE serial2_unique_check_pk_ref (
    serial2 serial2 PRIMARY KEY UNIQUE CHECK ( serial2 > 0 ) REFERENCES serial2_ref (serial2_ref)
);

DROP TABLE IF EXISTS serial2_unique_check_ref;
CREATE TABLE serial2_unique_check_ref (
    serial2 serial2 UNIQUE CHECK ( serial2 > 0 ) REFERENCES serial2_ref (serial2_ref)
);