select * from Temporada;
select Foto from Cliente where Username='Audrie8a.';
select * from Membresia;
select * from TipoMembresia;
select * from EstadoMembresia;
select * from Fase;
select * from Jornada;
select * from EventoDeportivo;
select * from Deporte;
select * from Prediccion;
select * from Quinela;

update Cliente set Foto='descarga.jpeg' where Username='Audrie8a.';
select length ((select Foto from Cliente where Username='Audrie8a.')) from dual;
--tamaString:= tamaString-2;
select substr((select Foto from Cliente where Username='Audrie8a.'),13,(select length ((select Foto from Cliente where Username='Audrie8a.')) from dual)) from dual;

drop sequence QuinelaSeq;
drop sequence PrediccionSeq;
drop sequence EventoDSeq;
drop sequence DeporteSeq;
drop sequence jornadaSec;
drop sequence FaseSeq;
drop sequence ClienteSeq;
drop sequence tipoMembresiaSeq;
drop sequence temporadaSec;
drop sequence membresiaSeq;
drop sequence estadoMembresiaSeq;
alter table Quinela drop constraint fk_Prediccion;
alter table Quinela drop constraint fk_UsuarioQuinela;
alter table Prediccion drop constraint fk_EventoDeportivo;
alter table EventoDeportivo drop constraint fk_JornadaEvento;
alter table EventoDeportivo drop constraint fk_Deporte;
alter table Cliente drop constraint fk_Tier;
alter table Membresia drop constraint fk_idTemporada;
alter table Membresia drop constraint fk_idEstadoMembresia;
alter table Membresia drop constraint fk_idTipoMembresia;
alter table Temporada drop constraint fk_idFase;
alter table Jornada drop constraint fk_idFaseJornada;
alter table Jornada drop constraint fk_TemporadaJornada;
drop table Deporte;
drop table EventoDeportivo;
drop table Jornada;
drop table Membresia;
drop table Cliente;
drop table TipoMembresia;
drop table EstadoMembresia;
drop table Temporada;
drop table Fase;
drop table Prediccion;
drop table Quinela;

commit;

ALTER SESSION SET nls_date_format = 'DD/MM/YYYY HH24:MI';




create sequence estadoMembresiaSeq start with 1 increment by 1;
create table EstadoMembresia(
    idEstadoMembresia number default estadoMembresiaSeq.nextval,
    Estado varchar(500),
    Descripcion varchar(500),
    primary key (idEstadoMembresia)
);

insert into EstadoMembresia (Estado, Descripcion) values ('Activa','La membresía está al día en pagos');
insert into EstadoMembresia (Estado, Descripcion) values ('Inactiva','Falta de pago');
insert into EstadoMembresia values (0, 'Cancelada', 'Cancelaron Membresia');

create sequence tipoMembresiaSeq start with 1 increment by 1;
create table TipoMembresia(
    idTipomembresia number default tipoMembresiaSeq.nextval,
    TipoM varchar(500),
    Precio binary_double,
    primary key (idTipoMembresia)
);

select * from TipoMembresia;
insert into TipoMembresia (TipoM, Precio) values ('Gold',900.00);
insert into TipoMembresia (TipoM, Precio) values ('Silver',450.00);
insert into TipoMembresia (TipoM, Precio) values ('Bronze',150.00);
insert into TipoMembresia values(0,'No Registrado',0);


create sequence FaseSeq start with 1 increment by 1;
create table Fase (
    idFase number default FaseSeq.nextval,
    NombreFase varchar(500),
    Descripcion varchar(500),
    primary key(idFase)
);

