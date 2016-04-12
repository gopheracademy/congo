// This module exports functions that give access to the congo API hosted at congo.gopheracademy.com.
// It uses the axios javascript library for making the actual HTTP requests.
define(['axios'] , function (axios) {
  function merge(obj1, obj2) {
    var obj3 = {};
    for (var attrname in obj1) { obj3[attrname] = obj1[attrname]; }
    for (var attrname in obj2) { obj3[attrname] = obj2[attrname]; }
    return obj3;
  }

  return function (scheme, host, timeout) {
    scheme = scheme || 'http';
    host = host || 'congo.gopheracademy.com';
    timeout = timeout || 20000;

    // Client is the object returned by this module.
    var client = axios;

    // URL prefix for all API requests.
    var urlPrefix = scheme + '://' + host;

  // Record new user
  // path is the request path, the format is "/api/admin/users"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createAdminuser = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Record new event
  // path is the request path, the format is "/api/tenants/:tenantID/events"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createEvent = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Record new series
  // path is the request path, the format is "/api/tenants/:tenantID/series"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createSeries = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Record new tenant
  // path is the request path, the format is "/api/tenants"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createTenant = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Record new user
  // path is the request path, the format is "/api/tenants/:tenantID/users"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.createUser = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // deleteAdminuser calls the delete action of the adminuser resource.
  // path is the request path, the format is "/api/admin/users/:userID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteAdminuser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // deleteEvent calls the delete action of the event resource.
  // path is the request path, the format is "/api/tenants/:tenantID/events/:eventID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteEvent = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // deleteSeries calls the delete action of the series resource.
  // path is the request path, the format is "/api/tenants/:tenantID/series/:seriesID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteSeries = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // deleteTenant calls the delete action of the tenant resource.
  // path is the request path, the format is "/api/tenants/:tenantID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteTenant = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // deleteUser calls the delete action of the user resource.
  // path is the request path, the format is "/api/tenants/:tenantID/users/:userID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteUser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // List all users
  // path is the request path, the format is "/api/admin/users"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listAdminuser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // List all events
  // path is the request path, the format is "/api/tenants/:tenantID/events"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listEvent = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // List all series
  // path is the request path, the format is "/api/tenants/:tenantID/series"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listSeries = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // List all tenants
  // path is the request path, the format is "/api/tenants"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listTenant = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // List all users for a tenant
  // path is the request path, the format is "/api/tenants/:tenantID/users"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listUser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Obtain a refreshed access token
  // path is the request path, the format is "/api/auth/refresh"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.refreshAuth = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Retrieve user with given id
  // path is the request path, the format is "/api/admin/users/:userID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showAdminuser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Retrieve event with given id
  // path is the request path, the format is "/api/tenants/:tenantID/events/:eventID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showEvent = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Retrieve series with given id
  // path is the request path, the format is "/api/tenants/:tenantID/series/:seriesID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showSeries = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Retrieve tenant with given id
  // path is the request path, the format is "/api/tenants/:tenantID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showTenant = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Retrieve user with given id
  // path is the request path, the format is "/api/tenants/:tenantID/users/:userID"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showUser = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Get Server Status
  // path is the request path, the format is "/api/healthz"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.statusHealthz = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Obtain an access token
  // path is the request path, the format is "/api/auth/token"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.tokenAuth = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // updateAdminuser calls the update action of the adminuser resource.
  // path is the request path, the format is "/api/admin/users/:userID"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateAdminuser = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'patch',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // updateEvent calls the update action of the event resource.
  // path is the request path, the format is "/api/tenants/:tenantID/events/:eventID"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateEvent = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'patch',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // updateSeries calls the update action of the series resource.
  // path is the request path, the format is "/api/tenants/:tenantID/series/:seriesID"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateSeries = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'patch',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // updateTenant calls the update action of the tenant resource.
  // path is the request path, the format is "/api/tenants/:tenantID"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateTenant = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'patch',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // updateUser calls the update action of the user resource.
  // path is the request path, the format is "/api/tenants/:tenantID/users/:userID"
  // data contains the action payload (request body)
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateUser = function (path, data, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'patch',
    data: data,
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // validate user email
  // path is the request path, the format is "/api/validate/:userID"
  // v is used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.validateValidate = function (path, v, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        v: v
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }
  return client;
  };
});
