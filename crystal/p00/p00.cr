require "socket"

server = TCPServer.new("::", 4242)
p "Smoke Test - Listening on ::4242 ..."
while sock = server.accept?
  spawn handle(sock)
end

def handle(sock)
    puts "Connected client: #{sock}"
    buf = Bytes.new(1024)
    loop do
      read_bytes = sock.read(buf)
      break if read_bytes.zero?
      sock.write buf[0...read_bytes]
    end
    sock.close
end 
