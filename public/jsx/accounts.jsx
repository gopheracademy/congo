var Root = React.createClass({
  getInitialState: function() {
    return {
      accounts: [],
      formAccountName: ""
    };
  },
  componentDidMount: function() {
    $.getJSON("/api/accounts", function(data) {
      this.setState({accounts: data});
    }.bind(this));
  },
  handleChange: function(evt) {
    evt.preventDefault();
    this.setState({formAccountName: evt.target.value});
  },
  handleSubmit: function(e) {
    e.preventDefault();
    $.ajax({
      url: "/api/accounts",
      data: JSON.stringify({"name": this.state.formAccountName}),
      method: "POST",
      success: function(data, textStatus, request) {
		  var newNames = this.state.accounts;
		  var href = request.getResponseHeader('Location');
		   newNames.push({"href": href, "id": href.substr(href.length -1), "name": this.state.formAccountName});
        this.setState({
          formAccountName: "",
          accounts: newNames
        });
      }.bind(this)
    });
  },
  render: function() {
    var listElts = [];
    for(var i = 0; i < this.state.accounts.length; i++) {
		var acct = this.state.accounts[i];
		var usersUrl = "/accounts/" + acct.id + "/users";
		var seriesUrl = "/accounts/" + acct.id + "/series";
	  listElts.push(<li key={acct.id} className="list-group-item">{acct.href} - {acct.name} <a href={usersUrl}>Show Users</a> <a href={seriesUrl}>Show Series</a></li>);
    }
    return (
      <div className="container accounts">
        <div className="row">
          <ul className="list-group">
            {listElts}
          </ul>
        </div>
        <div className="row">
          <h3>Create New Account</h3>
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
              <label htmlFor="accountName">Account Name</label>
              <input type="text" className="form-control" value={this.state.formAccountName} id="accountName" placeholder="Account Name" onChange={this.handleChange}/>
            </div>
            <button className="btn btn-default">Create</button>
          </form>
        </div>
      </div>
    );
  },
});

ReactDOM.render(<Root/>, document.getElementById("page-content"));
