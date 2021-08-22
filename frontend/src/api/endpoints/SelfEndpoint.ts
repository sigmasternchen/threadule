import Endpoint from "./Endpoint";
import {Client} from "../client";
import User from "../entities/User";

const API_PREFIX = "/api/self/"

class AuthenticationEndpoint extends Endpoint {
    constructor(client: Client) {
        super(client);
    }

    public async getUser(): Promise<User> {
        return this.get<User>(API_PREFIX)
    }
}

export default AuthenticationEndpoint