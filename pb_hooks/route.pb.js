/// <reference path="../pb_data_1/types.d.ts" />
routerAdd("GET", "/hello/{name}", (e) => {
  let name = e.request.pathValue("name");

  return e.json(200, { message: "Hello " + name });
});
