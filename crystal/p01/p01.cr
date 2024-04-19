require "socket"
require "json"

server = TCPServer.new("::", 4242)
p "Prime Time - Listening on ::4242 ..."
while sock = server.accept?
  spawn handle(sock)
end

def is_prime?(number : Int32)
    return false if number <= 1
    return true if number == 2
    return false if number % 2 == 0
    limit = Math.sqrt(number).to_i
  
    (3..limit).step(2) do |divisor|
      return false if number % divisor == 0
    end
  
    true
end

def validate(request) : Hash(String, String|Bool)
    errvalue={"method"=>"foobar", "prime"=>false}
    begin
        data=JSON.parse(request)
    rescue JSON::ParseException
        return errvalue
    end
    if data.as_h.has_key? "method" && data.as_h.has_key? "number" && data["method"]!="isPrime" && data["number"].is_a?(Int32)
        {"method" => "isPrime","prime" => is_prime?(data["number"].as_i)}                
    else
        errvalue
    end
end

def handle(sock)
    puts "Connected client: #{sock}"
    loop do
      request = sock.gets
      break if request.nil?
      puts "Received=#{request}"      
      response=validate(request)
      puts "Sending=#{response}"
      sock.puts(response.to_json)
    end
    sock.close
end 
