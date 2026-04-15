DROP TABLE IF EXISTS constraint_combo_ref;
CREATE TABLE constraint_combo_ref
(
    constraint_combo_ref real UNIQUE
);

DROP TABLE IF EXISTS constraint_combo_non_pk;
CREATE TABLE constraint_combo_non_pk
(
    not_null_unique_check_def_const real NOT NULL UNIQUE CHECK (not_null_unique_check_def_const > 0) DEFAULT 42,
    not_null_unique_ref_def_const   real NOT NULL UNIQUE DEFAULT 42 REFERENCES constraint_combo_ref(constraint_combo_ref),
    not_null_check_ref_def_const    real NOT NULL CHECK (not_null_check_ref_def_const > 0) DEFAULT 42 REFERENCES constraint_combo_ref(constraint_combo_ref),
    not_null_unique_def_const       real NOT NULL UNIQUE DEFAULT 42,
    not_null_check_def_const        real NOT NULL CHECK (not_null_check_def_const > 0) DEFAULT 42,
    not_null_ref_def_const          real NOT NULL DEFAULT 42 REFERENCES constraint_combo_ref(constraint_combo_ref)
);

DROP TABLE IF EXISTS constraint_combo_not_null_unique_pk_def_const;
CREATE TABLE constraint_combo_not_null_unique_pk_def_const
(
    constraint_combo_not_null_unique_pk_def_const real PRIMARY KEY NOT NULL UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS constraint_combo_not_null_check_pk_def_const;
CREATE TABLE constraint_combo_not_null_check_pk_def_const
(
    constraint_combo_not_null_check_pk_def_const real PRIMARY KEY NOT NULL CHECK (constraint_combo_not_null_check_pk_def_const > 0) DEFAULT 42
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_ref_def_const;
CREATE TABLE constraint_combo_not_null_pk_ref_def_const
(
    constraint_combo_not_null_pk_ref_def_const real PRIMARY KEY NOT NULL DEFAULT 42 REFERENCES constraint_combo_ref(constraint_combo_ref)
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_def_const;
CREATE TABLE constraint_combo_not_null_pk_def_const
(
    constraint_combo_not_null_pk_def_const real NOT NULL PRIMARY KEY DEFAULT 42
);
