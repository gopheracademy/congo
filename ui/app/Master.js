import React from 'react';
import { Link } from 'react-router';
import DocumentTitle from 'react-document-title';

import Header from './Header';

export default class Master extends React.Component {
        render() {
                return (
                        <DocumentTitle title='GopherCon 2016'>
			<div className="container">
                        	<div className='Master'>
                              		<Header auth={this.props.auth} />
                                	{ this.props.children }
                        	</div>
			</div>
                        </DocumentTitle>
                );
        }
}
