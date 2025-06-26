/// <reference path="../pb_data_1/types.d.ts" />
migrate(
  (app) => {
    const collection = app.findCollectionByNameOrId("pbc_1125843985");

    // update collection data
    unmarshal(
      {
        createRule: "",
        deleteRule: "",
        listRule: "",
        updateRule: "",
        viewRule: "",
      },
      collection
    );

    return app.save(collection);
  },
  (app) => {
    const collection = app.findCollectionByNameOrId("pbc_1125843985");

    // update collection data
    unmarshal(
      {
        createRule: null,
        deleteRule: null,
        listRule: null,
        updateRule: null,
        viewRule: null,
      },
      collection
    );

    return app.save(collection);
  }
);
