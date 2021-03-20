// package: scarab
// file: scarab.proto

import * as scarab_pb from "./scarab_pb";
import {grpc} from "@improbable-eng/grpc-web";

type ManagerRegisterProfile = {
  readonly methodName: string;
  readonly service: typeof Manager;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof scarab_pb.RegisterProfileRequest;
  readonly responseType: typeof scarab_pb.KeepAlive;
};

export class Manager {
  static readonly serviceName: string;
  static readonly RegisterProfile: ManagerRegisterProfile;
}

type WorkerReportLoad = {
  readonly methodName: string;
  readonly service: typeof Worker;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof scarab_pb.ReportLoadRequest;
  readonly responseType: typeof scarab_pb.LoadMetrics;
};

type WorkerRunJob = {
  readonly methodName: string;
  readonly service: typeof Worker;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof scarab_pb.RunJobRequest;
  readonly responseType: typeof scarab_pb.JobMetrics;
};

export class Worker {
  static readonly serviceName: string;
  static readonly ReportLoad: WorkerReportLoad;
  static readonly RunJob: WorkerRunJob;
}

type ManagerUIStartJob = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scarab_pb.StartJobRequest;
  readonly responseType: typeof scarab_pb.StartJobResponse;
};

type ManagerUIListProfiles = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof scarab_pb.ListProfilesRequest;
  readonly responseType: typeof scarab_pb.ListProfilesResponse;
};

type ManagerUIWatchActiveJobs = {
  readonly methodName: string;
  readonly service: typeof ManagerUI;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof scarab_pb.WatchActiveJobsRequest;
  readonly responseType: typeof scarab_pb.WatchActiveJobsResponse;
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

export class ManagerClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  registerProfile(requestMessage: scarab_pb.RegisterProfileRequest, metadata?: grpc.Metadata): ResponseStream<scarab_pb.KeepAlive>;
}

export class WorkerClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  reportLoad(requestMessage: scarab_pb.ReportLoadRequest, metadata?: grpc.Metadata): ResponseStream<scarab_pb.LoadMetrics>;
  runJob(metadata?: grpc.Metadata): BidirectionalStream<scarab_pb.RunJobRequest, scarab_pb.JobMetrics>;
}

export class ManagerUIClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  startJob(
    requestMessage: scarab_pb.StartJobRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scarab_pb.StartJobResponse|null) => void
  ): UnaryResponse;
  startJob(
    requestMessage: scarab_pb.StartJobRequest,
    callback: (error: ServiceError|null, responseMessage: scarab_pb.StartJobResponse|null) => void
  ): UnaryResponse;
  listProfiles(
    requestMessage: scarab_pb.ListProfilesRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: scarab_pb.ListProfilesResponse|null) => void
  ): UnaryResponse;
  listProfiles(
    requestMessage: scarab_pb.ListProfilesRequest,
    callback: (error: ServiceError|null, responseMessage: scarab_pb.ListProfilesResponse|null) => void
  ): UnaryResponse;
  watchActiveJobs(requestMessage: scarab_pb.WatchActiveJobsRequest, metadata?: grpc.Metadata): ResponseStream<scarab_pb.WatchActiveJobsResponse>;
}

