const { FeathersError } = require('@feathersjs/errors');

class CustomError extends FeathersError {
    constructor(message, name, code) {
        super(message, name, code);
    }

    toJSON() {
        return {
          error: {
            message: this.message,
            code: this.code,
          },
          meta: {
            status: this.code
          }
        }
    }
}

class UnauthorizedError extends CustomError {
    constructor() {
        super('Unauthorized', 'unauthorized', 401);
    }
}

class InternalServerError extends CustomError {
    constructor() {
        super('Internal Server Error', 'internal-server', 500);
    }
}

const error = {
    UnauthorizedError: () => {
        return new UnauthorizedError()
    },
    InternalServerError: () => {
        return new InternalServerError()
    }
}

module.exports = error