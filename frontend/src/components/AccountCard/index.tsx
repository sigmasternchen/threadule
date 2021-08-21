import {FunctionComponent, useState} from "react";
import {Avatar, Card, CardActions, CardContent, CardHeader, IconButton} from "@material-ui/core";
import Account from "../../api/entities/Account";
import ThreadList from "../ThreadList";

import AddIcon from '@material-ui/icons/Add';
import ThreadFormDialog from "../ThreadFormDialog";
import Thread, {ThreadStatus} from "../../api/entities/Thread";
import {TweetStatus} from "../../api/entities/Tweet";

export type AccountCardProps = {
    account: Account
}

const emptyThread = (): Thread => ({
    id: "",
    scheduled_for: new Date(),
    status: ThreadStatus.SCHEDULED,
    tweets: [
        {
            id: "",
            text: "",
            status: TweetStatus.SCHEDULED,
            tweet_id: null,
            error: null,
        }
    ],
    error: null,
})

const AccountCard: FunctionComponent<AccountCardProps> = ({account}) => {
    const [editThread, setEditThread] = useState<Thread | null>(null)

    const openNewForm = () => {
        setEditThread(emptyThread())
    }

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
                    <ThreadList threads={account.threads}/>
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
                initial={editThread ? editThread : emptyThread()}
                onSubmit={(thread) => {
                    account.threads.push(thread)

                    setEditThread(null)
                }}
                onCancel={() => {
                    // TODO show confirmation
                    setEditThread(null)
                }}
            />
        </>
    )
}

export default AccountCard