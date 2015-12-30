	var Profile= React.createClass({
  mixins: [React.addons.LinkedStateMixin],
  getInitialState: function() {
    return {
      users: []
    };
  },
  componentDidMount: function() {
	  console.log(this.props.userid);
	var u = "/api/users/21";
	  $.getJSON(u, function(data) {
		  this.setState({users: [data]});
		  console.log(data)
    }.bind(this));
  },
  handleSubmit: function(e) {
	e.preventDefault();
    $.ajax({
      url: "/api/users/21",
      data: JSON.stringify({"id": 21, "firstname": "Brian", "lastname": "Ketelsen", "email": "bketelsen@gmail.com"}),
      method: "PATCH",
      error: function(request, textStatus, response) {
	  var resp = JSON.parse(request.responseText);
         alert(resp[0].title +": " +  resp[0].msg);
        },
    });
  },
  render: function() {
    return (
      <div className="container users">
	      <div className="row">
          <form onSubmit={this.handleSubmit}>
                             {this.state.users.map(function(u) {
				     return (
					     <User key={u.id} user={u} />
                               )
                             })}
            <button className="btn btn-default">Update</button>
          </form>
        </div>
      </div>
    );
  },
	});

var User= React.createClass({
	render: function() {
		return (
	<div className="form-group">
              <label htmlFor="accountName">My Info</label>
		<input type="text" className="form-control" id="firstName" placeholder="First Name" value={this.props.user.firstname} /> 
		<input type="text" className="form-control" id="lastName" placeholder="Last Name"  value={this.props.user.lastname} />
		<input type="text" className="form-control" id="email" placeholder="email"  value={this.props.user.email} />
	</div>
	)}
});


