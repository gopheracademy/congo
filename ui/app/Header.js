import React from 'react';
import { Link } from 'react-router';
import jwt from 'jwt-decode';

export default class Header extends React.Component {
        render() {
                var loginButton;
		var admin;
                if (this.props.auth) {
                        loginButton = "Logout"
                } else {
                        loginButton = "Login"
                }
		if (this.props.auth.Admin) {
			admin = true
		} else {
			admin = false
		}
                return (
                        <nav className="navbar navbar-default navbar-static-top">
                                <div className="container">
                                        <div id="navbar-collapse" className="collapse navbar-collapse">
                                                <ul className="nav navbar-nav">
                                                        <li><Link to="/">Home</Link></li>

                                                </ul>
                                                <ul className="nav navbar-nav navbar-right">
                                                        <li><Link to="#">{loginButton}</Link></li>
							{admin ? <li>Admin</li> : ""}
                                                </ul>
                                        </div>
                                </div>
                        </nav>
                );
        }
}
