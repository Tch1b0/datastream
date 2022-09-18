# datastream

stream data without the fear of data corruption

## purpose

datastream transfers data in chunks via a websocket. 

For each chunk transferred its checksum is passed. If the checksum from the client and server doesnt match the chunk is simply getting refetched until it is right.

**datastream is only a proof of concept project**, so it is not made for being used in other projects.

## idea

This project was inspired by the [Storage Migration: Schnell und Einfach](https://chemnitzer.linux-tage.de/2022/de/programm/beitrag/147) presentation by [@RaphaelPour](https://github.com/RaphaelPour) at the [Chemnitzer Linux Tage 2022](https://chemnitzer.linux-tage.de/2022/de/).

The presentation explains the migration process between servers at [Hetzner](https://www.hetzner.com/), and how to optimize it.

## usage

server:
```sh
$ go run main.go <SRC_FILE>
```

client:
```sh
$ python ./client.py <DEST_FILE>
```

## example

**datafetch dialogue** between server and client (transferred byte data is displayed as a string for visualization purposes):
```yml
# the data being transferred is "this is a message from the server"
# the chunk size equals 9 Bytes

# ===[ Chunk 1 ]===

Server: CHUNK
Server: this is a
Server: 10b643f048c7b33a2e734fe583fce2c3

Client: ok

# ===[ Chunk 2 ]===

Server: CHUNK
Server: message f
Server: 3beceb452c297195dd5a05330295e4c9

Client: mismatch

Server: message f
Server: 3beceb452c297195dd5a05330295e4c9

Client: ok

# ===[ Chunk 3 ]===

Server: CHUNK
Server: rom the s
Server: ee13af32b46181b7c664fe0047595cc3

Client: ok

# ===[ Chunk 4 ]===

Server: CHUNK
Server: erver
Server: 57d08d606d880ce78867fb48c7f556ff

Client: ok

# ===[ End ]===

Server: EOF
```
