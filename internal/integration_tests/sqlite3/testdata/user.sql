DROP TABLE IF EXISTS "user";
/*
This is the user table.
Contains account information.
*/
CREATE TABLE "user"
(
    id          integer NOT NULL,
    user_id     integer NOT NULL, /* This is the identifier of the user. */
    email       text NOT NULL, /* This is the email of the user.
Use it for notifications. */
    website_url text /* This is the website URL of the user.
Optional field. */
);
