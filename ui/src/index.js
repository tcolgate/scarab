const {ListProfileRequest, ListProfilesResponse} = require('./ui_pb.js');
const {ManagerUIServiceClient} = require('./ui_grpc_web_pb.js');

var uiService = new ManagerUIServiceClient('http://localhost:8080');

var request = new ListProfilesResponse();
request.setMessage('Hello World!');

uiService.echo(request, {}, function(err, response) {
  // ...
});

function component() {
  const element = document.createElement('div');

  // Lodash, currently included via a script, is required for this line to work
  element.innerHTML = _.join(['Hello', 'webpack'], ' ');

  return element;
}

document.body.appendChild(component());
