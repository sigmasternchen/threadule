import {FunctionComponent, useState} from "react";
import {Avatar, Card, CardActions, CardContent, CardHeader, IconButton} from "@material-ui/core";
import Account from "../../api/entities/Account";
import ThreadList from "../ThreadList";

import AddIcon from '@material-ui/icons/Add';
import ThreadFormDialog from "../ThreadFormDialog";
import Thread, {ThreadStatus} from "../../api/entities/Thread";
import {TweetStatus} from "../../api/entities/Tweet";
import {useAuth} from "../../auth/AuthProvider";
import ThreadEndpoint from "../../api/endpoints/ThreadEndpoint";
import {MessageBox, MessageBoxProps} from "../MessageBox";

export type AccountCardProps = {
    account: Account,
    onUpdate: (account: Account) => void,
}

const emptyThread = (account: Account): Thread => ({
    scheduled_for: new Date(),
    status: ThreadStatus.SCHEDULED,
    account: {
        id: account.id
    },
    tweets: [
        {
            id: "new",
            text: "",
            status: TweetStatus.SCHEDULED,
            tweet_id: null,
            error: null,
            ordinal: 0,
        }
    ],
    error: null,
})

const AccountCard: FunctionComponent<AccountCardProps> = (
    {
        account,
        onUpdate
    }
) => {
    const {client} = useAuth()

    const [editThread, setEditThread] = useState<Thread | null>(null)

    const openNewForm = () => {
        setEditThread(emptyThread(account))
    }
    const openEditForm = (thread: Thread) => {
        setEditThread(thread)
    }

    const [message, setMessage] = useState<MessageBoxProps>({
        open: false,
        success: false,
        message: "",
        onClose: () => {
            setMessage({
                ...message,
                open: false
            })
        }
    })

    return (
        <>
            <Card>
                <CardHeader
                    avatar={
                        <Avatar alt={account.screen_name} src={account.avatar_url}/>
                    }
                    action={
                        <IconButton aria-label="settings">
                        </IconButton>
                    }
                    title={account.name}
                    subheader={account.screen_name}
                />
                <CardContent>
                    <ThreadList
                        threads={account.threads}
                        onSelect={openEditForm}
                    />
                </CardContent>
                <CardActions disableSpacing>
                    <IconButton aria-label="add" onClick={() => {
                        openNewForm()
                    }}>
                        <AddIcon/>
                    </IconButton>
                </CardActions>
            </Card>

            <ThreadFormDialog
                account={account}
                open={Boolean(editThread)}
                initial={editThread ? editThread : emptyThread(account)}
                onSubmit={(thread) => {
                    thread.tweets.forEach(t => {
                        t.id = undefined
                    })

                    const endpoint = new ThreadEndpoint(client)
                    const onSuccess = (result: Thread) => {
                        setEditThread(null)
                        if (!thread.id) {
                            setMessage({
                                ...message,
                                open: true,
                                success: true,
                                message: "Thread was added successfully."
                            })
                            account.threads.push(result)
                        } else {
                            setMessage({
                                ...message,
                                open: true,
                                success: true,
                                message: "Thread was updated successfully."
                            })
                            account.threads = account.threads.filter(t => t.id != result.id)
                            account.threads.push(result)
                        }
                        account.threads = account.threads.sort((a, b) => a.scheduled_for > b.scheduled_for ? 1 : -1)
                        onUpdate(account)
                    }
                    const onFailure = (error: any) => {
                        console.error(error)
                        setMessage({
                            ...message,
                            open: true,
                            success: false,
                            message: "Something went wrong."
                        })
                    }

                    if (!thread.id) {
                        endpoint.add(thread)
                            .then(onSuccess)
                            .catch(onFailure)
                    } else {
                        endpoint.update(thread)
                            .then(onSuccess)
                            .catch(onFailure)
                    }
                }}
                onCancel={() => {
                    // TODO show confirmation
                    setEditThread(null)
                }}
            />

            <MessageBox {...message} />
        </>
    )
}

export default AccountCard