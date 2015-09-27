var console = require("console"),
my_http = require("http");

my_http.createServer(function(request,response){
    console.log("Got kicked at:", request.url);
    response.writeHeader(200, {"Content-Type": "text/plain"});
    response.write("Hello World");
    response.end();
}).listen(8080);

console.log("Server Running on 8080");
