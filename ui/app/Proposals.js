import React from 'react';
import DocumentTitle from 'react-document-title';

export default class Proposals extends React.Component {
  render() {
    return (
      <DocumentTitle title={`Proposals`}>
        <h3>Your Proposals</h3>
        <p>
        {this.props.auth}
        </p>
      </DocumentTitle>
    );
  }
}
