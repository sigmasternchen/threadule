import {Snackbar} from "@material-ui/core";
import {FunctionComponent} from "react";
import {Alert} from "@material-ui/lab";

export type MessageBoxProps = {
    open: boolean
    success: boolean
    message: string
    onClose: () => void
}

export const ClosedMessageBox: MessageBoxProps = {
    open: false,
    success: false,
    message: "",
    onClose: () => {}
}

export const MessageBox: FunctionComponent<MessageBoxProps> = ({open, success, message, onClose}) => {
    return (
        <Snackbar open={open} autoHideDuration={5000} onClose={onClose}>
            <Alert onClose={onClose} severity={success ? "success": "error"}>
                {message}
            </Alert>
        </Snackbar>
    )
}