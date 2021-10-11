import axios from "axios";

function clientCall(method, location, { data, params }, options) {
  return axios({
    method: method,
    url: location,
    params: params ? params : null,
    data: data ? data : null,
    ...options,
  })
    .then((resp) => {
      console.log("clientCall ->", resp);
    })
    .catch((error) => {
      console.error("clientCall", error);
    });
}

export default {
  doPostRequest(location, payload, options) {
    console.debug("doPostRequest was called");
    return clientCall("post", location, { data: payload }, options);
  },
  doPutRequest(location, payload, options) {
    console.debug("doPutRequest was called");
    return clientCall("put", location, { data: payload }, options);
  },
  doGetRequest(location, payload, options) {
    console.debug("doGetRequest was called");
    if (payload.params) {
      return clientCall("get", location, { params: payload }, options);
    } else {
      return clientCall("get", location, null, options);
    }
  },
  doPatchRequest(location, payload, options) {
    console.debug("doPatchRequest was called");
    return clientCall("patch", location, payload, options);
  },
};
