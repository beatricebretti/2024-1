# Laboratorio 4 - Grupo 13

## Integrantes
- Beatrice Valdes - 201941556-5
- Tomas Garreton - 201823565-2

# Ejecución 
- `make all` para compilar el código y construir las imágenes de los contenedores Docker.
- Para ejecutar el DataNode, Director y NameNode:
     ```
     make run-datanode
     make run-director
     make run-namenode
     ```
- Para ejecutar el código de los mercenarios en las máquinas virtuales correspondientes:
  - Tenemos las siguientes maquinas virtuales `dist050`, `dist051` y `dist052`, en las que se debe ejecutar `make run-mercenarios`.
- `make clean` para detener los servicios de los contenedores Docker.






