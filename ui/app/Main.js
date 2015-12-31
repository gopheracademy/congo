import React from "react";
import ReactDOM from "react-dom";
import Master from "./Master"
import Login from "./Login"
import Profile from "./Profile"
import { Router, Route, hashHistory } from "react-router";

ReactDOM.render(
  <Router history={hashHistory}>
    <Route path="/" component={Master}>
      <Route path="/login" component={Login} />
      <Route path="/profile" component={Profile} />
    </Route>
  </Router>,
  document.getElementById("app-container")
);
