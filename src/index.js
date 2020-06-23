import Svelte from './App.svelte'
import { mapReady } from "./stores/common";

// eslint-disable-next-line no-new
new Svelte({
  target: document.body
})

window.initMap = function ready() {
  mapReady.set(true);
}