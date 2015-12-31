import React from 'react';
import DocumentTitle from 'react-document-title';

export default class Login extends React.Component {
  render() {
    return (
      <DocumentTitle title={`Sign In`}>
        <h3>Sign In With</h3>
        <p>
          <a class="btn btn-twitter" href="/auth/twitter"><i class="icon-twitter"></i>
          | Twitter
          </a>
          <a class="btn btn-github" href="/auth/github"><i class="icon-github"></i>
          | GitHub
          </a>
        </p>
      </DocumentTitle>
    );
  }
}
