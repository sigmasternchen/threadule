import {FunctionComponent, useState} from "react";
import {
    Avatar,
    IconButton,
    List,
    ListItem,
    ListItemAvatar,
    ListItemSecondaryAction,
    ListItemText,
    Typography
} from "@material-ui/core";
import Thread from "../../api/entities/Thread";
import CustomDate from "../../utils/CustomDate";
import styles from "./ThreadList.module.css"
import DeleteIcon from "@material-ui/icons/Delete";
import {useAuth} from "../../auth/AuthProvider";
import ThreadEndpoint from "../../api/endpoints/ThreadEndpoint";
import {MessageBox, MessageBoxProps} from "../MessageBox";


export type ThreadListProps = {
    threads: Thread[]
    onSelect: (thread: Thread) => void
    onDelete: (thread: Thread) => void
}


const ThreadList: FunctionComponent<ThreadListProps> = (
    {
        threads,
        onSelect,
        onDelete
    }
) => {
    const {client} = useAuth()

    const [message, setMessage] = useState<MessageBoxProps>({
        open: false,
        success: false,
        message: "",
        onClose: () => {
            setMessage({
                ...message,
                open: false,
            })
        }
    })

    const deleteThread = (thread: Thread) => {
        new ThreadEndpoint(client).remove(thread)
            .then(() => {
                onDelete(thread)
                setMessage({
                    ...message,
                    open: true,
                    success: true,
                    message: "Thread was deleted successfully."
                })
            })
            .catch((errors) => {
                console.error(errors)
                setMessage({
                    ...message,
                    open: true,
                    success: false,
                    message: "Thread couldn't be deleted."
                })
            })
    }

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
            <>
            <List>
                {threads.map(thread => {
                    return (
                        <ListItem
                            button
                            key={thread.id}
                            onClick={() => {
                                onSelect(thread)
                            }}
                        >
                            <ListItemAvatar>
                                <Avatar>
                                    {thread.tweets.length}
                                </Avatar>
                            </ListItemAvatar>
                            <ListItemText
                                primary={<span className={styles.title}>{thread.tweets[0].text}</span>}
                                secondary={new CustomDate(thread.scheduled_for).toLocalISOString(false, true)}
                            />
                            <ListItemSecondaryAction>
                                <IconButton
                                    edge="end"
                                    aria-label="delete"
                                    onClick={() => deleteThread(thread)}
                                >
                                    <DeleteIcon />
                                </IconButton>
                            </ListItemSecondaryAction>
                        </ListItem>
                    )
                })}
            </List>
                <MessageBox {...message} />
            </>
        )
    }
}

export default ThreadList