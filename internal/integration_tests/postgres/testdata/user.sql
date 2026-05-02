DROP TABLE IF EXISTS "user" CASCADE;
CREATE TABLE "user"
(
    id          INTEGER NOT NULL,
    user_id     INTEGER NOT NULL,
    email       VARCHAR NOT NULL,
    website_url VARCHAR
);

COMMENT ON TABLE "user" IS 'This is the user table.
Contains account information.';
COMMENT ON COLUMN "user".user_id IS 'This is the identifier of the user.';
COMMENT ON COLUMN "user".email IS 'This is the email of the user.
Use it for notifications.';
COMMENT ON COLUMN "user".website_url IS 'This is the website URL of the user.
Optional field.';
