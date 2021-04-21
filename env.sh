export DATABASE_DSN=audrie/audrie@localhost:1521/ORCLCDB.localdomain #String de conexión
# export TNS_ADMIN=/home/renato/Oracle/Credentials # Ruta hacia tu wallet
export LD_LIBRARY_PATH=$ORACLE_HOME:/opt/oracle/ #Si tienes Instant Client, reemplaza esta ruta por la dirección en la que descomprimiste las librerías
export PATH=$PATH:$LD_LIBRARY_PATH #No olvides incluir este directorio en tu PATH!

