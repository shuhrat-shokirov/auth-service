Create table if not exists users(
id Bigserial primary key ,
login Text UNIQUE not null,
password text not null
);

INSERT INTO users (login, password)
VALUES ('admin', 'pass'),
       ('qwe', 'pass');

INSERT INTO users (login, password) VALUES (?, ?);