// App.js
import React from "react";
import "./App.css";

const App = () => {
  return (
    <div className="app-container">
      <header className="navbar">
        <h1 className="logo">MapGo</h1>
        <nav>
          <ul>
            <li>Home</li>
            <li>Features</li>
            <li>About</li>
            <li>Contact</li>
          </ul>
        </nav>
      </header>

      <main className="main-content">
        <div className="intro-section">
          <h2>Discover Smarter Navigation</h2>
          <p>
            Real-time traffic updates, optimized routes, and alternate paths to save your time. 
            Navigate smarter with MapGo.
          </p>
          <button className="cta-button">Start Planning</button>
        </div>

        <div className="map-preview">
          <div className="map-placeholder">
            {/* This is where you can embed a map component */}
            <p>Interactive Map Preview</p>
          </div>
          <div className="route-planner">
            <h3>Plan Your Route</h3>
            <form>
              <input type="text" placeholder="Starting Point" />
              <input type="text" placeholder="Destination" />
              <button type="submit">Get Route</button>
            </form>
          </div>
        </div>
      </main>

      <footer className="footer">
        <p>© 2025 MapGo. All rights reserved.</p>
      </footer>
    </div>
  );
};

export default App;
