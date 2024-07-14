CREATE TABLE people
(
    id serial not null unique,
    name varchar(255) not null,
    surname  varchar(255) not null,
    patronymic varchar(255),
    address varchar(255) not null,
    passport_number int not null,
    passport_serie int not null 
);

CREATE TABLE tasks
(
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255),
    start_time varchar(255) not null,
    total_time varchar(255) not null,
    done boolean not null default false
);

CREATE TABLE people_task
(
    id serial not null unique,
    people_id int references people (id) not null,
    tasks_id int references tasks (id) not null
);

-- people tasks people_task