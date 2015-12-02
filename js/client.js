// This module exports functions that give access to the congo API hosted at api.gopheracademy.com.
// It uses the axios javascript library for making the actual HTTP requests.
define(['axios'] , function (axios) {
    return function (scheme, host, timeout) {
        scheme = scheme || 'http';
        host = host || 'api.gopheracademy.com';
        timeout = timeout || 20000;

        // Client is the object returned by this module.
        var client = axios;

        // URL prefix for all API requests.
        var urlPrefix = scheme + '://' + host;

        // Create new account
        // path is the request path, the format is "/api/accounts"
        // data contains the action payload (request body)
        // config is an optional object to be merged into the config built by the function prior to making the request.
        // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
        // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
        client.createAccount = function (path, data, config) {
            cfg = {
                timeout: timeout,
                url: urlPrefix + path,
                method: 'post',
                data: data,
                responseType: 'json'
            };
            if (config) {
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // Record new series
        // path is the request path, the format is "/api/accounts/:accountID/series"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // Record new user
        // path is the request path, the format is "/api/accounts/:accountID/users"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // deleteAccount calls the delete action of the account resource.
        // path is the request path, the format is "/api/accounts/:accountID"
        // config is an optional object to be merged into the config built by the function prior to making the request.
        // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
        // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
        client.deleteAccount = function (path, config) {
            cfg = {
                timeout: timeout,
                url: urlPrefix + path,
                method: 'delete',
                responseType: 'json'
            };
            if (config) {
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // deleteSeries calls the delete action of the series resource.
        // path is the request path, the format is "/api/accounts/:accountID/series/:seriesID"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // deleteUser calls the delete action of the user resource.
        // path is the request path, the format is "/api/accounts/:accountID/users/:userID"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // List all accounts
        // path is the request path, the format is "/api/accounts"
        // config is an optional object to be merged into the config built by the function prior to making the request.
        // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
        // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
        client.listAccount = function (path, config) {
            cfg = {
                timeout: timeout,
                url: urlPrefix + path,
                method: 'get',
                responseType: 'json'
            };
            if (config) {
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // List all series in account
        // path is the request path, the format is "/api/accounts/:accountID/series"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // List all users in account
        // path is the request path, the format is "/api/accounts/:accountID/users"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // Retrieve account with given id
        // path is the request path, the format is "/api/accounts/:accountID"
        // config is an optional object to be merged into the config built by the function prior to making the request.
        // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
        // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
        client.showAccount = function (path, config) {
            cfg = {
                timeout: timeout,
                url: urlPrefix + path,
                method: 'get',
                responseType: 'json'
            };
            if (config) {
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // Retrieve series with given id
        // path is the request path, the format is "/api/accounts/:accountID/series/:seriesID"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // Retrieve user with given id
        // path is the request path, the format is "/api/accounts/:accountID/users/:userID"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // Change account name
        // path is the request path, the format is "/api/accounts/:accountID"
        // data contains the action payload (request body)
        // config is an optional object to be merged into the config built by the function prior to making the request.
        // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
        // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
        client.updateAccount = function (path, data, config) {
            cfg = {
                timeout: timeout,
                url: urlPrefix + path,
                method: 'put',
                data: data,
                responseType: 'json'
            };
            if (config) {
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // updateSeries calls the update action of the series resource.
        // path is the request path, the format is "/api/accounts/:accountID/series/:seriesID"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }

        // updateUser calls the update action of the user resource.
        // path is the request path, the format is "/api/accounts/:accountID/users/:userID"
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
                cfg = utils.merge(cfg, config);
            }
            return client(cfg);
        }
        return client;
    };
});
