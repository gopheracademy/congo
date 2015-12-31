import React from 'react';
import DocumentTitle from 'react-document-title';

export default class Login extends React.Component {
  render() {
    return (
      <div>
      <DocumentTitle title={`Sign In`} />
      <h3>Sign In With</h3>
      <a className="btn btn-twitter" href="/auth/twitter"><i className="icon-twitter"></i> | Twitter </a>
      <a className="btn btn-github" href="/auth/github"><i className="icon-github"></i> | GitHub </a>
      </div>
    );
  }
}
