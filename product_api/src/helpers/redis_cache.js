/* eslint-disable no-unused-vars */
'use strict';

const fastStringify = require('fast-safe-stringify');
const JSONparse = require('fast-json-parse');
const redisPoolCon = require('redis-pool-connection');

/** load app feathers */
const feathers = require('@feathersjs/feathers');
const configuration = require('@feathersjs/configuration');
const app = feathers().configure(configuration());

/** connection ke redis */
const redisClient = redisPoolCon(app.get('redis'));

const redisCache = {
  /** get value by key */
  asyncGet: (key) => {
    return new Promise((resolve, reject) => {
      redisClient.get(key, (err, response) => {
        if (err) return reject(err);

        const data = JSONparse(response);
        resolve(data.value);
      });
    });
  },
  set: (key, data) => {
    redisClient.set(key, fastStringify(data));
  },
  /** set key name, expired time, value */
  setex: (key, ttl, data) => {
    const expiredTime = ttl || 60;

    redisClient.setex(key, expiredTime, fastStringify(data));
  },
  ttl: {
    FIVE_MINUTE: 300,
    TEN_MINUTE: 600,
    HALF_HOUR: 1800,
    ONE_HOUR: 3600,
    TWO_HOUR: 7200,
    SIX_HOUR: 21600
  }
};

module.exports = redisCache;