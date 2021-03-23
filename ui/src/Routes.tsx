import React from 'react';
import ProfilesTable from './ProfilesTable';
import Fab from '@material-ui/core/Fab';
import AddIcon from '@material-ui/icons/Add';

const Jobs: React.FC = () => {
    return (
          <div>
          <h1>Jobs</h1>
            <Fab color="primary" aria-label="add">
              <AddIcon />
            </Fab>
          </div>
        );
};

const Profiles: React.FC = () => {
    return (
          <div>
            <h1>Profiles</h1>
            <ProfilesTable />
          </div>
        );
};

const Archive: React.FC = () => {
    return (
          <h1>Archive</h1>
        );
};

const Routes = [
    {
      path: '/',
      sidebarName: 'Jobs',
      component: Jobs
    },
    {
      path: '/profiles',
      sidebarName: 'Profiles',
      component: Profiles
    },
    {
      path: '/archive',
      sidebarName: 'Archive',
      component: Archive
    },
];

export default Routes;
