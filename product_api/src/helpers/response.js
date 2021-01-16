const response = {
  Ok: () => {
    return (context) => {
      context.result = {
        result: context.result,
        meta: {
          status: 200
        }
      }
      return context
    }
  },
  Created: () => {
    return (context) => {
      context.result = {
        result: context.result,
        meta: {
          status: 201
        }
      }
      return context
    }
  }
}

module.exports = response