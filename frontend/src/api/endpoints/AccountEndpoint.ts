import {Client} from "../client";
import Account from "../entities/Account";
import Endpoint from "./Endpoint";

const API_PREFIX = "/api/account/"

class AccountEndpoint extends Endpoint {
    constructor(client: Client) {
        super(client)
        this.requireAuthenticated()
    }

    public async getAll(): Promise<Account[]> {
        return await this.get<Account[]>(API_PREFIX)
    }
}

export default AccountEndpoint