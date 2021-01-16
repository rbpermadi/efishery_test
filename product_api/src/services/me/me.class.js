/* eslint-disable no-unused-vars */
exports.Me = class Me {
  constructor (options) {
    this.options = options || {};
  }

  setup(app) {
    this.app = app;
  }

  async find (params) {
    delete params.user.exp
    return params.user;
  }
};
