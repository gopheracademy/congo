import React from "react";
import ReactDOM from "react-dom";
import Master from "./Master"
import Profile from "./Profile"
import Proposals from "./Proposals"
import { Router, IndexRoute, Route, browserHistory } from "react-router";

// Rewrite the URL to point to / post oauth redirect
if (window.location.href.indexOf("callback") > -1) {
	window.history.pushState({},"", "/");
 }

var root = document.getElementById("app-container")
var auth = root.getAttribute("data-auth")
var AuthedMaster = React.createClass({
	render: function() {
                var childrenWithAuth = React.Children.map(this.props.children, function(child) {
                        return React.cloneElement(child, { auth: auth });
                });
		return (
			<Master auth={auth}>
				{ childrenWithAuth }
			</Master>
		);
	}
});
ReactDOM.render(
	<Router history={browserHistory}>
		<Route path="/" component={AuthedMaster}>
			<IndexRoute component={Profile} />
			<Route path="proposals" component={Proposals} />
			<Route path="profile" component={Profile} />
		</Route>
	</Router>,
	root
);
