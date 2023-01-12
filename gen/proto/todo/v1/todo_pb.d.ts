import * as jspb from 'google-protobuf'

import * as validate_validate_pb from '../../../validate/validate_pb';


export class ListTodosRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTodosRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListTodosRequest): ListTodosRequest.AsObject;
  static serializeBinaryToWriter(message: ListTodosRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTodosRequest;
  static deserializeBinaryFromReader(message: ListTodosRequest, reader: jspb.BinaryReader): ListTodosRequest;
}

export namespace ListTodosRequest {
  export type AsObject = {
  }
}

export class Todo extends jspb.Message {
  getId(): number;
  setId(value: number): Todo;

  getContent(): string;
  setContent(value: string): Todo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Todo.AsObject;
  static toObject(includeInstance: boolean, msg: Todo): Todo.AsObject;
  static serializeBinaryToWriter(message: Todo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Todo;
  static deserializeBinaryFromReader(message: Todo, reader: jspb.BinaryReader): Todo;
}

export namespace Todo {
  export type AsObject = {
    id: number,
    content: string,
  }
}

export class ListTodosResponse extends jspb.Message {
  getTodosList(): Array<Todo>;
  setTodosList(value: Array<Todo>): ListTodosResponse;
  clearTodosList(): ListTodosResponse;
  addTodos(value?: Todo, index?: number): Todo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTodosResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListTodosResponse): ListTodosResponse.AsObject;
  static serializeBinaryToWriter(message: ListTodosResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTodosResponse;
  static deserializeBinaryFromReader(message: ListTodosResponse, reader: jspb.BinaryReader): ListTodosResponse;
}

export namespace ListTodosResponse {
  export type AsObject = {
    todosList: Array<Todo.AsObject>,
  }
}

export class CreateTodoRequest extends jspb.Message {
  getContent(): string;
  setContent(value: string): CreateTodoRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTodoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTodoRequest): CreateTodoRequest.AsObject;
  static serializeBinaryToWriter(message: CreateTodoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTodoRequest;
  static deserializeBinaryFromReader(message: CreateTodoRequest, reader: jspb.BinaryReader): CreateTodoRequest;
}

export namespace CreateTodoRequest {
  export type AsObject = {
    content: string,
  }
}

