import React from 'react';
import DocumentTitle from 'react-document-title';

export default class Profile extends React.Component {
        render() {
                return (
                        <DocumentTitle title={`Profile`}>
                        <div>
                        <h3>Your Profile</h3>
                        <p>
                                { this.props.auth }
                        </p>
                        </div>
                        </DocumentTitle>
                );
        }
}
