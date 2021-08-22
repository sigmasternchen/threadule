import {Client} from "../client";
import Endpoint from "./Endpoint";
import Thread from "../entities/Thread";

const API_PREFIX = "/api/thread/"

class ThreadEndpoint extends Endpoint {
    constructor(client: Client) {
        super(client)
        this.requireAuthenticated()
    }

    public async add(thread: Thread): Promise<Thread> {
        return await this.post<Thread, Thread>(API_PREFIX, thread)
    }

    public async update(thread: Thread): Promise<Thread> {
        return await this.put<Thread, Thread>(API_PREFIX + thread.id, thread)
    }
}

export default ThreadEndpoint