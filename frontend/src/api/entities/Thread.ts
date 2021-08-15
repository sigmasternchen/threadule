import Tweet from "./Tweet";

export enum ThreadStatus {
    SCHEDULED = "SCHEDULED",
    PROCESSING = "PROCESSING",
    FAILED = "FAILED",
    DONE = "DONE",
}

type Thread = {
    id: string,
    tweets: Tweet[],
    scheduled_for: Date,
    status: ThreadStatus,
    error: string|null,
}

export default Thread