CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    username TEXT UNIQUE,
    email TEXT UNIQUE,

    password_hash varchar(64) NOT NULL,
    password_salt varchar(32) NOT NULL,
    password_iterations int NOT NULL,

    tokens json DEFAULT '{"objects":[]}',

    dreams json DEFAULT '{"objects":[]}',
    goals json DEFAULT '{"objects":[]}'
);

CREATE TABLE dreams (
    dreamid SERIAL PRIMARY KEY,
    userid BIGINT NOT NULL,
    title TEXT,
    descrip TEXT
);

CREATE TABLE goals (
    goalid SERIAL PRIMARY KEY,
    userid BIGINT NOT NULL,
    title TEXT,
    descrip TEXT,
    enddate TEXT,
    tasks json DEFAULT '{"objects":[]}'
);

CREATE TABLE tasks (
    taskid SERIAL PRIMARY KEY,
    title TEXT,
    stat BOOLEAN
);