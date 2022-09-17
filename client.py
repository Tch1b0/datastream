from hashlib import md5
from websocket import create_connection

ws = create_connection("ws://localhost:8080/ws")
data = bytearray()
print("Start receiving...")
while ws.recv() != "EOF":
    print("Start receiving chunk...")
    while True:
        content: bytearray = ws.recv() # type: ignore
        checksum: str = ws.recv()
        print(f"Received chunk with content \"{content}\" and checksum {checksum}")
        if (generated_sum := md5(content).hexdigest()) != checksum:
            print(f"Checksum {generated_sum} not matching => Request new chunk")
            ws.send("mismatch")
        else:
            print("Chunk matching")
            ws.send("ok")
            data.extend(content)
            break

ws.close()

print(data)
