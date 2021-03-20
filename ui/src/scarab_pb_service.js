// package: scarab
// file: scarab.proto

var scarab_pb = require("./scarab_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Manager = (function () {
  function Manager() {}
  Manager.serviceName = "scarab.Manager";
  return Manager;
}());

Manager.RegisterProfile = {
  methodName: "RegisterProfile",
  service: Manager,
  requestStream: false,
  responseStream: true,
  requestType: scarab_pb.RegisterProfileRequest,
  responseType: scarab_pb.KeepAlive
};

exports.Manager = Manager;

function ManagerClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ManagerClient.prototype.registerProfile = function registerProfile(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Manager.RegisterProfile, {
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

exports.ManagerClient = ManagerClient;

var Worker = (function () {
  function Worker() {}
  Worker.serviceName = "scarab.Worker";
  return Worker;
}());

Worker.ReportLoad = {
  methodName: "ReportLoad",
  service: Worker,
  requestStream: false,
  responseStream: true,
  requestType: scarab_pb.ReportLoadRequest,
  responseType: scarab_pb.LoadMetrics
};

Worker.RunJob = {
  methodName: "RunJob",
  service: Worker,
  requestStream: true,
  responseStream: true,
  requestType: scarab_pb.RunJobRequest,
  responseType: scarab_pb.JobMetrics
};

exports.Worker = Worker;

function WorkerClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

WorkerClient.prototype.reportLoad = function reportLoad(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Worker.ReportLoad, {
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

WorkerClient.prototype.runJob = function runJob(metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.client(Worker.RunJob, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  client.onMessage(function (message) {
    listeners.data.forEach(function (handler) {
      handler(message);
    })
  });
  client.start(metadata);
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.WorkerClient = WorkerClient;

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
  requestType: scarab_pb.StartJobRequest,
  responseType: scarab_pb.StartJobResponse
};

ManagerUI.ListProfiles = {
  methodName: "ListProfiles",
  service: ManagerUI,
  requestStream: false,
  responseStream: false,
  requestType: scarab_pb.ListProfilesRequest,
  responseType: scarab_pb.ListProfilesResponse
};

ManagerUI.WatchActiveJobs = {
  methodName: "WatchActiveJobs",
  service: ManagerUI,
  requestStream: false,
  responseStream: true,
  requestType: scarab_pb.WatchActiveJobsRequest,
  responseType: scarab_pb.WatchActiveJobsResponse
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

