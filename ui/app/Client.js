import GoaClient from "./gen/client";
// requirejs(['./gen/client'], function (client) {
import Axios from "axios";

export default function Client(data) {
	var client = GoaClient(Axios);
	client.interceptors.request.use(function (config) {
		config.headers = { "Bearer": data.Auth.AccessToken };
		return config;
  	});
	console.log("CLIENT: " + client)
	return client;
}
