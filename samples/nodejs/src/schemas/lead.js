const Joi = require('@hapi/joi');

const bodySchema = Joi.object({
  leadOrigin: Joi.string().required(),
  timestamp: Joi.string().required(),
  originLeadId: Joi.string().required(),
  originListingId: Joi.string().required(),
  clientListingId: Joi.string().required(),
  name: Joi.string().required(),
  email: Joi.string().required().regex(/^[\w+._%-]{1,256}@[\w][\w-]{0,64}(\.[\w][\w-]{0,25})+$/),
  ddd: Joi.string().required(),
  phone: Joi.string().required(),
  message: Joi.string().required(),
});

module.exports = {
  bodySchema,
};
