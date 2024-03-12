# Implementation of Google File System / Hadoop
# for educational purposes

## ARCHITECTURE

### client

this is the front end app to interact with the file system,
there will be handlers using cli, grpc servers, and http. The client jobs is to request operation to nameNode and
then act accordingly. When client want to do write operation, it will query the nameNode and retrieve set of dataNode target
and then split the files into multiple blocks and replicate it to dataNode in paralel. When client want to do read operation,
it will query the nameNode to get which dataNodes contain the blocks / chunks and then aggregate those blocks into one unit.


### nameNode

this is the manager of file system cluster. It handles node allocation, metadata of files and directory, and transactions. It can be deployed multiple instances.


### dataNode

this is chunk servers. It stores blocks of byte and serves that to the client. It also responsible to replicate blocks to another dataNode until replication target is achieved

## consul

this is the service registry. It registers the address and ports of nameNode and dataNode services. This service routinely do health checking to each nodes.

## redis

this is in memory db, it is required to store caching, something like metadata caching or leased storage state for each dataNode.

## postgres

this is relational database to store metadata information of files