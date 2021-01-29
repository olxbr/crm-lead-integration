const request = require('supertest');
const express = require('express');

const healthcheckRouter = require('../../../src/routes/healthcheck')

const app = express();
app.use(healthcheckRouter);

describe('should check in the GET / health check successfully', () => {
  it('GET /', (done) => {
    request(app)
      .get('/')
      .set('Accept', 'application/json')
      .expect(200, done);
  });

  it('should check in the GET /health-check successfully', (done) => {
    request(app)
      .get('/health-check')
      .set('Accept', 'application/json')
      .expect(200, done);
  });
});
