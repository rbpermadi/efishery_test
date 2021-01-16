/* eslint-disable no-unused-vars */
const jwt = require('jsonwebtoken');
const _ = require('lodash');
const moment = require('moment');
const request = require('request-promise');
const redisCache = require('../../helpers/redis_cache');
const error = require('../../helpers/error');

exports.Storages = class Storages {
  constructor (options) {
    this.options = options || {};
  }

  setup(app) {
    this.app = app;

    /** make request untuk internal API dari micro service ke kong port 8001 */
    this.makeRequest = request.defaults({
      json: true
    });

    this.url = app.get('url');
  }

  async find () {
    try {
      const efisheryResult = await this.makeRequest({
        method: 'GET',
        uri:this.url.efishery
      });

      _.remove(efisheryResult, {
        uuid: null
      });

      const checkRedis = await redisCache.asyncGet('product_api:currency');

      let currency;
      if(!checkRedis){
        currency = await this.makeRequest({
          method: 'GET',
          uri:this.url.currconv
        });

        redisCache.setex('product_api:currency',600,currency);
      }else{
        currency = checkRedis;
      }

      const result = efisheryResult.map(res => {
        res.size = parseInt(res.size);

        res.tgl_parsed = moment(res.tgl_parsed).zone('+07').format('YYYY-MM-DD HH:mm ZZ');

        res.price = parseInt(res.price);
        res.price_usd = res.price * currency.IDR_USD;

        res.id = res.uuid;
        delete res.uuid;

        return res;
      });

      return result;
    }
    catch(err) {
      throw error.InternalServerError();
    }
  }
};
