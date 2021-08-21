import {Client} from "../client";
import Account from "../entities/Account";
import Endpoint from "./Endpoint";

const API_PREFIX = "/api/account/"

export type AddAccountResponse = {
    id: string
    url: string
}

export type AddAccountResolveParam = {
    pin: string
}

class AccountEndpoint extends Endpoint {
    constructor(client: Client) {
        super(client)
        this.requireAuthenticated()
    }

    public async getAll(): Promise<Account[]> {
        return await this.get<Account[]>(API_PREFIX)
    }

    public async addAccount(): Promise<AddAccountResponse> {
        return await this.post<any, AddAccountResponse>(API_PREFIX, null)
    }

    public async addAccountResolve(accoundId: string, pin: string): Promise<void> {
        return await this.post<AddAccountResolveParam, void>(API_PREFIX + "/" + accoundId, {
            pin: pin
        })
    }
}

export default AccountEndpoint