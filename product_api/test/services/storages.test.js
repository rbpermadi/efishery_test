const assert = require('assert');
const app = require('../../src/app');

describe('\'storages\' service', () => {
  it('registered the service', () => {
    const service = app.service('storages');

    assert.ok(service, 'Registered the service');
  });
});
