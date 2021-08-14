import React, {FunctionComponent, useContext, useState} from "react";

type AuthState = {
    loggedIn: boolean

    login: (username: string, password: string) => Promise<string>
}

const emptyAuthState = {
    loggedIn: false,

    login: async () => "not implemented"
}

const AuthContext = React.createContext<AuthState>(emptyAuthState)

export const useAuth = () => {
    return useContext(AuthContext)
}

type AuthProviderProps = {
}

export const AuthProvider: FunctionComponent<AuthProviderProps> = ({children}) => {
    const defaultAuthState = {
        ...emptyAuthState
    }

    const [authState, setAuthState] = useState<AuthState>(defaultAuthState)

    return (
        <AuthContext.Provider value={authState}>
            {children}
        </AuthContext.Provider>
    )
}
