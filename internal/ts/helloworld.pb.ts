/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "./fetch.pb"
import * as GoogleProtobufEmpty from "./google/protobuf/empty.pb"
export type PongResponse = {
  message?: string
}

export class Helloworld {
  static Ping(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<PongResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, PongResponse>(`/v1/ping?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}