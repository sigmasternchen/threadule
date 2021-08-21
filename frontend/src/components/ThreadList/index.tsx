import {FunctionComponent} from "react";
import {Avatar, List, ListItem, ListItemAvatar, ListItemText, Typography} from "@material-ui/core";
import Thread from "../../api/entities/Thread";
import CustomDate from "../../utils/CustomDate";
import styles from "./ThreadList.module.css"


export type ThreadListProps = {
    threads: Thread[]
}

const ThreadList: FunctionComponent<ThreadListProps> = ({threads}) => {
    if (threads.length == 0) {
        return (
            <Typography style={{
                color: "lightgrey"
            }}>
                No threads scheduled yet.
            </Typography>
        )
    } else {
        return (
            <List>
                {threads.map(thread => {
                    return (
                        <ListItem key={thread.id}>
                            <ListItemAvatar>
                                <Avatar>
                                    {thread.tweets.length}
                                </Avatar>
                            </ListItemAvatar>
                            <ListItemText
                                primary={<span className={styles.title}>{thread.tweets[0].text}</span>}
                                secondary={new CustomDate(thread.scheduled_for).toLocalISOString(false, true)}/>
                        </ListItem>
                    )
                })}
            </List>
        )
    }
}

export default ThreadList