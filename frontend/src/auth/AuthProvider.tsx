import React, {FunctionComponent, useContext, useEffect, useState} from "react";
import AuthenticationEndpoint from "../api/endpoints/AuthenticationEndpoint";
import {Client, getClient} from "../api/client";
import User from "../api/entities/User";
import {Backdrop, CircularProgress} from "@material-ui/core";

type AuthState = {
    loggedIn: boolean
    user: User | null

    login: (username: string, password: string) => Promise<User>
    logout: () => void
    client: Client
}

const emptyAuthState = {
    loggedIn: false,
    user: null,

    login: async () => {
        throw "not implemented"
    },
    logout: () => {},
    client: getClient(),
}

const AuthContext = React.createContext<AuthState>(emptyAuthState)

export const useAuth = (): AuthState => {
    return useContext(AuthContext)
}

type AuthProviderProps = {
}

const LOCAL_STORAGE_SESSION_TOKEN_KEY = "session_token"

const initialSessionToken = localStorage.getItem(LOCAL_STORAGE_SESSION_TOKEN_KEY)

const AuthProvider: FunctionComponent<AuthProviderProps> = ({children}) => {
    const [loading, setLoading] = useState<boolean>(true)

    const [authState, setAuthState] = useState<AuthState>({
        ...emptyAuthState,

        login: async (username: string, password: string): Promise<User> => {
            let authenticationEndpoint = new AuthenticationEndpoint(authState.client)

            const response = await authenticationEndpoint.login({
                username: username,
                password: password
            })

            localStorage.setItem(LOCAL_STORAGE_SESSION_TOKEN_KEY, response.token);
            const client = getClient(response.token)

            // local new authenticationEndpoint
            authenticationEndpoint = new AuthenticationEndpoint(client)
            const user = await authenticationEndpoint.getUser()

            setAuthState({
                ...authState,
                loggedIn: true,
                user: user,
                client: client
            })

            return user
        },
        logout: () => {
            localStorage.removeItem(LOCAL_STORAGE_SESSION_TOKEN_KEY)
            setAuthState({
                ...authState,
                loggedIn: false,
                user: null,
                client: getClient(),
            })

            // don't use DOM router
            window.location.replace("/login")
        }
    })


    useEffect(() => {
        if (initialSessionToken) {
            const client = getClient(initialSessionToken)
            const authenticationEndpoint = new AuthenticationEndpoint(client)
            authenticationEndpoint.getUser()
                .then(user => {
                    setAuthState({
                        ...authState,
                        loggedIn: true,
                        user: user,
                        client: client
                    })
                    setLoading(false)
                })
                .catch(_ => {
                    localStorage.removeItem(LOCAL_STORAGE_SESSION_TOKEN_KEY);
                    setLoading(false)
                })
        } else {
            setLoading(false)
        }
    }, [])

    if (loading) {
        return (
            <Backdrop open={true}>
                <CircularProgress color="inherit" />
            </Backdrop>
        )
    } else {
        return (
            <AuthContext.Provider value={authState}>
                {children}
            </AuthContext.Provider>
        )
    }
}

export default AuthProvider