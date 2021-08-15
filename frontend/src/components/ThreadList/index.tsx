import {FunctionComponent} from "react";
import {Avatar, List, ListItem, ListItemAvatar, ListItemText, Typography} from "@material-ui/core";
import Thread from "../../api/entities/Thread";


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
                {threads.map(thread => (
                    <ListItem key={thread.id}>
                        <ListItemAvatar>
                            <Avatar>
                                {thread.id}
                            </Avatar>
                        </ListItemAvatar>
                        <ListItemText primary="Thread" secondary={thread.scheduled_for}/>
                    </ListItem>
                ))}
            </List>
        )
    }
}

export default ThreadList