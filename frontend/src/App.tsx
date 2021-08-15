import React from 'react';
import './App.css';
import AuthProvider from "./auth/AuthProvider";
import {BrowserRouter, Redirect, Route, Switch} from "react-router-dom";
import PrivateRoute from "./auth/PrivateRoute";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import MenuBar from "./components/MenuBar";

function App() {
    return (
        <AuthProvider>
            <BrowserRouter>
                <Switch>
                    <Route path="/login">
                        <Login/>
                    </Route>
                    <PrivateRoute path="/">
                        <MenuBar />
                        <Dashboard />
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
