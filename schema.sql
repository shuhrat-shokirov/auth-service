Create table if not exists users
(
    id       Bigserial primary key,
    name     Text NOT NULL,
    login    Text UNIQUE not null,
    password text        not null
);

INSERT INTO users (login, password)
VALUES ('admin', 'pass'),
       ('qwe', 'pass');

INSERT INTO users (login, password)
VALUES (?, ?);

Create table if not exists mitings
(
    id       Bigserial primary key,
    status   bool,
    timeInHour int,
    timeInMinutes int,
    timeOutHour int,
    timeOutMinutes int
);

INSERT INTO mitings (status)
VALUES (?);

SELECT id, timeinhour, timeinminutes, timeouthour, timeoutminutes, filename FROM mitings;