// package: scarab
// file: scarab-common.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class WorkerDetails extends jspb.Message {
  getAddr(): string;
  setAddr(value: string): void;

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WorkerDetails.AsObject;
  static toObject(includeInstance: boolean, msg: WorkerDetails): WorkerDetails.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WorkerDetails, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WorkerDetails;
  static deserializeBinaryFromReader(message: WorkerDetails, reader: jspb.BinaryReader): WorkerDetails;
}

export namespace WorkerDetails {
  export type AsObject = {
    addr: string,
    name: string,
  }
}

export class RegisteredProfile extends jspb.Message {
  hasSpec(): boolean;
  clearSpec(): void;
  getSpec(): ProfileSpec | undefined;
  setSpec(value?: ProfileSpec): void;

  hasFirstregistration(): boolean;
  clearFirstregistration(): void;
  getFirstregistration(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setFirstregistration(value?: google_protobuf_timestamp_pb.Timestamp): void;

  clearWorkersList(): void;
  getWorkersList(): Array<WorkerDetails>;
  setWorkersList(value: Array<WorkerDetails>): void;
  addWorkers(value?: WorkerDetails, index?: number): WorkerDetails;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegisteredProfile.AsObject;
  static toObject(includeInstance: boolean, msg: RegisteredProfile): RegisteredProfile.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RegisteredProfile, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegisteredProfile;
  static deserializeBinaryFromReader(message: RegisteredProfile, reader: jspb.BinaryReader): RegisteredProfile;
}

export namespace RegisteredProfile {
  export type AsObject = {
    spec?: ProfileSpec.AsObject,
    firstregistration?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    workersList: Array<WorkerDetails.AsObject>,
  }
}

export class ProfileArg extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  hasDefault(): boolean;
  clearDefault(): void;
  getDefault(): JobArgValue | undefined;
  setDefault(value?: JobArgValue): void;

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
    pb_default?: JobArgValue.AsObject,
  }
}

export class ProfileSpec extends jspb.Message {
  getProfile(): string;
  setProfile(value: string): void;

  getVersion(): string;
  setVersion(value: string): void;

  getDescription(): string;
  setDescription(value: string): void;

  clearArgsList(): void;
  getArgsList(): Array<ProfileArg>;
  setArgsList(value: Array<ProfileArg>): void;
  addArgs(value?: ProfileArg, index?: number): ProfileArg;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ProfileSpec.AsObject;
  static toObject(includeInstance: boolean, msg: ProfileSpec): ProfileSpec.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ProfileSpec, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ProfileSpec;
  static deserializeBinaryFromReader(message: ProfileSpec, reader: jspb.BinaryReader): ProfileSpec;
}

export namespace ProfileSpec {
  export type AsObject = {
    profile: string,
    version: string,
    description: string,
    argsList: Array<ProfileArg.AsObject>,
  }
}

export class JobArgOption extends jspb.Message {
  clearOptionList(): void;
  getOptionList(): Array<string>;
  setOptionList(value: Array<string>): void;
  addOption(value: string, index?: number): string;

  getMultiple(): boolean;
  setMultiple(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JobArgOption.AsObject;
  static toObject(includeInstance: boolean, msg: JobArgOption): JobArgOption.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JobArgOption, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JobArgOption;
  static deserializeBinaryFromReader(message: JobArgOption, reader: jspb.BinaryReader): JobArgOption;
}

export namespace JobArgOption {
  export type AsObject = {
    optionList: Array<string>,
    multiple: boolean,
  }
}

export class JobArgValue extends jspb.Message {
  hasString(): boolean;
  clearString(): void;
  getString(): string;
  setString(value: string): void;

  hasFloat(): boolean;
  clearFloat(): void;
  getFloat(): number;
  setFloat(value: number): void;

  hasBool(): boolean;
  clearBool(): void;
  getBool(): boolean;
  setBool(value: boolean): void;

  hasOption(): boolean;
  clearOption(): void;
  getOption(): JobArgOption | undefined;
  setOption(value?: JobArgOption): void;

  getValueCase(): JobArgValue.ValueCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): JobArgValue.AsObject;
  static toObject(includeInstance: boolean, msg: JobArgValue): JobArgValue.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: JobArgValue, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): JobArgValue;
  static deserializeBinaryFromReader(message: JobArgValue, reader: jspb.BinaryReader): JobArgValue;
}

export namespace JobArgValue {
  export type AsObject = {
    string: string,
    pb_float: number,
    bool: boolean,
    option?: JobArgOption.AsObject,
  }

  export enum ValueCase {
    VALUE_NOT_SET = 0,
    STRING = 1,
    FLOAT = 2,
    BOOL = 3,
    OPTION = 4,
  }
}

export class JobArg extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  hasValue(): boolean;
  clearValue(): void;
  getValue(): JobArgValue | undefined;
  setValue(value?: JobArgValue): void;

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
    value?: JobArgValue.AsObject,
  }
}

export class StartJobRequest extends jspb.Message {
  getProfile(): string;
  setProfile(value: string): void;

  getVersion(): string;
  setVersion(value: string): void;

  clearArgsList(): void;
  getArgsList(): Array<JobArg>;
  setArgsList(value: Array<JobArg>): void;
  addArgs(value?: JobArg, index?: number): JobArg;

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
    profile: string,
    version: string,
    argsList: Array<JobArg.AsObject>,
  }
}

export class StartJobResponse extends jspb.Message {
  getId(): string;
  setId(value: string): void;

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
    id: string,
  }
}

