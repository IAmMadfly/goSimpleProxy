import socket
import time

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

sock.sendto(b"Some data", ("127.0.0.1", 5760))
sock.sendto(b"Some data again", ("127.0.0.1", 5770))
print("Data sent!")

counter = 0

while True:
    sock.sendto(
        bytes("MOAR DATA {}".format(counter), "utf-8"),
        ("127.0.0.1", 5760)
    )
    counter += 1
    time.sleep(1)
