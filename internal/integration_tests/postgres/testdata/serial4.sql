DROP TABLE IF EXISTS serial4_ref CASCADE;
CREATE TABLE serial4_ref (
    serial4_ref serial4 UNIQUE
);

DROP TABLE IF EXISTS serial4;
CREATE TABLE serial4 (
    serial4 serial4
);

DROP TABLE IF EXISTS serial4_notnull;
CREATE TABLE serial4_notnull (
    serial4 serial4 NOT NULL
);

DROP TABLE IF EXISTS serial4_unique;
CREATE TABLE serial4_unique (
    serial4 serial4 UNIQUE
);

DROP TABLE IF EXISTS serial4_check;
CREATE TABLE serial4_check (
    serial4 serial4 CHECK ( serial4 > 0 )
);

DROP TABLE IF EXISTS serial4_pk;
CREATE TABLE serial4_pk (
    serial4 serial4 PRIMARY KEY
);

DROP TABLE IF EXISTS serial4_references;
CREATE TABLE serial4_references (
    serial4 serial4 REFERENCES serial4_ref (serial4_ref)
);

DROP TABLE IF EXISTS serial4_unique_check;
CREATE TABLE serial4_unique_check (
    serial4 serial4 UNIQUE CHECK ( serial4 > 0 )
);

DROP TABLE IF EXISTS serial4_unique_check_pk;
CREATE TABLE serial4_unique_check_pk (
    serial4 serial4 PRIMARY KEY UNIQUE CHECK ( serial4 > 0 )
);

DROP TABLE IF EXISTS serial4_unique_check_pk_ref;
CREATE TABLE serial4_unique_check_pk_ref (
    serial4 serial4 PRIMARY KEY UNIQUE CHECK ( serial4 > 0 ) REFERENCES serial4_ref (serial4_ref)
);

DROP TABLE IF EXISTS serial4_unique_check_ref;
CREATE TABLE serial4_unique_check_ref (
    serial4 serial4 UNIQUE CHECK ( serial4 > 0 ) REFERENCES serial4_ref (serial4_ref)
);
