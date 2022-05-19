// import axios from "axios";

// function clientCall(method, location, data, options) {
//   return axios({
//     method: method,
//     url: location,
//     data: data,
//     ...options,
//   })
//     .then((resp) => {
//       console.log("clientCall ->", resp);
//       return resp;
//     })
//     .catch((error) => {
//       console.log({ method, location, data, options });
//       console.error("clientCall", error);
//       return -1;
//     });
// }

// export default {
//   doPostRequest(location, payload, options) {
//     console.debug("doPostRequest was called");
//     return clientCall("post", location, payload, options);
//   },
//   doPutRequest(location, payload, options) {
//     console.debug("doPutRequest was called");
//     return clientCall("put", location, { data: payload }, options);
//   },
//   doGetRequest(location, options) {
//     console.debug("doGetRequest was called");
//     return clientCall("get", location, null, options);
//   },
//   doPatchRequest(location, payload, options) {
//     console.debug("doPatchRequest was called");
//     return clientCall("patch", location, payload, options);
//   },
// };


/* File where all calls to backends can be stored */

import axios from 'axios';

const clientCall = (promise, location, requestType) =>{
  return promise
    .then((resp) => {
      console.log(requestType + ' ' + location + ':');
      return resp;
    })
    .catch((error) => {
      if (error.response) {
        //TODO

        return error.response;
      }
      return null;
    });
};

export default {
  async doGetRequest(location, queryData, config = {}) {
    console.log('config:', config);
    
    let qd = queryData ? `?${queryData}` : '';
    console.log('queryData:', queryData);

    return clientCall(axios.get(location + qd, config), location + qd, 'doGetRequest()');
  },
  async doPostRequest(location, queryData, formData = new FormData(), config = {}) {
    let qd = queryData ? `?${queryData}` : '';

    return clientCall(axios.post(location + qd, formData, config), location + qd, 'doPostRequest()');
  },
  async doPutRequest(location, queryData, formData = new FormData(), config = {}) {
    let qd = queryData ? `?${queryData}` : '';

    return clientCall(axios.put(location + qd, formData, config), location + qd, 'doPutRequest()');
  },
};
