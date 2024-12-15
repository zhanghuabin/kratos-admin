// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: user/service/v1/position.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type Timestamp } from "../../../google/protobuf/timestamp.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 职位 */
export interface Position {
  /** 职位ID */
  id?:
    | number
    | null
    | undefined;
  /** 职位名称 */
  name?:
    | string
    | null
    | undefined;
  /** 排序号 */
  orderNo?:
    | number
    | null
    | undefined;
  /** 职位值 */
  code?:
    | string
    | null
    | undefined;
  /** 状态 */
  status?:
    | string
    | null
    | undefined;
  /** 备注 */
  remark?:
    | string
    | null
    | undefined;
  /** 父节点ID */
  parentId?:
    | number
    | null
    | undefined;
  /** 子节点树 */
  children: Position[];
  /** 创建时间 */
  createTime?:
    | Timestamp
    | null
    | undefined;
  /** 更新时间 */
  updateTime?:
    | Timestamp
    | null
    | undefined;
  /** 删除时间 */
  deleteTime?: Timestamp | null | undefined;
}

/** 获取职位列表 - 答复 */
export interface ListPositionResponse {
  items: Position[];
  total: number;
}

/** 获取职位数据 - 请求 */
export interface GetPositionRequest {
  id: number;
}

/** 创建职位 - 请求 */
export interface CreatePositionRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  position: Position | null;
}

/** 更新职位 - 请求 */
export interface UpdatePositionRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  position:
    | Position
    | null;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除职位 - 请求 */
export interface DeletePositionRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  id: number;
}

/** 职位服务 */
export interface PositionService {
  /** 查询职位列表 */
  ListPosition(request: PagingRequest): Promise<ListPositionResponse>;
  /** 查询职位详情 */
  GetPosition(request: GetPositionRequest): Promise<Position>;
  /** 创建职位 */
  CreatePosition(request: CreatePositionRequest): Promise<Empty>;
  /** 更新职位 */
  UpdatePosition(request: UpdatePositionRequest): Promise<Empty>;
  /** 删除职位 */
  DeletePosition(request: DeletePositionRequest): Promise<Empty>;
}