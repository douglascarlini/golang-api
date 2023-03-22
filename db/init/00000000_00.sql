CREATE EXTENSION "uuid-ossp";

DROP TABLE IF EXISTS domains;
CREATE TABLE domains (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(20) NOT NULL UNIQUE,
    label VARCHAR(45) DEFAULT NULL
);

DROP TABLE IF EXISTS apps;
CREATE TABLE apps (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    secret UUID NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(20) NOT NULL UNIQUE,
    label VARCHAR(45) DEFAULT NULL
);

DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(20) NOT NULL UNIQUE,
    label VARCHAR(45) DEFAULT NULL
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(45) NOT NULL,
    username VARCHAR(20) NOT NULL UNIQUE,
    password VARCHAR(80) NOT NULL,
    role_id UUID NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles (id)
);

DROP TABLE IF EXISTS modules;
CREATE TABLE modules (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(20) NOT NULL UNIQUE,
    label VARCHAR(45) DEFAULT NULL
);

DROP TABLE IF EXISTS actions;
CREATE TABLE actions (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(20) NOT NULL UNIQUE,
    label VARCHAR(45) DEFAULT NULL
);

DROP TABLE IF EXISTS permissions;
CREATE TABLE permissions (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    role_id UUID NOT NULL,
    module_id UUID NOT NULL,
    action_id UUID NOT NULL,
    item_id UUID DEFAULT NULL,
    FOREIGN KEY (role_id) REFERENCES roles (id),
    FOREIGN KEY (module_id) REFERENCES modules (id),
    FOREIGN KEY (action_id) REFERENCES actions (id)
);