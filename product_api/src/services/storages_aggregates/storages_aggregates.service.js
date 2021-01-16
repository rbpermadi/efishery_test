// Initializes the `storages_aggregates` service on path `/storages-aggregates`
const { StoragesAggregates } = require('./storages_aggregates.class');
const hooks = require('./storages_aggregates.hooks');

module.exports = function (app) {
  const options = {
    paginate: app.get('paginate')
  };

  // Initialize our service with any options it requires
  app.use('/storages-aggregates', new StoragesAggregates(options, app));

  // Get our initialized service so that we can register hooks
  const service = app.service('storages-aggregates');

  service.hooks(hooks);
};
