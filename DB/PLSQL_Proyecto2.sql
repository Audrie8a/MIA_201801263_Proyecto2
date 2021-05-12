Set serverout on;

create or replace procedure CrearUsuario(
    Usuario in Cliente.Username%TYPE,
    Contra in Cliente.Password%TYPE,
    Name in Cliente.Nombre%TYPE,
    Last in Cliente.Apellido%TYPE,
    Memb in Cliente.Tier%TYPE,
    FNAC in Cliente.Nombre%TYPE,
    FREGISTRO in Cliente.Nombre%TYPE,
    email in Cliente.Correo%TYPE,
    photo in Cliente.Foto%TYPE
)
is 
    fechaN date;
    fechaR date;
begin
    select to_date(FNAC, 'DD/MM/YYYY HH24:MI') into fechaN from dual;
    select to_date(FREGISTRO, 'DD/MM/YYYY HH24:MI') into fechaR from dual;
    insert into Cliente (Username, Password, Nombre, Apellido, Tier, FechaNac, FechaRegistro, Correo, Foto) values (Usuario, Contra, Name, Last, Memb, fechaN, fechaR, email, photo);
end;

create or replace procedure UpdateUsuario(
    Usuario in Cliente.Username%TYPE,
    Contra in Cliente.Password%TYPE,
    Name in Cliente.Nombre%TYPE,
    Last in Cliente.Apellido%TYPE,
    --FNAC in Cliente.Nombre%TYPE,
    email in Cliente.Correo%TYPE,
    photo in Cliente.Foto%TYPE
)
is 
    --fechaN date;
begin
    --select to_date(FNAC, 'DD/MM/YYYY HH24:MI') into fechaN from dual;
    update Cliente set password =Contra, Nombre=Name, Apellido= Last, Correo=email, Foto=photo where Username=Usuario;
end;
create or replace procedure Insert_Cliente(
    nombre in Cliente.Nombre%TYPE,
    usuario in Cliente.Username%TYPE, 
    contra in Cliente.Password%TYPE, 
    apellido in Cliente.Apellido%TYPE, 
    correo in Cliente.Correo%TYPE)
is 

begin
    
    insert into Cliente (Username, Password, Nombre,Apellido, Correo) values (usuario,contra,nombre,apellido,correo);  
    
end;

create or replace procedure Insert_Resultados(    
    nombreTemporada in Temporada.Nombre%TYPE,
    tipo in TipoMembresia.TipoM%TYPE)
is
    idTemp int;
    tierC int;
    idMemb int;
    idTipoMemb int;
    usernameC varchar(500);
    contador int;
begin
    tierC:=0;
    --A temporada no le funciona el sequence
    select idTemporada into idTemp from (select * from Temporada order by idTemporada desc) where rownum=1;
    idTemp:=idTemp+1;
    insert into  Temporada (idTemporada, Nombre, idFase) values(idTemp, nombreTemporada,3);
    idTemp:=0;
    select idTemporada into idTemp from (select * from Temporada order by idTemporada desc) where rownum=1;
    select idTipoMembresia into idTipoMemb from TipoMembresia  where TipoM=tipo;
    insert into Membresia (idTemporada, idEstadoMembresia,idTipoMembresia) values (idTemp,2,idTipoMemb);
    select idMembresia into tierC from (select * from Membresia order by idMembresia desc) where rownum=1;
     DBMS_OUTPUT.PUT_LINE(tierC);
    select Username into usernameC from (select * from Cliente order by idCliente desc) where rownum=1;
     DBMS_OUTPUT.PUT_LINE(usernamec);
    update Cliente set Tier=tierC where Username=usernameC;
end;

create or replace procedure Insert_Jornadas (nombreJornada in Jornada.Nombre%TYPE)
is 
    idTemp int;
    respuesta varchar(500);
    contador int;
begin
    select idTemporada into idTemp from (select * from Temporada order by idTemporada desc) where rownum=1;
    select count(*) into contador from Jornada where Nombre=nombreJornada and idTemporada=idTemp;
    if contador=0 then
        insert into Jornada (Nombre,idTemporada,idFase) values (nombreJornada, idTemp, 3);
        respuesta:='Jornada Registrada!';
    else
        respuesta:='Ya existe una jornada con este nombre en la misma Temporada';
    end if;
    
     DBMS_OUTPUT.PUT_LINE(respuesta);
end;


create or replace procedure Insert_EventoDeportivoDeporte(
    NombreDeporte in Deporte.Nombre%TYPE,
    colorDeporte in Deporte.Color%TYPE,
    fechaED in Deporte.Nombre%TYPE,
    VisitanteED in EventoDeportivo.NombreVisitante%TYPE,
    LocalED in EventoDeportivo.NombreLocal%TYPE,
    PreVisitante in Prediccion.Visitante%TYPE,
    PreLocal in Prediccion.Local%TYPE,
    ResVisitante in EventoDeportivo.ResultadoVisitante%TYPE,
    ResLocal in EventoDeportivo.ResultadoLocal%TYPE
    )
is
    idSport int;
    idJor int;
    idED int;
    fechaX date;
    
    Puntos int;
    Ganador int;
    Condicion int;
    PrecioMemb binary_double;
    idPredic int;
    UsuarioCliente varchar(500);
    contador int;
