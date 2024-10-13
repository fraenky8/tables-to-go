DROP TABLE IF EXISTS smallserial_ref CASCADE;
CREATE TABLE smallserial_ref (
    smallserial_ref smallserial UNIQUE
);

DROP TABLE IF EXISTS smallserial;
CREATE TABLE smallserial (
    smallserial smallserial
);

DROP TABLE IF EXISTS smallserial_notnull;
CREATE TABLE smallserial_notnull (
    smallserial smallserial NOT NULL
);

DROP TABLE IF EXISTS smallserial_unique;
CREATE TABLE smallserial_unique (
    smallserial smallserial UNIQUE
);

DROP TABLE IF EXISTS smallserial_check;
CREATE TABLE smallserial_check (
    smallserial smallserial CHECK ( smallserial > 0 )
);

DROP TABLE IF EXISTS smallserial_pk;
CREATE TABLE smallserial_pk (
    smallserial smallserial PRIMARY KEY
);

DROP TABLE IF EXISTS smallserial_references;
CREATE TABLE smallserial_references (
    smallserial smallserial REFERENCES smallserial_ref (smallserial_ref)
);

DROP TABLE IF EXISTS smallserial_unique_check;
CREATE TABLE smallserial_unique_check (
    smallserial smallserial UNIQUE CHECK ( smallserial > 0 )
);

DROP TABLE IF EXISTS smallserial_unique_check_pk;
CREATE TABLE smallserial_unique_check_pk (
    smallserial smallserial PRIMARY KEY UNIQUE CHECK ( smallserial > 0 )
);

DROP TABLE IF EXISTS smallserial_unique_check_pk_ref;
CREATE TABLE smallserial_unique_check_pk_ref (
    smallserial smallserial PRIMARY KEY UNIQUE CHECK ( smallserial > 0 ) REFERENCES smallserial_ref (smallserial_ref)
);

DROP TABLE IF EXISTS smallserial_unique_check_ref;
CREATE TABLE smallserial_unique_check_ref (
    smallserial smallserial UNIQUE CHECK ( smallserial > 0 ) REFERENCES smallserial_ref (smallserial_ref)
);