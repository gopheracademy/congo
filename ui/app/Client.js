import GoaClient from "./gen/client-es6";
import Axios from "axios";

export default function Client(data) {
	var client = GoaClient();
	client.interceptors.request.use(function (config) {
		config.headers = { "Bearer": data.Auth.AccessToken };
		return config;
  	});
	return client;
}
