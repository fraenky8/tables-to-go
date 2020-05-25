DROP TABLE IF EXISTS bigserial_ref CASCADE;
CREATE TABLE bigserial_ref (
    bigserial_ref bigserial UNIQUE
);

DROP TABLE IF EXISTS bigserial;
CREATE TABLE bigserial (
    bigserial bigserial
);

DROP TABLE IF EXISTS bigserial_notnull;
CREATE TABLE bigserial_notnull (
    bigserial bigserial NOT NULL
);

DROP TABLE IF EXISTS bigserial_unique;
CREATE TABLE bigserial_unique (
    bigserial bigserial UNIQUE
);

DROP TABLE IF EXISTS bigserial_check;
CREATE TABLE bigserial_check (
    bigserial bigserial CHECK ( bigserial > 0 )
);

DROP TABLE IF EXISTS bigserial_pk;
CREATE TABLE bigserial_pk (
    bigserial bigserial PRIMARY KEY
);

DROP TABLE IF EXISTS bigserial_references;
CREATE TABLE bigserial_references (
    bigserial bigserial REFERENCES bigserial_ref (bigserial_ref)
);

DROP TABLE IF EXISTS bigserial_unique_check;
CREATE TABLE bigserial_unique_check (
    bigserial bigserial UNIQUE CHECK ( bigserial > 0 )
);

DROP TABLE IF EXISTS bigserial_unique_check_pk;
CREATE TABLE bigserial_unique_check_pk (
    bigserial bigserial PRIMARY KEY UNIQUE CHECK ( bigserial > 0 )
);

DROP TABLE IF EXISTS bigserial_unique_check_pk_ref;
CREATE TABLE bigserial_unique_check_pk_ref (
    bigserial bigserial PRIMARY KEY UNIQUE CHECK ( bigserial > 0 ) REFERENCES bigserial_ref (bigserial_ref)
);

DROP TABLE IF EXISTS bigserial_unique_check_ref;
CREATE TABLE bigserial_unique_check_ref (
    bigserial bigserial UNIQUE CHECK ( bigserial > 0 ) REFERENCES bigserial_ref (bigserial_ref)
);