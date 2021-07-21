CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE teams
(
    id serial not null unique,
    title varchar(255) not null
);

CREATE TABLE teamRoles
(
    id serial not null unique,
    teamId int references teams (id) on delete cascade not null,
    title varchar(255) not null,
    moderator boolean not null default false
);

CREATE TABLE roles
(
    id serial not null unique,
    userId int references users (id) on delete cascade not null,
    teamRoleId int references teamRoles (id) on delete cascade not null
);

CREATE TABLE tasks
(
    id serial not null unique,
    title varchar(255) not null,
    teamId int references teams (id) on delete cascade not null,
    description text,
    deadline int not null,
    isActual boolean not null default true
);

CREATE TABLE jobs
(
    id serial not null unique,
    taskId int references tasks (id) on delete cascade not null,
    teamRoleId int references teamRoles (id) on delete cascade not null
);

CREATE TABLE dones
(
    id serial not null unique,
    userId int references users (id) on delete cascade not null,
    jobId int references jobs (id) on delete cascade not null,
    isDone boolean not null default false
);

CREATE TABLE invites
(
    id serial not null unique,
    teamId int references teams (id) on delete cascade not null,
    isActive boolean not null default false
);