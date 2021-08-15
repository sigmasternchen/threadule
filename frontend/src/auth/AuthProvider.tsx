import React, {FunctionComponent, useContext, useState} from "react";
import AuthenticationEndpoint from "../api/endpoints/AuthenticationEndpoint";
import {Client, getClient} from "../api/client";
import User from "../api/entities/User";

type AuthState = {
    loggedIn: boolean

    login: (username: string, password: string) => Promise<User>
}

const emptyAuthState = {
    loggedIn: false,

    login: async () => {
        throw "not implemented"
    }
}

const AuthContext = React.createContext<AuthState>(emptyAuthState)

export const useAuth = (): AuthState => {
    return useContext(AuthContext)
}

type AuthProviderProps = {
}


const AuthProvider: FunctionComponent<AuthProviderProps> = ({children}) => {
    const [client, setClient] = useState<Client>(getClient())
    const authenticationEndpoint = new AuthenticationEndpoint(client)

    const [authState, setAuthState] = useState<AuthState>({
        ...emptyAuthState,

        login: async (username: string, password: string): Promise<User> => {
            const response = await authenticationEndpoint.login({
                username: username,
                password: password
            })

            // local new client
            const client = getClient(response.token)
            setClient(client)

            // local new authenticationEndpoint
            const tmpAuthenticationEndpoint = new AuthenticationEndpoint(client)
            return await tmpAuthenticationEndpoint.getUser()
        }
    })

    return (
        <AuthContext.Provider value={authState}>
            {children}
        </AuthContext.Provider>
    )
}

export default AuthProvider