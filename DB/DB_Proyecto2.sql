select * from Temporada;
select * from Cliente;
select * from Membresia;
select * from Prueba;
insert into Pureba values(1,"Gracias");
update membresia set idTipoMembresia=0 where idMembresia=0;
ALTER SESSION SET nls_date_format = 'DD-MM-YYYY HH24:MI';
select max(rownum)  from Cliente;

select Username from Cliente where rownum= 3;
select Username from (select * from Cliente order by idCliente desc) where rownum=1;
create table Cliente (
    Username varchar(500) primary key,
    Password varchar (500) not null,
    Nombre varchar(500),
    Apellido varchar(500),
    Tier int, 
    FechaNac date,
    FechaRegistro date,
    Correo varchar(500),
    Foto varchar(500),    
    constraint fk_Tier foreign key (Tier) references Membresia(idMembresia)
);

alter table Cliente modify Username varchar2 primary key;
select * from Cliente;

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

select * from EstadoMembresia;

create sequence tipoMembresiaSeq start with 1 increment by 1;
create table TipoMembresia(
    idTipomembresia number default tipoMembresiaSeq.nextval,
    TipoMembresia varchar(500),
    Precio binary_double,
    primary key (idTipoMembresia)
);
select * from TipoMembresia;
insert into TipoMembresia (TipoMembresia, Precio) values ('Gold',900.00);
insert into TipoMembresia (TipoMembresia, Precio) values ('Silver',450.00);
insert into TipoMembresia (TipoMembresia, Precio) values ('Bronze',150.00);
insert into TipoMembresia values(0,'No Registrado',0);

create sequence temporadaSec start with 1 increment by 1;
create table Temporada(
    idTemporada number default temporadaSec.nextval,
    FechaIni date,
    FechaFin date,
    Nombre varchar (500),
    primary key (idTemporada)
);



drop table Temporada;
select * from Temporada;
insert into  Temporada  values(2,'26-04-2020 15:10', '26-06-2020 15:10', 'TemporadaPrueba4');
update Temporada set FechaIni='01-05-2020 15:08', FechaFin='3-04-2020 15:08' where idTemporada=1;
select idTemporada from (select * from Temporada order by idTemporada desc) where rownum=1;


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

alter table Membresia drop constraint fk_idTemporada;
alter table Membresia add constraint fk_idTemporada  foreign key (idTipoMembresia) references TipoMembresia(idTipoMembresia);
select * from Membresia;
insert into Membresia (idTemporada, idEstadoMembresia,idTipoMembresia) values (0,2,1);
insert into Membresia values (0,0,2,1);


select * from Membresia;
insert into Cliente values ('201801263', '201801263', 'Audrie', 'del Cid', 0, '11-05-2000 23:27','26-04-2021 15:08', 'ann.audrie8a@gmail.com', 'foto'); 
insert into Cliente values ('Audrie8a.', 'Rodaudrie', 'Audrie', 'del Cid', 1, '11-05-2000 23:27','26-04-2021 15:08', 'ann.audrie8a@gmail.com', ''); 
insert into Cliente values ('Prueba.', 'Prueba', 'Prueba', 'Prueba', 0, '12-05-2000 23:27','2-05-2021 15:08', 'prueba@gmail.com', '');

Select * from Cliente where Username='201801263' and Password ='201801263';

update cliente set password='Rodaudrie',nombre='Audrie', apellido= 'del Cid', Tier=0, FechaNac='11-05-2000 23:27',FechaRegistro='26-04-2021 15:08', correo='ann.audrie8a@gmail.com', Foto= 'foto2' where username='Audrie8a.';
update cliente set password='201801263' where Username='201801263';
select count (*) from Temporada;
select * from (select * from Membresia order by idMembresia desc) where rownum=1;
select * from Cliente;
truncate table Cliente;
drop table Cliente;
Truncate table Cliente;

--PROCEDIMIENTO LOGIN
create or replace procedure Login_Usuario (
    Usuario in Cliente.Username%TYPE,
    Pass in Cliente.Password%TYPE )
IS
    respuesta varchar2(500);
    total int;
BEGIN 
    total:=0;
    respuesta:='Acceso Concedido!';
    Select count(*) into total from Cliente where Username=Usuario and Password=Pass;
    if total =0 then
        respuesta:='Acceso Denegado! No hay usuarios registrados con los datos ingresados';
    end if;
      
    DBMS_OUTPUT.PUT_LINE(respuesta);
   
END;
Set serverout on;
select * from Cliente where Username='201801263';
select * from Membresia;
select Membresia.idTipoMembresia from Cliente, Membresia where Cliente.Username='Audrie8a.' and Membresia.idMembresia=Cliente.Tier;
select idMembresia from (select * from Membresia order by idMembresia desc) where rownum=1;

-- PROCEDIMIENTO MEMBRESIA (TipoMembresia,EstadoMembresia y Usuario)
create or replace procedure Membresia_Usuario (
    idTipoMembresiaVar in Membresia.idTipoMembresia%TYPE,
    idEstadoMembresiaVar in Membresia.idEstadoMembresia%TYPE,
    Usuario in Cliente.Username%Type)
IS
    respuesta varchar2(500);
    idTemp int;
    idTipoMemb int;
    idMemb int;
BEGIN 
    idTemp:=0;
    select idTemporada into idTemp from (select * from Temporada order by idTemporada desc) where rownum=1;
    DBMS_OUTPUT.PUT_LINE('idTemporada: ' || idTemp);
    --Si está activa se actualiza (update) o se inserta
    if idEstadoMembresiaVar =1 then
        select Membresia.idTipoMembresia into idTipoMemb from Cliente, Membresia where Cliente.Username=Usuario and Membresia.idMembresia=Cliente.Tier;
        DBMS_OUTPUT.PUT_LINE('idTipoMembresia: ' || idTipoMemb);
        --Si Tipo Membresia es 0, se debe crear una nueva membresia
        if idTipoMemb =0 then
            insert into Membresia (idTemporada, idEstadoMembresia,idTipoMembresia) values (idTemp,1,idTipoMembresiaVar);
            select idMembresia into idMemb from (select * from Membresia order by idMembresia desc) where rownum=1;
            update cliente set Tier=idMemb where Username=Usuario;
        else  --Si Tipo Membresia es 1, se debe actualizar membresia            
            select Tier into idMemb from Cliente where Username=Usuario;
            DBMS_OUTPUT.PUT_LINE('idMembresia: ' || idMemb);
            update Membresia set idTipoMembresia=idTipoMembresiaVar, idTemporada=idTemp where idMembresia=idMemb;
            update Membresia set idTemporada=idTemp where idMembresia=idMemb;
        end if;
    elsif idEstadoMembresiaVar =0 then --Si EstadoMembresia=0, la estan cancelando
        select Tier into idMemb from Cliente where Username=Usuario;
        update Membresia set idEstadoMembresia=idEstadoMembresiaVar where idMembresia=idMemb;
    end if;
    respuesta:='Procedimiento Membresia Finalizado Correctamente!';
    DBMS_OUTPUT.PUT_LINE(respuesta);
   
END;
exec login_usuario('Audrie8a.','Rodaudrie');
exec Membresia_Usuario(1,1,'Prueba.');
select Membresia.idTipoMembresia  from Cliente, Membresia where Cliente.Username='Prueba.' and Membresia.idMembresia=Cliente.Tier;
select Username,Tier from Cliente;


create or replace trigger RegistroUsuarios 
after insert on A 
for each row
enable
declare
begin
end;

commit;