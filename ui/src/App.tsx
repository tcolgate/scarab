import React from 'react';
import Container from '@material-ui/core/Container';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import Link from '@material-ui/core/Link';
import ProTip from './ProTip';
import {grpc} from "@improbable-eng/grpc-web";
import { ManagerUI } from './ui_pb_service';
import { ListProfilesRequest, ListProfilesResponse, StartJobRequest, StartJobResponse } from './ui_pb';

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


function Copyright() {
  return (
    <Typography variant="body2" color="textSecondary" align="center">
      {'Copyright © '}
      <Link color="inherit" href="https://material-ui.com/">
        Your Website
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

export default function App() {
  listProfiles()
  return (
    <Container maxWidth="sm">
      <Box my={4}>
        <Typography variant="h4" component="h1" gutterBottom>
          Create React App v4-beta example with TypeScript
        </Typography>
        <ProTip />
        <Copyright />
      </Box>
    </Container>
  );
}
