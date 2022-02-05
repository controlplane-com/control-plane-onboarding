const express = require('express');
const app = express();

app.get('/', root);

app.listen(8080);

function root(req, res) {

	console.log('Request Received');
	console.log('Workload Name: ' + process.env.CPLN_WORKLOAD);
	console.log('Location: ' + process.env.CPLN_LOCATION);

	res.send('Hello!');
}
