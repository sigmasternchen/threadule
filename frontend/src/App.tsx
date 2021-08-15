import React from 'react';
import './App.css';
import AuthProvider from "./auth/AuthProvider";
import {BrowserRouter, Route, Switch} from "react-router-dom";
import PrivateRoute from "./auth/PrivateRoute";
import Login from "./components/Login";

function App() {
    return (
        <AuthProvider>
            <BrowserRouter>
                <Switch>
                    <Route path="/login">
                        <Login/>
                    </Route>
                    <PrivateRoute path="/">
                        <h1>Private</h1>
                    </PrivateRoute>
                    <Route path="/*">
                        <h1>Test</h1>
                    </Route>
                </Switch>
            </BrowserRouter>
        </AuthProvider>
    );
}

export default App;
