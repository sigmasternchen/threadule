import React from 'react';
import './App.css';
import {AuthProvider} from "./auth/AuthProvider";

function App() {
  return (
    <AuthProvider>
        <div className="App">
            <h1>Hi</h1>
        </div>
    </AuthProvider>
  );
}

export default App;
