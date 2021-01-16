const me = require('./me/me.service.js');
const storages = require('./storages/storages.service.js');
// eslint-disable-next-line no-unused-vars
module.exports = function (app) {
  app.configure(me);
  app.configure(storages);
};
