/**
 * @fileoverview gRPC-Web generated client stub for ui
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.ui = require('./ui_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ui.ManagerUIClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.ui.ManagerUIPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ui.StartJobRequest,
 *   !proto.ui.StartJobResponse>}
 */
const methodDescriptor_ManagerUI_StartJob = new grpc.web.MethodDescriptor(
  '/ui.ManagerUI/StartJob',
  grpc.web.MethodType.UNARY,
  proto.ui.StartJobRequest,
  proto.ui.StartJobResponse,
  /**
   * @param {!proto.ui.StartJobRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ui.StartJobResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ui.StartJobRequest,
 *   !proto.ui.StartJobResponse>}
 */
const methodInfo_ManagerUI_StartJob = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ui.StartJobResponse,
  /**
   * @param {!proto.ui.StartJobRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ui.StartJobResponse.deserializeBinary
);


/**
 * @param {!proto.ui.StartJobRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ui.StartJobResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ui.StartJobResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ui.ManagerUIClient.prototype.startJob =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ui.ManagerUI/StartJob',
      request,
      metadata || {},
      methodDescriptor_ManagerUI_StartJob,
      callback);
};


/**
 * @param {!proto.ui.StartJobRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ui.StartJobResponse>}
 *     Promise that resolves to the response
 */
proto.ui.ManagerUIPromiseClient.prototype.startJob =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ui.ManagerUI/StartJob',
      request,
      metadata || {},
      methodDescriptor_ManagerUI_StartJob);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.ui.ListProfilesRequest,
 *   !proto.ui.ListProfilesResponse>}
 */
const methodDescriptor_ManagerUI_ListProfiles = new grpc.web.MethodDescriptor(
  '/ui.ManagerUI/ListProfiles',
  grpc.web.MethodType.UNARY,
  proto.ui.ListProfilesRequest,
  proto.ui.ListProfilesResponse,
  /**
   * @param {!proto.ui.ListProfilesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ui.ListProfilesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.ui.ListProfilesRequest,
 *   !proto.ui.ListProfilesResponse>}
 */
const methodInfo_ManagerUI_ListProfiles = new grpc.web.AbstractClientBase.MethodInfo(
  proto.ui.ListProfilesResponse,
  /**
   * @param {!proto.ui.ListProfilesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.ui.ListProfilesResponse.deserializeBinary
);


/**
 * @param {!proto.ui.ListProfilesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.ui.ListProfilesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.ui.ListProfilesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.ui.ManagerUIClient.prototype.listProfiles =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/ui.ManagerUI/ListProfiles',
      request,
      metadata || {},
      methodDescriptor_ManagerUI_ListProfiles,
      callback);
};


/**
 * @param {!proto.ui.ListProfilesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.ui.ListProfilesResponse>}
 *     Promise that resolves to the response
 */
proto.ui.ManagerUIPromiseClient.prototype.listProfiles =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/ui.ManagerUI/ListProfiles',
      request,
      metadata || {},
      methodDescriptor_ManagerUI_ListProfiles);
};


module.exports = proto.ui;

