import {FunctionComponent} from "react";
import {useAuth} from "./AuthProvider";
import {Redirect, Route, RouteProps} from "react-router-dom";

type PrivateRouteProps = RouteProps & {
}

const PrivateRoute: FunctionComponent<PrivateRouteProps> = ({
    ...props
}) => {
    const { loggedIn } = useAuth();

    const ok = loggedIn

    if (ok) {
        return <Route {...props} />;
    } else {
        return <Redirect to="/login" />;
    }
}

export default PrivateRoute