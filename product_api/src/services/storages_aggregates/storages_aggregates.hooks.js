const { authenticate, isAdmin } = require('../../helpers/auth');
const response = require('../../helpers/response');

module.exports = {
  before: {
    all: [],
    find: [
      authenticate(),
      isAdmin()
    ],
    get: [],
    create: [],
    update: [],
    patch: [],
    remove: []
  },

  after: {
    all: [],
    find: [response.Ok()],
    get: [],
    create: [],
    update: [],
    patch: [],
    remove: []
  },

  error: {
    all: [],
    find: [],
    get: [],
    create: [],
    update: [],
    patch: [],
    remove: []
  }
};
