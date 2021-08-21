
export enum TweetStatus {
    SCHEDULED = "SCHEDULED",
    FAILED = "FAILED",
    DONE = "DONE",
}

type Tweet = {
    id?: string,
    text: string,
    status: TweetStatus,
    tweet_id: number|null,
    error: string|null
}

export default Tweet