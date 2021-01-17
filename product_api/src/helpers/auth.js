const jwt = require('jsonwebtoken');
const error = require('./error');

module.exports = {
  authenticate: () => {
    return (context) => {
      try {
        const jwtSecretKey = context.app.get("jwt_private_key");

        const {headers} = context.params;

        const token = headers.authorization.replace('Token ', '');

        const decoded = jwt.verify(token, jwtSecretKey);

        context.params.user = decoded

        return context;
      }
      catch(err) {
        throw error.UnauthorizedError();
      }
    };
  },

  isAdmin: () => {
    return (context) => {
      try {
        if (context.params.user.role != "admin") {
          throw error.UnauthorizedError();
        }
        return context
      }
      catch(err) {
        throw err;
      }
    };
  }
}