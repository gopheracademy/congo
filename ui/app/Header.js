import React from 'react';
import { Navbar, Nav, NavItem, NavDropdown, MenuItem } from 'react-bootstrap';
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
		if (this.props.auth.Admin == true) {
			admin = true
		} else {
			admin = false
		}

		return (
   <Navbar>
    <Navbar.Header>
      <Navbar.Brand>
        <a href="#">Congo</a>
      </Navbar.Brand>
    </Navbar.Header>
    <Nav pullRight>
      <NavItem eventKey={1} href="#">Link</NavItem>
      <NavItem eventKey={2} href="#">Link</NavItem>
      <NavDropdown eventKey={3} title="Menu" id="basic-nav-dropdown">
        <MenuItem eventKey={3.1}>{loginButton}</MenuItem>
        <MenuItem eventKey={3.2}>Another action</MenuItem>
        <MenuItem eventKey={3.3}>Something else here</MenuItem>
        <MenuItem divider />
        <MenuItem eventKey={3.3}>Separated link</MenuItem>
      </NavDropdown>
    </Nav>
  </Navbar>
		);

        }
}
