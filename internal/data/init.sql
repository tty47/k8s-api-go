CREATE TABLE IF NOT EXISTS users (
     id serial NOT NULL,
     user_name VARCHAR(150) NOT NULL UNIQUE,
     role VARCHAR(150) NOT NULL,
     CONSTRAINT pk_users PRIMARY KEY(id)
);
