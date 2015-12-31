import React from 'react';
import { Link } from 'react-router';
import DocumentTitle from 'react-document-title';

import Header from './Header';

export default class Master extends React.Component {
        render() {
                var childrenWithAuth = React.Children.map(this.props.children, function(child) {
                        return React.cloneElement(child, { auth: this.props.auth });
                });
                return (
                        <DocumentTitle title='GopherCon 2016'>
                        <div className='Master'>
                                <Header />
                                { childrenWithAuth }
                        </div>
                        </DocumentTitle>
                );
        }
}
