create database user_management;


create table users (
     user_id serial PRIMARY KEY,
     user_name VARCHAR ( 50 ) UNIQUE NOT NULL,
     password VARCHAR ( 500 ) NOT NULL,
     email VARCHAR ( 255 ) NOT NULL
);


insert into users (username, password, email)
values ('mitun_rahman', 'nemo', 'abc@gmail.com');

select * from users;



