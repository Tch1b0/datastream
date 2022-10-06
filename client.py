import sys
from binascii import crc32
from websocket import create_connection

def arg_or_default(idx: int, default: str) -> str:
    """
    :returns: the argument at the given index or the default value
    """
    return sys.argv[idx] if len(sys.argv) > idx else default

def short_hex(number: int) -> str:
    """
    :returns: the hexadecimal representation of the number :number:

    :example:
    ```py
    short_hex(18) # => "3f"
    ```
    """
    return hex(number)[2:]

base_url = arg_or_default(1, "localhost:8080")
destination_file = arg_or_default(2, "")

ws = create_connection(f"ws://{base_url}/ws")
data = bytearray()

# receive chunks while end-of-file wasn't reached
while ws.recv() != "EOF":
    print("receiving chunk")
    # keep requesting chunk until it is transferred correctly
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

# is the streamed data a binary or text content?
is_bin = False
text_content: str

try:
    text_content = data.decode()
except:
    # the streamed data can't be decoded to text => it has to be a binary
    is_bin = True

if destination_file != "":
    with open(sys.argv[1], "w" if not is_bin else "wb") as f:
        f.write(text_content if not is_bin else data) # type: ignore
else:
    print(text_content if not is_bin else data) # type: ignore

