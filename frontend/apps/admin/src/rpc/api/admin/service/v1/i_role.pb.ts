// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: admin/service/v1/i_role.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";
import {
  type CreateRoleRequest,
  type DeleteRoleRequest,
  type GetRoleRequest,
  type ListRoleResponse,
  type Role,
  type UpdateRoleRequest,
} from "../../../user/service/v1/role.pb";

/** 角色管理服务 */
export interface RoleService {
  /** 查询角色列表 */
  ListRole(request: PagingRequest): Promise<ListRoleResponse>;
  /** 查询角色详情 */
  GetRole(request: GetRoleRequest): Promise<Role>;
  /** 创建角色 */
  CreateRole(request: CreateRoleRequest): Promise<Empty>;
  /** 更新角色 */
  UpdateRole(request: UpdateRoleRequest): Promise<Empty>;
  /** 删除角色 */
  DeleteRole(request: DeleteRoleRequest): Promise<Empty>;
}