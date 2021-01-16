const http = require("../common/http");

const Client = (baseUrl) => (
  (client = http.Client(baseUrl)),
  {
    ListTablets: () => client.get("/tablets"),
    UpdateMachine: (id, isWorking) => client.patch(`/tablets`, { id, isWorking }),
  }
);

module.exports = { Client };