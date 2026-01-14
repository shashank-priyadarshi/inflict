# API

## Expectations

- Quick feedback iterations
- Ready-For-Use ASAP, hence, minimal time to production

## Decisions

- `Entity` being the foundational object, CRUD for entities should be first priority.
    ```
      Entity -> Members -> Wealth
             -> Family -> Members -> Wealth
             -> Wealth
    ```
  However, since relationship is like above, API creation will be Bottom-Up.
