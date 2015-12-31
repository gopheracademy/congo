import React from "react";
import ReactDOM from "react-dom";
import Master from "./Master"
import Profile from "./Profile"
import Proposals from "./Proposals"
import { Router, IndexRoute, Route, hashHistory } from "react-router";

var root = document.getElementById("app-container")
var jwt = root.getAttribute("data-jwt")
var AuthedMaster = React.createClass({
	render() {
		return <Master jwt={jwt} />;
	}
});
ReactDOM.render(
	<Router history={hashHistory}>
		<Route path="/" component={AuthedMaster}>
			<IndexRoute component={Profile} />
			<Route path="/proposals" component={Proposals} />
			<Route path="/profile" component={Profile} />
		</Route>
	</Router>,
	root
);
