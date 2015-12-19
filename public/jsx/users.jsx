var Root = React.createClass({
  mixins: [React.addons.LinkedStateMixin],
  getInitialState: function() {
    return {
      users: [],
	  firstName: "",
	  lastName: "",
	  email: "",
    };
  },
  componentDidMount: function() {
	var cookies = document.cookie;
	console.log(cookies)
      var path = this.getPath();
    var u = "/api/users"
	  $.getJSON(u, function(data) {
      this.setState({users: data});
    }.bind(this));
  },
 getUrl: function() {
    return '' + window.location;
  },
  getPath: function() {
    return window.location.pathname + window.location.search;
  },
  handleSubmit: function(e) {
	e.preventDefault();
      var path = this.getPath();
    var u = "/api/users"
    $.ajax({
	url: u,
      data: JSON.stringify({"firstname": this.state.firstName, "lastname": this.state.lastName, "email": this.state.email}),
      method: "POST",
      success: function(data, textStatus, request) {
		  var newNames = this.state.users;
		  var href = request.getResponseHeader('Location');
		   newNames.push({"href": href, "id": href.substr(href.length -1),"email": this.state.email,  "firstname": this.state.firstName, "lastname": this.state.lastName});
        this.setState({
			firstName: "",
			lastName: "",
			email: "",
          users: newNames
		});
	  }.bind(this),
	  error: function(request, textStatus, response) {
		  var resp = JSON.parse(request.responseText);
         alert(resp[0].title +": " +  resp[0].msg);
        },
    });
  },
  render: function() {

    var listElts = [];
    for(var i = 0; i < this.state.users.length; i++) {
      var u = this.state.users[i];
      listElts.push(<User key={i} user={u} />);
    }
    return (
      <div className="container users">
        <div className="row">
          <ul className="list-group">
            {listElts}
          </ul>
        </div>
        <div className="row">
          <h3>Create New User</h3>
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
              <label htmlFor="accountName">User Name</label>
			  <input type="text" className="form-control" id="firstName" placeholder="First Name" valueLink={this.linkState('firstName')}/> 
			  <input type="text" className="form-control" id="lastName" placeholder="Last Name"  valueLink={this.linkState('lastName')}/>
			  <input type="text" className="form-control" id="email" placeholder="email"  valueLink={this.linkState('email')}/>
            </div>
            <button className="btn btn-default">Create</button>
          </form>
        </div>
      </div>
    );
  },
});

ReactDOM.render(<Root/>, document.getElementById("page-content"));
var User= React.createClass({
  render: function() {
	  return (
      <li key={this.props.user.id} className="list-group-item">{this.props.user.href} - {this.props.user.firstname} {this.props.user.lastname}</li>)
  }
});


