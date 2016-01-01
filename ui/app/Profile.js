import React from 'react';
import DocumentTitle from 'react-document-title';
import Client from './Client';
import { Input, ButtonInput } from 'react-bootstrap';

var LinkedStateMixin = require('react-addons-linked-state-mixin');

var Profile = React.createClass({
	mixins: [LinkedStateMixin],
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
	handleSubmit: function(e) {
    	e.preventDefault();
                var current = this;
		var u = {firstname: this.state.firstname, lastname: this.state.lastname, email: this.state.email, city: this.state.city, state: this.state.city, country: this.state.country, bio: this.state.bio}
                Client(this.props.auth).updateUser("/api/users/" + this.state.id , u)
                        .then(function (resp) {
				if (resp.status > 400) {
					alert(resp.data);
				} else {
					alert("looks fine thanks");
				}
                        })
                        .catch(function(error) {
                                console.log(error);
                        });
  },
        render: function() {
                return (
                        <DocumentTitle title={`Profile`}>
                         <div className="container">
                    <h3>Your Profile</h3>
     
  <form onSubmit={this.handleSubmit}>
    <Input type="text" label="First Name" placeholder="First Name" valueLink={this.linkState('firstname')} />
    <Input type="text" label="Last Name" placeholder="Last Name" valueLink={this.linkState('lastname')} />
    <Input type="email" label="Email Address" placeholder="Email Address" valueLink={this.linkState('email')} />
    <Input type="text" label="City" placeholder="City" valueLink={this.linkState('city')} />
    <Input type="text" label="State" placeholder="State" valueLink={this.linkState('state')} />
    <Input type="text" label="Country" placeholder="Country" valueLink={this.linkState('country')} />
    <Input type="textarea" label="Bio" placeholder="Biography" valueLink={this.linkState('bio')} />
    <ButtonInput type="submit" value="Update Profile" />
  </form>
                       </div>
                        </DocumentTitle>
                );
        }
})
module.exports = Profile;
