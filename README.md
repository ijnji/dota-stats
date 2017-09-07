## Develop via local appengine instance

dev_appserver.py .

## Deploy to appengine

gcloud app deploy

## Manual update tasks

On a new major Dota version, update constants.js

## TODOs

- add tooltips to explain mu, sigma
- add estimated WL + Role stats to the heroes page
- currently the prior is Beta(75, 75), evaluate alternatives that have solid mathematical rationale. User feedback states the current value seems high.

### Visualizations

- replace vega-embed with vue-vega when it works
- add graphs conceptually similar to K+A vs. D
  - GPM, XPM?

- add interactivity
  - input to limit date range or patch versions
  - alternative facet dimensions, e.g. "str/agi/int" instead of "win/loss"
  - filtering games by hero role / player position (GPM as proxy for 1-5?)

- add timeseries of windowed estimates
  - estimated W/L ratio per hero or per role?
  - user-configurable window size, either #games or #days?
  - possibly show via cubism.js or by addining spark lines to vega-lite?

## DONTDOs

- add webpack
