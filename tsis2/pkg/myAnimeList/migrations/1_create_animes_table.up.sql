CREATE TABLE IF NOT EXISTS animes (
    id           serial     PRIMARY KEY,
    title        text       NOT NULL,
    release_date DATE,
    status       text       NOT NULL,
    type         text       NOT NULL,
    genre        text,
    studio       text,
    cover        text,
    episodes     integer,
    duration     integer
);

CREATE TABLE IF NOT EXISTS users (
    id           serial     PRIMARY KEY,
    username     text
);

CREATE TABLE IF NOT EXISTS users_and_animes (
    user_id  int NOT NULL REFERENCES users(id),
    anime_id int NOT NULL REFERENCES animes(id),
    status   text,
    rating   float,
    PRIMARY KEY(user_id, anime_id)
);

 