// package: scarab
// file: scarab-ui.proto

var scarab_ui_pb = require("./scarab-ui_pb");
var scarab_common_pb = require("./scarab-common_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var ManagerUI = (function () {
  function ManagerUI() {}
  ManagerUI.serviceName = "scarab.ManagerUI";
  return ManagerUI;
}());

ManagerUI.StartJob = {
  methodName: "StartJob",
  service: ManagerUI,
  requestStream: false,
  responseStream: false,
  requestType: scarab_common_pb.StartJobRequest,
  responseType: scarab_common_pb.StartJobResponse
};

ManagerUI.StopJob = {
  methodName: "StopJob",
  service: ManagerUI,
  requestStream: false,
  responseStream: false,
  requestType: scarab_ui_pb.StopJobRequest,
  responseType: scarab_ui_pb.StopJobResponse
};

ManagerUI.ListProfiles = {
  methodName: "ListProfiles",
  service: ManagerUI,
  requestStream: false,
  responseStream: false,
  requestType: scarab_ui_pb.ListProfilesRequest,
  responseType: scarab_ui_pb.ListProfilesResponse
};

ManagerUI.WatchActiveJobs = {
  methodName: "WatchActiveJobs",
  service: ManagerUI,
  requestStream: false,
  responseStream: true,
  requestType: scarab_ui_pb.WatchActiveJobsRequest,
  responseType: scarab_ui_pb.WatchActiveJobsResponse
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

ManagerUIClient.prototype.stopJob = function stopJob(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(ManagerUI.StopJob, {
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

