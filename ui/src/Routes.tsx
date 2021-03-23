import React from 'react';

const Jobs: React.FC = () => {
    return (
          <h1>Jobs</h1>
        );
};

const Profiles: React.FC = () => {
    return (
          <h1>Profiles</h1>
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
