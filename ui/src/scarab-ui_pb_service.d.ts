// package: scarab
// file: scarab-ui.proto

import * as scarab_ui_pb from "./scarab-ui_pb";
import * as scarab_common_pb from "./scarab-common_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ManagerUIStartJob = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scarab_common_pb.StartJobRequest;
  readonly responseType: typeof scarab_common_pb.StartJobResponse;
};

type ManagerUIStopJob = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scarab_ui_pb.StopJobRequest;
  readonly responseType: typeof scarab_ui_pb.StopJobResponse;
};

type ManagerUIListProfiles = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scarab_ui_pb.ListProfilesRequest;
  readonly responseType: typeof scarab_ui_pb.ListProfilesResponse;
};

type ManagerUIWatchActiveJobs = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof scarab_ui_pb.WatchActiveJobsRequest;
  readonly responseType: typeof scarab_ui_pb.WatchActiveJobsResponse;
};

export class ManagerUI {
  static readonly serviceName: string;
  static readonly StartJob: ManagerUIStartJob;
  static readonly StopJob: ManagerUIStopJob;
  static readonly ListProfiles: ManagerUIListProfiles;
  static readonly WatchActiveJobs: ManagerUIWatchActiveJobs;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class ManagerUIClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  startJob(
    requestMessage: scarab_common_pb.StartJobRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scarab_common_pb.StartJobResponse|null) => void
  ): UnaryResponse;
  startJob(
    requestMessage: scarab_common_pb.StartJobRequest,
    callback: (error: ServiceError|null, responseMessage: scarab_common_pb.StartJobResponse|null) => void
  ): UnaryResponse;
  stopJob(
    requestMessage: scarab_ui_pb.StopJobRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scarab_ui_pb.StopJobResponse|null) => void
  ): UnaryResponse;
  stopJob(
    requestMessage: scarab_ui_pb.StopJobRequest,
    callback: (error: ServiceError|null, responseMessage: scarab_ui_pb.StopJobResponse|null) => void
  ): UnaryResponse;
  listProfiles(
    requestMessage: scarab_ui_pb.ListProfilesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scarab_ui_pb.ListProfilesResponse|null) => void
  ): UnaryResponse;
  listProfiles(
    requestMessage: scarab_ui_pb.ListProfilesRequest,
    callback: (error: ServiceError|null, responseMessage: scarab_ui_pb.ListProfilesResponse|null) => void
  ): UnaryResponse;
  watchActiveJobs(requestMessage: scarab_ui_pb.WatchActiveJobsRequest, metadata?: grpc.Metadata): ResponseStream<scarab_ui_pb.WatchActiveJobsResponse>;
}

