create table emails
(
	id bigserial,
	email varchar(100) not null
);

comment on table emails is 'Почты';

alter table emails owner to "user";
