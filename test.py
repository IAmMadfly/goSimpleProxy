import socket
import time


s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

counter1 = 0
counter2 = 0

while True:
    counter1 += 1
    counter2 += 2
    s.sendto(bytes("Helo wold: {}".format(counter1), "utf-8"), ("127.0.0.1", 5760))
    s.sendto(bytes("Helo wold: {}".format(counter2), "utf-8"), ("127.0.0.1", 5770))
    time.sleep(1.2)
