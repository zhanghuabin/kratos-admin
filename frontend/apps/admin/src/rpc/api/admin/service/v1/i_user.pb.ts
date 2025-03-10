// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_user.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";
import {
  type CreateUserRequest,
  type DeleteUserRequest,
  type GetUserRequest,
  type ListUserResponse,
  type UpdateUserRequest,
  type User,
} from "../../../user/service/v1/user.pb";

/** 用户管理服务 */
export interface UserService {
  /** 获取用户列表 */
  ListUser(request: PagingRequest): Promise<ListUserResponse>;
  /** 获取用户数据 */
  GetUser(request: GetUserRequest): Promise<User>;
  /** 创建用户 */
  CreateUser(request: CreateUserRequest): Promise<Empty>;
  /** 更新用户 */
  UpdateUser(request: UpdateUserRequest): Promise<Empty>;
  /** 删除用户 */
  DeleteUser(request: DeleteUserRequest): Promise<Empty>;
}
