const leadRouter = require('../../../src/routes/lead')

const request = require('supertest');
const express = require('express');

const { lead } = require('../../fixtures/fixtures')

const authorizationHash = 'dml2YXJlYWw6NTk0RjgwM0IzODBBNDEzOTZFRDYzRENBMzk1MDM1NDI=';

const app = express();
app.use(leadRouter);

const env = Object.assign({}, process.env);

after(() => {
  process.env = env;
  process.env.SECRET_KEY = '594F803B380A41396ED63DCA39503542';
});

describe('POST /lead', () => {
  it('should receive the lead successfully', (done) => {
    request(app)
      .post('/lead')
      .send(lead)
      .set('Accept', 'application/json')
      .set('Authorization', `Basic ${authorizationHash}`)
      .expect(201, done);
  });

  it('should receive the lead with the status code 401 unsuccessfully', (done) => {
    process.env.SECRET_KEY = 'BLA';
    request(app)
      .post('/lead')
      .send(lead)
      .set('Accept', 'application/json')
      .set('Authorization', `Basic ${authorizationHash}`)
      .expect(401, done)
  });

  it('should receive the lead unsuccessfully', (done) => {
    request(app)
      .post('/lead')
      .send(lead)
      .set('Accept', 'application/json')
      .set('Authorization', null)
      .expect(500, done)
  });
});