insert into Fase (NombreFase, Descripcion) values ('Activa','Una jornada se considera activa inmediatamente luego de ser creada. Al momento de su creación la
duración por defecto de la jornada será de 1 semana. El administrador puede cambiar la fecha de finalización
en cualquier momento, media vez la jornada se encuentre activa.');
insert into Fase (NombreFase, Descripcion) values ('Calculo','Durante esta fase de la jornada se les mostrará a los clientes el mensaje de “Calculando” y se
refiere al perı́odo de tiempo en que una jornada ha finalizado su perı́odo activo, pero el administrador aún
no ha ingresado los resultados reales de cada evento deportivo en la jornada.');
insert into Fase (NombreFase, Descripcion) values ('Finalizada','Esta fase ocurre automáticamente tras haber ingresado los resultados reales de todos los eventos
deportivos de la jornada.');

create sequence temporadaSec start with 1 increment by 1;
create table Temporada(
    idTemporada number default temporadaSec.nextval,
    FechaIni date,
    FechaFin date,
    Nombre varchar (500),
    idFase int not null,
    primary key (idTemporada),
    constraint fk_idFase foreign key (idFase) references Fase(idFase)
);




select * from Temporada;
insert into  Temporada  values(0,'26/04/2020 15:10', '26/06/2020 15:10', 'TemporadaPrueba',3);

create sequence jornadaSec start with 1 increment by 1;
create table Jornada(
    idJornada number default jornadaSec.nextval,
    FechaIni date,
    FechaFin date,
    Nombre varchar (500),
    idFase int not null,
    idTemporada int not null,
    primary key (idJornada),
    constraint fk_TemporadaJornada foreign key (idTemporada) references Temporada(idTemporada),
    constraint fk_idFaseJornada foreign key (idFase) references Fase(idFase)
);


create sequence membresiaSeq start with 1 increment by 1;
create table Membresia(
    idMembresia number default membresiaSeq.nextval,
    idTemporada int not null,
    idEstadoMembresia int not null,
    idTipoMembresia int not null,
    primary key (idMembresia),
    constraint fk_idTemporada foreign key (idTemporada) references Temporada(idTemporada),
    constraint fk_idEstadoMembresia foreign key (idEstadoMembresia) references EstadoMembresia(idEstadoMembresia),
    constraint fk_idTipoMembresia foreign key (idTipoMembresia) references TipoMembresia(idTipoMembresia)
);

select * from Membresia;
insert into Membresia (idTemporada, idEstadoMembresia,idTipoMembresia) values (0,2,1);
insert into Membresia values (0,0,2,0); --Membresía quemada por defecto


select * from Membresia;

create sequence ClienteSeq start with 1 increment by 1;
create table Cliente (
    idCliente number default ClienteSeq.nextval,
    Username varchar(500) primary key,
    Password varchar (500) not null,
    Nombre varchar(500),
    Apellido varchar(500),
    Tier int, 
    FechaNac date null,
    FechaRegistro date null,
    Correo varchar(500),
    Foto varchar(500),    
    constraint fk_Tier foreign key (Tier) references Membresia(idMembresia)
);


insert into Cliente (Username,Password,Nombre,Apellido,Tier,FechaNac,FechaRegistro,Correo,Foto) values ('201801263', '201801263', 'Audrie', 'del Cid', 0, '11/05/2000 23:27','26/04/2021 15:08', 'ann.audrie8a@gmail.com', 'foto'); 
insert into Cliente (Username,Password,Nombre,Apellido,Tier,FechaNac,FechaRegistro,Correo,Foto) values ('Audrie8a.', 'Rodaudrie', 'Audrie', 'del Cid', 1, '11/05/2000 23:27','26/04/2021 15:08', 'ann.audrie8a@gmail.com', ''); 
insert into Cliente (Username,Password,Nombre,Apellido,Tier,FechaNac,FechaRegistro,Correo,Foto) values ('Prueba.', 'Prueba', 'Prueba', 'Prueba', 0, '12/05/2000 23:27','2/05/2021 15:08', 'prueba@gmail.com', '');
insert into Cliente (Username,Password,Nombre,Apellido,Tier,FechaNac,FechaRegistro,Correo,Foto) values ('Prueba2.', 'Prueba2', 'Prueba2', 'Prueba2', 0, '12/05/2000 23:27','2/05/2021 15:08', 'prueba2@gmail.com', '');
Select * from Cliente where Username='201801263' and Password ='201801263';
Select * from Cliente where Username='Audrie8a.' and Password= 'Rodaudrie';

create sequence DeporteSeq start with 1 increment by 1;
create table Deporte(
    idDeporte number default DeporteSeq.nextval,
    Nombre varchar(500),
    Imagen varchar(500),
    Color varchar(500),
    primary key (idDeporte)
);

create sequence EventoDSeq start with 1 increment by 1;
create table EventoDeportivo(
    idEventoDeportivo number default EventoDSeq.nextval,
    idJoranadaED int not null,
    idDeporte int not null,
    Fecha date,
    NombreVisitante varchar(500),
    NombreLocal varchar(500),
    ResultadoVisitante int,
    ResultadoLocal int,
    primary key (idEventoDeportivo), 
    constraint fk_JornadaEvento foreign key (idJoranadaED) references Jornada(idJornada),
    constraint fk_Deporte foreign key (idDeporte) references Deporte(idDeporte)    
);

create sequence PrediccionSeq start with 1 increment by 1;
create table Prediccion(
    idPrediccion number default PrediccionSeq.nextval,
    Local int, 
    Visitante int, 
    idEventoDeportivo int not null,
    primary key (idPrediccion),
    constraint fk_EventoDeportivo foreign key (idEventoDeportivo) references EventoDeportivo(idEventoDeportivo)
);

create sequence QuinelaSeq start with 1 increment by 1;
create table Quinela(
    idQuinela  number default QuinelaSeq.nextval,
    Precio binary_double, 
    idPrediccion int not null,
    UsernameCliente varchar(500),
    Puntaje int,
    primary key (idQuinela),
    constraint fk_Prediccion foreign key (idPrediccion) references Prediccion(idPrediccion),
    constraint fk_UsuarioQuinela foreign key (UsernameCliente) references Cliente(Username)
);

commit;

