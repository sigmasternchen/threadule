import {FunctionComponent, useState} from "react";
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Grid,
    IconButton,
    TextField
} from "@material-ui/core";
import DeleteIcon from '@material-ui/icons/Delete';
import Account from "../../api/entities/Account";
import Thread from "../../api/entities/Thread";
import AddIcon from "@material-ui/icons/Add";
import {TweetStatus} from "../../api/entities/Tweet";
import CustomDate from "../../utils/CustomDate";
import {Alert} from "@material-ui/lab";

export type ThreadFormProps = {
    open: boolean
    account: Account
    initial: Thread
    onSubmit: (thread: Thread) => void
    onCancel: () => void
}

const Index: FunctionComponent<ThreadFormProps> = (
    {
        open,
        account,
        initial,
        onSubmit,
        onCancel
    }
) => {
    const [thread, setThread] = useState<Thread>(initial)

    const [error, setError] = useState<string|null>(null)

    return (
        <Dialog open={open}>
            <DialogTitle title={"Thread"}/>
            <DialogContent>
                <Grid container spacing={2}>
                    <Grid item xs={12}>
                        <TextField
                            id="datetime-local"
                            label="Scheduled For"
                            type="datetime-local"
                            value={
                                new CustomDate(thread.scheduled_for).toLocalISOString(false, false)
                            }
                            onChange={event => {
                                setThread({
                                    ...thread,
                                    scheduled_for: new Date(event.target.value)
                                })
                            }}
                            InputLabelProps={{
                                shrink: true,
                            }}
                        />
                    </Grid>
                    {
                        thread.tweets.map((tweet, index) => (
                            <>
                                <Grid item xs={11}>
                                    <TextField
                                        id="outlined-multiline-static"
                                        label={"Tweet " + (index + 1)}
                                        multiline
                                        rows={3}
                                        value={tweet.text}
                                        onChange={event => {
                                            thread.tweets[index].text = event.target.value
                                            setThread({
                                                ...thread
                                            })
                                        }}
                                        variant="outlined"
                                        fullWidth
                                    />
                                </Grid>
                                <Grid item xs={1}>
                                    <IconButton
                                        aria-label="delete"
                                        color={"secondary"}
                                        disabled={thread.tweets.length <= 1}
                                        onClick={() => {
                                            thread.tweets.splice(index, 1)
                                            setThread({
                                                ...thread
                                            })
                                        }}
                                    >
                                        <DeleteIcon fontSize="medium" />
                                    </IconButton>
                                </Grid>
                            </>
                        ))
                    }
                    <Grid item xs={12}>
                        <IconButton aria-label="add" onClick={() => {
                            thread.tweets.push({
                                id: "",
                                text: "",
                                status: TweetStatus.SCHEDULED,
                                tweet_id: null,
                                error: null
                            })
                            setThread({
                                ...thread
                            })
                        }}>
                            <AddIcon/>
                        </IconButton>
                    </Grid>
                </Grid>
                { error && <Alert severity="error">{error}</Alert> }
            </DialogContent>
            <DialogActions>
                <Button onClick={() => {
                    // TODO show confirmation dialog
                    onCancel()
                }} color="secondary">
                    Cancel
                </Button>
                <Button onClick={() => {
                    if (!thread.tweets.every(t => t.text.trim().length != 0)) {
                        setError("Empty tweets are not allowed!")
                        return
                    }
                    setError(null)
                    onSubmit(thread)
                }} color="primary">
                    Save changes
                </Button>
            </DialogActions>
        </Dialog>
    )
}

export default Index