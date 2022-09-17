# datastream

securely donwload data

## purpose

datastream transfers data in chunks via a socket. For each chunk transfered it's checksum is passed. If the checksum from the client and server doesnt match the chunk is simply getting refetched until it is right.

## example

**datafetch dialogue** between server and client (transferred data is displayed as a string for visualization purposes):
```yml
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
