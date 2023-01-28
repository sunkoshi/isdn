import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import ProvidedApp from "./AuxComponents/ProvidedApp";
import "./index.css";
ReactDOM.render(
  <React.StrictMode>
    <ProvidedApp>
      <App />
    </ProvidedApp>
  </React.StrictMode>,
  document.getElementById("root")
);
