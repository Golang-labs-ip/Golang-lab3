const http = require("../common/http");

const Client = (baseUrl) => (
  (client = http.Client(baseUrl)),
  {
    ListTablets: () => client.get("/tablets"),
    UpdateTablet: (id) => client.patch("/tablets", {id}),
  }
);

module.exports = { Client };

