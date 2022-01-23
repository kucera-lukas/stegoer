/* tslint:disable */
/* eslint-disable */
// This file was automatically generated and should not be edited.

//==============================================================
// START Enums and Input Objects
//==============================================================

export enum Channel {
  BLUE = "BLUE",
  GREEN = "GREEN",
  GREEN_BLUE = "GREEN_BLUE",
  RED = "RED",
  RED_BLUE = "RED_BLUE",
  RED_GREEN = "RED_GREEN",
  RED_GREEN_BLUE = "RED_GREEN_BLUE",
}

export enum ErrorCode {
  AUTHORIZATION_ERROR = "AUTHORIZATION_ERROR",
  BAD_REQUEST_ERROR = "BAD_REQUEST_ERROR",
  DB_ERROR = "DB_ERROR",
  GRAPHQL_ERROR = "GRAPHQL_ERROR",
  INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR",
  NOT_FOUND_ERROR = "NOT_FOUND_ERROR",
  VALIDATION_ERROR = "VALIDATION_ERROR",
}

export enum ImageOrderField {
  CREATED_AT = "CREATED_AT",
  UPDATED_AT = "UPDATED_AT",
}

export enum OrderDirection {
  ASC = "ASC",
  DESC = "DESC",
}

export interface ImageOrder {
  direction: OrderDirection;
  field?: ImageOrderField | null;
}

/**
 * ImageWhereInput is used for filtering Image objects.
 * Input was generated by ent.
 */
export interface ImageWhereInput {
  not?: ImageWhereInput | null;
  and?: ImageWhereInput[] | null;
  or?: ImageWhereInput[] | null;
  createdAt?: any | null;
  createdAtNEQ?: any | null;
  createdAtIn?: any[] | null;
  createdAtNotIn?: any[] | null;
  createdAtGT?: any | null;
  createdAtGTE?: any | null;
  createdAtLT?: any | null;
  createdAtLTE?: any | null;
  updatedAt?: any | null;
  updatedAtNEQ?: any | null;
  updatedAtIn?: any[] | null;
  updatedAtNotIn?: any[] | null;
  updatedAtGT?: any | null;
  updatedAtGTE?: any | null;
  updatedAtLT?: any | null;
  updatedAtLTE?: any | null;
  channel?: Channel | null;
  channelNEQ?: Channel | null;
  channelIn?: Channel[] | null;
  channelNotIn?: Channel[] | null;
  id?: string | null;
  idNEQ?: string | null;
  idIn?: string[] | null;
  idNotIn?: string[] | null;
  idGT?: string | null;
  idGTE?: string | null;
  idLT?: string | null;
  idLTE?: string | null;
  hasUser?: boolean | null;
  hasUserWith?: UserWhereInput[] | null;
}

/**
 * UserWhereInput is used for filtering User objects.
 * Input was generated by ent.
 */
export interface UserWhereInput {
  not?: UserWhereInput | null;
  and?: UserWhereInput[] | null;
  or?: UserWhereInput[] | null;
  createdAt?: any | null;
  createdAtNEQ?: any | null;
  createdAtIn?: any[] | null;
  createdAtNotIn?: any[] | null;
  createdAtGT?: any | null;
  createdAtGTE?: any | null;
  createdAtLT?: any | null;
  createdAtLTE?: any | null;
  updatedAt?: any | null;
  updatedAtNEQ?: any | null;
  updatedAtIn?: any[] | null;
  updatedAtNotIn?: any[] | null;
  updatedAtGT?: any | null;
  updatedAtGTE?: any | null;
  updatedAtLT?: any | null;
  updatedAtLTE?: any | null;
  name?: string | null;
  nameNEQ?: string | null;
  nameIn?: string[] | null;
  nameNotIn?: string[] | null;
  nameGT?: string | null;
  nameGTE?: string | null;
  nameLT?: string | null;
  nameLTE?: string | null;
  nameContains?: string | null;
  nameHasPrefix?: string | null;
  nameHasSuffix?: string | null;
  nameEqualFold?: string | null;
  nameContainsFold?: string | null;
  password?: string | null;
  passwordNEQ?: string | null;
  passwordIn?: string[] | null;
  passwordNotIn?: string[] | null;
  passwordGT?: string | null;
  passwordGTE?: string | null;
  passwordLT?: string | null;
  passwordLTE?: string | null;
  passwordContains?: string | null;
  passwordHasPrefix?: string | null;
  passwordHasSuffix?: string | null;
  passwordEqualFold?: string | null;
  passwordContainsFold?: string | null;
  id?: string | null;
  idNEQ?: string | null;
  idIn?: string[] | null;
  idNotIn?: string[] | null;
  idGT?: string | null;
  idGTE?: string | null;
  idLT?: string | null;
  idLTE?: string | null;
  hasImages?: boolean | null;
  hasImagesWith?: ImageWhereInput[] | null;
}

//==============================================================
// END Enums and Input Objects
//==============================================================
