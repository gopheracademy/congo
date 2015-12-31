import React from 'react';
import ReactDOM from 'react-dom';
import { Router, IndexRoute, Route } from 'react-router';

ReactDOM.render(
  <Router>
    <Route path='/' component={Master}>
      <LoginRoute path='/login' component={Login} />
      <ProfileRoute path='/profile' component={Profile} />
    </Route>
  </Router>,
  document.getElementById('app-container')
);
