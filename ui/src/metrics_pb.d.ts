// package: io.prometheus.client
// file: metrics.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class LabelPair extends jspb.Message {
  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasValue(): boolean;
  clearValue(): void;
  getValue(): string | undefined;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LabelPair.AsObject;
  static toObject(includeInstance: boolean, msg: LabelPair): LabelPair.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: LabelPair, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LabelPair;
  static deserializeBinaryFromReader(message: LabelPair, reader: jspb.BinaryReader): LabelPair;
}

export namespace LabelPair {
  export type AsObject = {
    name?: string,
    value?: string,
  }
}

export class Gauge extends jspb.Message {
  hasValue(): boolean;
  clearValue(): void;
  getValue(): number | undefined;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Gauge.AsObject;
  static toObject(includeInstance: boolean, msg: Gauge): Gauge.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Gauge, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Gauge;
  static deserializeBinaryFromReader(message: Gauge, reader: jspb.BinaryReader): Gauge;
}

export namespace Gauge {
  export type AsObject = {
    value?: number,
  }
}

export class Counter extends jspb.Message {
  hasValue(): boolean;
  clearValue(): void;
  getValue(): number | undefined;
  setValue(value: number): void;

  hasExemplar(): boolean;
  clearExemplar(): void;
  getExemplar(): Exemplar | undefined;
  setExemplar(value?: Exemplar): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Counter.AsObject;
  static toObject(includeInstance: boolean, msg: Counter): Counter.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Counter, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Counter;
  static deserializeBinaryFromReader(message: Counter, reader: jspb.BinaryReader): Counter;
}

export namespace Counter {
  export type AsObject = {
    value?: number,
    exemplar?: Exemplar.AsObject,
  }
}

export class Quantile extends jspb.Message {
  hasQuantile(): boolean;
  clearQuantile(): void;
  getQuantile(): number | undefined;
  setQuantile(value: number): void;

  hasValue(): boolean;
  clearValue(): void;
  getValue(): number | undefined;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Quantile.AsObject;
  static toObject(includeInstance: boolean, msg: Quantile): Quantile.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Quantile, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Quantile;
  static deserializeBinaryFromReader(message: Quantile, reader: jspb.BinaryReader): Quantile;
}

export namespace Quantile {
  export type AsObject = {
    quantile?: number,
    value?: number,
  }
}

export class Summary extends jspb.Message {
  hasSampleCount(): boolean;
  clearSampleCount(): void;
  getSampleCount(): number | undefined;
  setSampleCount(value: number): void;

  hasSampleSum(): boolean;
  clearSampleSum(): void;
  getSampleSum(): number | undefined;
  setSampleSum(value: number): void;

  clearQuantileList(): void;
  getQuantileList(): Array<Quantile>;
  setQuantileList(value: Array<Quantile>): void;
  addQuantile(value?: Quantile, index?: number): Quantile;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Summary.AsObject;
  static toObject(includeInstance: boolean, msg: Summary): Summary.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Summary, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Summary;
  static deserializeBinaryFromReader(message: Summary, reader: jspb.BinaryReader): Summary;
}

export namespace Summary {
  export type AsObject = {
    sampleCount?: number,
    sampleSum?: number,
    quantileList: Array<Quantile.AsObject>,
  }
}

export class Untyped extends jspb.Message {
  hasValue(): boolean;
  clearValue(): void;
  getValue(): number | undefined;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Untyped.AsObject;
  static toObject(includeInstance: boolean, msg: Untyped): Untyped.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Untyped, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Untyped;
  static deserializeBinaryFromReader(message: Untyped, reader: jspb.BinaryReader): Untyped;
}

export namespace Untyped {
  export type AsObject = {
    value?: number,
  }
}

export class Histogram extends jspb.Message {
  hasSampleCount(): boolean;
  clearSampleCount(): void;
  getSampleCount(): number | undefined;
  setSampleCount(value: number): void;

  hasSampleSum(): boolean;
  clearSampleSum(): void;
  getSampleSum(): number | undefined;
  setSampleSum(value: number): void;

  clearBucketList(): void;
  getBucketList(): Array<Bucket>;
  setBucketList(value: Array<Bucket>): void;
  addBucket(value?: Bucket, index?: number): Bucket;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Histogram.AsObject;
  static toObject(includeInstance: boolean, msg: Histogram): Histogram.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Histogram, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Histogram;
  static deserializeBinaryFromReader(message: Histogram, reader: jspb.BinaryReader): Histogram;
}

export namespace Histogram {
  export type AsObject = {
    sampleCount?: number,
    sampleSum?: number,
    bucketList: Array<Bucket.AsObject>,
  }
}

export class Bucket extends jspb.Message {
  hasCumulativeCount(): boolean;
  clearCumulativeCount(): void;
  getCumulativeCount(): number | undefined;
  setCumulativeCount(value: number): void;

