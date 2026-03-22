DROP TABLE IF EXISTS constraint_combo_ref;
CREATE TABLE constraint_combo_ref (
    constraint_combo_ref double UNIQUE
);

DROP TABLE IF EXISTS constraint_combo_non_pk;
CREATE TABLE constraint_combo_non_pk (
    not_null_unique_check_def_const double NOT NULL UNIQUE CHECK (not_null_unique_check_def_const > 0) DEFAULT 42,
    not_null_unique_check_def_func double NOT NULL UNIQUE CHECK (not_null_unique_check_def_func > 0) DEFAULT (pi()),
    not_null_unique_ref_def_const double NOT NULL UNIQUE DEFAULT 42 REFERENCES constraint_combo_ref(constraint_combo_ref),
    not_null_unique_ref_def_func double NOT NULL UNIQUE DEFAULT (pi()) REFERENCES constraint_combo_ref(constraint_combo_ref),
    not_null_check_ref_def_const double NOT NULL CHECK (not_null_check_ref_def_const > 0) DEFAULT 42 REFERENCES constraint_combo_ref(constraint_combo_ref),
    not_null_check_ref_def_func double NOT NULL CHECK (not_null_check_ref_def_func > 0) DEFAULT (pi()) REFERENCES constraint_combo_ref(constraint_combo_ref),
    not_null_unique_def_const double NOT NULL UNIQUE DEFAULT 42,
    not_null_unique_def_func double NOT NULL UNIQUE DEFAULT (pi()),
    not_null_check_def_const double NOT NULL CHECK (not_null_check_def_const > 0) DEFAULT 42,
    not_null_check_def_func double NOT NULL CHECK (not_null_check_def_func > 0) DEFAULT (pi()),
    not_null_ref_def_const double NOT NULL DEFAULT 42 REFERENCES constraint_combo_ref(constraint_combo_ref),
    not_null_ref_def_func double NOT NULL DEFAULT (pi()) REFERENCES constraint_combo_ref(constraint_combo_ref)
);

DROP TABLE IF EXISTS constraint_combo_not_null_unique_pk_def_const;
CREATE TABLE constraint_combo_not_null_unique_pk_def_const (
    constraint_combo_not_null_unique_pk_def_const double PRIMARY KEY NOT NULL UNIQUE DEFAULT 42
);

DROP TABLE IF EXISTS constraint_combo_not_null_unique_pk_def_func;
CREATE TABLE constraint_combo_not_null_unique_pk_def_func (
    constraint_combo_not_null_unique_pk_def_func double PRIMARY KEY NOT NULL UNIQUE DEFAULT (pi())
);

DROP TABLE IF EXISTS constraint_combo_not_null_check_pk_def_const;
CREATE TABLE constraint_combo_not_null_check_pk_def_const (
    constraint_combo_not_null_check_pk_def_const double PRIMARY KEY NOT NULL CHECK (constraint_combo_not_null_check_pk_def_const > 0) DEFAULT 42
);

DROP TABLE IF EXISTS constraint_combo_not_null_check_pk_def_func;
CREATE TABLE constraint_combo_not_null_check_pk_def_func (
    constraint_combo_not_null_check_pk_def_func double PRIMARY KEY NOT NULL CHECK (constraint_combo_not_null_check_pk_def_func > 0) DEFAULT (pi())
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_ref_def_const;
CREATE TABLE constraint_combo_not_null_pk_ref_def_const (
    constraint_combo_not_null_pk_ref_def_const double PRIMARY KEY NOT NULL DEFAULT 42 REFERENCES constraint_combo_ref(constraint_combo_ref)
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_ref_def_func;
CREATE TABLE constraint_combo_not_null_pk_ref_def_func (
    constraint_combo_not_null_pk_ref_def_func double PRIMARY KEY NOT NULL DEFAULT (pi()) REFERENCES constraint_combo_ref(constraint_combo_ref)
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_def_const;
CREATE TABLE constraint_combo_not_null_pk_def_const (
    constraint_combo_not_null_pk_def_const double NOT NULL PRIMARY KEY DEFAULT 42
);

DROP TABLE IF EXISTS constraint_combo_not_null_pk_def_func;
CREATE TABLE constraint_combo_not_null_pk_def_func (
    constraint_combo_not_null_pk_def_func double NOT NULL PRIMARY KEY DEFAULT (pi())
);
