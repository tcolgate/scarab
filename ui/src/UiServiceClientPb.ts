/**
 * @fileoverview gRPC-Web generated client stub for ui
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as ui_pb from './ui_pb';


export class ManagerUIClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoStartJob = new grpcWeb.AbstractClientBase.MethodInfo(
    ui_pb.StartJobResponse,
    (request: ui_pb.StartJobRequest) => {
      return request.serializeBinary();
    },
    ui_pb.StartJobResponse.deserializeBinary
  );

  startJob(
    request: ui_pb.StartJobRequest,
    metadata: grpcWeb.Metadata | null): Promise<ui_pb.StartJobResponse>;

  startJob(
    request: ui_pb.StartJobRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ui_pb.StartJobResponse) => void): grpcWeb.ClientReadableStream<ui_pb.StartJobResponse>;

  startJob(
    request: ui_pb.StartJobRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: ui_pb.StartJobResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/ui.ManagerUI/StartJob',
        request,
        metadata || {},
        this.methodInfoStartJob,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/ui.ManagerUI/StartJob',
    request,
    metadata || {},
    this.methodInfoStartJob);
  }

  methodInfoListProfiles = new grpcWeb.AbstractClientBase.MethodInfo(
    ui_pb.ListProfilesResponse,
    (request: ui_pb.ListProfilesRequest) => {
      return request.serializeBinary();
    },
    ui_pb.ListProfilesResponse.deserializeBinary
  );

  listProfiles(
    request: ui_pb.ListProfilesRequest,
    metadata: grpcWeb.Metadata | null): Promise<ui_pb.ListProfilesResponse>;

  listProfiles(
    request: ui_pb.ListProfilesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: ui_pb.ListProfilesResponse) => void): grpcWeb.ClientReadableStream<ui_pb.ListProfilesResponse>;

  listProfiles(
    request: ui_pb.ListProfilesRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: ui_pb.ListProfilesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/ui.ManagerUI/ListProfiles',
        request,
        metadata || {},
        this.methodInfoListProfiles,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/ui.ManagerUI/ListProfiles',
    request,
    metadata || {},
    this.methodInfoListProfiles);
  }

}

