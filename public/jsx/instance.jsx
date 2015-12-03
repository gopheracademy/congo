var Root = React.createClass({
  mixins: [React.addons.LinkedStateMixin],
  getInitialState: function() {
    return {
      instances: [],
	  name: "",
    };
  },
  componentDidMount: function() {
      var path = this.getPath();
	  var account = this.getAccount(path);
   	var series = this.getSeries(path);
      
	console.log(account);
	console.log(series)
	var u = "/api/accounts/" + account + "/series/" + series + "/instances"
	  $.getJSON(u, function(data) {
      this.setState({instances: data});
    }.bind(this));
  },
 getUrl: function() {
    return '' + window.location;
  },
  getPath: function() {
    return window.location.pathname + window.location.search;
  },
  getAccount: function(path) {
    var urlPieces = path.split("/");
    var account = (urlPieces[2]) ? urlPieces[2] : "";
    return account;
  },
  getSeries: function(path) {
    var urlPieces = path.split("/");
    var series= (urlPieces[4]) ? urlPieces[4] : "";
    return series;
  },
  handleSubmit: function(e) {
	e.preventDefault();
      var path = this.getPath();
	  var account = this.getAccount(path);
	  var series = this.getSeries(path);
	  var u = "/api/accounts/" + account + "/series/" + series + "/instances"
    $.ajax({
	url: u,
      data: JSON.stringify({ "name": this.state.name}),
      method: "POST",
      success: function(data, textStatus, request) {
		  var newNames = this.state.instances;
		  var href = request.getResponseHeader('Location');
		   newNames.push({"href": href, "id": href.substr(href.length -1), "name": this.state.name});
        this.setState({
			name: "",
          instances: newNames
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
    for(var i = 0; i < this.state.instances.length; i++) {
      var instance = this.state.instances[i];
      listElts.push(<li key={instance.id} className="list-group-item">{instance.href} - {instance.name}</li>);
    }
    return (
      <div className="container instances">
        <div className="row">
          <ul className="list-group">
            {listElts}
          </ul>
        </div>
        <div className="row">
          <h3>Create New Instance</h3>
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
              <label htmlFor="instanceName">Instance Name</label>
			  <input type="text" className="form-control" id="name" placeholder="Name" valueLink={this.linkState('name')}/> 
            </div>
            <button className="btn btn-default">Create</button>
          </form>
        </div>
      </div>
    );
  },
});

ReactDOM.render(<Root/>, document.getElementById("page-content"));
