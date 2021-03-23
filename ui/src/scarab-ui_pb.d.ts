// package: scarab
// file: scarab-ui.proto

import * as jspb from "google-protobuf";
import * as scarab_common_pb from "./scarab-common_pb";

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
  clearRegisteredList(): void;
  getRegisteredList(): Array<scarab_common_pb.RegiteredProfile>;
  setRegisteredList(value: Array<scarab_common_pb.RegiteredProfile>): void;
  addRegistered(value?: scarab_common_pb.RegiteredProfile, index?: number): scarab_common_pb.RegiteredProfile;

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
    registeredList: Array<scarab_common_pb.RegiteredProfile.AsObject>,
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

