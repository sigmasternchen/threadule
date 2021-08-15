import React, {FunctionComponent, useContext, useEffect, useState} from "react";
import AuthenticationEndpoint from "../api/endpoints/AuthenticationEndpoint";
import {Client, getClient} from "../api/client";
import User from "../api/entities/User";
import {Backdrop, CircularProgress} from "@material-ui/core";

type AuthState = {
    loggedIn: boolean
    user: User | null

    login: (username: string, password: string) => Promise<User>
    getClient: () => Client
}

const emptyAuthState = {
    loggedIn: false,
    user: null,

    login: async () => {
        throw "not implemented"
    },
    getClient: () => getClient()
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

    const [client, setClient] = useState<Client>(getClient(initialSessionToken ? initialSessionToken : undefined))
    const authenticationEndpoint = new AuthenticationEndpoint(client)

    const [authState, setAuthState] = useState<AuthState>({
        ...emptyAuthState,

        login: async (username: string, password: string): Promise<User> => {
            const response = await authenticationEndpoint.login({
                username: username,
                password: password
            })

            // local new client
            const tmpClient = getClient(response.token)
            setClient(tmpClient)

            localStorage.setItem(LOCAL_STORAGE_SESSION_TOKEN_KEY, response.token);

            // local new authenticationEndpoint
            const tmpAuthenticationEndpoint = new AuthenticationEndpoint(tmpClient)
            const user = await tmpAuthenticationEndpoint.getUser()

            setAuthState({
                ...authState,
                loggedIn: true,
                user: user
            })

            return user
        },
        getClient: () => {
            return client
        }
    })


    useEffect(() => {
        if (initialSessionToken) {
            authenticationEndpoint.getUser()
                .then(user => {
                    setAuthState({
                        ...authState,
                        loggedIn: true,
                        user: user
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