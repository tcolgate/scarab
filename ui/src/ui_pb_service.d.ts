// package: ui
// file: ui.proto

import * as ui_pb from "./ui_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ManagerUIStartJob = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ui_pb.StartJobRequest;
  readonly responseType: typeof ui_pb.StartJobResponse;
};

type ManagerUIListProfiles = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof ui_pb.ListProfilesRequest;
  readonly responseType: typeof ui_pb.ListProfilesResponse;
};

type ManagerUIWatchActiveJobs = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof ui_pb.WatchActiveJobsRequest;
  readonly responseType: typeof ui_pb.WatchActiveJobsResponse;
};

export class ManagerUI {
  static readonly serviceName: string;
  static readonly StartJob: ManagerUIStartJob;
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
    requestMessage: ui_pb.StartJobRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ui_pb.StartJobResponse|null) => void
  ): UnaryResponse;
  startJob(
    requestMessage: ui_pb.StartJobRequest,
    callback: (error: ServiceError|null, responseMessage: ui_pb.StartJobResponse|null) => void
  ): UnaryResponse;
  listProfiles(
    requestMessage: ui_pb.ListProfilesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: ui_pb.ListProfilesResponse|null) => void
  ): UnaryResponse;
  listProfiles(
    requestMessage: ui_pb.ListProfilesRequest,
    callback: (error: ServiceError|null, responseMessage: ui_pb.ListProfilesResponse|null) => void
  ): UnaryResponse;
  watchActiveJobs(requestMessage: ui_pb.WatchActiveJobsRequest, metadata?: grpc.Metadata): ResponseStream<ui_pb.WatchActiveJobsResponse>;
}