  hasUpperBound(): boolean;
  clearUpperBound(): void;
  getUpperBound(): number | undefined;
  setUpperBound(value: number): void;

  hasExemplar(): boolean;
  clearExemplar(): void;
  getExemplar(): Exemplar | undefined;
  setExemplar(value?: Exemplar): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Bucket.AsObject;
  static toObject(includeInstance: boolean, msg: Bucket): Bucket.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Bucket, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Bucket;
  static deserializeBinaryFromReader(message: Bucket, reader: jspb.BinaryReader): Bucket;
}

export namespace Bucket {
  export type AsObject = {
    cumulativeCount?: number,
    upperBound?: number,
    exemplar?: Exemplar.AsObject,
  }
}

export class Exemplar extends jspb.Message {
  clearLabelList(): void;
  getLabelList(): Array<LabelPair>;
  setLabelList(value: Array<LabelPair>): void;
  addLabel(value?: LabelPair, index?: number): LabelPair;

  hasValue(): boolean;
  clearValue(): void;
  getValue(): number | undefined;
  setValue(value: number): void;

  hasTimestamp(): boolean;
  clearTimestamp(): void;
  getTimestamp(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTimestamp(value?: google_protobuf_timestamp_pb.Timestamp): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Exemplar.AsObject;
  static toObject(includeInstance: boolean, msg: Exemplar): Exemplar.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Exemplar, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Exemplar;
  static deserializeBinaryFromReader(message: Exemplar, reader: jspb.BinaryReader): Exemplar;
}

export namespace Exemplar {
  export type AsObject = {
    labelList: Array<LabelPair.AsObject>,
    value?: number,
    timestamp?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Metric extends jspb.Message {
  clearLabelList(): void;
  getLabelList(): Array<LabelPair>;
  setLabelList(value: Array<LabelPair>): void;
  addLabel(value?: LabelPair, index?: number): LabelPair;

  hasGauge(): boolean;
  clearGauge(): void;
  getGauge(): Gauge | undefined;
  setGauge(value?: Gauge): void;

  hasCounter(): boolean;
  clearCounter(): void;
  getCounter(): Counter | undefined;
  setCounter(value?: Counter): void;

  hasSummary(): boolean;
  clearSummary(): void;
  getSummary(): Summary | undefined;
  setSummary(value?: Summary): void;

  hasUntyped(): boolean;
  clearUntyped(): void;
  getUntyped(): Untyped | undefined;
  setUntyped(value?: Untyped): void;

  hasHistogram(): boolean;
  clearHistogram(): void;
  getHistogram(): Histogram | undefined;
  setHistogram(value?: Histogram): void;

  hasTimestampMs(): boolean;
  clearTimestampMs(): void;
  getTimestampMs(): number | undefined;
  setTimestampMs(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Metric.AsObject;
  static toObject(includeInstance: boolean, msg: Metric): Metric.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Metric, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Metric;
  static deserializeBinaryFromReader(message: Metric, reader: jspb.BinaryReader): Metric;
}

export namespace Metric {
  export type AsObject = {
    labelList: Array<LabelPair.AsObject>,
    gauge?: Gauge.AsObject,
    counter?: Counter.AsObject,
    summary?: Summary.AsObject,
    untyped?: Untyped.AsObject,
    histogram?: Histogram.AsObject,
    timestampMs?: number,
  }
}

export class MetricFamily extends jspb.Message {
  hasName(): boolean;
  clearName(): void;
  getName(): string | undefined;
  setName(value: string): void;

  hasHelp(): boolean;
  clearHelp(): void;
  getHelp(): string | undefined;
  setHelp(value: string): void;

  hasType(): boolean;
  clearType(): void;
  getType(): MetricTypeMap[keyof MetricTypeMap] | undefined;
  setType(value: MetricTypeMap[keyof MetricTypeMap]): void;

  clearMetricList(): void;
  getMetricList(): Array<Metric>;
  setMetricList(value: Array<Metric>): void;
  addMetric(value?: Metric, index?: number): Metric;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MetricFamily.AsObject;
  static toObject(includeInstance: boolean, msg: MetricFamily): MetricFamily.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MetricFamily, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MetricFamily;
  static deserializeBinaryFromReader(message: MetricFamily, reader: jspb.BinaryReader): MetricFamily;
}

export namespace MetricFamily {
  export type AsObject = {
    name?: string,
    help?: string,
    type?: MetricTypeMap[keyof MetricTypeMap],
    metricList: Array<Metric.AsObject>,
  }
}

export interface MetricTypeMap {
  COUNTER: 0;
  GAUGE: 1;
  SUMMARY: 2;
  UNTYPED: 3;
  HISTOGRAM: 4;
}

export const MetricType: MetricTypeMap;

