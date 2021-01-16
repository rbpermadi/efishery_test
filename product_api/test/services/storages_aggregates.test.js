const assert = require('assert');
const app = require('../../src/app');

describe('\'storages_aggregates\' service', () => {
  it('registered the service', () => {
    const service = app.service('storages-aggregates');

    assert.ok(service, 'Registered the service');
  });
});
