// @generated by protoc-gen-connect-web v0.6.0 with parameter "target=ts"
// @generated from file proto/todo/v1/todo.proto (package todo.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateTodoRequest, ListTodosRequest, ListTodosResponse, Todo } from "./todo_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service todo.v1.TodoService
 */
export const TodoService = {
  typeName: "todo.v1.TodoService",
  methods: {
    /**
     * @generated from rpc todo.v1.TodoService.ListTodos
     */
    listTodos: {
      name: "ListTodos",
      I: ListTodosRequest,
      O: ListTodosResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc todo.v1.TodoService.CreateTodo
     */
    createTodo: {
      name: "CreateTodo",
      I: CreateTodoRequest,
      O: Todo,
      kind: MethodKind.Unary,
    },
  }
} as const;