begin
    select to_date(fechaED, 'DD/MM/YYYY HH24:MI') into fechaX from dual;
    select count(*) into contador from Deporte where Nombre=NombreDeporte;
    if contador=0 then
        insert into Deporte (Nombre, Color) values (NombreDeporte, colorDeporte);
        select idDeporte into idSport from (select *from Deporte order by idDeporte desc) where rownum=1;
    else
        select idDeporte into idSport from Deporte where Nombre=NombreDeporte;
    end if;
    
    
    select idJornada into idJor from (select *from Jornada order by idJornada desc) where rownum=1;
    insert into EventoDeportivo (idJoranadaED, idDeporte, Fecha, NombreVisitante, NombreLocal, ResultadoVisitante,ResultadoLocal) values (idJor, idSport, fechaX, VisitanteED, LocalED, ResVisitante, ResLocal);
    select idEventoDeportivo into idED from (select *from EventoDeportivo order by idEventoDeportivo desc) where rownum=1;
    insert into Prediccion (Local, Visitante, idEventoDeportivo) values (PreLocal, PreVisitante, idED);
    
    --Insertando datos a quinela
    DBMS_OUTPUT.PUT_LINE('Visitante R: '||ResVisitante ||' Visitante P:  ' || PreVisitante ||' Local R: ' || ResLocal||' Local P:  '||PreLocal);
    if ResVisitante=PreVisitante and ResLocal=PreLocal then
        Puntos:=10;
        DBMS_OUTPUT.PUT_LINE('Acerto Exacto');
    elsif ResVisitante!=PreVisitante and ResLocal!=PreLocal then
        Puntos:=0;
        DBMS_OUTPUT.PUT_LINE('Empate');
    else 
        DBMS_OUTPUT.PUT_LINE('Otro');
        if ResVisitante> ResLocal then
            DBMS_OUTPUT.PUT_LINE('Ganador: Visitante');
            Ganador:= ResVisitante;
            Condicion:= ResLocal-PreLocal;
            DBMS_OUTPUT.PUT_LINE(Ganador || 'condicion: '||Condicion);
            if Ganador=PreVisitante and (Condicion<=2 or Condicion>=-2)  then
                Puntos:=5;
            elsif Ganador=PreVisitante and (Condicion>2 or Condicion<-2) then
                Puntos:=3;
            else 
                Puntos:=0;
            end if;
        elsif ResVisitante<ResLocal then
            DBMS_OUTPUT.PUT_LINE('Ganador: Local');
            Ganador:=ResLocal;
            Condicion:= ResVisitante-PreVisitante;
            if Ganador=PreLocal and (Condicion<=2 or Condicion>=-2)  then
                Puntos:=5;
            elsif Ganador=PreLocal and (Condicion>2 or Condicion<-2) then
                Puntos:=3;
            else 
                Puntos:=0;
            end if;
        else
            Puntos:=0;
        end if;
        DBMS_OUTPUT.PUT_LINE(Puntos);
        
    end if;
    select Username into UsuarioCliente from (select *from Cliente order by idCliente desc) where rownum=1;
    select TipoMembresia.Precio into PrecioMemb from Membresia, TipoMembresia, Cliente where Cliente.Tier=Membresia.idMembresia and Membresia.idTipoMembresia=TipoMembresia.idTipoMembresia and Cliente.Username=UsuarioCliente;
    select idPrediccion into idPredic from (select *from Prediccion order by idPrediccion desc) where rownum=1;
    insert into Quinela (Precio, idPrediccion, UsernameCliente, Puntaje) values (PrecioMemb,idPredic,UsuarioCliente,Puntos);
    
end;

select To_Number(TipoMembresia.Precio)  from Membresia, TipoMembresia, Cliente where Cliente.Tier=Membresia.idMembresia and Membresia.idTipoMembresia=TipoMembresia.idTipoMembresia and Cliente.Username='S8';

create or replace procedure Update_Jornada(
    fechaInicio in Deporte.Nombre%TYPE,
    fechaFinal in Deporte.Nombre%TYPE
)
is
    idJor int;
    
    fechaY date;
    fechaX date;
begin
    select to_date(fechaInicio, 'DD/MM/YYYY HH24:MI') into fechaX from dual;
    select to_date(fechaFinal, 'DD/MM/YYYY HH24:MI') into fechaX from dual;
    select idJornada into idJor from (select *from Jornada order by idJornada desc) where rownum=1;
    
    update Jornada set FechaIni=fechaY, FechaFin=fechaX where idJornada=idJor;
    
end;

create or replace procedure Update_Temporada(
    fechaInicio in Deporte.Nombre%TYPE,
    fechaFinal in Deporte.Nombre%TYPE
)
is
    idTemp int;
    
    fechaY date;
    fechaX date;
begin
    select to_date(fechaInicio, 'DD/MM/YYYY HH24:MI') into fechaX from dual;
    select to_date(fechaFinal, 'DD/MM/YYYY HH24:MI') into fechaX from dual;
    select idTemporada into idTemp from (select *from Temporada order by idTemporada desc) where rownum=1;
    
    update Temporada set FechaIni=fechaY, FechaFin=fechaX where idTemporada=idTemp;
end;




exec Insert_Cliente('Sergio','S8','SergioC','delCid','seresdo');
exec Insert_Resultados('TempSergio','Silver');
exec Insert_Jornadas('J4');
exec Insert_EventoDeportivoDeporte ('Futbol', 'rojo','13/09/2018 11:42', 'Barca','Real',3,2,3,0);
select Username from (select * from Cliente order by idCliente desc) where rownum=1;
select * from Cliente where username=(select Username from (select * from Cliente order by idCliente desc) where rownum=1);
update Cliente set Tier=4 where Username='S8';
select * from Cliente where Username='A47';
commit;