// Initializes the `storages` service on path `/storages`
const { Storages } = require('./storages.class');
const hooks = require('./storages.hooks');

module.exports = function (app) {
  const options = {
    paginate: app.get('paginate')
  };

  // Initialize our service with any options it requires
  app.use('/storages', new Storages(options, app));

  // Get our initialized service so that we can register hooks
  const service = app.service('storages');

  service.hooks(hooks);
};
