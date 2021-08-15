import axios, {AxiosInstance} from "axios";

const AUTHORIZATION_TYPE = "Bearer"

export type Client = {
    axios: AxiosInstance
    authorized: boolean
}

export const getClient = (authToken?: string): Client => {
    if (authToken) {
        console.log(authToken)
        return {
            axios: axios.create({
                headers: {
                    "Authorization": AUTHORIZATION_TYPE + " " + authToken
                }
            }),
            authorized: true
        }
    } else {
        return {
            axios: axios,
            authorized: false
        }
    }
}