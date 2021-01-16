// Initializes the `me` service on path `/me`
const { Me } = require('./me.class');
const hooks = require('./me.hooks');

module.exports = function (app) {
  const options = {
    paginate: app.get('paginate')
  };

  // Initialize our service with any options it requires
  app.use('/me', new Me(options, app));

  // Get our initialized service so that we can register hooks
  const service = app.service('me');

  service.hooks(hooks);
};
