import React from 'react';
import { Link } from 'react-router';

export default class Header extends React.Component {
  render() {
    return (
      <nav className="navbar navbar-default navbar-static-top">
        <div className="container">
          <div id="navbar-collapse" className="collapse navbar-collapse">
            <ul className="nav navbar-nav">
              <li><Link to="/">Home</Link></li>
            </ul>
            <ul className="nav navbar-nav navbar-right">
            </ul>
          </div>
        </div>
      </nav>
    );
  }
}
