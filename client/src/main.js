import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "./registerServiceWorker";

Vue.config.productionTip = false;

const requireComponent = require.context(
  // The relative path of the components folder
  "./components/formik",
  // Whether or not to look in subfolders
  false,
  // The regular expression used to match form component filenames
  /Form[A-Z]\w+\.(vue|js)$/
);

requireComponent.keys().forEach(fileName => {
  // Get component config
  const componentConfig = requireComponent(fileName);
  // console.log('fileName',fileName
  //   .split('/')
  //   .pop()
  //   .replace(/\.\w+$/, ''))

  // Get PascalCase name of component
  const componentName = fileName
    .split("/")
    .pop()
    .replace(/\.\w+$/, "");

  // Register component globally
  Vue.component(
    componentName,
    // Look for the component options on `.default`, which will
    // exist if the component was exported with `export default`,
    // otherwise fall back to module's root.
    componentConfig.default || componentConfig
  );
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
