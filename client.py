import sys
from binascii import crc32
from websocket import create_connection

ws = create_connection("ws://localhost:8080/ws")
data = bytearray()

def short_hex(number: int) -> str:
    return hex(number)[2:]

while ws.recv() != "EOF":
    print("receiving chunk")
    while True:
        content: bytearray = ws.recv() # type: ignore
        checksum: str = ws.recv()
        if (generated_sum := short_hex(crc32(content))) != checksum:
            print(f"mismatch ({checksum} != {generated_sum}) => refetching")
            ws.send("mismatch")
        else:
            print("chunk ok")
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

