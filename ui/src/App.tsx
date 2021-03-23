import React , { MouseEvent, KeyboardEvent, useState } from 'react';
import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import Link from '@material-ui/core/Link';

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import {
    AppBar,
    Drawer,
    Toolbar,
    IconButton,
   } from '@material-ui/core';
import MenuIcon from '@material-ui/icons/Menu';


import ProTip from './ProTip';
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


const useStyles = makeStyles((theme: Theme) =>
    createStyles({
          root: {
                  flexGrow: 1,
                  },
          menuButton: {
                  marginRight: theme.spacing(2),
                  },
          title: {
                  flexGrow: 1,
                  },
        }),
);

const App: React.FC = () => {
    const classes = useStyles();

    const [isOpen, setIsOpen] = useState(false);
    const toggleDrawer = (open: boolean) => (
          event: KeyboardEvent | MouseEvent,
        ) => {
              if (
                  event.type === 'keydown' &&
                  ((event as React.KeyboardEvent).key === 'Tab' ||
                    (event as React.KeyboardEvent).key === 'Shift')
                 ) {
                 return;
              };
              setIsOpen(open);
            };

    return (
          <div>
            <div className={classes.root}>
              <AppBar position="static">
                <Toolbar>
                  <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu" onClick={toggleDrawer(true)}>
                    <MenuIcon />
                  </IconButton>
                  <Typography variant="h6" className={classes.title}>
                    Scarab
                  </Typography>
                </Toolbar>
              </AppBar>
            </div>
            <Drawer open={isOpen} onClose={toggleDrawer(false)}>
              Hello Drawer!
            </Drawer>
          </div>
        );
}

export default App;
