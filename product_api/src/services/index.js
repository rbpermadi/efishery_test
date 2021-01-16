const me = require('./me/me.service.js');
const storages = require('./storages/storages.service.js');
const storagesAggregates = require('./storages_aggregates/storages_aggregates.service.js');
// eslint-disable-next-line no-unused-vars
module.exports = function (app) {
  app.configure(me);
  app.configure(storages);
  app.configure(storagesAggregates);
};
