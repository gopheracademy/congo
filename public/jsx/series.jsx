var Root = React.createClass({
  mixins: [React.addons.LinkedStateMixin],
  getInitialState: function() {
    return {
      series: [],
	  name: "",
    };
  },
  componentDidMount: function() {
      var path = this.getPath();
      var account = this.getAccount(path);
      console.log(account);
    var u = "/api/accounts/" + account + "/series"
	  $.getJSON(u, function(data) {
      this.setState({series: data});
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
  handleSubmit: function(e) {
	e.preventDefault();
      var path = this.getPath();
      var account = this.getAccount(path);
    var u = "/api/accounts/" + account + "/series"
    $.ajax({
	url: u,
      data: JSON.stringify({"name": this.state.name}),
      method: "POST",
      success: function(data, textStatus, request) {
		  var newNames = this.state.series;
		  var href = request.getResponseHeader('Location');
		   newNames.push({"href": href, "id": href.substr(href.length -1), "name": this.state.name});
        this.setState({
			name: "",
          series: newNames
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
    for(var i = 0; i < this.state.series.length; i++) {
      var series = this.state.series[i];
      listElts.push(<li key={series.id} className="list-group-item">{series.href} - {series.name}</li>);
    }
    return (
      <div className="container series">
        <div className="row">
          <ul className="list-group">
            {listElts}
          </ul>
        </div>
        <div className="row">
          <h3>Create New Series</h3>
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
              <label htmlFor="seriesName">Series Name</label>
			  <input type="text" className="form-control" id="name" placeholder="Name"  valueLink={this.linkState('name')}/>
            </div>
            <button className="btn btn-default">Create</button>
          </form>
        </div>
      </div>
    );
  },
});

ReactDOM.render(<Root/>, document.getElementById("page-content"));
