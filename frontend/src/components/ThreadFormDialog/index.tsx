import React, {FunctionComponent, useEffect, useState} from "react";
import {
    Button,
    Dialog,
    DialogActions,
    DialogContent,
    DialogTitle,
    Grid,
    IconButton,
    InputAdornment,
    TextField
} from "@material-ui/core";
import DeleteIcon from '@material-ui/icons/Delete';
import Account from "../../api/entities/Account";
import Thread from "../../api/entities/Thread";
import AddIcon from "@material-ui/icons/Add";
import {TweetStatus} from "../../api/entities/Tweet";
import CustomDate from "../../utils/CustomDate";
import {Alert} from "@material-ui/lab";
import {DragDropContext, Draggable, Droppable} from "react-beautiful-dnd";

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
        initial,
        onSubmit,
        onCancel
    }
) => {
    const [_idCounter, _setIdCounter] = useState<number>(0)
    const getId = () => {
        _setIdCounter(_idCounter + 1)
        return _idCounter
    }

    const [_initial, _setInitial] = useState<Thread>(initial)
    const [thread, setThread] = useState<Thread>(initial)

    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        if (initial != _initial) {
            // initial changed
            _setInitial(initial)
            setThread(initial)
            setError(null)
            _setIdCounter(0)
        }
    })

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
                    <DragDropContext onDragEnd={() => {
                    }}>
                        <Droppable droppableId={"1"}>
                            {(provided, snapshot) => (
                                <Grid
                                    container
                                    item
                                    xs={12}
                                    {...provided.droppableProps}
                                    ref={provided.innerRef}
                                >
                                    {
                                        thread.tweets.map((tweet, index) => (
                                            <Draggable draggableId={tweet.id!} index={index}>
                                                {(provided, snapshot) => (
                                                    <Grid
                                                        container
                                                        item
                                                        xs={12}
                                                        ref={provided.innerRef}
                                                        {...provided.draggableProps}
                                                        {...provided.dragHandleProps}
                                                        style={{
                                                            ...provided.draggableProps.style,
                                                            background: snapshot.isDragging ? "lightgrey" : undefined
                                                        }}
                                                    >
                                                        <Grid item xs={11}>
                                                            <TextField
                                                                error={
                                                                    thread.tweets[index].text.trim().length == 0 ||
                                                                    thread.tweets[index].text.length > 280
                                                                }
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
                                                                InputProps={{
                                                                    endAdornment: (
                                                                        <InputAdornment position="end">
                                                                            {280 - thread.tweets[index].text.length}
                                                                        </InputAdornment>
                                                                    )
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
                                                                <DeleteIcon fontSize="medium"/>
                                                            </IconButton>
                                                        </Grid>
                                                    </Grid>
                                                )}
                                            </Draggable>
                                        ))
                                    }
                                    {provided.placeholder}
                                </Grid>
                            )}
                        </Droppable>
                    </DragDropContext>
                    <Grid item xs={12}>
                        <IconButton aria-label="add" onClick={() => {
                            thread.tweets.push({
                                id: "new" + getId(),
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
                {error && <Alert severity="error">{error}</Alert>}
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
                    if (!thread.tweets.every(t => t.text.length <= 280)) {
                        setError("Tweets can't be longer than 280 characters!")
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