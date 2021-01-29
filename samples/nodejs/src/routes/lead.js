const router = require('express').Router();
const { body } = require('express-joi-validation').createValidator();

const { bodySchema } = require('../schemas/lead');

router.post('/lead', body(bodySchema), async (req, res, next) => {
  try {
    const authorization = req.headers.authorization;
    const base64 = authorization.split(' ')[1];
    const value = Buffer.from(base64, 'base64').toString('utf-8');
    const secretKey = value.split(':')[1];
    if (secretKey !== process.env.SECRET_KEY) {
      return res.sendStatus(401);
    }
    return res.sendStatus(201);
  } catch (err) {
    return next(err);
  }
});

module.exports = router;
