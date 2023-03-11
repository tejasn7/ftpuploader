import sys
from pyftpdlib.authorizers import DummyAuthorizer
from pyftpdlib.handlers import FTPHandler
from pyftpdlib.servers import FTPServer

try:
    port = int(sys.argv[1])
except IndexError as e:
    port = 2121

authorizer = DummyAuthorizer()
authorizer.add_anonymous("./", perm="elradfmwM")
handler = FTPHandler
handler.authorizer = authorizer

handler.banner = "Server Ready.."

address = ("",port)
server = FTPServer(address, handler)

server.max_cons = 10
server.serve_forever()