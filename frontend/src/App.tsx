import React from 'react';
import './App.css';
import AuthProvider from "./auth/AuthProvider";
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom";
import PrivateRoute from "./auth/PrivateRoute";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import MenuBar from "./components/MenuBar";
import Settings from "./pages/Settings";

function App() {
    return (
        <AuthProvider>
            <BrowserRouter>
                <Switch>
                    <Route path="/login">
                        <Login/>
                    </Route>
                    <PrivateRoute exact path="/">
                        <MenuBar pageName="Dashboard" />
                        <Dashboard />
                    </PrivateRoute>
                    <PrivateRoute path="/settings">
                        <MenuBar pageName="Settings" />
                        <Settings />
                    </PrivateRoute>
                    <Route path="/*">
                        <Redirect to={"/login"} />
                    </Route>
                </Switch>
            </BrowserRouter>
        </AuthProvider>
    );
}

export default App;
