import React from "react";
import ReactDOM from "react-dom/client"; // Import from 'react-dom/client'
import App from "./app";

// Create a root to mount the app
const root = ReactDOM.createRoot(document.getElementById("root"));

// Render the app inside the root
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
