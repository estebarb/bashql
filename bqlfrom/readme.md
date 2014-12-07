# bqlfrom

## Nombre
: bqlfrom - Lee un archivo CSV, lo normaliza y lo escribe en stdout.

## Sinopsis
: bqlfrom [opciones] archivo

## Descripción
: Lee un archivo CSV según el formato especificado en las opciones (o bien el formato estandar), lo transforma al formato estandar y lo escribe en stdout.

## Opciones

* -d Indica el caracter que delimita las columnas del archivo CSV. Ejm: «-d ';'» o «-d=';'». Por defecto se utiliza una coma («,») como separador.

* -c Indica el caracter de inicio de un comentario. Las líneas comentadas son ignoradas por el lector de CSV. Ejm: «-c '%'» o «-c='%'». Por defecto los CSV no tienen comentarios.

## Ejemplo
```
$ cat personal.csv
% Personal de la tienda
id;nombre;edad;puesto
1;Antonio;25;Cajas
2;María;26;Cajas
$ bqlfrom -d ';' -c='%' personal.csv
id,nombre,edad,puesto
1,Antonio,25,Cajas
2,María,26,Cajas
```