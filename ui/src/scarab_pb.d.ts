// package: scarab
// file: scarab.proto

import * as jspb from "google-protobuf";
import * as metrics_pb from "./metrics_pb";

export class ProfileArgOption extends jspb.Message {
  clearOptionList(): void;
  getOptionList(): Array<string>;
  setOptionList(value: Array<string>): void;
  addOption(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProfileArgOption.AsObject;
  static toObject(includeInstance: boolean, msg: ProfileArgOption): ProfileArgOption.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProfileArgOption, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProfileArgOption;
  static deserializeBinaryFromReader(message: ProfileArgOption, reader: jspb.BinaryReader): ProfileArgOption;
}

export namespace ProfileArgOption {
  export type AsObject = {
    optionList: Array<string>,
  }
}

export class ProfileArg extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  hasString(): boolean;
  clearString(): void;
  getString(): string;
  setString(value: string): void;

  hasFloat(): boolean;
  clearFloat(): void;
  getFloat(): number;
  setFloat(value: number): void;

  hasOption(): boolean;
  clearOption(): void;
  getOption(): ProfileArgOption | undefined;
  setOption(value?: ProfileArgOption): void;

  getDefaultCase(): ProfileArg.DefaultCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProfileArg.AsObject;
  static toObject(includeInstance: boolean, msg: ProfileArg): ProfileArg.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProfileArg, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProfileArg;
  static deserializeBinaryFromReader(message: ProfileArg, reader: jspb.BinaryReader): ProfileArg;
}

export namespace ProfileArg {
  export type AsObject = {
    name: string,
    description: string,
    string: string,
    pb_float: number,
    option?: ProfileArgOption.AsObject,
  }

  export enum DefaultCase {
    DEFAULT_NOT_SET = 0,
    STRING = 3,
    FLOAT = 4,
    OPTION = 5,
  }
}

export class RegisterProfileRequest extends jspb.Message {
  getProfile(): string;
  setProfile(value: string): void;

  getVersion(): string;
  setVersion(value: string): void;

  getAddr(): string;
  setAddr(value: string): void;

  getName(): string;
  setName(value: string): void;

  clearArgsList(): void;
  getArgsList(): Array<ProfileArg>;
  setArgsList(value: Array<ProfileArg>): void;
  addArgs(value?: ProfileArg, index?: number): ProfileArg;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisterProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegisterProfileRequest): RegisterProfileRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisterProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisterProfileRequest;
  static deserializeBinaryFromReader(message: RegisterProfileRequest, reader: jspb.BinaryReader): RegisterProfileRequest;
}

export namespace RegisterProfileRequest {
  export type AsObject = {
    profile: string,
    version: string,
    addr: string,
    name: string,
    argsList: Array<ProfileArg.AsObject>,
  }
}

export class KeepAlive extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): KeepAlive.AsObject;
  static toObject(includeInstance: boolean, msg: KeepAlive): KeepAlive.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: KeepAlive, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): KeepAlive;
  static deserializeBinaryFromReader(message: KeepAlive, reader: jspb.BinaryReader): KeepAlive;
}

export namespace KeepAlive {
  export type AsObject = {
  }
}

export class ReportLoadRequest extends jspb.Message {
  getInterval(): number;
  setInterval(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReportLoadRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReportLoadRequest): ReportLoadRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReportLoadRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReportLoadRequest;
  static deserializeBinaryFromReader(message: ReportLoadRequest, reader: jspb.BinaryReader): ReportLoadRequest;
}

export namespace ReportLoadRequest {
  export type AsObject = {
    interval: number,
  }
}

export class LoadMetrics extends jspb.Message {
  clearMetricsList(): void;
  getMetricsList(): Array<metrics_pb.MetricFamily>;
  setMetricsList(value: Array<metrics_pb.MetricFamily>): void;
  addMetrics(value?: metrics_pb.MetricFamily, index?: number): metrics_pb.MetricFamily;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LoadMetrics.AsObject;
  static toObject(includeInstance: boolean, msg: LoadMetrics): LoadMetrics.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LoadMetrics, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LoadMetrics;
  static deserializeBinaryFromReader(message: LoadMetrics, reader: jspb.BinaryReader): LoadMetrics;
}

export namespace LoadMetrics {
  export type AsObject = {
    metricsList: Array<metrics_pb.MetricFamily.AsObject>,
  }
}

