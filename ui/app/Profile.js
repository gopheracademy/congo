import React from 'react';
import DocumentTitle from 'react-document-title';
import Client from './Client'

var Profile = React.createClass({
        getInitialState: function() {
                return {
                        bio: "",
                        city: "",
                        country: "",
                        email: "",
                        firstname: "",
                        href: "",
                        id: 0,
                        lastname: "",
                        state: ""
                };
        },
        componentDidMount: function() {
                var current = this;
                Client(this.props.auth).showUser("/api/users/" + this.props.auth.UserID)
                        .then(function (resp) {
                                current.setState(resp.data);
                        })
                        .catch(function(error) {
                                console.log(error);
                        });
        },
        render: function() {
                return (
                        <DocumentTitle title={`Profile`}>
                        <div>
                        <h3>Your Profile</h3>
                        <p>
                                bio: { this.state.bio }
                        </p>
                        <p>
                                city: { this.state.city }
                        </p>
                        <p>
                                country: { this.state.country }
                        </p>
                        <p>
                                email: { this.state.email }
                        </p>
                        <p>
                                firstname: { this.state.firstname }
                        </p>
                        <p>
                                href: { this.state.href }
                        </p>
                        <p>
                                id: { this.state.id }
                        </p>
                        <p>
                                lastname: { this.state.lastname }
                        </p>
                        <p>
                                state: { this.state.state }
                        </p>
                        </div>
                        </DocumentTitle>
                );
        }
})
module.exports = Profile;
