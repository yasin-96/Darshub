import axios from "axios";

function clientCall(method, location, data, options) {
  return axios({
    method: method,
    url: location,
    data: data,
    ...options,
  })
    .then((resp) => {
      console.log("clientCall ->", resp);
      return resp;
    })
    .catch((error) => {
      console.log({method, location, data, options })
      console.error("clientCall", error);
      return -1;
    });
}

export default {
  doPostRequest(location, payload, options) {
    console.debug("doPostRequest was called");
    return clientCall("post", location, payload , options);
  },
  doPutRequest(location, payload, options) {
    console.debug("doPutRequest was called");
    return clientCall("put", location, { data: payload }, options);
  },
  doGetRequest(location, options) {
    console.debug("doGetRequest was called");
    return clientCall("get", location, null, options);
  },
  doPatchRequest(location, payload, options) {
    console.debug("doPatchRequest was called");
    return clientCall("patch", location, payload, options);
  },
};
