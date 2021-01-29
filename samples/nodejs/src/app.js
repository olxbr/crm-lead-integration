const express = require('express');
const bodyParser = require('body-parser');

const healthcheck = require('./routes/healthcheck');
const lead = require('./routes/lead');

const port = process.env.PORT || 3334;

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

app.use(healthcheck);
app.use(lead);

app.listen(port, () => {
  console.log(`api is running on port ${port}`);
});
