import Thread from "./Thread";

type Account = {
    id: string
    name: string
    screen_name: string
    twitter_id: string
    avatar_url: string
    threads: Thread[]
}

export default Account