const {HelloRequest, HelloReply} = require('../pb/js/helloworld_pb.js');
const {GreeterClient} = require('../pb/js/helloworld_grpc_web_pb.js');

var client = new GreeterClient('http://0.0.0.0:8080');

var request = new HelloRequest();
request.setName('World');

client.sayHello(request, {}, (err, response) => {
  console.log(response.getMessage());
});
