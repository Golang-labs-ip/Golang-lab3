
drop table if exists fullList;
drop table if exists tablet;
drop table if exists telemetry;



create table telemetry(
id int auto_increment,
battery int not null,
deviceTime varchar(60) not null,
timeStamp varchar(60) not null,
currentVideo varchar(60),
primary key(id)
);

create table tablet(
id int auto_increment,
name varchar(60) not null,
primary key(id)
);

create table fullList(
tabletId int not null,
telemetryId int not null,
foreign key (tabletId) references tablet (id),
foreign key (telemetryId) references telemetry (id)
);
