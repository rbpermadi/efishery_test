/* eslint-disable no-unused-vars */
const _ = require('lodash');
const moment = require('moment');
const request = require('request-promise');
const error = require('../../helpers/error');

exports.StoragesAggregates = class StoragesAggregates {
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

  async find (params) {
    try {
      const efisheryResult = await this.makeRequest({
        method: 'GET',
        uri:this.url.efishery
      });

      _.remove(efisheryResult, {
        uuid: null
      });

      const provinceAggs = new Map()
      const weeklyAggs = new Map()
      efisheryResult.forEach(item => {
        if (item.price != null) {
          let price = parseInt(item.price)

          if (item.area_provinsi != null) {
            let key = item.area_provinsi.toLowerCase().replace(" ", "_");
            if (provinceAggs.has(key)) {
              let temp = provinceAggs.get(key)
              if (temp.value.min > price) {
                temp.value.min = price
              } else if (temp.value.max < price) {
                temp.value.max = price
              }
              temp.value.count++
              temp.value.total += price
              temp.value.median = (temp.value.max + temp.value.min)/2
              temp.value.avg = temp.value.total / temp.value.count

              provinceAggs.set(key, temp)
            } else {
              provinceAggs.set(
                key,
                {
                  key: item.area_provinsi,
                  value: {
                    min: price,
                    max: price,
                    median: price,
                    avg: price,
                    total: price,
                    count: 1,
                  }
                }
              )
            }
          }

          if (item.tgl_parsed != null) {
            let firstDay = moment(item.tgl_parsed).startOf('isoWeek').format('YYYYMMDD')
            let lastDay = moment(item.tgl_parsed).endOf('isoWeek').format('YYYYMMDD')
            let weeklyRange = firstDay + "_" + lastDay
            if (weeklyAggs.has(weeklyRange)) {
              let temp = weeklyAggs.get(weeklyRange)
              if (temp.value.min > price) {
                temp.value.min = price
              } else if (temp.value.max < price) {
                temp.value.max = price
              }
              temp.value.count++
              temp.value.total += price
              temp.value.median = (temp.value.max + temp.value.min)/2
              temp.value.avg = temp.value.total / temp.value.count

              weeklyAggs.set(weeklyRange, temp)
            } else {
              weeklyAggs.set(
                weeklyRange,
                {
                  key: weeklyRange,
                  value: {
                    min: price,
                    max: price,
                    median: price,
                    avg: price,
                    total: price,
                    count: 1,
                  }
                }
              )
            }
          }
        }
      })

      return {
        province: Array.from(provinceAggs.values()),
        weekly: Array.from(weeklyAggs.values()),
      };
    }
    catch(err) {
      console.log(err)
      throw error.InternalServerError();
    }
  }
};
