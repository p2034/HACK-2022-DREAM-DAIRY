CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    email TEXT UNIQUE,

    password_hash varchar(64) NOT NULL,
    password_salt varchar(32) NOT NULL,
    password_iterations int NOT NULL,

    tokens json
);

CREATE TABLE dreams (
    
);