import Endpoint from "./Endpoint";
import {Client} from "../client";
import {LoginParams, LoginResponse} from "../entities/login";
import User from "../entities/User";

const API_PREFIX = "http://localhost:8080/authentication/"

class AuthenticationEndpoint extends Endpoint {
    constructor(client: Client) {
        super(client);
    }

    public async login(params: LoginParams): Promise<LoginResponse> {
        return this.post<LoginResponse>(API_PREFIX, {
            data: params
        })
    }

    public async getUser(): Promise<User> {
        return this.get<User>(API_PREFIX)
    }
}

export default AuthenticationEndpoint