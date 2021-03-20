// package: ui
// file: ui.proto

var ui_pb = require("./ui_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ManagerUI = (function () {
  function ManagerUI() {}
  ManagerUI.serviceName = "ui.ManagerUI";
  return ManagerUI;
}());

ManagerUI.StartJob = {
  methodName: "StartJob",
  service: ManagerUI,
  requestStream: false,
  responseStream: false,
  requestType: ui_pb.StartJobRequest,
  responseType: ui_pb.StartJobResponse
};

ManagerUI.ListProfiles = {
  methodName: "ListProfiles",
  service: ManagerUI,
  requestStream: false,
  responseStream: false,
  requestType: ui_pb.ListProfilesRequest,
  responseType: ui_pb.ListProfilesResponse
};

ManagerUI.WatchActiveJobs = {
  methodName: "WatchActiveJobs",
  service: ManagerUI,
  requestStream: false,
  responseStream: true,
  requestType: ui_pb.WatchActiveJobsRequest,
  responseType: ui_pb.WatchActiveJobsResponse
};

exports.ManagerUI = ManagerUI;

function ManagerUIClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ManagerUIClient.prototype.startJob = function startJob(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ManagerUI.StartJob, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ManagerUIClient.prototype.listProfiles = function listProfiles(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ManagerUI.ListProfiles, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

ManagerUIClient.prototype.watchActiveJobs = function watchActiveJobs(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(ManagerUI.WatchActiveJobs, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.ManagerUIClient = ManagerUIClient;
