import React , { MouseEvent, KeyboardEvent, useState } from 'react';
import { Switch, Route } from 'react-router-dom';
import Routes from './Routes';
import NavigationBar from './NavigationBar';

import {grpc} from "@improbable-eng/grpc-web";
import { ManagerUI } from './scarab-ui_pb_service';
import { ListProfilesRequest, ListProfilesResponse } from './scarab-ui_pb';
import { StartJobRequest, StartJobResponse } from './scarab-common_pb';

const host = "http://localhost:8081";

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


const App: React.FC = () => {
    return (
          <div>
            <NavigationBar />
            <Switch>
              {Routes.map((route: any) => (
                          <Route exact path={route.path} key={route.path}>
                            <route.component />
                          </Route>
                        ))}
            </Switch>
          </div>
        );
}

export default App;
