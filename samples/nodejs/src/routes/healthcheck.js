const router = require('express').Router();

router.get('/', function (req, res) {
  res.send('OK')
});

router.get('/health-check', function (req, res) {
  res.send('OK')
});

module.exports = router;
