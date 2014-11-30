# BashQL

BashQL es un conjunto de herramientas que permiten hacer consultas
a archivos CSV.

En lugar de crear un programa enorme, con su propio lenguaje de consultas,
BashQL consiste en una serie de programas pequeños, que cumplen una y solamente
una labor específica, y la cumplen bien.

## bqlfrom

```
bqlfrom [-d delimitador=","] [-c caracter-comentario=""] archivo
```
Lee un archivo CSV Permite cambiar el formato de un CSV al “estandar” (coma como separador y comillas dobles para texto con comas adentro).

## bqlselect
```
bqlselect colA...
```
Recibe un archivo CSV por STDIN y devuelve una nueva tabla que contiene las columnas dadas como parámetro (en el orden dado como parámetro).

## bqlwhere y bqlwhenever
```
bqlwhere columna operador string|numero|columna …
bqlwhenever columna operador string|numero|columna …
```
Devuelve por STDOUT las filas que cumplen la condición. Cada condición se concatena con un AND si es bqlwhere y con un OR si es bqlwhenever.

### Operadores disponibles:

* c=,<,>, !=, >= y >= compara números o strings
* c=, c<, c>, c!=, c>= y c>= compara números o strings contra una columna
* like, unlike compara contra una regex



## bqljoin
```
bqljoin [-d1 delim] [-d2 delim] csv1 columna csv2 columna 
bqljoin [-d delim] columnaSTDIN csv1 columna 
```
El primero lee dos CSV y hace el join por la columna y los tira a STDOUT
El segundo toma una de las tablas desde STDIN

## bqlgroupby
```
bqlgroupby (-g colAgrupar)+ [-c columnaReducir -f funcionReduccion]
```
Agrupa una tabla por la columnaGroup y acumula el resto de las columnas según el acumulador dado…

## Acumuladores: sum,count, avg, max, min, distinct
Serían programas que reciben un montón de valores por STDIN y devuelven un único valor por STDOUT cuando la entrada se acaba.

#Ejemplos
```
bqljoin bqljoin/personal.csv carnet bqljoin/carreras.csv carnet
bqlfrom bqlwhere/example.csv | bqlselect carrera edad | bqlgroupby -g carrera -c edad -f bqlcount
bqlfrom bqlwhere/example.csv | bqlselect carrera edad | bqlgroupby -g carrera -c edad -f bqlavg
```
