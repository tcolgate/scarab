import {grpc} from "@improbable-eng/grpc-web";
import { ManagerUI } from './ui_pb_service';
import { ListProfilesRequest, ListProfilesResponse, StartJobRequest, StartJobResponse } from './ui_pb';

declare const USE_TLS: boolean;
const host = USE_TLS ? "https://localhost:9091" : "http://localhost:9090";

function listProfiles() {
    const lpRequest = new ListProfilesRequest();
    grpc.unary(ManagerUI.ListProfiles, {
          request: lpRequest,
          host: host,
          onEnd: res => {
                  const { status, statusMessage, headers, message, trailers } = res;
                  console.log("listProfiles.onEnd.status", status, statusMessage);
                  console.log("listProfiles.onEnd.headers", headers);
                  if (status === grpc.Code.OK && message) {
                            console.log("getBook.onEnd.message", message.toObject());
                          }
                  console.log("getBook.onEnd.trailers", trailers);
                }
        });
}

