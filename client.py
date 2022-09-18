import sys
from hashlib import md5
from websocket import create_connection

ws = create_connection("ws://localhost:8080/ws")
data = bytearray()

while ws.recv() != "EOF":
    while True:
        content: bytearray = ws.recv() # type: ignore
        checksum: str = ws.recv()
        if (generated_sum := md5(content).hexdigest()) != checksum:
            ws.send("mismatch")
        else:
            ws.send("ok")
            data.extend(content)
            break

ws.close()

is_bin = False
text_content: str

try:
    text_content = data.decode()
except:
    is_bin = True

if len(sys.argv) > 1:
    with open(sys.argv[1], "w" if not is_bin else "wb") as f:
        f.write(text_content if not is_bin else data) # type: ignore
else:
    print(text_content if not is_bin else data) # type: ignore

