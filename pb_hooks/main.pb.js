/// <reference path="../pb_data_1/types.d.ts" />
routerAdd("GET", "/anotherRoute/hello", (e) => {
  return e.json({ message: "Hello, world!" });
});
