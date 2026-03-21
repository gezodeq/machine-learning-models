// types.ts

import { Primitive } from 'machine-learning-models/primitives';

export type ModelType = 'regression' | 'classification';
export type ModelConfiguration = {
  epochs?: number;
  batchSize?: number;
  optimizer?: string;
  lossFunction?: string;
  metrics?: (string | string[])[];
};

export type InputData = {
  features: Primitive[][];
  labels: Primitive[][];
};

export type ModelOutput = {
  predictions: Primitive[][];
  metrics: { [key: string]: number } | null;
};

export type Model = {
  type: ModelType;
  configuration: ModelConfiguration;
  train: (input: InputData) => Promise<ModelOutput>;
  predict: (input: InputData) => Promise<ModelOutput>;
};