export class JobArg extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  hasString(): boolean;
  clearString(): void;
  getString(): string;
  setString(value: string): void;

  hasFloat(): boolean;
  clearFloat(): void;
  getFloat(): number;
  setFloat(value: number): void;

  getValueCase(): JobArg.ValueCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JobArg.AsObject;
  static toObject(includeInstance: boolean, msg: JobArg): JobArg.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JobArg, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JobArg;
  static deserializeBinaryFromReader(message: JobArg, reader: jspb.BinaryReader): JobArg;
}

export namespace JobArg {
  export type AsObject = {
    name: string,
    string: string,
    pb_float: number,
  }

  export enum ValueCase {
    VALUE_NOT_SET = 0,
    STRING = 2,
    FLOAT = 3,
  }
}

export class RunJobRequest extends jspb.Message {
  getProfile(): string;
  setProfile(value: string): void;

  getInterval(): number;
  setInterval(value: number): void;

  getMaxrps(): number;
  setMaxrps(value: number): void;

  clearArgsList(): void;
  getArgsList(): Array<JobArg>;
  setArgsList(value: Array<JobArg>): void;
  addArgs(value?: JobArg, index?: number): JobArg;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RunJobRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RunJobRequest): RunJobRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RunJobRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RunJobRequest;
  static deserializeBinaryFromReader(message: RunJobRequest, reader: jspb.BinaryReader): RunJobRequest;
}

export namespace RunJobRequest {
  export type AsObject = {
    profile: string,
    interval: number,
    maxrps: number,
    argsList: Array<JobArg.AsObject>,
  }
}

export class JobMetrics extends jspb.Message {
  clearMetricsList(): void;
  getMetricsList(): Array<metrics_pb.MetricFamily>;
  setMetricsList(value: Array<metrics_pb.MetricFamily>): void;
  addMetrics(value?: metrics_pb.MetricFamily, index?: number): metrics_pb.MetricFamily;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JobMetrics.AsObject;
  static toObject(includeInstance: boolean, msg: JobMetrics): JobMetrics.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JobMetrics, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JobMetrics;
  static deserializeBinaryFromReader(message: JobMetrics, reader: jspb.BinaryReader): JobMetrics;
}

export namespace JobMetrics {
  export type AsObject = {
    metricsList: Array<metrics_pb.MetricFamily.AsObject>,
  }
}

export class StartJobRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartJobRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StartJobRequest): StartJobRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartJobRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartJobRequest;
  static deserializeBinaryFromReader(message: StartJobRequest, reader: jspb.BinaryReader): StartJobRequest;
}

export namespace StartJobRequest {
  export type AsObject = {
  }
}

export class StartJobResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StartJobResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StartJobResponse): StartJobResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StartJobResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StartJobResponse;
  static deserializeBinaryFromReader(message: StartJobResponse, reader: jspb.BinaryReader): StartJobResponse;
}

export namespace StartJobResponse {
  export type AsObject = {
  }
}

export class ListProfilesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListProfilesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListProfilesRequest): ListProfilesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListProfilesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListProfilesRequest;
  static deserializeBinaryFromReader(message: ListProfilesRequest, reader: jspb.BinaryReader): ListProfilesRequest;
}

export namespace ListProfilesRequest {
  export type AsObject = {
  }
}

export class ListProfilesResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListProfilesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListProfilesResponse): ListProfilesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListProfilesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListProfilesResponse;
  static deserializeBinaryFromReader(message: ListProfilesResponse, reader: jspb.BinaryReader): ListProfilesResponse;
}

export namespace ListProfilesResponse {
  export type AsObject = {
  }
}

export class WatchActiveJobsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WatchActiveJobsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: WatchActiveJobsRequest): WatchActiveJobsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WatchActiveJobsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WatchActiveJobsRequest;
  static deserializeBinaryFromReader(message: WatchActiveJobsRequest, reader: jspb.BinaryReader): WatchActiveJobsRequest;
}

export namespace WatchActiveJobsRequest {
  export type AsObject = {
  }
}

export class WatchActiveJobsResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WatchActiveJobsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: WatchActiveJobsResponse): WatchActiveJobsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WatchActiveJobsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WatchActiveJobsResponse;
  static deserializeBinaryFromReader(message: WatchActiveJobsResponse, reader: jspb.BinaryReader): WatchActiveJobsResponse;
}

export namespace WatchActiveJobsResponse {
  export type AsObject = {
  }
}
